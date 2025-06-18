import React from 'react';
import { SeatSelection, BackendPassenger } from '../../models/SeatModel';
import { Trash2, AlertCircle } from 'lucide-react';

interface SelectionSummaryProps {
  selectedSeats: SeatSelection[];
  totalPrice: number;
  currency: string;
  onClearSelection: () => void;
  passengers: BackendPassenger[];
}

export const SelectionSummary: React.FC<SelectionSummaryProps> = ({
  selectedSeats,
  totalPrice,
  currency,
  onClearSelection,
  passengers,
}) => {
  const maxSeats = passengers.length;
  const isComplete = selectedSeats.length === maxSeats;
  const needsMoreSeats = selectedSeats.length < maxSeats;

  if (selectedSeats.length === 0) {
    return (
      <div className="bg-gray-50 rounded-lg border p-4 text-center">
        <p className="text-gray-500 text-sm">
          Select seats for your {passengers.length} passenger{passengers.length > 1 ? 's' : ''}
        </p>
        <p className="text-xs text-gray-400 mt-1">
          You need to select {maxSeats} seat{maxSeats > 1 ? 's' : ''}
        </p>
      </div>
    );
  }

  return (
    <div className="bg-white rounded-lg shadow-sm border p-4">
      <div className="flex items-center justify-between mb-3">
        <h3 className="text-sm font-semibold text-gray-800">Selected Seats</h3>
        <button
          onClick={onClearSelection}
          className="text-red-500 hover:text-red-700 p-1 rounded-md hover:bg-red-50 transition-colors"
          title="Clear selection"
        >
          <Trash2 size={16} />
        </button>
      </div>
      
      {/* Selection Status */}
      <div className="mb-3">
        {needsMoreSeats && (
          <div className="flex items-center gap-2 p-2 bg-amber-50 border border-amber-200 rounded-lg">
            <AlertCircle size={16} className="text-amber-600 flex-shrink-0" />
            <p className="text-xs text-amber-700">
              Select {maxSeats - selectedSeats.length} more seat{maxSeats - selectedSeats.length > 1 ? 's' : ''} for all passengers
            </p>
          </div>
        )}
        {isComplete && (
          <div className="flex items-center gap-2 p-2 bg-green-50 border border-green-200 rounded-lg">
            <div className="w-4 h-4 bg-green-500 rounded-full flex items-center justify-center flex-shrink-0">
              <span className="text-white text-xs">✓</span>
            </div>
            <p className="text-xs text-green-700">
              All seats selected for your passengers
            </p>
          </div>
        )}
      </div>
      
      <div className="space-y-2">
        {selectedSeats.map((selection, index) => (
          <div key={index} className="flex justify-between items-center text-sm">
            <span className="font-medium text-gray-800">
              Seat {selection.seat.code}
            </span>
            <span className="text-gray-600">
              {selection.seat.freeOfCharge 
                ? 'Free' 
                : `${currency} ${selection.seat.total?.alternatives?.[0]?.[0]?.amount || 0}`
              }
            </span>
          </div>
        ))}
        
        <div className="border-t pt-2 mt-3">
          <div className="flex justify-between items-center font-semibold text-lg">
            <span className="text-gray-800">Total</span>
            <span className="text-blue-600">{currency} {totalPrice.toFixed(2)}</span>
          </div>
        </div>
      </div>
      
      <button 
        className={`w-full mt-4 font-semibold py-2 px-4 rounded-lg transition-colors ${
          isComplete 
            ? 'bg-blue-600 hover:bg-blue-700 text-white' 
            : 'bg-gray-300 text-gray-500 cursor-not-allowed'
        }`}
        disabled={!isComplete}
      >
        {isComplete ? 'Continue to Payment' : `Select ${maxSeats - selectedSeats.length} more seat${maxSeats - selectedSeats.length > 1 ? 's' : ''}`}
      </button>
    </div>
  );
};