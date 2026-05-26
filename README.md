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

##  Project Overview

This project is a **fully operational cloud-native Go backend application** demonstrating **GitOps principles** in action. It showcases a complete production-ready setup with:

- **Containerized microservices** using Docker with multi-stage builds
- **Kubernetes deployment** with Helm charts (v0.1.0)
- **Automated CI/CD pipelines** via GitHub Actions with Docker Hub integration
- **Built-in observability** with Prometheus & Grafana (monitoring stack included)
- **Infrastructure as Code (IaC)** best practices across all layers
- **Security-first approach** with Gitleaks scanning, secret management, and container hardening
- **🎯 GitOps Integration** with ArgoCD for fully declarative, git-driven deployments (ACTIVE)

### Current Status:  PRODUCTION-READY WITH ACTIVE ARGOCD

The project is now **fully integrated with ArgoCD** for continuous deployment workflows. All infrastructure is version-controlled in Git, and changes automatically synchronize to the Kubernetes cluster.

### Purpose

This application serves as a **learning platform** and **reference architecture** for:
- Building scalable cloud-native applications with production patterns
- Implementing complete GitOps workflows end-to-end
- Managing containerized deployments across multiple environments
- Implementing observability and monitoring in production systems
- Automated CI/CD pipelines with security scanning
- Declarative infrastructure management with ArgoCD

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

##  Architecture

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

##  Features

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

##  Quick Start

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

##  Project Structure

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

### What's Been Built (7 Major Components)

#### 1. **Go Backend Application** ✅
- Lightweight HTTP server using Chi router framework
- Three endpoints fully functional:
  - `GET /` - Serves static HTML landing page
  - `GET /health` - Returns JSON status with PORT and SECRET values
  - `GET /metrics` - Prometheus-compatible metrics endpoint
- Automatic Prometheus metrics collection via middleware
- Custom metrics: `api_http_request_total`, `api_http_request_error_total`
- Configuration management with environment variables (PORT, SECRET)
- Graceful error handling and logging

#### 2. **Containerization** ✅
- Multi-stage Dockerfile for 90%+ image size reduction
- Golang 1.24-alpine builder stage → Alpine runtime
- Alpine Linux base for security and minimal footprint
- Proper signal handling and container best practices
- Health check integration in docker-compose (30s interval, 5 retries)
- Binary compiled with Go 1.24.3

#### 3. **Local Development Setup** ✅
- Complete docker-compose stack with 3 services:
  - Go API on port 8000 with hot-reload watch mode
  - Prometheus v2.55.0 on port 9090
  - Grafana 11.3.0 on port 3000 (admin/password)
- Pre-configured monitoring stack
- Volume mappings for configuration files
- Custom bridge network isolation (`go-network`)
- Health checks and service dependency ordering

#### 4. **Kubernetes Deployment** ✅
- Complete Helm chart (v0.1.0, app v1.16.0)
- Separate manifests: Deployment, Service, ConfigMap, Secret
- Resource limits: 128Mi/100m requests, 256Mi/500m limits
- LoadBalancer service for external access
- Port mapping: 80 (external) → 8000 (internal)
- Proper labeling for service discovery (`app: gitops`)
- ConfigMap injection for environment variables
- Secret injection for sensitive data (base64-encoded)

#### 5. **Monitoring Infrastructure** ✅
- Prometheus scrape configuration (15s interval)
- Prometheus targets: API `/metrics` endpoint
- Grafana with Prometheus datasource (proxy mode)
- Custom application metrics collection
- Pre-configured datasource for dashboards
- Metrics visualization ready

#### 6. **CI/CD Pipeline (GitHub Actions)** ✅
- Automated builds triggered on `dev` branch pushes
- Security scanning with Gitleaks (prevents secret commits)
- Go 1.21 setup and `go mod tidy` (⚠️ see Issues below)
- Docker build with multi-layer caching
- Automated Docker image tagging:
  - `:latest` for latest build
  - `:{SHA}` for specific commit tracking
