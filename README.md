# analysis_service

`analysis_service` is the backend service for the **main Chrome extension**.
It is implemented in **Go** and exposes a **GraphQL API** for fast, structured communication with the extension.

This service is responsible for **business logic**, **data analysis**, and **API orchestration** required by the extension.

---

## 🧱 Tech Stack

* **Go**
* **GraphQL** – API layer
* **gqlgen** (or equivalent) – GraphQL server implementation
* **net/http** – HTTP server
* **GitHub Actions** – CI/CD
* **Docker** – Containerization (optional but recommended)

## 🔐 Usage with Chrome Extension

* The Chrome extension communicates with this service via **GraphQL**
* This backend performs:

  * Data analysis
  * Validation
  * Aggregation
  * Response shaping
* Optimized for **fast responses** and **minimal payloads**

