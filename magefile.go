//go:build mage
// +build mage

// This is a magefile, and is a "makefile for go".
// See https://magefile.org/
package main

import (
	"fmt"
	"get.porter.sh/example-bundles/mage/examples"
	"get.porter.sh/example-bundles/mage/setup"
	"get.porter.sh/magefiles/porter"
	"github.com/carolynvs/magex/mgx"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	// mage:import
	_ "get.porter.sh/magefiles/ci"
	"github.com/carolynvs/magex/shx"
	"github.com/hashicorp/go-multierror"
	"github.com/magefile/mage/mg"
)

func Build() {
	mg.Deps(BuildExamples)
}

// BuildExamples builds every example bundle
func BuildExamples() error {
	var bigErr *multierror.Error
	names, err := examples.List(".")
	if err != nil {
		return err
	}

	for _, bundleName := range names {
		if err := BuildExample(bundleName); err != nil {
			// Keep trying to build all the bundles, don't stop on the first one
			bigErr = multierror.Append(bigErr, fmt.Errorf("error building bundle %s: %w", bundleName, err))
			continue
		}
	}

	return bigErr.ErrorOrNil()
}

// BuildExample builds the specified example bundle
func BuildExample(name string) error {
	mg.SerialDeps(setup.InstallMixins)

	fmt.Println("\n==========================")
	fmt.Printf("Building example bundle: %s\n", name)

	if customBuildFlags, err := ioutil.ReadFile(filepath.Join(name, "build-args.txt")); err == nil {
		customBuildArgs := strings.Split(string(customBuildFlags), " ")
		buildArgs := append([]string{"build"}, customBuildArgs...)
		return shx.Command("porter", buildArgs...).
			CollapseArgs().In(name).RunV()
	}
	// Always build for amd64 even if on an arm host
	// This is a bit of a hack until we have multi-arch support
	return shx.Command("porter", "build", "--debug", "--verbose").
		In(name).Env("DOCKER_DEFAULT_PLATFORM=linux/amd64").RunV()
}

func Publish() {
	mg.Deps(PublishExamples)
}

// PublishExamples publishes every example bundle
func PublishExamples() error {
	var bigErr *multierror.Error
	names, err := examples.List(".")
	if err != nil {
		return err
	}

	for _, bundleName := range names {
		if err := PublishExample(bundleName); err != nil {
			// Keep trying to publish all the bundles, don't stop on the first one
			bigErr = multierror.Append(bigErr, fmt.Errorf("error publishing bundle %s: %w", bundleName, err))
			continue
		}
	}

	return bigErr.ErrorOrNil()
}

// PublishExample publishes the specified example bundle
func PublishExample(name string) error {
	mg.SerialDeps(porter.UseBinForPorterHome, porter.EnsurePorter)

	fmt.Println("\n==========================")

	registryFlag := ""
	registry := os.Getenv("PORTER_REGISTRY")
	if registry != "" {
		registryFlag = "--registry=" + registry
	}

	// Check if the bundle already is published
	bundleRef, err := examples.GetBundleRef(name, registry)
	mgx.Must(err)

	// Do not overwrite an already published bundle
	// See https://github.com/getporter/porter/issues/2017
	if err := shx.RunS("porter", "explain", "-r", bundleRef); err == nil {
		fmt.Printf("Skipping publish for example bundle: %s. The bundle is already published to %s.\n", name, bundleRef)
		return nil
	}

	fmt.Printf("Publishing example bundle: %s\n", name)
	return shx.Command("porter", "publish", registryFlag).CollapseArgs().In(name).RunV()
}
