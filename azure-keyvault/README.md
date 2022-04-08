# Azure Key Vault

## Running the Bundle

### Setup

Generate credentials by going through the interactive cli.

```sh
 porter generate credentials -r ghcr.io/asalbers/hello-keyvault:v0.1.2
```
Sample output from generate credentials shows sourcing from environment variables.

```sh
Building bundle ===>
Copying porter runtime ===>
Copying mixins ===>
Copying mixin exec ===>
Copying mixin az ===>

Generating Dockerfile =======>

Writing Dockerfile =======>

Starting Invocation Image Build (ghcr.io/asalbers/hello-keyvault-installer:v0.1.2) =======>
Generating new credential hello-keyvault from bundle hello-keyvault
==> 4 credentials required for bundle hello-keyvault
? How would you like to set credential "AZURE_CLIENT_ID"
   [Use arrows to move, space to select, type to filter]
  secret
  specific value
> environment variable
  file path
  shell command
 Enter the environment variable that will be used to set credential "AZURE_CLIENT_ID"
  AZURE_CLIENT_ID
? How would you like to set credential "AZURE_SP_PASSWORD"
  environment variable
? Enter the environment variable that will be used to set credential "AZURE_SP_PASSWORD"
  AZURE_SP_PASSWORD
? How would you like to set credential "AZURE_SUBSCRIPTION_ID"
  environment variable
? Enter the environment variable that will be used to set credential "AZURE_SUBSCRIPTION_ID"
  AZURE_SUB_ID
? How would you like to set credential "AZURE_TENANT_ID_OR_DNS"
  environment variable
? Enter the environment variable that will be used to set credential "AZURE_TENANT_ID_OR_DNS"
  AZURE_TENANT_ID_OR_DNS

porter credentials list
NAME             MODIFIED
hello-keyvault   2 minutes ago

```

### Install the bundle

The command below installs the porter bundle. Replace the items in <> with your values.

```sh
porter install <name of bundle> -c <credentials name> -r ghcr.io/asalbers/hello-keyvault:v0.1.2
```
Sample output: The error at the end can be ignored for now as the keyvault successfully deploys

```sh
Status: Downloaded newer image for ghcr.io/asalbers/hello-keyvault@sha256:7808c26698bdc1a61b538667177e42be1d8ef5a9223fc59e995d2d37b222f448
executing install action from hello-keyvault (installation: andrew-test)
Logging into Azure...
Setting the Azure subscription....
Creating or using the Azure resource group....
hellokeyvault
Setting random string....
uaxkx9w8
Creating the KeyVault...
execution completed successfully!
```

### Uninstalling the Bundle

```sh
porter uninstall <bundle name> -c <credential name> -r ghcr.io/asalbers/hello-keyvault:v0.1.2
```

```sh
uninstalling andrew-test...
executing uninstall action from hello-keyvault (installation: andrew-test)
Logging into Azure...
Setting the Azure subscription....
Deleting the KeyVault...
execution completed successfully!
```

# Contents

## porter.yaml

### Parameters

| Name                      | Description                                                                                                    | Type   | Default       | Required | Applies to         |
|---------------------------|----------------------------------------------------------------------------------------------------------------|--------|---------------|----------|--------------------|
| AZURE_GROUP_REGION        | The azure resource group to use or create for the cluster resources.                                           | string | eastus2       | false    | All Actions        |
| AZURE_RESOURCE_GROUP_NAME | The azure resource group to use or create for the cluster resources.                                           | string | hellokeyvault | false    | ALL Actions        |
| keyvault_name             | The name of the created KeyVault.                                                                              | <nil>  | <nil>         | true     | upgrade, uninstall |
| random_string             | The random string used to create resources. This is automatically generated and reused, but can be overridden. | <nil>  | <nil>         | true     | upgrade, uninstall |

### Credentials

| Name                | Description | Required | Applies to |
|---------------------------|----------------------------------------------------------------------------------|----------|-------------|
| AZURE_CLIENT_ID           | The client id for the service principal used to automate the bundle's actions.   | true     | All Actions |
| AZURE_SP_PASSWORD         | The service principal password that is used to log into Azure inside the bundle. | true     | All Actions |
| AZURE_SUBSCRIPTION_ID     | The Azure subscription into which to deploy.                                     | true     | All Actions |
| AZURE_TENANT_ID_OR_DNS    | The tenant identity in which the service principal resides.                      | true     | All Actions |

## helpers.sh

random_string() 
Generates a random string to be used in the keyvault name.

## upsertGroup.sh

Script that checks for an existing azure resource group and creates it if it doesn't exist.

## Dockerfile.tmpl

This is a template Dockerfile for the bundle's invocation image. You can
customize it to use different base images, install tools and copy configuration
files. Porter will use it as a template and append lines to it for the mixin and to set
the CMD appropriately for the CNAB specification. You can delete this file if you don't
need it.

Add the following line to **porter.yaml** to enable the Dockerfile template:

```yaml
dockerfile: Dockerfile.tmpl
```

By default, the Dockerfile template is disabled and Porter automatically copies
all of the files in the current directory into the bundle's invocation image. When
you use a custom Dockerfile template, you must manually copy files into the bundle
using COPY statements in the Dockerfile template.

## .gitignore

This is a default file that we provide to help remind you which files are
generated by Porter, and shouldn't be committed to source control. You can
delete it if you don't need it.

## .dockerignore

This is a default file that controls which files are copied into the bundle's
invocation image by default. You can delete it if you don't need it.