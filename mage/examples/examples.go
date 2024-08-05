package examples

import (
	"fmt"
	"os"
	"path/filepath"

	"gopkg.in/yaml.v3"
)

// GetBundleRef builds the reference to the specified example bundle, given a registry override.
func GetBundleRef(bundleDir string, registryOverride string) (string, error) {
	manifestPath := filepath.Join(bundleDir, "porter.yaml")
	contents, err := os.ReadFile(manifestPath)
	if err != nil {
		return "", fmt.Errorf("error reading porter manifest at %s: %w", manifestPath, err)
	}

	manifest := map[string]interface{}{}
	if err = yaml.Unmarshal(contents, &manifest); err != nil {
		return "", fmt.Errorf("error parsing porter manifest at %s: %w", manifestPath, err)
	}

	name, ok := manifest["name"].(string)
	if !ok {
		return "", fmt.Errorf("name was not defined in %s", manifestPath)
	}

	version, ok := manifest["version"].(string)
	if !ok {
		return "", fmt.Errorf("version was not defined in %s", manifestPath)
	}

	registry, ok := manifest["registry"].(string)
	if !ok {
		return "", fmt.Errorf("registry was not defined in %s", manifestPath)
	}

	if registryOverride != "" {
		registry = registryOverride
	}

	return fmt.Sprintf("%s/%s:v%s", registry, name, version), nil
}

// List returns the names of all example bundles in the specified directory.
func List(dir string) ([]string, error) {
	results, err := os.ReadDir(dir)
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
