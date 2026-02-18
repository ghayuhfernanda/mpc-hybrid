export default function NodeStatus({ statuses }) {
  return (
    <div className="bg-gray-900 p-6 rounded-xl shadow-lg">
      <h2 className="text-xl mb-4 font-semibold">MPC Nodes</h2>

      {statuses.map((status, index) => (
        <div
          key={index}
          className="flex justify-between mb-2 text-sm"
        >
          <span>Node {index + 1}</span>

          <span
            className={`font-medium ${
              status === "online"
                ? "text-green-400"
                : "text-red-400"
            }`}
          >
            {status === "online" ? "● Online" : "● Offline"}
          </span>
        </div>
      ))}

      <div className="mt-4 text-xs text-gray-400">
        Threshold Required: 2 of 3
      </div>
    </div>
  );
}
