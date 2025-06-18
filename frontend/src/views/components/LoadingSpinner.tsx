import React from 'react';
import { Plane } from 'lucide-react';

export const LoadingSpinner: React.FC = () => {
  return (
    <div className="min-h-screen bg-gradient-to-br from-blue-50 to-indigo-100 flex items-center justify-center">
      <div className="text-center">
        <div className="relative">
          <Plane className="text-blue-600 animate-pulse mx-auto mb-4" size={48} />
          <div className="absolute -top-2 -right-2 w-4 h-4 bg-blue-600 rounded-full animate-ping"></div>
        </div>
        <h2 className="text-xl font-semibold text-gray-800 mb-2">Loading Seat Map</h2>
        <p className="text-gray-600">Please wait while we fetch the latest seat availability...</p>
      </div>
    </div>
  );
};