import React, { useState } from 'react';
import { useSeatController } from '../controllers/SeatController';
import { SeatRow } from './components/SeatRow';
import { SeatLegend } from './components/SeatLegend';
import { SelectionSummary } from './components/SelectionSummary';
import { ApiStatusIndicator } from './components/ApiStatusIndicator';
import { LoadingSpinner } from './components/LoadingSpinner';
import { ErrorDisplay } from './components/ErrorDisplay';
import { PassengerSummary } from './components/PassengerSummary';
import { FlightInfoModal } from './components/FlightInfoModal';
import { Plane, Info } from 'lucide-react';
import { SeatModel } from '../models/SeatModel';

interface SeatMapViewProps {
  flightId?: string;
}

export const SeatMapView: React.FC<SeatMapViewProps> = ({ flightId }) => {
  const [isFlightInfoOpen, setIsFlightInfoOpen] = useState(false);
  
  const {
    seatData,
    seatColumns,
    passengers,
    flightInfo,
    selectedSeats,
    hoveredSeat,
    loading,
    error,
    apiStatus,
    selectSeat,
    clearSelection,
    getTotalPrice,
    getCurrency,
    isSeatSelected,
    handleSeatHover,
    refreshData,
    canSelectMoreSeats,
    getRemainingSeats,
  } = useSeatController(flightId);

  if (loading) {
    return <LoadingSpinner />;
  }

  if (error) {
    return <ErrorDisplay error={error} onRetry={refreshData} loading={loading} />;
  }

  const selectedSeatCodes = selectedSeats.map(s => s.seat.code || '');
  const remainingSeats = getRemainingSeats();

  return (
    <div className="min-h-screen bg-gradient-to-br from-blue-50 to-indigo-100 p-4">
      <div className="max-w-6xl mx-auto">
        {/* Header */}
        <div className="text-center mb-6">
          <div className="flex items-center justify-center gap-2 mb-2">
            <Plane className="text-blue-600" size={32} />
            <h1 className="text-3xl font-bold text-gray-800">Select Your Seats</h1>
          </div>
          <p className="text-gray-600">Choose your preferred seats for a comfortable journey</p>
          
          {/* Flight Info Button */}
          {flightInfo && (
            <div className="flex items-center justify-center gap-4 mt-3">
              <p className="text-sm text-gray-500">
                Flight {SeatModel.formatFlightNumber(flightInfo)} • {flightInfo.origin} → {flightInfo.destination}
              </p>
              <button
                onClick={() => setIsFlightInfoOpen(true)}
                className="flex items-center gap-1 text-blue-600 hover:text-blue-700 text-sm font-medium hover:bg-blue-50 px-2 py-1 rounded-lg transition-colors"
              >
                <Info size={16} />
                Flight Details
              </button>
            </div>
          )}
          
          {remainingSeats !== null && (
            <p className="text-sm font-medium text-blue-600 mt-2">
              {remainingSeats > 0 
                ? `You can select ${remainingSeats} more seat${remainingSeats > 1 ? 's' : ''}`
                : 'All seats selected for your passengers'
              }
            </p>
          )}
        </div>

        {/* API Status */}
        <div className="flex justify-center mb-6">
          <ApiStatusIndicator
            apiStatus={apiStatus}
            onRefresh={refreshData}
            loading={loading}
          />
        </div>

        <div className="grid lg:grid-cols-3 gap-8">
          {/* Seat Map */}
          <div className="lg:col-span-2">
            <div className="bg-white rounded-xl shadow-lg overflow-hidden">
              {/* Aircraft nose */}
              <div className="bg-gradient-to-r from-blue-600 to-blue-700 text-white p-4 text-center">
                <div className="text-sm font-semibold">Aircraft Cabin</div>
                <div className="text-xs opacity-90 mt-1">Front of Aircraft</div>
                {flightInfo && (
                  <div className="text-xs opacity-75 mt-1">
                    {flightInfo.equipment} • {flightInfo.cabinClass}
                  </div>
                )}
              </div>
              
              {/* Seat Columns Header */}
              {seatColumns.length > 0 && (
                <div className="bg-gray-50 px-6 py-2 border-b">
                  <div className="flex items-center justify-center gap-1 md:gap-2">
                    <div className="w-8 text-center text-xs font-medium text-gray-500">Row</div>
                    <div className="flex items-center gap-1 md:gap-2">
                      {seatColumns.map((column, index) => (
                        <div key={index} className="w-10 md:w-12 text-center text-xs font-medium text-gray-500">
                          {column === 'LEFT_SIDE' || column === 'RIGHT_SIDE' ? '' : 
                           column === 'AISLE' ? '|' : column}
                        </div>
                      ))}
                    </div>
                    <div className="w-8 text-center text-xs font-medium text-gray-500">Row</div>
                  </div>
                </div>
              )}
              
              {/* Seat map */}
              <div className="p-6">
                {seatData.length === 0 ? (
                  <div className="text-center py-8 text-gray-500">
                    No seat data available
                  </div>
                ) : (
                  <div className="space-y-1">
                    {seatData.map((row, index) => (
                      <SeatRow
                        key={row.rowNumber}
                        row={row}
                        selectedSeats={selectedSeatCodes}
                        hoveredSeat={hoveredSeat}
                        onSeatSelect={selectSeat}
                        onSeatHover={handleSeatHover}
                        canSelectMore={canSelectMoreSeats()}
                      />
                    ))}
                  </div>
                )}
              </div>
              
              {/* Aircraft tail */}
              <div className="bg-gradient-to-r from-gray-600 to-gray-700 text-white p-2 text-center">
                <div className="text-xs opacity-90">Rear of Aircraft</div>
              </div>
            </div>
          </div>

          {/* Sidebar */}
          <div className="space-y-6">
            {passengers.length > 0 && <PassengerSummary passengers={passengers} />}
            <SeatLegend />
            <SelectionSummary
              selectedSeats={selectedSeats}
              totalPrice={getTotalPrice()}
              currency={getCurrency()}
              onClearSelection={clearSelection}
              passengers={passengers}
            />
          </div>
        </div>
      </div>

      {/* Flight Info Modal */}
      {flightInfo && (
        <FlightInfoModal
          isOpen={isFlightInfoOpen}
          onClose={() => setIsFlightInfoOpen(false)}
          flightInfo={flightInfo}
        />
      )}
    </div>
  );
};