# Headless CRM - Indonesian Telco

A Cloud-Native, Headless CRM designed for the Indonesian Telecommunications market, leverage Google Cloud Platform (GCP) and TM Forum Open APIs.

## Architecture Overview

This project follows a microservices architecture with a headless design, separating the backend logic (TMF APIs) from the frontend channels (Mobile App, Web Portal).

### Tech Stack
- **Cloud**: Google Cloud Platform (GCP)
- **Container Orchestration**: GKE (Google Kubernetes Engine)
- **Backend Languages**: Go / Java / Python (Service dependent)
- **Frontend**: Next.js (Web), Flutter (Mobile)
- **AI/Agent**: Vertex AI (Gemini 2.5), MCP Protocol
- **Database**: Cloud SQL, Firestore, BigQuery

## Project Structure

```
├── services/               # Backend Microservices (TMF APIs)
│   ├── customer-mgmt/      # TMF629/632 Customer & Party
│   ├── product-catalog/    # TMF620 Product Catalog
│   ├── order-mgmt/         # TMF622/641 Product & Service Order
│   ├── usage-balance/      # TMF635/654 Usage & Balance
│   ├── resource-inventory/ # TMF639 Resource Inventory
│   └── kyc-service/        # Indonesian KYC (Dukcapil)
├── frontend/               # channel Applications
│   ├── web-portal/         # Next.js Agent/Admin Portal
│   ├── mobile-app/         # Flutter Customer App
├── infrastructure/         # IaC and Config
│   ├── k8s/                # Kubernetes Manifests (Helm/Kustomize)
│   ├── terraform/          # GKE, Pub/Sub, CloudSQL provisioning
│   └── istio/              # Service Mesh Config
├── data/                   # Data Platform
│   ├── bigquery/           # Schemas & Queries
│   └── dataflow/           # Beam Pipelines
└── docs/                   # Architecture & Design Documents
```

## Getting Started

*(Instructions to be added as services are implemented)*

## License

Private / Proprietary
