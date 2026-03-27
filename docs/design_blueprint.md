# Headless CRM Design Blueprint for Indonesian Telco

## 1. Executive Summary
This document outlines the architectural blueprint for a cloud-native, headless CRM designed for a telecommunications provider in Indonesia. The system leverages a microservices architecture, adhering to TM Forum Open APIs for interoperability and industry alignment. Key objectives include agility, scalability, and seamless integration with the local digital ecosystem (e.g., e-wallets, regulatory compliance).

## 2. Business Architecture
The CRM supports the full customer lifecycle from acquisition to retention.

*   **Customer Management**: 360-degree view, KYC (Prepaid/Postpaid), Profile Management.
*   **Product Catalog**: Digital & Connectivity products, Bundling.
*   **Order Management**: Capture, Validation, Decomposition, Fulfillment tracking.
*   **Billing & Payments**: Real-time balance, Postpaid invoicing, Top-ups, Payment Gateway integration.
*   **Engagement**: Omnichannel support (Mobile App, Web, WhatsApp, USSD).

## 3. Application Architecture (Microservices & TMF Mapping)
The backend is composed of domain-driven microservices exposing TMF Open APIs. Each service is autonomous, owning its data and logic.

### Domain: Customer Management
*   **Customer Service** (`TMF629`, `TMF632`)
    *   **Purpose**: managing individual and organization customer data.
    *   **Role**: The "Single Source of Truth" for customer profiles. it centralizes demographic data, preferences, and contact details, ensuring consistency across all channels.
*   **Party Management Service** (`TMF632`)
    *   **Purpose**: managing innovative parties (partners, vendors, dealers).
    *   **Role**: Handles the ecosystem interacting with the Telco, distinct from end-customers. Essential for B2B2C models or dealer networks.
*   **KYC Service**
    *   **Purpose**: Specialized service for Indonesian regulatory validation (Dukcapil).
    *   **Role**: Ensures compliance with Kominfo regulations by validating NIK/KK against the government database *before* service activation.

### Domain: Product & Catalog
*   **Product Catalog Service** (`TMF620`)
    *   **Purpose**: Manages product definitions, offerings, and specifications.
    *   **Role**: The commercial brain. It defines *what* you sell, including bundles, pricing rules, and eligibility criteria.
*   **Product Inventory Service** (`TMF637`)
    *   **Purpose**: Tracks customer's subscribed products.
    *   **Role**: Tells you *what* a specific customer owns/uses right now (e.g., "Active 5GB Data Plan").

### Domain: Order Management
*   **Product Order Service** (`TMF622`)
    *   **Purpose**: Handles order capture and lifecycle.
    *   **Role**: The conductor. It receives the "I want to buy X" request, validates eligibility, and orchestrates the flow until completion.
*   **Service Order Service** (`TMF641`)
    *   **Purpose**: Manages technical service activation requests.
    *   **Role**: Translates a commercial order (Product) into technical instructions (Service) for the network (e.g., "Provision HSS profile").

### Domain: Usage, Balance & Assurance
*   **Trouble Ticket Service** (`TMF621`)
    *   **Purpose**: Manages customer complaints and technical issues.
    *   **Role**: Records, tracks, and resolves problems reported by customers (e.g., "Network slow", "Billing error").
*   **Usage Management Service** (`TMF635`)
    *   **Purpose**: Tracks data/voice/SMS usage.
    *   **Role**: Provides visibility into consumption history (CDR logs) for customers and agents.
*   **Account Balance Service** (`TMF654`)
    *   **Purpose**: Real-time balance checks and management.
    *   **Role**: The wallet. It holds the prepaid credit or monetary balance, queried in real-time before purchase.

### Domain: Resource & Service
*   **Resource Inventory Service** (`TMF639`)
    *   **Purpose**: Manages logical/physical resources (SIM, MSISDN).
    *   **Role**: Inventory keeper for physical assets (SIM Cards) and logical numbers (MSISDNs), ensuring no double-assignment.

## 4. Technology Architecture
**Strictly Google Cloud Platform (GCP) Native** stack for maximum integration, managed operations, and AI capabilities.

*   **Compute: Google Kubernetes Engine (GKE)**
    *   **Purpose**: Managed container orchestration.
    *   **Role**: Runs all microservices with auto-scaling and self-healing capabilities.
*   **Service Mesh: Cloud Service Mesh (managed Istio)**
    *   **Purpose**: Traffic management, security, and observability.
    *   **Role**: Handles mTLS encryption, rate limiting, and circuit breaking transparently, so code doesn't have to.
*   **API Gateway: Apigee X**
    *   **Purpose**: Advanced API management and monetization.
    *   **Role**: The front door for external partners. It secures, throttles, and potentially monetizes API access.
