# DevOps Task Manager – Kubernetes Monitoring Project

This project demonstrates a **containerized Go REST API deployed on Kubernetes with a full monitoring stack** using Prometheus and Grafana.

## Tech Stack

- Go (Golang)
- Chi Router
- PostgreSQL
- Docker
- Kubernetes (Minikube)
- Prometheus
- Grafana
- Helm
- GitHub Actions

---

# Architecture

User
 ↓
Kubernetes Service
 ↓
Go API Pod
 ↓
PostgreSQL Pod

Monitoring Stack
Prometheus → Collect metrics
Grafana → Visual dashboards
Alertmanager → Alerts

---

# Features

- Containerized Go REST API
- Kubernetes deployment
- PostgreSQL database
- CI/CD pipeline with GitHub Actions
- Prometheus monitoring
- Grafana dashboards
- ServiceMonitor for automatic metric scraping

---

# API Endpoints

Health check
/api/v1/health


Users
GET /api/v1/users
POST /api/v1/users

Metrics endpoint
/metrics


---

# Kubernetes Deployment

Apply resources
kubectl apply -f k8s/

Check pods
kubectl get pods


---

# Monitoring Stack

Install monitoring stack
helm install monitoring prometheus-community/kube-prometheus-stack -n monitoring


Access Grafana
kubectl port-forward svc/monitoring-grafana -n monitoring 3000:80


---

# Screenshots

Grafana dashboards showing Kubernetes and application metrics.

---

# Author

Debanjan Bhuinya
