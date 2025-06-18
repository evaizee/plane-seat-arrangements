import React from 'react';
import { Users, User, Mail, MapPin } from 'lucide-react';
import { BackendPassenger, SeatModel } from '../../models/SeatModel';

interface PassengerSummaryProps {
  passengers: BackendPassenger[];
}

export const PassengerSummary: React.FC<PassengerSummaryProps> = ({ passengers }) => {
  return (
    <div className="bg-white rounded-lg shadow-sm border p-4">
      <div className="flex items-center gap-2 mb-3">
        <Users className="text-blue-600" size={20} />
        <h3 className="text-sm font-semibold text-gray-800">Passengers</h3>
        <span className="bg-blue-100 text-blue-800 text-xs font-medium px-2 py-1 rounded-full">
          {passengers.length}
        </span>
      </div>
      
      <div className="space-y-3 max-h-64 overflow-y-auto">
        {passengers.map((passenger, index) => (
          <div key={passenger.passengerIndex} className="bg-gray-50 rounded-lg p-3">
            <div className="flex items-start gap-2">
              <User size={16} className="text-gray-500 flex-shrink-0 mt-0.5" />
              <div className="min-w-0 flex-1">
                <p className="text-sm font-medium text-gray-800">
                  {SeatModel.formatPassengerName(passenger)}
                </p>
                <p className="text-xs text-gray-500 mb-2">
                  {passenger.passengerNameNumber} • {passenger.passengerInfo.type}
                </p>
                
                {/* Contact Info */}
                {passenger.passengerInfo.emails.length > 0 && (
                  <div className="flex items-center gap-1 mb-1">
                    <Mail size={12} className="text-gray-400" />
                    <p className="text-xs text-gray-600 truncate">
                      {passenger.passengerInfo.emails[0]}
                    </p>
                  </div>
                )}
                
                {/* Address */}
                <div className="flex items-center gap-1">
                  <MapPin size={12} className="text-gray-400" />
                  <p className="text-xs text-gray-600 truncate">
                    {passenger.passengerInfo.address.city}, {passenger.passengerInfo.address.country}
                  </p>
                </div>
                
                {/* Document Info */}
                <div className="mt-2 pt-2 border-t border-gray-200">
                  <p className="text-xs text-gray-500">
                    {passenger.documentInfo.documentType === 'P' ? 'Passport' : 
                     passenger.documentInfo.documentType === 'F' ? 'ID Card' : 'Document'} • 
                    {passenger.documentInfo.nationality}
                  </p>
                </div>
              </div>
            </div>
          </div>
        ))}
      </div>
    </div>
  );
};