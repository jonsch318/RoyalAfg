resource "kubernetes_manifest" "royalafg_auth_ingress_route" {
  depends_on = [helm_release.traefik]
  manifest = {
    "apiVersion" = "traefik.containo.us/v1alpha1"
    "kind" = "IngressRoute"
    "metadata" = {
      "name" = "royalafg-auth"
      "namespace" = "${kubernetes_namespace.royalafg.metadata[0].name}"
    }
    "spec" = {
      "entryPoints" = [
        "web",
      ]
      "routes" = [
        {
          "kind" = "Rule"
          "match" = "PathPrefix(`/api/auth`)"
          "services" = [
            {
              "name" = "royalafg-auth"
              "port" = 8080
            },
          ]
        },
      ]
    }
  }
}

resource "kubernetes_service" "royalafg_auth_service" {
  metadata {
    name = "royalafg-auth"
    labels = {
      app = "royalafg-auth"
    }
  }
  spec {
    selector = {
      app = "royalafg-auth"
    }
    port {
      name = "web"
      port = 8080
      target_port = 8080
    }
  }
}

resource "kubernetes_service_account" "royalafg_auth_service_account" {
  metadata {
    name = "royalafg-auth"
  }
}

resource "kubernetes_config_map" "royalafg_auth_config_map" {
  metadata {
    name = "royalafg-auth"
    labels = {
      app = "royalafg-auth"
    }
  }
  data = {
    "config.yaml" = <<-EOF
      jwt_signing_key: testsecret
      userservice_url: localhost:5200
      cors_enabled: false
      rabbitmq_url: royalafg-rabbitmq.default.svc.cluster.local:5672
    EOF
  }
}

resource "kubernetes_deployment" "royalafg_auth" {
    metadata {
    name = "royalafg-auth"
    labels = {
      app = "royalafg-auth"
    }
  }
  
  spec {
    replicas = 1
    selector {
      match_labels = {
        app = "royalafg-auth"
      }
    }
    template {
      metadata {
        labels = {
          app = "royalafg-auth"
          service = "auth"
        }
        name = "royalafg-auth"
        annotations = {
          "consul.hashicorp.com/connect-inject" =  "true"
          "consul.hashicorp.com/connect-service-upstreams" = "royalafg-user:5200"

          "consul.hashicorp.com/sidecar-proxy-cpu-request" = "10m"
          "consul.hashicorp.com/sidecar-proxy-cpu-limit" =  "25m"
          "consul.hashicorp.com/sidecar-proxy-memory-request" =  "20Mi"
          "consul.hashicorp.com/sidecar-proxy-memory-limit" = "50Mi"
          
          "vault.hashicorp.com/agent-limits-cpu" = "150m"
          "vault.hashicorp.com/agent-limits-mem" = "50Mi"
          "vault.hashicorp.com/agent-requests-cpu" = "20m"
          "vault.hashicorp.com/agent-requests-mem" = "25Mi"

          "vault.hashicorp.com/agent-inject" =  "true"
          "vault.hashicorp.com/agent-inject-secret-rabbitmq.yaml" = "rabbitmq/creds/royalafg"
          "vault.hashicorp.com/agent-inject-template-rabbitmq.yaml" = <<-EOF
            {{- with secret "rabbitmq/creds/royalafg" -}} 
            rabbitmq_username: {{.Data.username}}
            rabbitmq_password: {{.Data.password}}
            {{- end }}
            EOF
          "vault.hashicorp.com/agent-inject-secret-session.yaml" = "kv-v2/session-secret"
          "vault.hashicorp.com/agent-inject-template-session.yaml" = <<-EOF
          {{- with secret "kv-v2/session-secret" -}}
          jwt_signing_key: {{.Data.data.secret}}
          {{- end }}
          EOF
          "vault.hashicorp.com/role" = "auth-role"
        }
      }
      spec {
        service_account_name = "${kubernetes_service_account.royalafg_auth_service_account.metadata[0].name}"
        container {
          name = "royalafg-auth"
          image = "docker.io/johnnys318/royalafg_auth:latest"
          env {
            name = "GRPC_GO_LOG_SEVERITY_LEVEL"
            value = "info"
          }
          resources {
            requests = {
              memory = "10Mi"
              cpu = "10m"
            }
            limits = {
              memory = "50Mi"
              cpu = "25m"
            }
          }
          port {
            protocol = "TCP"
            container_port = 8080
          }
          volume_mount {
            name = "royalafg-auth-config"
            mount_path = "/etc/royalafg-auth"
          }
        }
        volume {
          name = "royalafg-auth-config"
          config_map {
            name = "${kubernetes_config_map.royalafg_auth_config_map.metadata[0].name}"
          }
        }
      }
    }
  }
}

