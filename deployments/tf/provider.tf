provider "kubernetes" {
  config_path = "~/.kube/config"
  experiments {
    manifest_resource = true
  }
}

provider "helm" {
  kubernetes {
    config_path = "~/.kube/config"
  }  
}
