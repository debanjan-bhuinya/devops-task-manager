# Loki Logging Stack

Installed using Helm.

## Namespace
logging

## Installation

helm repo add grafana https://grafana.github.io/helm-charts
helm repo update

helm install loki grafana/loki-stack \
  --namespace logging \
  --create-namespace \
  --set grafana.enabled=false

