# Bundle with sensitive data

This bundle demonstrates how porter works with bundle that contains sensitive
data. 

It requires users to set up a secret store in their porter configuration
file in order to work with this bundle.

## Try it out 

Follow the steps in the [Change default secret plugin](https://release-v1.porter.sh/end-users/configuration/#change-the-default-secrets-plugin) documentation
to setup [filesystem](https://release-v1.porter.sh/plugins/filesystem/) plugin.
After setting up the secret plugin, run to install the bundle:
```
porter install --reference ghcr.io/getporter/examples/porter-sensitive-data:v0.1.0 --param password=test
```


