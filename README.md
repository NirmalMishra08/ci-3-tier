# GitOps-Ready Go Backend - Complete Cloud-Native Application

A production-ready Go backend application demonstrating modern cloud-native development practices with GitOps workflow, Kubernetes orchestration, comprehensive monitoring, and CI/CD automation.

---

## 📋 Table of Contents

- [Project Overview](#project-overview)
- [Technology Stack](#technology-stack)
- [Architecture](#architecture)
- [Features](#features)
- [Quick Start](#quick-start)
- [Project Structure](#project-structure)
- [Current Implementation](#current-implementation)
- [Project Progression](#project-progression)
- [Future Roadmap](#future-roadmap)
- [Deployment Guide](#deployment-guide)
- [Monitoring & Observability](#monitoring--observability)
- [Contributing](#contributing)

---

## 📌 Project Overview

This project is a **cloud-native Go backend application** designed to demonstrate and implement **GitOps principles** from ground up. It showcases a complete production-ready setup with:

- **Containerized microservices** using Docker
- **Kubernetes deployment** with Helm charts
- **Automated CI/CD pipelines** via GitHub Actions
- **Built-in observability** with Prometheus & Grafana
- **Infrastructure as Code (IaC)** best practices
- **Security-first approach** with secret management and container scanning
- **Prepared for ArgoCD** GitOps controller integration

### Purpose

This application serves as a **learning platform** and **reference architecture** for:
- Building scalable cloud-native applications
- Implementing GitOps workflows
- Managing containerized deployments across environments
- Implementing observability in production systems
- CI/CD automation with GitHub Actions

---

## 🛠 Technology Stack

### Backend
- **Language**: Go 1.24.3
- **Framework**: Chi Router (lightweight HTTP router)
- **Monitoring**: Prometheus client library
- **Configuration**: godotenv (environment variable management)

### Containerization & Orchestration
- **Docker**: Multi-stage builds for optimized images
- **Kubernetes**: Container orchestration platform
- **Helm**: Package manager for Kubernetes (v0.1.0)

### Observability
- **Prometheus v2.55.0**: Metrics collection and storage
- **Grafana 11.3.0**: Metrics visualization and dashboards
- **Custom Metrics**: Application-level request tracking

### CI/CD & DevOps
- **GitHub Actions**: Automated build and push pipeline
- **Docker Hub**: Container registry (nirmal08/ci-3-tier)
- **Gitleaks**: Secret scanning and prevention
- **Git**: Version control with GitOps workflow

### Development
- **docker-compose**: Local development environment orchestration
- **kubectl**: Kubernetes command-line client
- **Helm**: Kubernetes package management

---

## 🏗 Architecture

```
┌─────────────────────────────────────────────────────────┐
│                    Users/Clients                        │
└────────────────────────┬────────────────────────────────┘
                         │
         ┌───────────────┼───────────────┐
         │               │               │
    ┌────▼─────┐   ┌─────▼────┐   ┌────▼──────┐
    │  / (Root)│   │ /health  │   │ /metrics  │
    │   Page   │   │  Status  │   │ Prometheus│
    └──────────┘   └──────────┘   └───────────┘
         │
    ┌────▼────────────────────────┐
    │   Go HTTP Server (Port 8000)│
    │  ┌──────────────────────┐   │
    │  │ Request Middleware   │   │
    │  │ (Prometheus Metrics) │   │
    │  └──────────────────────┘   │
    └────────────────────────┬────┘
         │
    ┌────▼────────────────────────┐
    │  Config Management           │
    │  • PORT (env)                │
    │  • SECRET (env)              │
    └──────────────────────────────┘
         │
    ┌────▼──────────────────────────────────┐
    │      Deployment Environments           │
    │  ┌──────────────┐   ┌──────────────┐  │
    │  │Docker-Compose│   │  Kubernetes  │  │
    │  │  (Dev/Test)  │   │ (Production) │  │
    │  └──────────────┘   └──────────────┘  │
    └─────────────────────────────────────────┘
         │
    ┌────▼──────────────────────────────────┐
    │     Monitoring Stack                   │
    │  ┌──────────┐   ┌──────────────────┐  │
    │  │Prometheus│   │    Grafana       │  │
    │  │ :9090    │   │    :3000         │  │
    │  └──────────┘   └──────────────────┘  │
    └────────────────────────────────────────┘
```

### Deployment Flow

```
Code Push → GitHub Actions Workflow
    ↓
Security Scan (Gitleaks)
    ↓
Docker Build & Tag
    ↓
Push to Docker Hub
    ↓
Kubernetes Manifest Deployment (via kubectl/ArgoCD)
    ↓
Service LoadBalancer (Port 80 → 8000)
    ↓
Prometheus Metrics Collection
    ↓
Grafana Visualization
```

---

## ✨ Features

### 1. **HTTP API Endpoints**
- `GET /` - Serves static HTML landing page with greeting message
- `GET /health` - Returns application health status with PORT and SECRET values
- `GET /metrics` - Prometheus-compatible metrics endpoint

### 2. **Prometheus Metrics**
**Request Tracking:**
- `api_http_request_total{path, status}` - Total successful requests (status < 400)
- `api_http_request_error_total{path, status}` - Total error requests (status >= 400)

All metrics include path and HTTP status code labels for detailed analysis.

### 3. **Configuration Management**
- **Environment Variables**: `PORT`, `SECRET`
- **Default Values**: PORT=8000
- **Multi-file Support**: Loads from `.env` in current, parent, and root directories
- **Runtime Loading**: Configuration loaded on application startup

### 4. **Container Security**
- **Multi-stage Dockerfile**: Reduces final image size by 90%+
- **Alpine Base**: Security-focused minimal OS
- **CA Certificates**: Enables HTTPS certificate validation
- **Non-root Execution**: Container runs with minimal privileges

### 5. **Kubernetes Deployment**
- **ConfigMap Injection**: Environment-based configuration
- **Secret Management**: Base64-encoded sensitive data
- **Resource Limits**: CPU and memory constraints
- **Health Checks**: Liveness and readiness probes
- **LoadBalancer Service**: External access to the application

### 6. **Local Development**
- **docker-compose Stack**: API + Prometheus + Grafana
- **Hot Reload**: File watch for automatic container rebuild
- **Health Checks**: Automatic container health monitoring
- **Isolated Networks**: Separate networking for security

### 7. **Monitoring & Observability**
- **Prometheus**: Time-series metrics database
- **Grafana**: Pre-configured dashboards and alerting
- **Custom Metrics**: Application-level instrumentation
- **Visualization**: HTTP request patterns and error rates

### 8. **CI/CD Pipeline**
- **Automated Builds**: On push to `dev` branch
- **Secret Scanning**: Gitleaks integration to prevent credential leaks
- **Docker Registry**: Automated image push with SHA tagging
- **Versioning**: Multiple tags (`:latest`, `:{SHA}`)

---

## 🚀 Quick Start

### Prerequisites
```bash
- Docker & Docker Compose
- Go 1.24+
- kubectl
- Helm 3+
- Git
```

### 1. Local Development (docker-compose)

```bash
cd backend
docker-compose up -d
```

**Access Points:**
- API: http://localhost:8000
- Prometheus: http://localhost:9090
- Grafana: http://localhost:3000 (admin/password)

**Health Check:**
```bash
curl http://localhost:8000/health
# Response: {"status":200,"PORT":"8000","secret":""}
```

### 2. Build & Run Locally (Native Go)

```bash
cd backend
go mod download
go run main.go

# In another terminal
curl http://localhost:8000/
curl http://localhost:8000/health
curl http://localhost:8000/metrics
```

### 3. Docker Build

```bash
cd backend
docker build -t my-go-app:latest .
docker run -p 8000:8000 -e PORT=8000 my-go-app:latest
```

### 4. Kubernetes Deployment

```bash
cd backend/k8s/gitops
helm install my-app . -f values.yaml
kubectl get pods
kubectl get svc
```

---

## 📁 Project Structure

```
gitops/
├── backend/
│   ├── main.go                    # HTTP server with Chi router + Prometheus middleware
│   ├── go.mod                     # Go module dependencies
│   ├── config/
│   │   └── config.go              # Configuration loader (env variables)
│   ├── Dockerfile                 # Multi-stage Docker build
│   ├── docker-compose.yml         # Local dev stack (API + Prometheus + Grafana)
│   ├── static/
│   │   ├── index.html             # Landing page
│   │   └── index.css              # Styling
│   ├── Docker/
│   │   ├── prometheus.yml         # Prometheus configuration
│   │   └── grafana.yml            # Grafana datasource config
│   ├── k8s/
│   │   └── gitops/                # Helm chart directory
│   │       ├── Chart.yaml         # Helm chart metadata (v0.1.0)
│   │       ├── values.yaml        # Default values (replicas, image, ports)
│   │       └── templates/
│   │           ├── deployment.yml # Kubernetes deployment manifest
│   │           ├── service.yml    # LoadBalancer service
│   │           ├── configmaps.yml # ConfigMap for environment config
│   │           └── secret.yml     # Secret for sensitive data
│   ├── .github/
│   │   └── workflows/
│   │       └── ci.yml             # GitHub Actions CI/CD pipeline
│   └── README.md                  # This file
└── .gitignore                      # Git ignore patterns
```

---

## 🔧 Current Implementation

### What's Been Built

#### 1. **Go Backend Application**
- Lightweight HTTP server using Chi router
- Three endpoints: `/` (static), `/health` (status), `/metrics` (Prometheus)
- Middleware for automatic Prometheus metrics collection
- Configuration management with environment variables
- Graceful error handling

#### 2. **Containerization**
- Multi-stage Dockerfile for optimized image size
- Alpine Linux for security and minimal footprint
- Proper signal handling and container best practices
- Health check integration in docker-compose

#### 3. **Local Development Setup**
- docker-compose with API, Prometheus, and Grafana
- Pre-configured monitoring stack
- Volume mappings for configuration files
- Health checks and dependency ordering

#### 4. **Kubernetes Deployment**
- Complete Helm chart for production deployment
- Separate manifests: Deployment, Service, ConfigMap, Secret
- Resource limits and requests defined
- LoadBalancer service for external access
- Proper labeling for service discovery

#### 5. **Monitoring Infrastructure**
- Prometheus scrape configuration
- Grafana with Prometheus datasource
- Custom application metrics
- Pre-built dashboards configuration

#### 6. **CI/CD Pipeline (GitHub Actions)**
- Automated builds on `dev` branch pushes
- Security scanning with Gitleaks
- Docker image build with SHA tagging
- Automated push to Docker Hub
- Versioning strategy (`:latest` + `:{SHA}`)

#### 7. **Security**
- Base64 secret encoding in Kubernetes
- Environment-based configuration
- Container security scanning
- Minimal Alpine base images
- Secret scanning in CI/CD pipeline

---

## 📈 Project Progression

### Phase 1: Foundation ✅ (COMPLETED)
- [x] Set up Go project structure with basic HTTP server
- [x] Create Chi router with `/`, `/health`, `/metrics` endpoints
- [x] Implement Prometheus metrics collection
- [x] Create static HTML/CSS landing page
- [x] Environment configuration management

### Phase 2: Containerization ✅ (COMPLETED)
- [x] Write multi-stage Dockerfile for optimized builds
- [x] Create docker-compose for local development
- [x] Add Prometheus configuration and data collection
- [x] Add Grafana with datasource configuration
- [x] Set up health checks in docker-compose

### Phase 3: Kubernetes & Orchestration ✅ (COMPLETED)
- [x] Create Kubernetes manifests (Deployment, Service, ConfigMap, Secret)
- [x] Set up Helm chart structure (v0.1.0)
- [x] Configure resource limits and requests
- [x] Implement LoadBalancer service for external access
- [x] Define proper labels and selectors

### Phase 4: CI/CD Pipeline ✅ (COMPLETED)
- [x] Set up GitHub Actions workflow on `dev` branch
- [x] Implement Gitleaks secret scanning
- [x] Configure Docker build and push to registry
- [x] Set up image tagging strategy (`:latest` + `:{SHA}`)
- [x] Automated deployment workflow

### Phase 5: Production Readiness ✅ (COMPLETED)
- [x] Implement comprehensive error handling
- [x] Add health check endpoints
- [x] Configure security best practices
- [x] Set up proper logging
- [x] Create comprehensive documentation

---

## 🗺️ Future Roadmap

### Phase 6: ArgoCD Integration & GitOps (NEXT)

#### 6.1 ArgoCD Installation & Setup
- [ ] Deploy ArgoCD controller in Kubernetes cluster (`argocd` namespace)
- [ ] Configure ArgoCD server and UI access
- [ ] Set up authentication (GitHub OAuth, RBAC)
- [ ] Install ArgoCD CLI tools

#### 6.2 GitOps Repository Structure
```
gitops-repo/
├── argocd/
│   ├── application.yaml        # ArgoCD Application manifest
│   └── appproject.yaml         # Project RBAC and restrictions
├── environments/
│   ├── dev/
│   │   ├── values.yaml
│   │   └── secrets.yaml
│   ├── staging/
│   │   ├── values.yaml
│   │   └── secrets.yaml
│   └── prod/
│       ├── values.yaml
│       └── secrets.yaml
└── helm/
    └── backend-chart/
        ├── Chart.yaml
        ├── values.yaml
        └── templates/
```

#### 6.3 ArgoCD Application Manifests
- [ ] Create ArgoCD Application CRD for each environment
- [ ] Configure auto-sync policies
- [ ] Set up health checks and status monitoring
- [ ] Implement notifications (Slack, email)
- [ ] Create separate AppProjects for environment isolation

#### 6.4 Environment Promotion Pipeline
- [ ] **Development**: Auto-sync on every git commit
- [ ] **Staging**: Manual approval for promotion from dev
- [ ] **Production**: Scheduled syncs with manual gates
- [ ] Rollback strategies and disaster recovery
- [ ] Canary deployments and progressive rollouts

#### 6.5 Multi-Environment Configuration
- [ ] Kustomize overlays for environment-specific configs
- [ ] Separate values.yaml per environment
- [ ] Secret management (Sealed Secrets or External Secrets)
- [ ] ConfigMap templating and variable substitution
- [ ] Resource quotas per environment

#### 6.6 Advanced GitOps Features
- [ ] **Sync Waves**: Ordered deployment of resources
- [ ] **Post-Sync Hooks**: Smoke tests and validation
- [ ] **Notifications**: Deployment status updates
- [ ] **Metrics Integration**: ArgoCD metrics in Prometheus
- [ ] **Policy Enforcement**: OPA/Kyverno policies

#### 6.7 Secret Management
- [ ] Integrate with HashiCorp Vault
- [ ] Sealed Secrets for git-friendly encryption
- [ ] External Secrets Operator for cloud secret backends
- [ ] Rotation policies and audit logs
- [ ] Multi-cluster secret synchronization

#### 6.8 Observability Enhancement
- [ ] ArgoCD metrics in Prometheus
- [ ] Deployment timelines and SLOs
- [ ] Application health dashboards
- [ ] GitOps workflow audit trails
- [ ] Performance metrics and optimization

### Phase 7: Multi-Cluster & High Availability (FUTURE)

- [ ] Secondary cluster setup and synchronization
- [ ] Cross-cluster failover strategies
- [ ] Cluster health monitoring
- [ ] Load balancing across clusters
- [ ] Disaster recovery procedures

### Phase 8: Advanced Monitoring & Observability

- [ ] OpenTelemetry tracing integration
- [ ] Distributed tracing with Jaeger
- [ ] Log aggregation (ELK, Loki)
- [ ] Alert rules and incident response
- [ ] Custom dashboards and reporting

### Phase 9: Security Hardening

- [ ] Network policies and segmentation
- [ ] Pod Security Policies
- [ ] Image scanning (Trivy, Snyk)
- [ ] Runtime security (Falco)
- [ ] Compliance monitoring (PCI, HIPAA, SOC2)

### Phase 10: Performance Optimization

- [ ] Horizontal Pod Autoscaler (HPA)
- [ ] Vertical Pod Autoscaler (VPA)
- [ ] Cluster autoscaling
- [ ] Performance benchmarking
- [ ] Cost optimization strategies

---

## 📦 Deployment Guide

### Prerequisites
- Kubernetes cluster (v1.20+)
- kubectl configured with cluster access
- Helm 3.x installed
- Docker registry access (Docker Hub or private)

### Step 1: Build & Push Docker Image

```bash
cd backend

# Build with tag
docker build -t nirmal08/ci-3-tier:v1.0.0 .

# Push to registry
docker push nirmal08/ci-3-tier:v1.0.0
```

### Step 2: Update Helm Values

```bash
# Edit k8s/gitops/values.yaml
image:
  repository: nirmal08/ci-3-tier
  tag: v1.0.0
  pullPolicy: IfNotPresent

replicas: 3
```

### Step 3: Deploy with Helm

```bash
# Dry run first
helm install my-backend ./k8s/gitops --dry-run --debug

# Actual deployment
helm install my-backend ./k8s/gitops

# Or upgrade existing release
helm upgrade my-backend ./k8s/gitops
```

### Step 4: Verify Deployment

```bash
# Check pods
kubectl get pods -l app=backend

# Check service
kubectl get svc my-backend

# Get external IP
kubectl get svc my-backend -o wide

# Test health endpoint
curl http://<EXTERNAL-IP>/health
```

### Step 5: Monitor with Prometheus/Grafana

```bash
# Port forward to Prometheus
kubectl port-forward svc/prometheus 9090:9090

# Port forward to Grafana
kubectl port-forward svc/grafana 3000:3000

# Access: http://localhost:3000 (admin/password)
```

---

## 📊 Monitoring & Observability

### Prometheus Metrics

#### Request Metrics
```
# Total successful requests
api_http_request_total{path="/", status="200"}
api_http_request_total{path="/health", status="200"}

# Total error requests
api_http_request_error_total{path="/nonexistent", status="404"}
```

### Querying Metrics

```promql
# Request rate (requests per second)
rate(api_http_request_total[1m])

# Error rate percentage
rate(api_http_request_error_total[5m]) / rate(api_http_request_total[5m]) * 100

# Requests by endpoint
sum(rate(api_http_request_total[5m])) by (path)

# Error requests by status code
sum(api_http_request_error_total) by (status)
```

### Grafana Dashboards

Pre-configured dashboards show:
- Request volume and trends
- Error rates and percentages
- Response time distribution
- Endpoint-specific metrics
- System resource usage

### Health Checks

**Kubernetes Probes:**
```yaml
livenessProbe:
  httpGet:
    path: /health
    port: 8000
  initialDelaySeconds: 10
  periodSeconds: 10

readinessProbe:
  httpGet:
    path: /health
    port: 8000
  initialDelaySeconds: 5
  periodSeconds: 5
```

---

## 🔒 Security Considerations

### Current Security Measures
- ✅ Multi-stage Docker builds for minimal attack surface
- ✅ Alpine Linux base for fewer vulnerabilities
- ✅ Gitleaks scanning in CI/CD pipeline
- ✅ Base64 secret encoding in Kubernetes
- ✅ Environment variable isolation
- ✅ Resource limits and requests
- ✅ Health checks for availability

### Recommended Enhancements
- [ ] Implement RBAC for Kubernetes access
- [ ] Use Sealed Secrets or External Secrets for git-safe secret management
- [ ] Network policies for pod-to-pod communication
- [ ] Pod Security Policies/Standards
- [ ] Regular container image scanning
- [ ] Runtime security monitoring with Falco
- [ ] Audit logging for all API calls

---

## 🧪 Testing

### Unit Tests
```bash
cd backend
go test ./...
```

### Health Endpoint Test
```bash
curl -X GET http://localhost:8000/health
```

### Metrics Endpoint Test
```bash
curl http://localhost:8000/metrics | grep api_http_request
```

### Load Testing with Docker Compose
```bash
docker-compose up -d
# Generate traffic
ab -n 1000 -c 10 http://localhost:8000/

# View metrics
curl http://localhost:9090/graph
```

---

## 📝 Configuration

### Environment Variables

| Variable | Default | Purpose |
|----------|---------|---------|
| `PORT` | 8000 | Server listen port |
| `SECRET` | (empty) | Application secret for health endpoint |

### Loading Configuration

Configuration is loaded in the following order (last one wins):
1. `.env` in current directory
2. `../.env` (parent directory)
3. `.env` in root directory
4. System environment variables
5. Hardcoded defaults

---

## 🤝 Contributing

### Development Workflow
1. Create feature branch from `dev`
2. Make changes and test locally
3. Push to feature branch
4. Create pull request to `dev`
5. GitHub Actions CI runs automatically
6. Merge after review and checks pass

### Git Commit Guidelines
```
feat: Add new feature
fix: Fix bug
docs: Update documentation
test: Add tests
refactor: Code refactoring
chore: Maintenance tasks
```

---

## 📚 Additional Resources

### Documentation
- [Go Documentation](https://golang.org/doc/)
- [Chi Router](https://github.com/go-chi/chi)
- [Prometheus Docs](https://prometheus.io/docs/)
- [Kubernetes Docs](https://kubernetes.io/docs/)
- [Helm Charts](https://helm.sh/docs/)

### Learning Resources
- [GitOps Best Practices](https://www.gitops.tech/)
- [Cloud Native Computing Foundation](https://www.cncf.io/)
- [Kubernetes The Hard Way](https://github.com/kelseyhightower/kubernetes-the-hard-way)
- [ArgoCD Documentation](https://argo-cd.readthedocs.io/)

---

## 📄 License

This project is licensed under the MIT License.

---

## 📞 Support

For issues, questions, or suggestions:
- Create an issue in the GitHub repository
- Check existing documentation
- Review code comments and examples

---

## 🎯 Summary

This GitOps-ready Go backend demonstrates a **complete cloud-native development lifecycle** from local development through production Kubernetes deployment. The project is structured to progressively introduce GitOps concepts, with ArgoCD integration as the next major milestone, enabling fully declarative, git-driven infrastructure management and continuous deployment workflows.

**Current Status:** ✅ Production-ready backend with containerization, Kubernetes deployment, CI/CD automation, and comprehensive monitoring.

**Next Steps:** 🚀 Implement ArgoCD for true GitOps workflow, environment promotion pipelines, and multi-cluster management.
