# Headless CRM - Telecommunications

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

## Implemented Services (Phase 2 Completed)

| Service | TMF API | Description | Status | Documentation |
| :--- | :--- | :--- | :--- | :--- |
| **Account Balance** | TMF654 | Balance Management & Adjustments | :white_check_mark: | [Walkthrough](./services/usage-balance/WALKTHROUGH.md) |
| **Customer Mgmt** | TMF629 | Customer Profile & Party | :white_check_mark: | [Walkthrough](./services/customer-mgmt/WALKTHROUGH.md) |
| **KYC Service** | - | Indonesian Identity Validation | :white_check_mark: | [Walkthrough](./services/kyc-service/WALKTHROUGH.md) |
| **Product Catalog** | TMF620 | Product Offerings & Specifications | :white_check_mark: | [Walkthrough](./services/product-catalog/WALKTHROUGH.md) |
| **Product Order** | TMF622 | Order Capture & Lifecycle | :white_check_mark: | [Walkthrough](./services/order-mgmt/WALKTHROUGH.md) |
| **Usage Mgmt** | TMF635 | Usage Records (CDR) & Rating | :white_check_mark: | [Walkthrough](./services/usage-mgmt/WALKTHROUGH.md) |

## Getting Started

See the [Walkthrough](./walkthrough.md) for a comprehensive guide on verifying all services.

## License

Private / Proprietary
