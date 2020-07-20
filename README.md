# Optimizing resource usage in a multi-tenant GKE cluster using node auto-provisioning
This repo demonstrates the use of [node auto-provisioning](https://cloud.google.com/kubernetes-engine/docs/how-to/node-auto-provisioning) to auto-scale a multi-tenant [Google Kubernetes Engine (GKE)](https://cloud.google.com/kubernetes-engine/) cluster. Tenant workloads are kept separate, and [workload identity](https://cloud.google.com/kubernetes-engine/docs/how-to/workload-identity) is used to control tenant access to GCP resources like Cloud Storage buckets. 

Cluster and Job configuration is managed with Kustomize.

See the Google Cloud [tutorial](https://cloud.google.com/solutions/optimizing-resources-in-multi-tenant-gke-clusters-with-auto-provisioning) for full details.
