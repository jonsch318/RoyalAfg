
# Deployment

Because of the micro-service architecture of this project it uses kubernetes to orchestrate all services and dependencies. Because of the many services and technologies used and showcased by this project it is more complicated to deploy it to a final state where all services are working together.



## Prerequisites

This deployment guide requires the user to install following tools to be installed



- A configured kubernetes cluster that can be accessed with [kubectl](https://kubernetes.io/de/docs/tasks/tools/install-kubectl/).

- for example [Minikube](https://minikube.sigs.k8s.io/docs/)

- A valid [helm](https://helm.sh/) installation

## 1: Start Minikube Cluster

When you are using minikube to provide the kubernetes cluster use the following command, if your cluster is already accessable with kubectl feel free to go to step 2:

```bash

minikube start --memory=8192 --extra-config=kubelet.authentication-token-webhook=true --extra-config=kubelet.authorization-mode=Webhook

```



Though you might consider using another [driver](https://minikube.sigs.k8s.io/docs/drivers/) than the normal vm driver. If you have docker installed on your machine (or WSL 2) then run `minikube start --driver=docker` for one time or `minikube config set driver docker` for global config.



## 2: Add Helm charts

Many services used by this project can be deployed easily with helm. To do that run the following command to add the used repositories to your local helm repository.



```bash
helm repo add jetstack https://charts.jetstack.io
helm repo add agones https://agones.dev/chart/stable
helm repo add hashicorp https://helm.releases.hashicorp.com
helm repo add elastic https://helm.elastic.co
helm repo add eventstore https://eventstore.github.io/EventStore.Charts
helm repo add prometheus-community https://prometheus-community.github.io/helm-charts
```
Finally run

    helm repo update

To update the local repo.

# Installation Procedure
Be sure to be in the `cd ./deployments` directory of the git repo. I will continue to improve this guide to deploy the royalafg services.

Install Consul, Traefik and the Agones Library
```
helm install consul -f ./consul-values.yaml hashicorp/consul

kubectl apply -f ./traefik.yaml

helm install agones --namespace agones-system --create-namespace agones/agones
```
Then install both rabbitmq, eventstore and mongodb
```
kubectl apply -f ./rabbitmq/operator.yaml
kubectl apply -f ./rabbitmq/rabbitmq.yaml

kubectl apply -f ./eventstore.yaml 

kubectl create namespace mongodb-systems
kubectl apply -f ./mongodb/rbac/ -n default
kubectl apply -f ./mongodb/crds/
kubectl apply -f ./mongodb/community_operator.yaml
kubectl apply -f ./mongodb/royalafg-user.yaml 
```

Then install and configure vault

    helm install vault -f ./vault-values.yaml hashicorp/vault

	kubectl exec -it vault-0 -- /bin/sh
	=> $ new shell...
	vault operator init \
	    -key-shares=1 \
	    -key-threshold=1 \
	=> note the two keys
	=> Unseal Key
	=> Root Token
	
	vault operator unseal
	//Provide the Unseal Key
	
	vault login 
	//Provide the Root token
	
	//Enable the kubernetes authentication method 
	vault auth enable kubernetes
	
	//And configure it
	vault write auth/kubernetes/config \	
	    token_reviewer_jwt="$(cat /var/run/secrets/kubernetes.io/serviceaccount/token)" \
	    kubernetes_host=https://${KUBERNETES_PORT_443_TCP_ADDR}:443 \
	    kubernetes_ca_cert=@/var/run/secrets/kubernetes.io/serviceaccount/ca.crt
    
	//enable the database secret
	vault secrets enable database

	//enable the rabbitmq secret (dynamic keys)
	vault secrets enable rabbitmq

Get the rabbitmq credentials

    username => kubectl get secret royalafg-rabbitmq-default-user -o jsonpath='{.data.username}' | base64 --decode
	
	password => kubectl get secret royalafg-rabbitmq-default-user -o jsonpath='{.data.password}' | base64 --decode

Configure the vault rabbitmq secret

    kubectl exec -it vault-0 -- /bin/sh
    
	vault write rabbitmq/config/connection \
	   connection_uri="http://royalafg-rabbitmq.default.svc.cluster.local:15672" \
	   username="{rabbitmq username}" \
	   password="{rabbitmq password}"
	//Fill in username and password
	
	vault write rabbitmq/roles/royalafg \
	    vhosts='{"/":{"write": ".*", "read": ".*", "configure": ".*"}}'

The next configurations can be done using the vault ui

    kubectl port-forward service/vault 8200
Then go to http://localhost:8200 and enter the root token
Head to policies and create the following policies:

--- Name: royalafg-bank

    path "rabbitmq/*" {
    	capabilities=["read"]
    }
    path "kv-v2/data/session-secret" {
	    capabilities=["read"]
    }

--- royalafg-auth

    path "rabbitmq/*" {
     	capabilities=["read"] 
    }
    path "kv-v2/data/session-secret" {
	    capabilities=["read"]
    }

--- royalafg-user

    path "database/*" {
	    capabilities=["read"]
    }
    path "kv-v2/data/session-secret" {
	    capabilities=["read"]
    }
    path "kv-v2/data/royalafg-user-mongo" {
	    capabilities=["read"]
    }

--- royalafg-poker

    path "rabbitmq/*" {
     	capabilities=["read"] 
    }
    path "kv-v2/data/session-secret" {
	    capabilities=["read"]
    }
    path "kv-v2/data/poker" {
	    capabilities=["read"]
    }

Then go to Auth => Kubernetes => edit roles => Roles and add the following roles:
--- auth-role

    Bound service account names: royalafg-auth
    Bound service account namespaces: default
    Generated Token's Policies: royalafg-auth

--- user-role

    Bound service account names: royalafg-user
    Bound service account namespaces: default
    Generated Token's Policies: royalafg-user

--- bank-role

    Bound service account names: royalafg-bank
    Bound service account namespaces: default
    Generated Token's Policies: royalafg-bank

--- poker-role

    Bound service account names: royalafg-poker, royalafg-poker-matchmaker
    Bound service account namespaces: default
    Generated Token's Policies: royalafg-poker

Then create a kv secret
go to overview => create secret => kv (version 2)
Path: kv-v2
Add following secrets:
--- poker

    -> signing_key: (e.g. matchmakingkey)

--- royalafg-user-mongo

    -> url: mongodb://vault:vaultmongodb@royalafg-user-mongodb-svc.default.svc.cluster.local:27017/?authSource=admin&replicaSet=royalafg-user-mongodb&tls=false

--- session-secret

    -> secret: (e.g. testsecret)

You are now ready to deploy the royalafg services

    kubectl apply -f ./royalafg/royalafg-web.yaml
    kubectl apply -f ./royalafg/royalafg-bank.yaml
    kubectl apply -f ./royalafg/royalafg-user.yaml
    kubectl apply -f ./royalafg/royalafg-auth.yaml
    kubectl apply -f ./royalafg/royalafg-poker-matchmaker.yaml
    kubectl apply -f ./games/poker/gameServer.yaml

You can open the kubernetes dashboard to see these pods deployed

    minikube dashboard

To access the service get the port of the traefik service

    minikube service traefik --url 
    => 1st is the ingress access
    => 2nd is the traefik dashboard
    => 3rd is the secure ingress access (not used currently)

The royalafg services should correctly deploy (see minikube dashboard)