- Automated push to Docker Hub (`nirmal08/ci-3-tier`)
- **🎯 Helm values auto-update**: Workflow automatically updates `helm/gitops/values.yaml` with new image tag
- Git commit and push back to dev branch for GitOps sync

#### 7. **GitOps Integration (ArgoCD)** ✅ NEW!
- ArgoCD application manifest configured (`argocd-app.yaml`)
- Git repository source: `https://github.com/NirmalMishra08/ci-3-tier.git`
- Target branch: `dev`
- Helm chart path: `helm/gitops`
- Destination: Default Kubernetes cluster, default namespace
- **Automated Sync Enabled**: Auto-prune and self-healing
- Changes to Git automatically deploy to cluster
- Full declarative GitOps workflow operational

#### 8. **Security Implementation** ✅
- Base64 secret encoding in Kubernetes manifests
- Environment-based configuration isolation
- Container security scanning via Gitleaks in CI/CD
- Minimal Alpine base images (reduced attack surface)
- Secret scanning prevents credential leaks in Git
- Multi-stage builds eliminate build tools from runtime image

---

##  Current Issues & Precautions

### Critical Issues Found

#### 1. **Docker Image Tagging - IMPORTANT** 🏷️
**Current Behavior**: Pipeline tags images with `:latest` AND `:{SHA}`

** PROBLEM**: 
- `:latest` tag is **unreliable in production**
- `:latest` can break at any time if someone pushes a new image
- No version control, no rollback capability
- Helm values auto-update to `:latest` is **risky**

** SOLUTION - Semantic Versioning Required**:
```bash
# CORRECT APPROACH - Use semantic versioning
docker build -t nirmal08/ci-3-tier:v1.0.0 .
docker build -t nirmal08/ci-3-tier:v1.0.0-$(git rev-parse --short HEAD) .  # Include commit hash
docker push nirmal08/ci-3-tier:v1.0.0

# Only use :latest for testing, NEVER for production
docker build -t nirmal08/ci-3-tier:latest .  # Only for dev/staging
```

**Recommended GitHub Actions Updates**:
```yaml
# Calculate version from git tag or bump version
- name: Determine Image Tag
  run: |
    VERSION=$(git describe --tags --always)
    echo "IMAGE_TAG=v${VERSION}" >> $GITHUB_ENV
    
- name: Docker Build & Push
  run: |
    docker build -t nirmal08/ci-3-tier:${IMAGE_TAG} .
    docker push nirmal08/ci-3-tier:${IMAGE_TAG}
    # Only add :latest for feature branches, not for main/prod
```

**Helm values.yaml Must Use Specific Versions**:
```yaml
#  BAD (production risk)
image:
  tag: latest

#  GOOD (production safe)
image:
  tag: v1.0.0  # Always specify exact version
```

---

#### 2. **Go Version Mismatch** 
- **Issue**: GitHub Actions uses Go 1.21, but code requires Go 1.24.3
- **Solution**: Update workflow to use Go 1.24 to match `go.mod`
```yaml
- uses: actions/setup-go@v5
  with:
    go-version: '1.24.3'  # Match go.mod
```

---

