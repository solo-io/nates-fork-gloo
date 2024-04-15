package export

import (
	"github.com/solo-io/gloo/pkg/utils/cmdutils"
	"os"
	"path/filepath"
)

func RunCommandOutputToFile(cmd cmdutils.Cmd, path string) func() error {
	return func() error {
		f, err := fileOnHost(path)
		if err != nil {
			return err
		}
		defer func() {
			_ = f.Close()
		}()
		// We intentionally do not output stderr to the output file
		// Otherwise, the output file will contain errors that may be misleading to users
		// For example, if a curl request failed first, and then succeeded, we do not want to
		// populate the failures in the output file
		return cmd.WithStdout(f).Run().Cause()
	}
}

// FileOnHost is a helper to create a file at path even if the parent directory doesn't exist
// in which case it will be created with ModePerm
func fileOnHost(path string) (*os.File, error) {
	if err := os.MkdirAll(filepath.Dir(path), os.ModePerm); err != nil {
		return nil, err
	}
	return os.Create(path)
}
