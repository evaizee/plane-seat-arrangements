import React from 'react';
import { AlertTriangle, RefreshCw } from 'lucide-react';

interface ErrorDisplayProps {
  error: string;
  onRetry: () => void;
  loading: boolean;
}

export const ErrorDisplay: React.FC<ErrorDisplayProps> = ({
  error,
  onRetry,
  loading,
}) => {
  return (
    <div className="min-h-screen bg-gradient-to-br from-blue-50 to-indigo-100 flex items-center justify-center p-4">
      <div className="bg-white rounded-xl shadow-lg p-8 max-w-md w-full text-center">
        <AlertTriangle className="text-red-500 mx-auto mb-4" size={48} />
        <h2 className="text-xl font-semibold text-gray-800 mb-2">Unable to Load Seat Map</h2>
        <p className="text-gray-600 mb-6">{error}</p>
        <button
          onClick={onRetry}
          disabled={loading}
          className="bg-blue-600 hover:bg-blue-700 disabled:opacity-50 text-white font-semibold py-2 px-6 rounded-lg transition-colors flex items-center gap-2 mx-auto"
        >
          <RefreshCw size={16} className={loading ? 'animate-spin' : ''} />
          Try Again
        </button>
      </div>
    </div>
  );
};