export default function SignPanel({ onSign, loading }) {
  return (
    <div className="bg-gray-900 p-6 rounded-xl shadow-lg">
      <h2 className="text-xl mb-4 font-semibold">
        Signing Console
      </h2>

      <button
        onClick={onSign}
        disabled={loading}
        className="bg-blue-600 hover:bg-blue-700 px-4 py-2 rounded-lg transition"
      >
        {loading ? "Signing..." : "Sign Transaction"}
      </button>

      <div className="mt-4 text-xs text-gray-400">
        Real Threshold ECDSA (GG18)
      </div>
    </div>
  );
}
