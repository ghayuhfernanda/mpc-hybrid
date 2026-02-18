export default function LogsPanel({ logs }) {
  return (
    <div className="bg-gray-900 p-6 rounded-xl shadow-lg">
      <h2 className="text-xl mb-4 font-semibold">
        Signing Logs
      </h2>

      <div className="bg-gray-800 p-4 rounded text-xs h-48 overflow-y-auto">
        {logs.map((log, index) => (
          <div key={index} className="mb-1">
            {log}
          </div>
        ))}
      </div>
    </div>
  );
}
