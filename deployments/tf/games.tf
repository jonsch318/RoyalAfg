resource "kubernetes_config_map" "royalafg_poker_config_map" {
  metadata {
    name = "royalafg-poker"
  }
  data = {
    "config.yaml" = <<EOF
      http_port: 7654
      rabbitmq_url: royalafg-rabbitmq.default.svc.cluster.local:5672
      matchmaker_signing_key: matchmakingkey
    EOF
  }
}

resource "kubernetes_service_account" "royalafg_poker_service_account" {
  metadata {
    name = "royalafg-poker"
  }  
}

resource "kubernetes_role_binding" "royalafg_poker_role_binding" {
  metadata {
    name = "royalafg-poker"
  }
  role_ref {
    api_group = "rbac.authorization.k8s.io"
    kind = "ClusterRole"
    name = "agones-sdk"
  }
  subject {
    kind = "ServiceAccount"
    name = "${kubernetes_service_account.royalafg_poker_service_account.metadata[0].name}"
  }
}

resource "kubernetes_manifest" "royalafg_poker_autoscaler" {
  depends_on = [helm_release.agones]
  manifest = {
    "apiVersion" = "autoscaling.agones.dev/v1"
    "kind" = "FleetAutoscaler"
    "metadata" = {
      "name" = "royalafg-poker-autoscaler"
      "namespace" = "${kubernetes_namespace.royalafg.metadata[0].name}"
    }
    "spec" = {
      "fleetName" = "royalafg-poker"
      "policy" = {
        "buffer" = {
          "bufferSize" = 1
          "maxReplicas" = 5
          "minReplicas" = 1
        }
        "type" = "Buffer"
      }
    }
  }
}


resource "kubernetes_manifest" "royalafg_poker" {
  depends_on = [helm_release.agones]
  manifest = {
    "apiVersion" = "agones.dev/v1"
    "kind" = "Fleet"
    "metadata" = {
      "labels" = {
        "game" = "poker"
        "name" = "royalafg-poker"
      }
      "name" = "royalafg-poker"
      "namespace" = "${kubernetes_namespace.royalafg.metadata[0].name}"
    }
    "spec" = {
      "replicas" = 1
      "template" = {
        "metadata" = {
          "labels" = {
            "game" = "poker"
            "name" = "royalafg-poker"
          }
        }
        "spec" = {
          "health" = {
            "failureThreshold" = 10
            "initialDelaySeconds" = 20
            "periodSeconds" = 60
          }
          "ports" = [
            {
              "containerPort" = 7654
              "name" = "default"
              "portPolicy" = "Dynamic"
              "protocol" = "TCP"
            },
          ]
          "template" = {
            "metadata" = {
              "annotations" = {
                "consul.hashicorp.com/connect-inject" = "false"
                "vault.hashicorp.com/agent-inject" = "true"
                "vault.hashicorp.com/agent-inject-secret-poker.yaml" = "kv-v2/poker"
                "vault.hashicorp.com/agent-inject-secret-rabbitmq.yaml" = "rabbitmq/creds/royalafg"
                "vault.hashicorp.com/agent-inject-template-poker.yaml" = <<-EOT
                {{- with secret "kv-v2/poker" -}}
                matchmaker_signing_key: {{.Data.data.signing_key}}
                {{- end}}
                
                EOT
                "vault.hashicorp.com/agent-inject-template-rabbitmq.yaml" = <<-EOT
                {{- with secret "rabbitmq/creds/royalafg" -}}
                rabbitmq_username: {{.Data.username}}
                rabbitmq_password: {{.Data.password}}
                {{- end}}
                
                EOT
                "vault.hashicorp.com/agent-limits-cpu" = "200m"
                "vault.hashicorp.com/agent-limits-mem" = "100Mi"
                "vault.hashicorp.com/agent-requests-cpu" = "20m"
                "vault.hashicorp.com/agent-requests-mem" = "25Mi"
                "vault.hashicorp.com/role" = "poker-role"
              }
              "labels" = {
                "app" = "royalafg-poker-gameserver"
              }
            }
            "spec" = {
              "containers" = [
                {
                  "image" = "docker.io/johnnys318/royalafg_poker:latest"
                  "name" = "royalafg-poker"
                  "resources" = {
                    "limits" = {
                      "cpu" = "25m"
                      "memory" = "50Mi"
                    }
                    "requests" = {
                      "cpu" = "10m"
                      "memory" = "10Mi"
                    }
                  }
                  "volumeMounts" = [
                    {
                      "mountPath" = "/etc/royalafg-poker"
                      "name" = "royalafg-poker-config"
                    },
                  ]
                },
              ]
              "serviceAccountName" = "${kubernetes_service_account.royalafg_poker_service_account.metadata[0].name}"
              "volumes" = [
                {
                  "configMap" = {
                    "name" = "${kubernetes_config_map.royalafg_poker_config_map.metadata[0].name}"
                  }
                  "name" = "royalafg-poker-config"
                },
              ]
            }
          }
        }
      }
    }
  }
}
