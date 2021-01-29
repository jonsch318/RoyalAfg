# Deployment
Because of the micro-service architecture of this project it uses kubernetes to orchestrate all services and dependencies. Because of the many services and technologies used and showcased by this project it is more complicated to deploy it to a final state where all services are working together.

## Prerequisites
This deployment guide requires the user to install following tools to be installed

 - A configured kubernetes cluster that can be accessed with [kubectl](https://kubernetes.io/de/docs/tasks/tools/install-kubectl/). 
	 -  for example [Minikube](https://minikube.sigs.k8s.io/docs/)  
 - A valid [helm](https://helm.sh/) installation
 
 ## 1: Start Minikube Cluster
 When you are using minikube to provide the kubernetes cluster use the following command, if your cluster is already accessable with kubectl feel free to go to step 2:
 
```bash
minikube start
```

Though you might consider using another [driver](https://minikube.sigs.k8s.io/docs/drivers/) than the normal vm driver. If you have docker installed on your machine (or WSL 2) then run `minikube start --driver=docker` for one time or `minikube config set driver docker` for global config.

## 2: Add Helm charts
Many  services used by this project can be deployed easily with helm. To do that run the following command to add the used repositories to your local helm repository.

```bash
helm repo add jetstack https://charts.jetstack.io
helm repo add agones https://agones.dev/chart/stable
helm repo add hashicorp https://helm.releases.hashicorp.com
helm repo add elastic https://helm.elastic.co
```

Finally run 

    helm repo update

To update the local repo.

## 3: Configure Vault and create Cluster CA
At first the Vault service will be installed which will be configured to serve as a certificate authority for the server to enable secure TLS and mTLS communication. 

### Installation and Unsealing of Vault

```bash
helm install vault hashicorp/vault
```
Next get the pods running on the cluster, which will look something like this:
```bash
kubectl get pods  
---
NAME READY STATUS RESTARTS AGE 
vault-0 0/1 Running 0 87s
```
Upon installation Vault is sealed and no communication is accepted. To unseal vault use initialize it:

```bash
kubectl exec vault-0 -- vault operator init -key-shares=1 -key-threshold=1 \
	-format=json > init-keys.json
``` 
This creates a file with the name `init-keys.json` which contain the root token and unseal key. **1 Key Share is not secure and should not be used in production**

To unseal vault read `the unseal_keys_b64` value may be something like this: `hmeMLoRiX/trBTx/xPZHjCcZ7c4H8OCt2Njkrv2yXZY=` an then unseal vault.

    kubectl exec vault-0 -- vault operator unseal <previous vault key here>

Now login to vault using the `root_token` value from `init-keys.json`.

    kubectl exec vault-0 -- vault login <root_token here>

Vault is now initialized and ready for configuration.

### Login
