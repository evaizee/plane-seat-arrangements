import React from 'react';

export const SeatLegend: React.FC = () => {
  const legendItems = [
    { icon: '🪟', label: 'Window', color: 'bg-green-50 border-green-300 text-green-800' },
    { icon: '🚶', label: 'Aisle', color: 'bg-blue-50 border-blue-300 text-blue-800' },
    { icon: '💺', label: 'Middle', color: 'bg-gray-50 border-gray-300 text-gray-800' },
    { icon: '🆓', label: 'Free', color: 'bg-emerald-50 border-emerald-300 text-emerald-800' },
    { icon: '✓', label: 'Selected', color: 'bg-blue-500 border-blue-600 text-white' },
    { icon: '✗', label: 'Occupied', color: 'bg-gray-300 border-gray-400 text-gray-500' },
  ];

  return (
    <div className="bg-white rounded-lg shadow-sm border p-4">
      <h3 className="text-sm font-semibold text-gray-800 mb-3">Seat Legend</h3>
      <div className="grid grid-cols-2 md:grid-cols-3 gap-2">
        {legendItems.map((item, index) => (
          <div key={index} className="flex items-center gap-2">
            <div className={`w-6 h-6 rounded border-2 flex items-center justify-center text-xs ${item.color}`}>
              {item.icon}
            </div>
            <span className="text-xs text-gray-600">{item.label}</span>
          </div>
        ))}
      </div>
    </div>
  );
};