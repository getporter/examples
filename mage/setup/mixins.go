package setup

import (
	"get.porter.sh/magefiles/porter"
	"github.com/magefile/mage/mg"
	"golang.org/x/sync/errgroup"
)

// InstallMixins used by the example bundles.
// If you add an example that uses a new mixin, update this function to install it.
func InstallMixins() error {
	mg.SerialDeps(porter.UseBinForPorterHome, porter.EnsurePorter)

	mixins := []porter.InstallMixinOptions{
		{Name: "arm"},
		{Name: "az"},
		{Name: "docker"},
		{Name: "docker-compose"},
		{Name: "exec"},
		// Use a build of helm3 that supports nonroot
		// https://github.com/MChorfa/porter-helm3/pull/42
		{Name: "helm3", Feed: "https://mchorfa.github.io/porter-helm3/atom.xml", Version: "v1.0.0-rc.1"},
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
