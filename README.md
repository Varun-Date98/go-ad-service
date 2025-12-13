# Go Ad Service

A **stateless, distributed ad-serving backend** written in Go that performs real-time ad targeting and enforces **user-level frequency capping** using Redis.

The service is designed to scale horizontally behind a load balancer and mirrors the core architecture of modern ad platforms (e.g. Twitch / Amazon Ads).

---

## 🚀 Features

- Real-time ad targeting (country, age, placement, creator, device, language, interests)
- Campaign → Creative separation
- Redis-backed **frequency capping** using atomic `SETNX + TTL`
- Stateless service, safe for horizontal scaling
- Typed SQL access via `sqlc`
- Schema migrations via `goose`
- Fail-open behavior if Redis is unavailable

---

## 🧠 Architecture

Client → Load Balancer → Go Ad Service (N replicas)  
                             │  
                             ├── Postgres (campaigns, users, targeting)  
                             └── Redis (frequency capping)

All state required for correctness lives outside the service.

---

## 🔁 Ad Serving Flow

1. Receive request with user and placement context
2. Fetch user profile from Postgres
3. Query eligible campaign + creative candidates
4. Score and select the best candidate
5. Atomically cap the winning campaign in Redis
6. Return ad or `204 No Content`

---

## 🧰 Tech Stack

- Go (net/http, chi)
- PostgreSQL
- Redis
- sqlc
- goose
- Docker & Docker Compose

---

## 🐳 Running Locally

```bash
docker compose up --build
```

Run migrations and seed data:

```bash
goose -dir sql/schema postgres "$DB_URL" up
goose -dir sql/seed postgres "$DB_URL" up
```

---

## 📡 API

### Get an Ad
```
GET /v1/ad?user_id=...&placement_id=...&creator_id=...
```

Returns:
- `200 OK` with ad payload
- `204 No Content` if no eligible ad

### Health Check
```
GET /v1/healthz
```

---

## 📈 Scaling

Run multiple stateless API replicas behind a load balancer:

```bash
docker compose up --scale api=3
```

All replicas share Postgres and Redis.

---

## 🔒 Frequency Capping

- Key: `cap:{user_id}:{campaign_id}`
- Enforced with Redis `SETNX + TTL`
- Prevents repeated ad exposure across replicas

---

## 📄 License

MIT