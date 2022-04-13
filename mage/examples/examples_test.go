package examples

import (
	"github.com/carolynvs/magex/mgx"
	"github.com/carolynvs/magex/shx"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"io/ioutil"
	"os"
	"testing"
)

func TestListExampleBundles(t *testing.T) {
	tmp, err := ioutil.TempDir("", "example-bundles")
	require.NoError(t, err)
	defer os.RemoveAll(tmp)

	mgx.Must(shx.Copy("../../hello", tmp, shx.CopyRecursive))
	mgx.Must(shx.Copy("../../mage", tmp, shx.CopyRecursive))

	names, err := List(tmp)
	require.NoError(t, err)

	assert.Equal(t, []string{"hello"}, names)
}

func TestGetBundleRef(t *testing.T) {
	ref, err := GetBundleRef("../../hello", "localhost:5000")
	require.NoError(t, err)
	assert.Equal(t, "localhost:5000/examples/hello:v0.2.0", ref)
}
