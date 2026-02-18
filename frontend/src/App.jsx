import { useEffect, useState } from "react";
import axios from "axios";
import NodeStatus from "./components/NodeStatus";
import SignPanel from "./components/SignPanel";
import LogsPanel from "./components/LogsPanel";

export default function App() {
  const [statuses, setStatuses] = useState([]);
  const [logs, setLogs] = useState([]);
  const [loading, setLoading] = useState(false);

  const API = "http://localhost:5000";

  const addLog = (msg) => {
    setLogs((prev) => [...prev, msg]);
  };

  const checkHealth = async () => {
    try {
      const res = await axios.get(`${API}/health`);
      setStatuses(res.data.statuses);
    } catch {
      addLog("Coordinator unreachable");
    }
  };

  const signTransaction = async () => {
    setLoading(true);
    addLog("Initiating signing request...");

    try {
      const res = await axios.post(`${API}/sign`, {
        message: "Hybrid MPC Demo",
      });

      addLog("Partial signatures collected");
      addLog("Final signature generated");
      addLog("Signature:");
      addLog(JSON.stringify(res.data.signature));
    } catch (err) {
      addLog("Signing failed");
    }

    setLoading(false);
  };

  useEffect(() => {
    checkHealth();
    const interval = setInterval(checkHealth, 3000);
    return () => clearInterval(interval);
  }, []);

  return (
    <div className="p-10 min-h-screen">
      <h1 className="text-3xl font-bold mb-10">
        Hybrid Threshold MPC Wallet
      </h1>

      <div className="grid grid-cols-3 gap-6">
        <NodeStatus statuses={statuses} />
        <SignPanel onSign={signTransaction} loading={loading} />
        <LogsPanel logs={logs} />
      </div>
    </div>
  );
}