*   **Agentic Framework**
    *   **Exposure**: Open APIs (App) & MCP/A2A (Agent).
    *   **AI Platform: Vertex AI (Gemini 2.5)**
        *   **Purpose**: Central brain for agentic reasoning.
        *   **Role**: Understands natural language intentions (e.g., "Why is my bill high?") and calls the appropriate TMF APIs or Agentic Interfaces to resolve them.
*   **Frontend / Channels**
    *   **Web Portal: Next.js (React)** on Cloud Run.
        *   **Purpose**: Dynamic Agent/Admin dashboard.
        *   **Role**: Provides a comprehensive interface for internal staff to serve customers efficiently.
    *   **Mobile App: Flutter**
        *   **Purpose**: B2C Customer App.
        *   **Role**: The primary digital touchpoint for customers to manage their lifestyle and telco needs.
*   **Event Bus: Cloud Pub/Sub**
    *   **Purpose**: Global asynchronous messaging.
    *   **Role**: Decouples services. When an order is placed, an event is published here so Billing, Analytics, and Notification systems can react independently without blocking the Order service.
*   **Data Architecture**
    *   **Operational Database**:
        *   **Cloud SQL (PostgreSQL)**: For relational data (Customer info, Order transactions) requiring ACID compliance.
        *   **Firestore**: For flexible document data (Product Catalog hierarchy).
        *   **Memorystore (Redis)**: For high-speed caching (Session management, token storage).
    *   **Unified Data Platform**:
        *   **Ingestion (Dataflow)**: Transforming real-time streams into analytic-ready data.
        *   **Warehouse (BigQuery)**: Enterprise Data Warehouse. Stores petabytes of data for deep analytics, reporting, and AI model training.

## 5. Integration Architecture

### Northbound (Channels)
*   **BFF (Backend for Frontend)**
    *   **Purpose**: GraphQL aggregation layer for Mobile App and Web.
    *   **Role**: Acts as a smart proxy that aggregates data from multiple microservices into a single response, optimized for specific UI screens (e.g., "Home Screen" query). Prevents over-fetching and reduces network round-trips.
*   **WhatsApp Bot**
    *   **Purpose**: Lightweight interaction channel.
    *   **Role**: Meets customers where they are. Handles high-frequency, low-complexity tasks (Balance check, simple top-up) via chat.

### Southbound (Core Telco & BSS)
*   **OCS (Online Charging System)**
    *   **Purpose**: Real-time credit control.
    *   **Role**: The "Cash Register". It approves or denies service usage in real-time based on balance (Prepaid).
*   **Billing System (Postpaid)**
    *   **Purpose**: Cycle processing and invoicing.
    *   **Role**: Aggregates usage at the end of the month, applies discounts, and generates the PDF bill.
*   **Order Management / Fulfillment (SOM/COM)**
    *   **Purpose**: Network provisioning orchestration.
    *   **Role**: The heavy lifter that actually talks to the network elements to turn the lights on.
*   **Voucher Management System (VMS)**
    *   **Purpose**: Lifecycle of physical/digital vouchers.
    *   **Role**: Generates and validates the scratch-card PINs for top-ups.
*   **Top-up System**
    *   **Purpose**: Electronic/Physical redemption.
    *   **Role**: Bridges banking/retail channels to the Telco balance.
*   **HSS/HLR**
    *   **Purpose**: Subscriber Database.
    *   **Role**: The network's list of valid SIMs and their allowed services (4G/5G, Roaming).

### East/West (Third-Party)
*   **Payment Gateways** (DOKU, Midtrans)
    *   **Role**: Processes credit card and e-wallet transactions securely.
*   **Dukcapil** (Gov ID Database)
    *   **Role**: The authoritative source for Indonesian citizen identity verification.

## 6. Indonesian Localization & Compliance
*   **Prepaid Registration**: Mandatory NIK/KK validation to prevent anonymous usage.
*   **Data Residency**: Ensuring PII stays within Indonesian GCP regions (Jakarta `asia-southeast2`) to comply with PP No. 71/2019.
*   **Payment Ecosystem**: Native integration with QRIS and local E-Wallets (GoPay, OVO) which dominate the market.

## 7. Security Architecture
*   **Identity Platform (Firebase Auth)**: Manages customer login/signup (OTP, Social Login).
*   **Cloud KMS**: Manages cryptographic keys for encrypting sensitive fields (NIK, Credit Card tokens).

## 8. Channel Architecture (Frontend)

### Web CRM Portal (Next.js)
*   **Audience**: Customer Service Agents, Administrators, Dealers.
*   **Features**:
    *   **Agentic Chat**: Powered by **Gemini 2.5**. Can autonomously navigate pages and execute TMF/MCP actions.
    *   **360 View**: Aggregated dashboard of Customer Profile + Products + Interaction History.

### Mobile App (Flutter)
*   **Audience**: End Customers (B2C).
*   **Features**: Self-service dashboard, Biometric login, One-tap buying.

## 9. Diagram (Mermaid)

