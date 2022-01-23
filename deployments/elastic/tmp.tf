resource "kubernetes_manifest" "kibana_search" {
  manifest = {
    "apiVersion" = "kibana.k8s.elastic.co/v1"
    "kind" = "Kibana"
    "metadata" = {
      "name" = "search"
    }
    "spec" = {
      "count" = 1
      "elasticsearchRef" = {
        "name" = "search"
      }
      "version" = "7.10.2"
    }
  }
}

resource "kubernetes_manifest" "secret_search_kibana_user" {
  manifest = {
    "apiVersion" = "v1"
    "kind" = "Secret"
    "metadata" = {
      "name" = "search-kibana-user"
    }
    "stringData" = {
      "users" = "admin:admin"
      "users_roles" = "admin:admin"
    }
  }
}
