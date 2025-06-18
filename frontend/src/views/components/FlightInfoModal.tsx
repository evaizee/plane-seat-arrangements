import React from 'react';
import { X, Plane, Clock, MapPin, Calendar } from 'lucide-react';
import { Segment, SeatModel } from '../../models/SeatModel';

interface FlightInfoModalProps {
  isOpen: boolean;
  onClose: () => void;
  flightInfo: Segment;
}

export const FlightInfoModal: React.FC<FlightInfoModalProps> = ({
  isOpen,
  onClose,
  flightInfo,
}) => {
  if (!isOpen) return null;

  const formatDuration = (hours: number) => {
    const h = Math.floor(hours);
    const m = Math.round((hours - h) * 60);
    return `${h}h ${m}m`;
  };

  return (
    <div className="fixed inset-0 bg-black bg-opacity-50 flex items-center justify-center z-50 p-4">
      <div className="bg-white rounded-xl shadow-2xl max-w-2xl w-full max-h-[90vh] overflow-y-auto">
        {/* Header */}
        <div className="flex items-center justify-between p-6 border-b border-gray-200">
          <div className="flex items-center gap-3">
            <Plane className="text-blue-600" size={24} />
            <h2 className="text-xl font-bold text-gray-800">Flight Details</h2>
          </div>
          <button
            onClick={onClose}
            className="p-2 hover:bg-gray-100 rounded-lg transition-colors"
          >
            <X size={20} className="text-gray-500" />
          </button>
        </div>

        {/* Flight Info */}
        <div className="p-6 space-y-6">
          {/* Flight Number & Aircraft */}
          <div className="bg-blue-50 rounded-lg p-4">
            <div className="flex items-center justify-between">
              <div>
                <h3 className="text-lg font-semibold text-blue-800">
                  {SeatModel.formatFlightNumber(flightInfo)}
                </h3>
                <p className="text-sm text-blue-600">
                  Operated by {flightInfo.flight.operatingAirlineCode} {flightInfo.flight.operatingFlightNumber}
                </p>
              </div>
              <div className="text-right">
                <p className="text-sm text-blue-600">Aircraft</p>
                <p className="font-semibold text-blue-800">{flightInfo.equipment}</p>
              </div>
            </div>
          </div>

          {/* Route Information */}
          <div className="grid md:grid-cols-3 gap-4">
            {/* Departure */}
            <div className="text-center">
              <div className="bg-green-50 rounded-lg p-4">
                <MapPin className="text-green-600 mx-auto mb-2" size={20} />
                <h4 className="font-semibold text-gray-800 text-lg">{flightInfo.origin}</h4>
                <p className="text-sm text-gray-600">{flightInfo.flight.departureTerminal}</p>
                <p className="text-lg font-bold text-green-600 mt-2">
                  {SeatModel.formatTime(flightInfo.departure)}
                </p>
                <p className="text-xs text-gray-500">
                  {SeatModel.formatDate(flightInfo.departure)}
                </p>
              </div>
            </div>

            {/* Duration */}
            <div className="flex items-center justify-center">
              <div className="text-center">
                <Clock className="text-gray-400 mx-auto mb-2" size={20} />
                <p className="text-sm text-gray-600">Duration</p>
                <p className="font-semibold text-gray-800">
                  {formatDuration(flightInfo.duration)}
                </p>
                <div className="w-16 h-0.5 bg-gray-300 mx-auto mt-2"></div>
              </div>
            </div>

            {/* Arrival */}
            <div className="text-center">
              <div className="bg-red-50 rounded-lg p-4">
                <MapPin className="text-red-600 mx-auto mb-2" size={20} />
                <h4 className="font-semibold text-gray-800 text-lg">{flightInfo.destination}</h4>
                <p className="text-sm text-gray-600">{flightInfo.flight.arrivalTerminal}</p>
                <p className="text-lg font-bold text-red-600 mt-2">
                  {SeatModel.formatTime(flightInfo.arrival)}
                </p>
                <p className="text-xs text-gray-500">
                  {SeatModel.formatDate(flightInfo.arrival)}
                </p>
              </div>
            </div>
          </div>

          {/* Additional Details */}
          <div className="grid md:grid-cols-2 gap-4">
            <div className="bg-gray-50 rounded-lg p-4">
              <h4 className="font-semibold text-gray-800 mb-3">Booking Details</h4>
              <div className="space-y-2 text-sm">
                <div className="flex justify-between">
                  <span className="text-gray-600">Cabin Class:</span>
                  <span className="font-medium">{flightInfo.cabinClass}</span>
                </div>
                <div className="flex justify-between">
                  <span className="text-gray-600">Booking Class:</span>
                  <span className="font-medium">{flightInfo.bookingClass}</span>
                </div>
                <div className="flex justify-between">
                  <span className="text-gray-600">Fare Basis:</span>
                  <span className="font-medium">{flightInfo.fareBasis}</span>
                </div>
              </div>
            </div>

            <div className="bg-gray-50 rounded-lg p-4">
              <h4 className="font-semibold text-gray-800 mb-3">Flight Information</h4>
              <div className="space-y-2 text-sm">
                <div className="flex justify-between">
                  <span className="text-gray-600">Miles:</span>
                  <span className="font-medium">{flightInfo.segmentOfferInformation?.flightsMiles || 'N/A'}</span>
                </div>
                <div className="flex justify-between">
                  <span className="text-gray-600">Layover:</span>
                  <span className="font-medium">
                    {flightInfo.layoverDuration > 0 ? `${flightInfo.layoverDuration}h` : 'None'}
                  </span>
                </div>
                <div className="flex justify-between">
                  <span className="text-gray-600">Stops:</span>
                  <span className="font-medium">
                    {flightInfo.flight.stopAirports.length > 0 
                      ? flightInfo.flight.stopAirports.join(', ')
                      : 'Non-stop'
                    }
                  </span>
                </div>
              </div>
            </div>
          </div>

          {/* Reference */}
          <div className="bg-yellow-50 rounded-lg p-4">
            <div className="flex items-center gap-2 mb-2">
              <Calendar className="text-yellow-600" size={16} />
              <h4 className="font-semibold text-yellow-800">Reference Information</h4>
            </div>
            <p className="text-sm text-yellow-700">
              Segment Reference: <span className="font-mono">{flightInfo.segmentRef}</span>
            </p>
            {flightInfo.subjectToGovernmentApproval && (
              <p className="text-xs text-yellow-600 mt-1">
                * Subject to government approval
              </p>
            )}
          </div>
        </div>

        {/* Footer */}
        <div className="border-t border-gray-200 p-6">
          <button
            onClick={onClose}
            className="w-full bg-blue-600 hover:bg-blue-700 text-white font-semibold py-2 px-4 rounded-lg transition-colors"
          >
            Close
          </button>
        </div>
      </div>
    </div>
  );
};