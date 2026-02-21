const express = require("express");
const cors = require("cors");
const axios = require("axios");

const app = express();
app.use(cors());
app.use(express.json());

const nodes = [
  "http://localhost:8001",
  "http://localhost:8002",
  "http://localhost:8003"
];

app.get("/health", async (req, res) => {
  const statuses = await Promise.all(
    nodes.map(async (n) => {
      try {
        await axios.get(n + "/health");
        return "online";
      } catch {
        return "offline";
      }
    })
  );

  res.json({ statuses });
});

app.post("/sign", async (req, res) => {
  const message = Buffer.from(req.body.message);

  const sigs = await Promise.all([
    axios.post(nodes[0] + "/sign", { message }),
    axios.post(nodes[1] + "/sign", { message })
  ]);

  const combined = combinedSignatures(sigs[0].data, sigs[1].data);

  res.json({ signature: combined });
});

app.listen(5000, () =>
  console.log("Coordinator running on 5000")
);
