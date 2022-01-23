resource "kubernetes_manifest" "rabbitmqcluster_royalafg_rabbitmq" {
  manifest = {
    "apiVersion" = "rabbitmq.com/v1beta1"
    "kind" = "RabbitmqCluster"
    "metadata" = {
      "annotations" = {
        "consul.hashicorp.com/connect-inject" = "true"
      }
      "name" = "royalafg-rabbitmq"
    }
    "spec" = {
      "replicas" = 1
      "resources" = {
        "limits" = {
          "cpu" = "800m"
          "memory" = "1Gi"
        }
        "requests" = {
          "cpu" = "500m"
          "memory" = "1Gi"
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