resource "kubernetes_manifest" "royalafg_bank_ingress_route" {
  depends_on = [helm_release.traefik]
  manifest = {
    "apiVersion" = "traefik.containo.us/v1alpha1"
    "kind" = "IngressRoute"
    "metadata" = {
      "name" = "royalafg-bank"
      "namespace" = "${kubernetes_namespace.royalafg.metadata[0].name}"
    }
    "spec" = {
      "entryPoints" = [
        "web",
        "websecure"
      ]
      "routes" = [
        {
          "kind" = "Rule"
          "match" = "PathPrefix(`/api/bank`)"
          "services" = [
            {
              "name" = "royalafg-bank"
              "port" = 8080
            },
          ]
        },
      ]
      "tls" = {
        "certResolver" = "myresolver" 
      }
    }
  }
}

resource "kubernetes_service" "royalafg_bank_service" {
  metadata {
    name = "royalafg-bank"
    labels = {
      app = "royalafg-bank"
    }
  }
  spec {
    selector = {
      app = "royalafg-bank"
    }
    port {
      name = "web"
      port = 8080
      target_port = 8080
    }
  }
}

resource "kubernetes_service_account" "royalafg_bank_service_account" {
  metadata {
    name = "royalafg-bank"
  }
}

resource "kubernetes_config_map" "royalafg_bank_config_map" {
  metadata {
    name = "royalafg-bank"
    labels = {
      app = "royalafg-bank"
    }
  }
  data = {
    "config.yaml" = <<-EOF
      eventstore_url: http://localhost:2113?tls=false
      rabbitmq_url: royalafg-rabbitmq.default.svc.cluster.local:5672
      Prod: true
    EOF
  }
}

resource "kubernetes_deployment" "royalafg_bank" {
    metadata {
    name = "royalafg-bank"
    labels = {
      app = "royalafg-bank"
    }
  }
  
  spec {
    replicas = 1
    selector {
      match_labels = {
        app = "royalafg-bank"
      }
    }
    template {
      metadata {
        labels = {
          app = "royalafg-bank"
          service = "bank"
        }
        name = "royalafg-bank"
        annotations = {
          "consul.hashicorp.com/connect-inject" =  "true"
          "consul.hashicorp.com/connect-service-upstreams" = "eventstore:2113"

          "consul.hashicorp.com/sidecar-proxy-cpu-request" = "10m"
          "consul.hashicorp.com/sidecar-proxy-cpu-limit" =  "25m"
          "consul.hashicorp.com/sidecar-proxy-memory-request" =  "20Mi"
          "consul.hashicorp.com/sidecar-proxy-memory-limit" = "50Mi"
          
          "vault.hashicorp.com/agent-limits-cpu" = "150m"
          "vault.hashicorp.com/agent-limits-mem" = "50Mi"
          "vault.hashicorp.com/agent-requests-cpu" = "20m"
          "vault.hashicorp.com/agent-requests-mem" = "25Mi"

          "vault.hashicorp.com/agent-inject" =  "true"
          "vault.hashicorp.com/agent-inject-secret-rabbitmq.yaml" = "rabbitmq/creds/royalafg"
          "vault.hashicorp.com/agent-inject-template-rabbitmq.yaml" = <<-EOF
            {{- with secret "rabbitmq/creds/royalafg" -}} 
            rabbitmq_username: {{.Data.username}}
            rabbitmq_password: {{.Data.password}}
            {{- end }}
            EOF
          "vault.hashicorp.com/agent-inject-secret-session.yaml" = "kv-v2/session-secret"
          "vault.hashicorp.com/agent-inject-template-session.yaml" = <<-EOF
          {{- with secret "kv-v2/session-secret" -}}
          jwt_signing_key: {{.Data.data.secret}}
          {{- end }}
          EOF
          "vault.hashicorp.com/role" = "bank-role"
        }
      }
      spec {
        service_account_name = "${kubernetes_service_account.royalafg_bank_service_account.metadata[0].name}"
        container {
          name = "royalafg-bank"
          image = "docker.io/johnnys318/royalafg_bank:latest"
          env {
            name = "GRPC_GO_LOG_SEVERITY_LEVEL"
            value = "info"
          }
          resources {
            requests = {
              memory = "10Mi"
              cpu = "10m"
            }
            limits = {
              memory = "50Mi"
              cpu = "25m"
            }
          }
          port {
            protocol = "TCP"
            container_port = 8080
          }
          volume_mount {
            name = "royalafg-bank-config"
            mount_path = "/etc/royalafg-bank"
          }
        }
        volume {
          name = "royalafg-bank-config"
          config_map {
            name = "${kubernetes_config_map.royalafg_bank_config_map.metadata[0].name}"
          }
        }
      }
    }
  }
}
