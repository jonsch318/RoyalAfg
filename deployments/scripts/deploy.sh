minikube start --memory 4096

helm repo add hashicorp https://helm.releases.hashicorp.com
helm repo add agones https://agones.dev/chart/stable

helm repo update



helm install -f ./../consul-values.yaml hashicorp hashicorp/consul
helm install -f ./../vault-values.yaml vault hashicorp/vault

helm install my-release agones/agones --set "gameservers.namespaces={default}" --namespace agones-system