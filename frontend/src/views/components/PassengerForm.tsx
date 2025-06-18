import React from 'react';
import { Passenger, PassengerModel } from '../../models/PassengerModel';
import { User, CreditCard } from 'lucide-react';

interface PassengerFormProps {
  passenger: Passenger;
  index: number;
  errors: string[];
  onUpdate: (passengerId: string, updates: Partial<Passenger>) => void;
}

export const PassengerForm: React.FC<PassengerFormProps> = ({
  passenger,
  index,
  errors,
  onUpdate,
}) => {
  const identityTypes: Array<{ value: Passenger['identityType']; label: string }> = [
    { value: 'passport', label: 'Passport' },
    { value: 'national_id', label: 'National ID' },
    { value: 'driving_license', label: 'Driving License' },
  ];

  return (
    <div className="bg-white rounded-lg border border-gray-200 p-6 shadow-sm">
      <div className="flex items-center gap-2 mb-4">
        <User className="text-blue-600" size={20} />
        <h3 className="text-lg font-semibold text-gray-800">
          Passenger {index + 1}
        </h3>
      </div>

      <div className="grid md:grid-cols-2 gap-4">
        {/* Name Field */}
        <div>
          <label htmlFor={`name-${passenger.id}`} className="block text-sm font-medium text-gray-700 mb-2">
            Full Name *
          </label>
          <input
            type="text"
            id={`name-${passenger.id}`}
            value={passenger.name}
            onChange={(e) => onUpdate(passenger.id, { name: e.target.value })}
            className={`w-full px-3 py-2 border rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-blue-500 transition-colors ${
              errors.some(e => e.includes('Name')) 
                ? 'border-red-300 bg-red-50' 
                : 'border-gray-300'
            }`}
            placeholder="Enter full name as per ID"
          />
        </div>

        {/* Identity Type Field */}
        <div>
          <label htmlFor={`identity-type-${passenger.id}`} className="block text-sm font-medium text-gray-700 mb-2">
            Identity Type *
          </label>
          <div className="relative">
            <CreditCard className="absolute left-3 top-1/2 transform -translate-y-1/2 text-gray-400" size={16} />
            <select
              id={`identity-type-${passenger.id}`}
              value={passenger.identityType}
              onChange={(e) => onUpdate(passenger.id, { identityType: e.target.value as Passenger['identityType'] })}
              className="w-full pl-10 pr-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-blue-500 transition-colors appearance-none bg-white"
            >
              {identityTypes.map(type => (
                <option key={type.value} value={type.value}>
                  {type.label}
                </option>
              ))}
            </select>
          </div>
        </div>

        {/* Identity Number Field */}
        <div className="md:col-span-2">
          <label htmlFor={`identity-number-${passenger.id}`} className="block text-sm font-medium text-gray-700 mb-2">
            {PassengerModel.getIdentityTypeLabel(passenger.identityType)} Number *
          </label>
          <input
            type="text"
            id={`identity-number-${passenger.id}`}
            value={passenger.identityNumber}
            onChange={(e) => onUpdate(passenger.id, { identityNumber: e.target.value })}
            className={`w-full px-3 py-2 border rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-blue-500 transition-colors ${
              errors.some(e => e.includes('Identity number')) 
                ? 'border-red-300 bg-red-50' 
                : 'border-gray-300'
            }`}
            placeholder={`Enter ${PassengerModel.getIdentityTypeLabel(passenger.identityType).toLowerCase()} number`}
          />
        </div>
      </div>

      {/* Error Messages */}
      {errors.length > 0 && (
        <div className="mt-4 p-3 bg-red-50 border border-red-200 rounded-lg">
          <ul className="text-sm text-red-600 space-y-1">
            {errors.map((error, index) => (
              <li key={index} className="flex items-center gap-1">
                <span className="w-1 h-1 bg-red-600 rounded-full"></span>
                {error}
              </li>
            ))}
          </ul>
        </div>
      )}
    </div>
  );
};