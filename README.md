# analysis_service

`analysis_service` is the backend service for the **main Chrome extension**.
It is built using **FastAPI** and exposes a **GraphQL API** powered by **Ariadne**.

This service is responsible for handling business logic, data processing, and API communication required by the extension.

---

## 🧱 Tech Stack

* **Python**
* **FastAPI** – high-performance async web framework
* **Ariadne** – GraphQL server library
* **Uvicorn** – ASGI server
* **GraphQL** – API layer
* **requirements.txt** – dependency management

---

## 🚀 Getting Started

### 1️⃣ Clone the Repository

```bash
git clone "https://github.com/FitRang/analysis_service"
cd analysis_service
```

---

### 2️⃣ Create & Activate Virtual Environment (Recommended)

```bash
python -m venv venv
source venv/bin/activate  # Linux / Mac
venv\Scripts\activate     # Windows
```

---

### 3️⃣ Install Dependencies

```bash
pip install -r requirements.txt
```

---

### 4️⃣ Run the Server

```bash
uvicorn app.main:app --reload
```

The service will be available at:

* **REST base**: `http://localhost:8000`
* **GraphQL endpoint**: `http://localhost:8000/graphql`

---

## 🔮 GraphQL

This service uses **Ariadne** to expose a GraphQL API.

### Features:

* Schema-first GraphQL design
* Modular resolvers
* Async resolver support
* Integrated with FastAPI

You can test GraphQL queries using:

* GraphiQL (if enabled)
* Postman
* Apollo Client
* Chrome extension frontend

---

## 📦 Dependencies

All dependencies are managed via `requirements.txt`.

Example:

```txt
fastapi
uvicorn
ariadne
pydantic
```
---

## 🔐 Usage with Chrome Extension

* The Chrome extension communicates with this service via GraphQL.
* This backend handles data analysis, validation, and response formatting.
* Designed to be lightweight and fast for extension use-cases.

---

## 🛠 Development Notes

* Follow async patterns for better performance.
* Keep resolvers thin; move logic into service layers.
* Use Pydantic models for validation where applicable.
