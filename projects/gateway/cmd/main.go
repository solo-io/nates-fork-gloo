package main

import (
	"github.com/solo-io/solo-kit/pkg/utils/log"
	"github.com/solo-io/solo-kit/projects/gateway/pkg/setup"
	gloosetup "github.com/solo-io/solo-kit/projects/gloo/pkg/setup"
	"github.com/solo-io/solo-kit/projects/gloo/pkg/syncer"
)

func main() {
	if err := run(); err != nil {
		log.Fatalf("err in main: %v", err.Error())
	}
}

func run() error {
	errs := make(chan error)
	go func() {
		errs <- runGloo()
	}()
	go func() {
		errs <- runGateway()
	}()
	return <-errs
}

func runGloo() error {
	opts, err := gloosetup.DefaultKubernetesConstructOpts()
	if err != nil {
		return err
	}
	return syncer.RunGloo(opts)
}

func runGateway() error {
	opts, err := setup.DefaultKubernetesConstructOpts()
	if err != nil {
		return err
	}
	return setup.RunGateway(opts)
}
