//go:build mage
// +build mage

// This is a magefile, and is a "makefile for go".
// See https://magefile.org/
package main

import (
	"fmt"
	"get.porter.sh/magefiles/porter"
	"golang.org/x/sync/errgroup"
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
	mg.SerialDeps(BuildExamples)
}

// BuildExamples builds every example bundle
func BuildExamples() error {
	results, err := ioutil.ReadDir(".")
	if err != nil {
		return err
	}

	var bigErr *multierror.Error
	for _, result := range results {
		if result.IsDir() {
			bundleName := result.Name()
			if _, err := os.Stat(filepath.Join(bundleName, "porter.yaml")); err == nil {
				err := BuildExample(bundleName)
				if err != nil {
					// Keep trying to build all the bundles, don't stop on the first one
					bigErr = multierror.Append(bigErr, fmt.Errorf("error building bundle %s: %w", bundleName, err))
					continue
				}
			}
		}
	}

	return bigErr.ErrorOrNil()
}

// BuildExample builds the specified example bundle
func BuildExample(name string) error {
	mg.SerialDeps(installMixins)

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
	return shx.Command("porter", "build").
		In(name).Env("DOCKER_DEFAULT_PLATFORM=linux/amd64").RunV()
}

func installMixins() error {
	mg.Deps(porter.UseBinForPorterHome)

	mixins := []porter.InstallMixinOptions{
		{Name: "arm"},
		{Name: "az"},
		{Name: "docker"},
		{Name: "docker-compose"},
		{Name: "exec"},
		{Name: "helm3", Feed: "https://mchorfa.github.io/porter-helm3/atom.xml", Version: "v0.1.14"},
		{Name: "kubernetes"},
		{Name: "terraform"},
	}
	var errG errgroup.Group
	for _, mixin := range mixins {
		mixin := mixin
		errG.Go(func() error {
			return porter.EnsureMixin(mixin)
		})
	}
	return errG.Wait()
}
