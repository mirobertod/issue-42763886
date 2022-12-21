# issue-42763886

Steps to reproduce the issue:

1. Create GKE standard cluster, all default settings except:
    - Select 2 zones for the default node pool
    - Number of nodes per zone: 1

2. Download `istioctl` 1.15.3 from the [official GitHub releases page](https://github.com/istio/istio/releases/tag/1.15.3)

3. Connect to the newly created GKE cluster (i.e. `USE_GKE_GCLOUD_AUTH_PLUGIN=True gcloud container clusters get-credentials cluster-1 --zone us-central1-c --project my-dev-project`)

4. Install istio with the following command: `istioctl install --filename ./spec-development.yml`

5. Create a self-signed SSL certificate: `./create-cert.sh`

6. Create a gateway resource: `kubectl apply --filename ./gateway.yml`

7. Create the nginx deployment: `kubectl apply --filename ./test-project.yaml`

8. Create a proxied DNS record (orange icon) on your Cloudflare account with the IP of your GLB

9. Check if the pods are properly spread across the two zones: `kubectl get pods -A -o wide | egrep 'istio|test-project' | grep Running`

10. Launch the script which make an HTTP GET request: `go run main.go "https://my-url.example.com/"`

11. Simulate a preemption process as stated [here in the docs](https://cloud.google.com/compute/docs/instances/preemptible#preemption-process): `gcloud --project my-dev-project compute instances stop gke-cluster-1-default-pool-f093f804-xlkf`

The running script will be terminated with the error code `HTTP 520`.  
If that doesn't happen repeat the process (starting from step 9.) a few times or try to stop the other instance.

Consideration:
- Remember to replace `my-url.example.com` with your proper DNS record
