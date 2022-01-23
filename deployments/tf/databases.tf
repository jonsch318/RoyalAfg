resource "kubernetes_manifest" "rabbitmqcluster_royalafg_rabbitmq" {
  depends_on = [
    helm_release.rabbitmq_operator
  ]
  manifest = {
    "apiVersion" = "rabbitmq.com/v1beta1"
    "kind" = "RabbitmqCluster"
    "metadata" = {
      "annotations" = {
        "consul.hashicorp.com/connect-inject" = "true"
      }
      "name" = "royalafg-rabbitmq"
      "namespace" = "${kubernetes_namespace.royalafg.metadata[0].name}"
    }
    "spec" = {
      "replicas" = 1
      "resources" = {
        "limits" = {
          "cpu" = "800m"
          "memory" = "1Gi"
        }
        "requests" = {
          "cpu" = "150m"
          "memory" = "250Mi"
        }
      }
    }
  }
}

resource "kubernetes_manifest" "podmonitor_rabbitmq" {
  manifest = {
    "apiVersion" = "monitoring.coreos.com/v1"
    "kind" = "PodMonitor"
    "metadata" = {
      "labels" = {
        "app" = "rabbitmq"
        "release" = "kube-prom"
      }
      "name" = "rabbitmq"
      "namespace" = "${kubernetes_namespace.royalafg.metadata[0].name}"
    }
    "spec" = {
      "namespaceSelector" = {
        "any" = true
      }
      "podMetricsEndpoints" = [
        {
          "interval" = "15s"
          "port" = "prometheus"
        },
      ]
      "selector" = {
        "matchLabels" = {
          "app.kubernetes.io/component" = "rabbitmq"
        }
      }
    }
  }
}

resource "kubernetes_secret" "royalafg_user_mongodb_password" {
  metadata {
    name = "royalafg-user-mongodb-password"
  }
  type = "Opaque"
  data = {
    "password" = "mongodbtestpassword"
  }
}

resource "kubernetes_secret" "royalafg_user_mongodb_vault_password" {
  metadata {
    name = "royalafg-user-mongodb-vault-password"
  }
  type = "Opaque"
  data = {
    "password" = "vaultmongodb"
  }
}

resource "kubernetes_manifest" "royalafg_user_mongodb" {
  depends_on = [helm_release.mongodb_operator]
  manifest = {
    "apiVersion" = "mongodbcommunity.mongodb.com/v1"
    "kind" = "MongoDBCommunity"
    "metadata" = {
      "name" = "royalafg-user-mongodb"
      "namespace" = "${kubernetes_namespace.royalafg.metadata[0].name}"
    }
    "spec" = {
      "additionalMongodConfig" = {
        "storage.wiredTiger.engineConfig.journalCompressor" = "zlib"
      }
      "members" = 1
      "security" = {
        "authentication" = {
          "modes" = [
            "SCRAM",
          ]
        }
      }
      "type" = "ReplicaSet"
      "users" = [
        {
          "db" = "admin"
          "name" = "admin"
          "passwordSecretRef" = {
            "name" = "${kubernetes_secret.royalafg_user_mongodb_password.metadata[0].name}"
          }
          "roles" = [
            {
              "db" = "admin"
              "name" = "clusterAdmin"
            },
            {
              "db" = "admin"
              "name" = "userAdminAnyDatabase"
            },
            {
              "db" = "RoyalafgUser"
              "name" = "readWrite"
            },
          ]
          "scramCredentialsSecretName" = "my-scram"
        },
        {
          "db" = "admin"
          "name" = "vault"
          "passwordSecretRef" = {
            "name" = "${kubernetes_secret.royalafg_user_mongodb_vault_password.metadata[0].name}"
          }
          "roles" = [
            {
              "db" = "admin"
              "name" = "clusterAdmin"
            },
            {
              "db" = "admin"
              "name" = "userAdminAnyDatabase"
            },
            {
              "db" = "admin"
              "name" = "readWrite"
            },
            {
              "db" = "RoyalafgUser"
              "name" = "readWrite"
            },
          ]
          "scramCredentialsSecretName" = "vault-scram"
        },
      ]
      "version" = "4.4.0"
    }
  }
}

resource "kubernetes_manifest" "elasticsearch_search" {
  depends_on = [helm_release.elastic_operator]
  manifest = {
    "apiVersion" = "elasticsearch.k8s.elastic.co/v1"
    "kind" = "Elasticsearch"
    "metadata" = {
      "name" = "search"
      "namespace" = "${kubernetes_namespace.royalafg.metadata[0].name}"
    }
    "spec" = {
      "nodeSets" = [
        {
          "config" = {
            "node.store.allow_mmap" = false
          }
          "count" = 1
          "name" = "default"
          "podTemplate" = {
            "metadata" = {
              "annotations" = {
                "consul.hashicorp.com/consul-inject" = "true"
              }
            }
          }
        },
      ]
      "version" = "7.16.2"
    }
  }
}

resource "kubernetes_secret_v1" "kibana_search_user_secret" {
  metadata {
    name = "search-kibana-user"
  }
  data = {
    "users" = <<EOF
      admin:admin
    EOF
    "users_roles" = <<EOF
      admin:admin
    EOF
  }
}

resource "kubernetes_manifest" "kibana_search" {
  depends_on=[helm_release.elastic_operator]
  manifest = {
    "apiVersion" = "kibana.k8s.elastic.co/v1"
    "kind" = "Kibana"
    "metadata" = {
      "name" = "search"
      "namespace" = "${kubernetes_namespace.royalafg.metadata[0].name}"
    }
    "spec" = {
      "count" = 1
      "elasticsearchRef" = {
        "name" = "${kubernetes_manifest.elasticsearch_search.manifest.metadata[0].name}"
      }
      "version" = "7.16.2"
    }
  }
}