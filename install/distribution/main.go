package main

import (
	"context"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"os/exec"
	"path"
	"path/filepath"
	"strings"
	"sync"
	"time"

	"go.uber.org/zap"

	"github.com/google/uuid"

	"github.com/ghodss/yaml"
	"github.com/solo-io/go-utils/docker"
	v1 "k8s.io/api/apps/v1"
	coreV1 "k8s.io/api/core/v1"
)

const (
	glooe              = "glooe"
	glooeYaml          = glooe + "-distribution.yaml"
	manifest           = "install/manifest/" + glooeYaml
	scriptsDir         = "install/distribution/scripts"
	output             = "_output"
	distribution       = output + "/distribution"
	deployment         = "Deployment"
	glooctl            = "glooctl"
	zipExt             = ".zip"
	tarExt             = ".tar"
	setupScriptName    = "setup"
	imageNameSeparator = "_"
)

var (
	version               string
	outputDistributionDir string
	id                    uuid.UUID

	logger *zap.SugaredLogger

	projectId = os.ExpandEnv("PROJECT_ID")
)

func init() {
	var err error
	devLogger, err := zap.NewDevelopment()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	logger = devLogger.Sugar()
	id, err = uuid.NewRandom()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func main() {
	defer logger.Sync()
	if len(os.Args) < 2 {
		panic("Must provide version as argument")
	} else {
		version = os.Args[1]
		outputDistributionDir = filepath.Join(distribution, version)
	}
	if err := prepareWorkspace(); err != nil {
		logger.Fatal(err.Error())
	}

	if err := prepareFiles(); err != nil {
		logger.Fatal(err.Error())
	}

	distBucketCli, err := newDistributionBucketClient()
	if err != nil {
		logger.Fatal(err.Error())
	}

	if err := syncDataToBucket(distBucketCli); err != nil {
		logger.Fatal(err.Error())
	}

}

func prepareFiles() error {
	specs, err := readManifestIntoParts()
	if err != nil {
		return err
	}
	deployments, err := extractContainerImagesFromSpecs(specs)
	if err != nil {
		return err
	}

	if err := saveImages(deployments); err != nil {
		return err
	}

	if err := copyManifest(); err != nil {
		return err
	}

	if err := copyBinaries(); err != nil {
		return err
	}

	if err := copySetupScripts(); err != nil {
		return err
	}
	return nil
}

func prepareWorkspace() error {
	directories := []string{outputDistributionDir}
	for _, v := range directories {
		if _, err := os.Stat(v); os.IsNotExist(err) {
			err = os.MkdirAll(v, 0755)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

func readManifestIntoParts() ([]string, error) {
	bytes, err := ioutil.ReadFile(manifest)
	if err != nil {
		return nil, err
	}
	return strings.Split(string(bytes), "---"), nil
}

func extractContainerImagesFromSpecs(specs []string) ([]v1.Deployment, error) {
	deployments := make([]v1.Deployment, 0, len(specs))
	for _, spec := range specs {
		var dpl v1.Deployment
		err := yaml.Unmarshal([]byte(spec), &dpl)
		if err != nil {
			return nil, err
		}
		if dpl.Kind == deployment {
			deployments = append(deployments, dpl)
		}
	}
	return deployments, nil
}

func savedImageName(fullImageName string) string {
	imageName := strings.Replace(path.Base(fullImageName), ":", imageNameSeparator, -1)
	return filepath.Join(outputDistributionDir, imageName) + tarExt
}

// Pull and save the images of of containers and initContainers of the given deployments
func saveImages(deployments []v1.Deployment) error {
	wg := &sync.WaitGroup{}
	errch := make(chan error)

	for _, dpl := range deployments {
		containers := append(dpl.Spec.Template.Spec.Containers, dpl.Spec.Template.Spec.InitContainers...)
		for _, image := range containers {
			wg.Add(1)
			go func(image coreV1.Container, wg *sync.WaitGroup) {
				defer wg.Done()
				var err error
				_, err = docker.PullIfNotPresent(context.TODO(), image.Image, 0)
				if err != nil {
					errch <- err
					return
				}
				err = docker.Save(image.Image, savedImageName(image.Image))
				if err != nil {
					errch <- err
					return
				}
				logger.Infof("Successfully saved image: (%s)", image.Image)
			}(image, wg)
		}
	}

	go func(wg *sync.WaitGroup, errch chan error) {
		wg.Wait()
		close(errch)
	}(wg, errch)

	select {
	case err := <-errch:
		if err != nil {
			return err
		}
	case <-time.After(10 * time.Minute):
		return fmt.Errorf("go routine timed out after 10 minutes")
	}
	return nil

}

func copyFile(source, dest string) error {
	cmd := exec.Command("cp", source, dest)
	if err := cmd.Run(); err != nil {
		return err
	}
	return nil
}

func copyManifest() error {
	destinationManifest := filepath.Join(outputDistributionDir, glooeYaml)
	if err := copyFile(manifest, destinationManifest); err != nil {
		return err
	}
	logger.Infof("Successfully copied local manifest to distribution directory")
	return nil
}

func copyBinaries() error {
	info, err := ioutil.ReadDir(output)
	if err != nil {
		return err
	}
	for _, file := range info {
		if strings.Contains(file.Name(), glooctl) {
			source := filepath.Join(output, file.Name())
			dest := filepath.Join(outputDistributionDir, file.Name())
			if err := copyFile(source, dest); err != nil {
				return err
			}
			logger.Infof("Successfully copied binary: (%s) to distribution directory", file.Name())
		}
	}
	return nil
}

// Copy setup scripts from install/distribution/scripts to output/distribution
func copySetupScripts() error {
	for _, extension := range []string{"sh", "bat"} {
		filename := strings.Join([]string{setupScriptName, extension}, ".")
		source := filepath.Join(scriptsDir, filename)
		dest := filepath.Join(outputDistributionDir, filename)
		if err := copyFile(source, dest); err != nil {
			return err
		}
		logger.Infof("Successfully copied setup script (%s) to distribution directory", filename)
	}
	return nil
}

func writeDistributionArchive(wr io.Writer, fileName string) error {
	// zip -r ARCHIVE_NAME DIR_TO_COMPRESS
	cmd := exec.Command("zip", "-r", fileName, version)
	cmd.Dir = distribution
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	logger.Infof("creating zip file...")
	if err := cmd.Run(); err != nil {
		return err
	}
	logger.Info("successfully created zip file")

	r, err := os.Open(filepath.Join(distribution, fileName))
	if err != nil {
		return err
	}
	_, err = io.Copy(wr, r)
	return err
}
