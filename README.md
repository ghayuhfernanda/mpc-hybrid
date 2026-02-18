# Hybrid Threshold MPC Wallet PoC

> This project demonstrates a Proof of Concept (PoC) for a **Hybrid Threshold MPC (Multiparty Computation) Wallet** architecture using real threshold ECDSA (GG18).  
It includes:

> 1️⃣ **3 Go MPC nodes** simulating threshold key generation and signing (dummy handlers for now)  
> 2️⃣ **Node.js coordinator** to orchestrate signing between nodes  
> 3️⃣ **React + Tailwind CSS frontend** with live node status and signing UI  

---

## Architecture Overview

```
[Frontend (React)] <--> [Coordinator (Node.js)] <--> [3 MPC Nodes (Go)]
```
✅ Frontend provides user interface and visual feedback  
✅ Coordinator manages MPC protocol flow, contacting MPC nodes  
✅ MPC nodes hold secret shares and perform partial signing  

---

## Prerequisites

✅ [Go 1.18+](https://golang.org/dl/)  
✅ [Node.js 16+](https://nodejs.org/en/download/)  
✅ [npm](https://www.npmjs.com/get-npm)  
✅ (Optional) [Docker & Docker Compose](https://docs.docker.com/get-docker/)

---

## Setup & Run Instructions

### 1. Run MPC Nodes

```
cd mpc-node
go mod tidy
```
> Run one node per terminal, change port in main.go each time (:8001, :8002, :8003)
```
go run *.go
```
> The nodes expose REST APIs: /health, /keygen, /sign

### 2. Run Coordinator
```
cd coordinator
npm install
```
> Make sure nodes URLs in index.js use localhost and ports 8001-8003
```
node index.js
```
Coordinator runs on http://localhost:5000

### 3. Run Frontend
```
cd frontend
npm install
npm run dev
```
> Frontend runs on http://localhost:5173

> Connects to coordinator API at http://localhost:5000

### Usage
```
- Open http://localhost:5173 in your browser
- Confirm all MPC nodes show online status
- Click Sign Transaction to initiate threshold signing
- View signing logs and status updates
- Troubleshooting
- Coordinator unreachable
- Ensure coordinator is running (node index.js)
- Confirm coordinator CORS is enabled (app.use(cors()))
- Verify frontend API URL points to http://localhost:5000
```

### Project Structure
```
mpc-hybrid/
├── coordinator/                    # Node.js coordinator orchestrating MPC nodes
├── frontend/                       # React + Tailwind frontend UI 
├── mpc-node/                       # Go MPC nodes (3 nodes run with different ports)
├── docker-compose.yml (optional)   # For running all components in Docker containers
└── README.md
```

### Run Docker Build
```
docker-compose up --build
```

### Next Steps
```
📃 Replace dummy keygen/sign handlers with real GG18 threshold ECDSA logic using tss-lib
📃 Add visual signing share progress and status updates in frontend
📃 Deploy to cloud environments or container orchestration systems
```

### License
This project is provided as-is for educational/demo purposes.