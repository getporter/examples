# Deploy a multi-arch image with Helm

While Porter bundles can only run on amd64 hosts, you can deploy mult-arch images to other platforms, such as an ARM64 environment.
This example demonstrates how to deploy a helm chart that uses a multi-arch image to an ARM64 kubernetes cluster.

1. Create a file named "mycluster.yaml" with the following contents. Edit the path to the kubeconfig file to the location where your kubeconfig file.

    ```
    schemaType: CredentialSet
    schemaVersion: 1.0.1
    name: mycluster
    credentials:
    - name: kubeconfig
      source:
        path: $HOME/.kind/config # TODO: Edit this path
    ```
2. Apply the credential set
    ```
    porter credentials apply mycluster.yaml
    ```
3. Install the bundle
    ```
    porter install -r ghcr.io/getporter/examples/helm-multiarch:v0.1.0 -c mycluster
    ```

Now that the bundle is installed successfully, let's verify that the deployed image was ARM64.

1. Ensure that your KUBECONFIG environment variable is pointing to the cluster to which you just deployed.
2. Get the image that was deployed and save that value to an environment variable
    ```
    IMG=`k get pods -o yaml -l app.kubernetes.io/instance=porter-helm-nginx -o=jsonpath='{.items[0].status.containerStatuses[0].imageID}'`
    ```
3. Pull the deployed image so that you can inspect it
    ```
    docker pull $IMG
    ```
4. Inspect it to see the architecture of the image
    ```
    $ docker image inspect $IMG -f "{{.Architecture}}"
    arm64
    ```

As you can see, while Porter bundles can't run directly on a ARM64 host, they can deploy to ARM64 environments using multi-arch images! 🎉
