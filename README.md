#  ParkingStatus Microservice

This microservice is part of the **SysPark** system and belongs to the **Parking** domain.  
It provides information about available parking slots.

---

##  Features

- Exposes a single HTTP GET endpoint:
  - `GET /api/parking/disponibles`
- Returns a list of available parking slots.
- Currently returns mock (simulated) data.
- Dockerized and ready for local testing or QA environments.

---

##  Technologies

- Language: [Go (Golang)](https://go.dev/) `v1.22+`
- Architecture: Clean modular structure (`cmd/`, `internal/`)
- Container: Docker + docker-compose

---

##  Getting Started

### 1. Clone the repository

```bash
git clone https://github.com/Jimb055/parking-status.git
cd parking-status
```

### 2. Build and run with Docker Compose

```bash
docker-compose up --build
```

### 3. Access the endpoint

```bash
GET http://localhost:8080/api/parking/disponibles
```

Response:
```json
[
  {
    "id": 1,
    "location": "Zone A - 01",
    "status": "available"
  },
  {
    "id": 2,
    "location": "Zone A - 03",
    "status": "available"
  }
]
```

---

## 📂 Project Structure

```
parking-status/
├── cmd/                    # Application entry point
│   └── main.go
├── internal/
│   ├── handler/            # HTTP handlers
│   │   └── parking_handler.go
│   └── service/            # Business logic
│       └── parking_service.go
├── go.mod
├── Dockerfile
├── docker-compose.yml
├── .gitignore
└── README.md
```

---

## 📌 Notes

- This service currently returns **simulated data** via hardcoded logic in `parking_service.go`.
- Future versions may integrate PostgreSQL or another data source.

---

## 🧑‍💻 Author

Developed by [Jimb055](https://github.com/Jimb055) as part of the SysPark distributed architecture project.
