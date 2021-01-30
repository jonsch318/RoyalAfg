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
At first the Vault Secret Service get installed which will be configured to serve as a certificate authority for the server to enable secure TLS and mTLS communication. 

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

### Configure PKI secret engine
Start an interactive shell session in the `vault-0` pod
```bash
kubectl exec --stdin=true --tty=true vault-0 -- /bin/sh
---
/ $
```
with the new prompt enable the pki secret engine

```bash
vault secrets enable pki
---
Success! Enabled the pki secret engine at: pki/
```

Next configure the max lease time for certificates to 1 year 8760h.

```bash
vault secrets tune -max-lease-ttl=8760h pki
---
Success! Tuned the secrets engine at: /pki
```

Now generate the root key pairs for the certificate authority and configure the pki secret engine to the vault service.

```bash
vault write pki/root/generate/internal \ 
	common_name=example.com \ 
	ttl=8760h
```

```bash
vault write pki/config/urls \
	issuing_certificates="http://vault.default:8200/v1/pki/ca" \
	crl_distribution_points="http://vault.default:8200/v1/pki/crl"
```
Then create a `royalafg-dot-games` role that enables the create of certificates of `royalafg.games` domain with any subdomains.

```bash
vault write pki/roles/royalafg-dot-games \ 
	allowed_domains=royalafg.games \ 
	allow_subdomains=true \ 
	max_ttl=72h
```
Then finally create a policy for the pki engine.

```bash
vault policy write pki - <<EOF 
path "pki*" { capabilities = ["read", "list"] } 
path "pki/roles/royalafg-dot-games" { capabilities = ["create", "update"] } 
path "pki/sign/royalafg-dot-games" { capabilities = ["create", "update"] } 
path "pki/issue/royalafg-dot-games" { capabilities = ["create"] } 
EOF
```

### Configure Vault Kubernetes authentication
To access and authenticate to the vault service use the Kubernetes authentication method of vault that will use Service Account Tokens to authenticate a service. 

Use the interactive shell from before, or create a new session
```bash
kubectl exec --stdin=true --tty=true vault-0 -- /bin/sh
```
Enable the kubernetes auth method
```bash
vault auth enable kubernetes
```
and configure it.
```bash
vault write auth/kubernetes/config \ 
	token_reviewer_jwt="$(cat /var/run/secrets/kubernetes.io/serviceaccount/token)" \ 
	kubernetes_host="https://$KUBERNETES_PORT_443_TCP_ADDR:443" \ 
	kubernetes_ca_cert=@/var/run/secrets/kubernetes.io/serviceaccount/ca.crt
```
 Finally create a auth role that binds the pki policy.
```bash
vault write auth/kubernetes/role/issuer \ 
	bound_service_account_names=issuer \ 
	bound_service_account_namespaces=default \ 
	policies=pki \ 
	ttl=20m
```
Then exit the shell session
```bash
exit
```
## Installation of Cert-Manager
To install the cert-manager run:
```bash
helm install \
  cert-manager jetstack/cert-manager \
  --namespace cert-manager \
  --version v1.1.0 \
  --set installCRDs=true
```