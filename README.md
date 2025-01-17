# Config Store

A simple key-value store application built with **Go** and **PostgreSQL**. It provides RESTful APIs to manage key-value pairs.

---

## Setup Using Docker

### 1. Clone the Repository

```bash
git clone https://github.com/basebandit/config-store.git
cd config-store
```

### 2. Set Up Environment Variables

Create a `.env` file in the root directory:

```bash
# .env
DB_URL=postgresql://${DB_USER}:${DB_PASSWORD}@db:5432/${DB_NAME}
```

Replace `your_db_user`, `your_db_password`, and `your_db_name` with your PostgreSQL credentials.

### 3. Build and Run the Application locally

Start the application and database using Docker Compose:

```bash
docker-compose up --build
```
Using docker to build for multiple platforms and push to docker hub

```bash
DOCKER_BUILDKIT=1 docker buildx build --platform linux/arm64,linux/amd64 --push -t baseband1t/config-store:1.0.0 . 
```

The application will be available at [http://localhost:3000](http://localhost:3000).

---

## Helm Chart

A Helm chart is available to deploy this application on Kubernetes. The chart is hosted in the `basebandit/config-store` repository.

### 1. Add the Helm Repository

```bash
helm repo add basebandit https://basebandit.github.io/config-store
helm repo update
```

### 2. Deploy the Application

Install the chart:

```bash
helm install config-store basebandit/config-store
```

Verify the deployment:

```bash
kubectl get pods
```

### 3. Access the Application

Use the service IP or configure an ingress to access the application.

---

## API Documentation

### Base URL

```bash
http://localhost:3000/api
```

### Endpoints

- **Create Key-Value Pair**: `POST /kv`
- **Get All Key-Value Pairs**: `GET /kv`
- **Get Specific Key-Value Pair**: `GET /kv/:key`
- **Update Key-Value Pair**: `PUT /kv/:key`
- **Delete Key-Value Pair**: `DELETE /kv/:key`