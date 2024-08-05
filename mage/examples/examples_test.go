package examples

import (
	"os"
	"testing"

	"github.com/carolynvs/magex/mgx"
	"github.com/carolynvs/magex/shx"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestListExampleBundles(t *testing.T) {
	tmp, err := os.MkdirTemp("", "example-bundles")
	require.NoError(t, err)
	defer os.RemoveAll(tmp)

	mgx.Must(shx.Copy("../../hello", tmp, shx.CopyRecursive))
	mgx.Must(shx.Copy("../../mage", tmp, shx.CopyRecursive))

	names, err := List(tmp)
	require.NoError(t, err)

	assert.Equal(t, []string{"hello"}, names)
}

func TestGetBundleRef(t *testing.T) {
	t.Run("valid manifest", func(t *testing.T) {
		ref, err := GetBundleRef("../../hello", "localhost:5000")
		require.NoError(t, err)
		assert.Equal(t, "localhost:5000/examples/porter-hello:v0.2.0", ref)
	})

	t.Run("missing registry", func(t *testing.T) {
		_, err := GetBundleRef("testdata/missing-registry", "localhost:5000")
		assert.EqualError(t, err, "registry was not defined in testdata/missing-registry/porter.yaml")
	})

	t.Run("missing name", func(t *testing.T) {
		_, err := GetBundleRef("testdata/missing-name", "localhost:5000")
		assert.EqualError(t, err, "name was not defined in testdata/missing-name/porter.yaml")
	})

	t.Run("missing version", func(t *testing.T) {
		_, err := GetBundleRef("testdata/missing-version", "localhost:5000")
		assert.EqualError(t, err, "version was not defined in testdata/missing-version/porter.yaml")
	})
}