```mermaid
flowchart TD
    %% Diagram Updated
    classDef micro fill:#e1f5fe,stroke:#01579b,stroke-width:2px;
    classDef legacy fill:#ffebee,stroke:#b71c1c,stroke-width:2px;
    classDef channel fill:#f3e5f5,stroke:#4a148c,stroke-width:2px;
    classDef ai fill:#e8f5e9,stroke:#1b5e20,stroke-width:2px,stroke-dasharray: 5 5;
    classDef data fill:#fff3e0,stroke:#e65100,stroke-width:2px;

    subgraph Frontend ["Frontend Layer (Channels)"]
        direction TB
        Portal["Web CRM Portal<br/>(Next.js)"]:::channel
        MobileApp["Mobile App<br/>(Flutter)"]:::channel
        WA["WhatsApp Bot"]:::channel
    end

    subgraph API_GW ["API Gateway / BFF"]
        APIGW[API Gateway]
        BFF[GraphQL BFF]
    end

    subgraph Microservices ["Headless CRM Microservices (TMF APIs & MCP)"]
        direction TB
        Cust["Customer Mgmt<br/>TMF629/632<br/>(MCP)"]:::micro
        Prod["Product Catalog<br/>TMF620<br/>(MCP)"]:::micro
        Order["Product Order<br/>TMF622<br/>(MCP)"]:::micro
        Inv["Service Inventory<br/>TMF638<br/>(MCP)"]:::micro
        Bal["Balance Mgmt<br/>TMF654<br/>(MCP)"]:::micro
        KYC["KYC Service<br/>(MCP)"]:::micro
        Tkt["Trouble Ticket<br/>TMF621<br/>(MCP)"]:::micro
    end

    subgraph Agentic ["Agentic & Data Platform (Vertex AI)"]
        Vertex["Vertex AI Agent"]:::ai
        BQ["BigQuery EDW"]:::data
    end

    subgraph Southbound ["Southbound Systems (Legacy/Core)"]
        OCS[OCS]:::legacy
        Bill[Billing System]:::legacy
        M_Cat[Master Catalog]:::legacy
        SOM[Order Fulfillment]:::legacy
        VMS[Voucher Mgmt]:::legacy
        TopUp[Top-up System]:::legacy
        HSS[HSS/HLR]:::legacy
    end

    subgraph External ["External/Reg"]
        Dukcapil["Dukcapil (NIK/KK)"]
        PayGW[Payment Gateway]
    end

    Portal --> BFF
    MobileApp --> BFF
    WA --> APIGW
    BFF --> APIGW
    APIGW --> Cust & Prod & Order & Inv & Bal & KYC & Tkt

    Vertex -.->|MCP/A2A| Cust & Prod & Order & Inv & Bal & KYC & Tkt
    Cust & Prod & Order & Inv & Bal & KYC & Tkt -.->|Data Pipeline| BQ

    Cust --> KYC
    KYC --> Dukcapil
    Order --> SOM
    Order --> PayGW
    Bal --> OCS
    Bal --> TopUp
    TopUp --> VMS
    Inv --> Bill

## 10. GCP Services Catalog

A comprehensive list of all Google Cloud Platform services used across the solution.

| Service Category | GCP Service | Role in Architecture |
| :--- | :--- | :--- |
| **Compute** | Google Kubernetes Engine (GKE) | Runs microservices with auto-scaling and self-healing. |
| **Compute** | Cloud Run | Serverless compute for Next.js web portal. |
| **Storage & Databases** | Cloud SQL (PostgreSQL) | Relational database for transactional data (ACID). |
| **Storage & Databases** | Firestore | Flexible document storage for Product Catalog hierarchy. |
| **Storage & Databases** | Memorystore (Redis) | High-speed caching for sessions and tokens. |
| **Storage & Databases** | Cloud Storage (GCS) | Object storage for static assets, reports, and backups. |
| **Networking** | Apigee X | API Management and gateway for external partners. |
| **Networking** | Cloud Service Mesh | Managed Istio for inter-service communication and tracking. |
| **Networking** | Cloud Load Balancing | Global traffic routing and edge caching. |
| **AI & ML** | Vertex AI (Gemini 2.5) | Generative AI engine for agentic reasoning and NLP. |
| **Analytics** | BigQuery | Enterprise Data Warehouse for analytics and reporting. |
| **Integration** | Cloud Pub/Sub | Asynchronous messaging for decoupling services. |
| **Integration** | Dataflow | Real-time stream processing and data ingestion. |
| **Security** | Firebase Auth | Identity platform for user login/signup. |
| **Security** | Cloud KMS | Cryptographic key management for sensitive data encryption. |
| **Security** | Secret Manager | Secure storage for API keys and credentials. |
| **DevOps & Operations** | Cloud Operations Suite | Logging and Monitoring for full-stack observability. |
| **DevOps & Operations** | Cloud Build | CI/CD automation for building container images. |
| **DevOps & Operations** | Artifact Registry | Secure storage for container images and packages. |