#### 3. **HTML File References** 
- **Issue**: `static/index.html` references `style.css` and `app.js` (don't exist)
- **Current**: `index.css` exists but is named incorrectly in HTML
- **Fix Required**: Update HTML to reference correct filenames

---

#### 4. **Kubernetes Deployment Port Mismatch** 
- **Issue**: `deployment.yml` container exposes port 80, but app runs on 8000
- **Current**: Service correctly maps 80 → 8000, so it works
- **Recommendation**: Update deployment to expose correct port (8000) for clarity

---

#### 5. **Hardcoded Secret in Manifests** 
- **Issue**: `secret.yml` has hardcoded base64-encoded secret: `c2VjcmV0` (="secret")
- **Production Problem**: Secrets should NOT be in Git
- **Solution Options**:
  - Use Sealed Secrets for git-safe encryption
  - Use External Secrets Operator with cloud secret managers
  - Use ArgoCD + Kyverno policies to prevent secret commits

---

#### 6. **Replica Count Mismatch** 
- **Issue**: `values.yaml` specifies `replicas: 2`, but `deployment.yml` hardcodes `replicas: 1`
- **Fix**: Remove hardcoded value from deployment and use Helm templating

---

### Precautions & Best Practices

####  Docker Tagging Hierarchy (For Production)

```
Production-Ready Tagging Strategy:

nirmal08/ci-3-tier:v1.0.0          # Production release (immutable)
nirmal08/ci-3-tier:v1.0.0-beta     # Pre-release version
nirmal08/ci-3-tier:v1.0.0-alpha    # Alpha version
nirmal08/ci-3-tier:v1.0.0-rc1      # Release candidate
nirmal08/ci-3-tier:v1               # Minor version (points to v1.x.y)
nirmal08/ci-3-tier:v1.0             # Patch version (points to v1.0.y)
nirmal08/ci-3-tier:latest           # ONLY for dev/test branches
```

#### 📋 Deployment Checklist

Before pushing to production:
- [ ] Use specific version tag (e.g., `v1.0.0`), never `:latest`
- [ ] Update `helm/gitops/values.yaml` with exact version tag
- [ ] Run `helm lint` to validate chart
- [ ] Test deployment in staging environment first
- [ ] Run smoke tests after deployment
- [ ] Verify all pods are healthy with `kubectl get pods`
- [ ] Check logs: `kubectl logs -f deployment/gitops-3-tier`
- [ ] Verify metrics endpoint: `curl http://<IP>/metrics`
- [ ] Confirm Prometheus is scraping: Check Prometheus UI port 9090

####  Secret Management Precautions

- [ ] **NEVER** commit `.env` files to Git
- [ ] **NEVER** commit base64-encoded secrets to manifests (they're visible in Git history)
- [ ] Use Sealed Secrets or External Secrets for production
- [ ] Rotate secrets regularly
- [ ] Audit secret access logs
- [ ] Use RBAC for secret access control

####  ArgoCD Best Practices

- [ ] Always sync from specific Git branches (`dev`, `staging`, `prod`)
- [ ] Use separate ArgoCD Applications per environment
- [ ] Enable auto-prune to clean up deleted resources
- [ ] Set up notifications (Slack, email) for deployments
- [ ] Use sync waves for ordered deployments
- [ ] Implement health checks for applications
- [ ] Review deployment history regularly
- [ ] Test rollback procedures

---

## 📈 Project Progression

### Phase 1: Foundation  (COMPLETED)
- [x] Set up Go project structure with basic HTTP server
- [x] Create Chi router with `/`, `/health`, `/metrics` endpoints
- [x] Implement Prometheus metrics collection
- [x] Create static HTML/CSS landing page
- [x] Environment configuration management

### Phase 2: Containerization  (COMPLETED)
- [x] Write multi-stage Dockerfile for optimized builds
- [x] Create docker-compose for local development
- [x] Add Prometheus configuration and data collection
- [x] Add Grafana with datasource configuration
- [x] Set up health checks in docker-compose

### Phase 3: Kubernetes & Orchestration  (COMPLETED)
- [x] Create Kubernetes manifests (Deployment, Service, ConfigMap, Secret)
- [x] Set up Helm chart structure (v0.1.0)
- [x] Configure resource limits and requests
- [x] Implement LoadBalancer service for external access
- [x] Define proper labels and selectors

### Phase 4: CI/CD Pipeline  (COMPLETED)
- [x] Set up GitHub Actions workflow on `dev` branch
- [x] Implement Gitleaks secret scanning
- [x] Configure Docker build and push to registry
- [x] Set up image tagging strategy (`:latest` + `:{SHA}`)
- [x] Automated deployment workflow

### Phase 5: Production Readiness  (COMPLETED)
- [x] Implement comprehensive error handling
- [x] Add health check endpoints
- [x] Configure security best practices
- [x] Set up proper logging
- [x] Create comprehensive documentation

### Phase 6: ArgoCD Integration & GitOps  (COMPLETED)
- [x] Deploy ArgoCD controller in Kubernetes cluster (`argocd` namespace)
- [x] Create ArgoCD Application manifest (`argocd-app.yaml`)
- [x] Configure automated sync policies (auto-prune + self-healing)
- [x] Set up Git repository as source of truth (`dev` branch)
- [x] Helm chart integration for templating
- [x] Enable continuous deployment from Git changes
- [x] Verify ArgoCD pods and services (`kubectl get pods -n argocd`)

---

## 🗺️ Future Roadmap

### Phase 7: Multi-Environment GitOps (NEXT)

#### 7.1 Environment Separation
- [ ] Create separate ArgoCD Applications for dev/staging/prod
- [ ] Implement environment-specific values files
- [ ] Set up AppProject for RBAC and resource restrictions
- [ ] Configure different sync policies per environment:
  - Dev: Auto-sync (auto-deploy on git push)
  - Staging: Manual approval required
  - Prod: Scheduled syncs with manual gates

#### 7.2 Kustomize Overlays for Multi-Environment
```
helm/
├── gitops/
│   ├── base/
│   │   ├── kustomization.yaml
│   │   ├── deployment.yaml
│   │   └── service.yaml
│   └── overlays/
│       ├── dev/
│       │   ├── kustomization.yaml
│       │   └── values.yaml (dev-specific)
│       ├── staging/
│       │   ├── kustomization.yaml
│       │   └── values.yaml (staging-specific)
│       └── prod/
│           ├── kustomization.yaml
│           └── values.yaml (prod-specific, hardened)
```

#### 7.3 Secret Management for Production
- [ ] Implement Sealed Secrets for git-safe encryption
- [ ] OR use External Secrets Operator with:
  - AWS Secrets Manager
  - HashiCorp Vault
  - Azure Key Vault
- [ ] Remove hardcoded secrets from Git
- [ ] Automate secret rotation
- [ ] Set up audit logging for secret access

### Phase 8: Advanced ArgoCD Features

#### 8.1 Deployment Strategies
- [ ] **Sync Waves**: Order critical deployments (database → app → ingress)
- [ ] **Post-Sync Hooks**: Run smoke tests after deployment
- [ ] **Pre-Sync Hooks**: Backup database before changes
- [ ] **Canary Deployments**: Gradual rollout with traffic splitting
- [ ] **Blue-Green Deployments**: Zero-downtime updates

#### 8.2 ArgoCD Notifications & Monitoring
- [ ] Configure webhooks for Slack notifications
- [ ] Email notifications on deployment status
- [ ] Discord integration for team alerts
- [ ] Expose ArgoCD metrics to Prometheus
- [ ] Create Grafana dashboards for deployment metrics
- [ ] Set up alerts for failed deployments

#### 8.3 GitOps Workflow Enhancements
- [ ] Implement pull request preview deployments
- [ ] Automated changelog generation
- [ ] Release notes automation
- [ ] Deployment timelines and SLOs tracking
- [ ] Automatic rollback on health check failures

### Phase 9: Multi-Cluster & High Availability

- [ ] Set up ArgoCD with multiple clusters
- [ ] Cross-cluster failover strategies
- [ ] Cluster health monitoring
- [ ] Load balancing across clusters
- [ ] Disaster recovery procedures
- [ ] Backup and restore automation

### Phase 10: Security Hardening & Compliance

- [ ] Network Policies for pod-to-pod communication
- [ ] Pod Security Policies (PSP) or Pod Security Standards
- [ ] Container image scanning (Trivy, Snyk)
- [ ] Runtime security monitoring (Falco)
- [ ] RBAC for cluster access control
- [ ] Audit logging for all API calls
- [ ] Compliance monitoring (PCI DSS, HIPAA, SOC2)
- [ ] OPA/Kyverno policy enforcement

### Phase 11: Performance Optimization

- [ ] Horizontal Pod Autoscaler (HPA) configuration
- [ ] Vertical Pod Autoscaler (VPA) for resource optimization
- [ ] Cluster autoscaling for node management
- [ ] Performance benchmarking
- [ ] Cost optimization strategies
- [ ] Resource quota enforcement per namespace

### Phase 12: Observability at Scale

- [ ] OpenTelemetry tracing integration
- [ ] Distributed tracing with Jaeger
- [ ] Log aggregation (ELK Stack or Loki)
- [ ] Centralized application metrics
- [ ] Custom business metrics
- [ ] Alert rules and incident response automation


---

## 📦 Deployment Guide

### Prerequisites
- Kubernetes cluster (v1.20+)
- kubectl configured with cluster access
- Helm 3.x installed
- Docker registry access (Docker Hub or private)
- ArgoCD installed in cluster (`argocd` namespace) - ALREADY CONFIGURED

### Step 1: Verify Current Setup

```bash
# Check if ArgoCD is running
kubectl get pods -n argocd

# Check ArgoCD Application status
kubectl get application -n argocd
argocd app list

# Check backend pods
kubectl get pods
```

### Step 2: Manual Docker Build & Push (Optional - CI/CD Does This)

```bash
cd backend

# ✅ CORRECT: Build with semantic version tag
docker build -t nirmal08/ci-3-tier:v1.0.0 .

# Tag with commit hash for traceability
docker build -t nirmal08/ci-3-tier:v1.0.0-$(git rev-parse --short HEAD) .

# Push to registry
docker push nirmal08/ci-3-tier:v1.0.0

#  AVOID: Using :latest in production
# docker build -t nirmal08/ci-3-tier:latest .  # Only for dev/test
```

### Step 3: Update Helm Values for Deployment

```bash
# Edit helm/gitops/values.yaml with EXACT version tag
image:
  repository: nirmal08/ci-3-tier
  tag: v1.0.0              # ✅ Use exact version, not :latest
  pullPolicy: IfNotPresent

replicas: 2
```

### Step 4: Deploy with ArgoCD (Recommended)

```bash
# ArgoCD automatically syncs when Git is updated
# Just commit your changes to dev branch
git add helm/gitops/values.yaml
git commit -m "chore: Update image tag to v1.0.0"
git push origin dev

# ArgoCD will automatically deploy within sync interval (default: 3 mins)
# Or manually sync:
argocd app sync backend-app

# Monitor deployment
argocd app watch backend-app
kubectl get pods -w
```

### Step 5: Manual Helm Deployment (Without ArgoCD)

```bash
# Dry run first to validate
helm lint helm/gitops/
helm install my-backend helm/gitops/ --dry-run --debug

# Actual deployment
helm install my-backend helm/gitops/

# Or upgrade existing release
helm upgrade my-backend helm/gitops/
```

### Step 6: Verify Deployment Success

```bash
# Check pods are running
kubectl get pods -l app=gitops
kubectl get pods -w  # Watch for readiness

# Check service endpoint
kubectl get svc gitops-nodeport

# Get external IP/LB endpoint
kubectl get svc gitops-nodeport -o wide

# Test health endpoint
curl http://<EXTERNAL-IP>/health

# View logs
kubectl logs -f deployment/gitops-3-tier
```

### Step 7: Monitor with Prometheus/Grafana

```bash
# Port forward to Prometheus
kubectl port-forward -n default svc/prometheus 9090:9090 &

# Port forward to Grafana
kubectl port-forward -n default svc/grafana 3000:3000 &

# Access URLs:
# Prometheus: http://localhost:9090
# Grafana: http://localhost:3000 (admin/password)

# Monitor metrics
curl http://localhost:8000/metrics
```

### Step 8: Rollback to Previous Version (If Needed)

```bash
# With ArgoCD - revert the Git commit
git revert HEAD
git push origin dev
# ArgoCD will sync automatically

# With Helm - rollback to previous release
helm history my-backend
helm rollback my-backend 1  # Rollback to release 1
```

---

## 🏷️ Docker Image Tagging Strategy (CRITICAL FOR PRODUCTION)

### Understanding the Problem with :latest

The `:latest` tag is **unreliable in production environments** because:
- It's **mutable**: Anyone can push a new image and overwrite it
- It **provides no version control**: You can't easily rollback to a specific version
- It **breaks on redeploys**: Pod restarts pull whatever :latest points to currently
- It **violates GitOps principles**: Your deployed version doesn't match what's in Git

### Recommended Semantic Versioning Scheme

```bash
# Primary Tags - Use ONE of these for deployments
nirmal08/ci-3-tier:v1.0.0              # ✅ PRODUCTION - Specific patch version
nirmal08/ci-3-tier:v1.0.0-rc.1         # Release candidate
nirmal08/ci-3-tier:v1.0.0-beta.1       # Beta version
nirmal08/ci-3-tier:v1.0.0-alpha.1      # Alpha version

# Secondary Tags - For convenience only
nirmal08/ci-3-tier:v1.0                # Points to latest v1.0.x
nirmal08/ci-3-tier:v1                  # Points to latest v1.x.x
nirmal08/ci-3-tier:latest              # ❌ ONLY for dev/test, never production

# Commit-based Tags - For debugging
nirmal08/ci-3-tier:v1.0.0-abc123d      # Include git SHA
```

### Implementation in CI/CD

**Current GitHub Actions Workflow Update Needed**:

```yaml
# ❌ CURRENT (risky)
- name: Build and Push
  run: |
    docker build -t nirmal08/ci-3-tier:latest .
    docker push nirmal08/ci-3-tier:latest
    
# ✅ RECOMMENDED
- name: Determine Version
  run: |
    # Extract from git tag (recommended)
    VERSION=$(git describe --tags --always --abbrev=7)
    echo "VERSION=${VERSION}" >> $GITHUB_ENV
    
- name: Build and Push with Version
  run: |
    docker build -t nirmal08/ci-3-tier:${VERSION} .
    docker build -t nirmal08/ci-3-tier:latest .  # Also tag latest for dev
    docker push nirmal08/ci-3-tier:${VERSION}
    docker push nirmal08/ci-3-tier:latest
```

### Helm Values Configuration

```yaml
# ❌ RISKY: Using :latest
image:
  repository: nirmal08/ci-3-tier
  tag: latest        # Will pull different image on each pod restart
  
# ✅ SAFE: Using specific version
image:
  repository: nirmal08/ci-3-tier
  tag: v1.0.0        # Always pulls same image
  imagePullPolicy: IfNotPresent  # Use cached version if available
```

### Version Bumping Strategy

```bash
# In GitHub Actions or locally:

# 1. Patch version (bug fixes)
git tag v1.0.1
git push origin v1.0.1

# 2. Minor version (new features, backwards compatible)
git tag v1.1.0
git push origin v1.1.0

# 3. Major version (breaking changes)
git tag v2.0.0
git push origin v2.0.0

# CI/CD automatically builds and pushes with this tag
```

### Rollback Procedure (With Versioned Tags)

```bash
# Easy rollback because versions are immutable
helm rollback my-backend 1            # Rollback to previous release
kubectl rollout undo deployment/gitops-3-tier  # Undo last deployment

# View history with versions
helm history my-backend
# Revision  Updated                   Status      Chart            Version  Description
# 3         Mon May 27 12:00:00 2026  deployed   gitops-0.1.0     v1.0.2   Upgrade complete
# 2         Mon May 27 11:00:00 2026  superseded gitops-0.1.0     v1.0.1   Upgrade complete
# 1         Mon May 27 10:00:00 2026  superseded gitops-0.1.0     v1.0.0   Install complete
```

### Precautions Checklist

- [ ] **Never use `:latest` in production deployments**
- [ ] **Always specify exact version tags** (e.g., v1.0.0)
- [ ] **Pin versions in Helm values.yaml** - don't let them auto-update
- [ ] **Use ImagePullPolicy: IfNotPresent** - prevents pulling new :latest on pod restart
- [ ] **Test new versions in dev/staging first** before production deployment
- [ ] **Maintain a CHANGELOG** to track what changed between versions
- [ ] **Tag stable releases** in Git (e.g., `git tag v1.0.0`)
- [ ] **Keep old images in registry** for quick rollbacks
- [ ] **Document version in commit message** when bumping versions
- [ ] **Use SemVer strictly** (Major.Minor.Patch: v1.0.0)

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

## 🎯 Project Status & Summary

### Current Status: ✅ PRODUCTION-READY WITH ACTIVE ARGOCD

This is a **fully operational GitOps-enabled cloud-native Go backend** demonstrating complete DevOps best practices:

**Completed Implementations (6 Phases):**
1. ✅ **Phase 1**: Go HTTP server with Prometheus metrics
2. ✅ **Phase 2**: Docker containerization with multi-stage builds
3. ✅ **Phase 3**: Kubernetes orchestration with Helm charts
4. ✅ **Phase 4**: GitHub Actions CI/CD with security scanning
5. ✅ **Phase 5**: Production-ready configuration and documentation
6. ✅ **Phase 6**: ArgoCD GitOps integration with automated deployments

**Active Features:**
- 🎯 **GitOps Enabled**: Changes to Git automatically deploy to Kubernetes
- 📊 **Full Observability**: Prometheus metrics + Grafana dashboards
- 🔒 **Security Hardened**: Multi-stage builds, Gitleaks scanning, secret management
- 🚀 **Automated Deployment**: GitHub Actions → Docker Hub → ArgoCD → Kubernetes
- 💾 **Infrastructure as Code**: All resources defined in Git (helm charts, k8s manifests)
- 🔄 **Auto-Healing**: ArgoCD continuously reconciles desired state

### Critical Precautions (⚠️ DO NOT SKIP)

#### 1. **Docker Image Tagging** (Most Important)
```bash
# ❌ NEVER use this in production:
docker build -t nirmal08/ci-3-tier:latest .

# ✅ ALWAYS use semantic versioning:
docker build -t nirmal08/ci-3-tier:v1.0.0 .

# Update Helm values with exact version:
image:
  tag: v1.0.0  # Not "latest"
```

**Why?** `:latest` tag is mutable and breaks GitOps principles. You lose version control, cannot rollback reliably, and violate infrastructure-as-code practices.

#### 2. **Issues to Address**
- [ ] Update GitHub Actions to use Go 1.24.3 (currently 1.21)
- [ ] Fix HTML static file references (style.css → index.css)
- [ ] Update deployment.yml port exposure to match app port
- [ ] Move hardcoded secrets to Sealed Secrets or External Secrets Operator
- [ ] Fix replica count mismatch between values.yaml and deployment
- [ ] Implement secret rotation policies

#### 3. **Before Production Deployment**
- [ ] Use version tags (v1.0.0), NOT :latest
- [ ] Test in staging environment first
- [ ] Run `helm lint` to validate charts
- [ ] Verify all pods reach ready state
- [ ] Check metrics are flowing to Prometheus
- [ ] Test rollback procedures
- [ ] Review ArgoCD sync status
- [ ] Monitor logs for errors

### Next Steps (Phase 7+)

**Immediate:**
- Fix issues identified above
- Implement multi-environment setup (dev/staging/prod)
- Move secrets to proper secret management solution

**Short Term:**
- Add Sealed Secrets for git-safe secret storage
- Set up environment promotion pipeline
- Implement sync waves for ordered deployments

**Long Term:**
- Multi-cluster management
- Advanced observability (tracing, log aggregation)
- Security hardening (RBAC, network policies, PSP)
- Performance optimization (HPA, VPA, cluster autoscaling)

### Key Achievements

This project demonstrates:
- **End-to-End GitOps**: From code push to production deployment
- **Production Patterns**: Multi-stage builds, health checks, resource limits
- **Security Best Practices**: Secret scanning, minimal images, RBAC-ready
- **Observability**: Built-in metrics collection and visualization
- **Automation**: Fully automated from Git to Kubernetes
- **Version Control**: Infrastructure defined entirely in Git
- **Declarative Management**: ArgoCD ensures desired state = actual state

### Recommended Reading

- [GitOps Best Practices](https://www.gitops.tech/)
- [Semantic Versioning](https://semver.org/)
- [Docker Best Practices](https://docs.docker.com/develop/dev-best-practices/)
- [Helm Chart Best Practices](https://helm.sh/docs/chart_best_practices/)
- [ArgoCD Documentation](https://argo-cd.readthedocs.io/)
- [Kubernetes Production Best Practices](https://kubernetes.io/docs/concepts/configuration/overview/)
