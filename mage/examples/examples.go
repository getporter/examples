package examples

import (
	"fmt"
	"gopkg.in/yaml.v3"
	"io/ioutil"
	"os"
	"path/filepath"
)

// GetBundleRef builds the reference to the specified example bundle, given a registry override.
func GetBundleRef(bundleDir string, registryOverride string) (string, error) {
	manifestPath := filepath.Join(bundleDir, "porter.yaml")
	contents, err := ioutil.ReadFile(manifestPath)
	if err != nil {
		return "", fmt.Errorf("error reading porter manifest at %s: %w", manifestPath, err)
	}

	manifest := map[string]interface{}{}
	if err = yaml.Unmarshal(contents, &manifest); err != nil {
		return "", fmt.Errorf("error parsing porter manifest at %s: %w", manifestPath, err)
	}

	name := manifest["name"].(string)
	version := manifest["version"].(string)
	registry := manifest["registry"].(string)
	if registryOverride != "" {
		registry = registryOverride
	}

	return fmt.Sprintf("%s/%s:v%s", registry, name, version), nil
}

// List returns the names of all example bundles in the specified directory.
func List(dir string) ([]string, error) {
	results, err := ioutil.ReadDir(dir)
	if err != nil {
		return nil, fmt.Errorf("error listing example bundles in current directory: %w", err)
	}

	exampleNames := make([]string, 0, len(results))
	for _, result := range results {
		if result.IsDir() {
			bundleName := result.Name()
			if _, err := os.Stat(filepath.Join(dir, bundleName, "porter.yaml")); err == nil {
				exampleNames = append(exampleNames, bundleName)
			}
		}
	}

	return exampleNames, nil
}
