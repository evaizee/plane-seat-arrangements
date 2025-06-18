import React from 'react';
import { usePassengerController } from '../controllers/PassengerController';
import { PassengerForm } from './components/PassengerForm';
import { Users, ArrowRight, Plane } from 'lucide-react';
import { PassengerFormData } from '../models/PassengerModel';

interface PassengerInfoViewProps {
  onContinue: (passengerData: PassengerFormData) => void;
}

export const PassengerInfoView: React.FC<PassengerInfoViewProps> = ({ onContinue }) => {
  const {
    numberOfPassengers,
    passengers,
    errors,
    isFormValid,
    updateNumberOfPassengers,
    updatePassenger,
    getPassengerFormData,
  } = usePassengerController();

  const handleContinue = () => {
    if (isFormValid) {
      onContinue(getPassengerFormData());
    }
  };

  return (
    <div className="min-h-screen bg-gradient-to-br from-blue-50 to-indigo-100 p-4">
      <div className="max-w-4xl mx-auto">
        {/* Header */}
        <div className="text-center mb-8">
          <div className="flex items-center justify-center gap-2 mb-2">
            <Plane className="text-blue-600" size={32} />
            <h1 className="text-3xl font-bold text-gray-800">Passenger Information</h1>
          </div>
          <p className="text-gray-600">Please provide details for all passengers</p>
        </div>

        {/* Number of Passengers */}
        <div className="bg-white rounded-xl shadow-lg p-6 mb-8">
          <div className="flex items-center gap-2 mb-4">
            <Users className="text-blue-600" size={24} />
            <h2 className="text-xl font-semibold text-gray-800">Number of Passengers</h2>
          </div>
          
          <div className="flex items-center gap-4">
            <label htmlFor="passenger-count" className="text-sm font-medium text-gray-700">
              How many passengers will be traveling?
            </label>
            <div className="flex items-center gap-2">
              <button
                onClick={() => updateNumberOfPassengers(numberOfPassengers - 1)}
                disabled={numberOfPassengers <= 1}
                className="w-8 h-8 rounded-full border border-gray-300 flex items-center justify-center hover:bg-gray-50 disabled:opacity-50 disabled:cursor-not-allowed transition-colors"
              >
                -
              </button>
              <input
                type="number"
                id="passenger-count"
                min="1"
                max="10"
                value={numberOfPassengers}
                onChange={(e) => updateNumberOfPassengers(parseInt(e.target.value) || 1)}
                className="w-16 text-center py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-blue-500"
              />
              <button
                onClick={() => updateNumberOfPassengers(numberOfPassengers + 1)}
                disabled={numberOfPassengers >= 10}
                className="w-8 h-8 rounded-full border border-gray-300 flex items-center justify-center hover:bg-gray-50 disabled:opacity-50 disabled:cursor-not-allowed transition-colors"
              >
                +
              </button>
            </div>
            <span className="text-sm text-gray-500">(Maximum 10 passengers)</span>
          </div>
        </div>

        {/* Passenger Forms */}
        <div className="space-y-6 mb-8">
          {passengers.map((passenger, index) => (
            <PassengerForm
              key={passenger.id}
              passenger={passenger}
              index={index}
              errors={errors[passenger.id] || []}
              onUpdate={updatePassenger}
            />
          ))}
        </div>

        {/* Continue Button */}
        <div className="bg-white rounded-xl shadow-lg p-6">
          <div className="flex items-center justify-between">
            <div>
              <h3 className="text-lg font-semibold text-gray-800">Ready to select seats?</h3>
              <p className="text-sm text-gray-600 mt-1">
                You'll be able to select up to {numberOfPassengers} seat{numberOfPassengers > 1 ? 's' : ''}
              </p>
            </div>
            <button
              onClick={handleContinue}
              disabled={!isFormValid}
              className="bg-blue-600 hover:bg-blue-700 disabled:bg-gray-300 disabled:cursor-not-allowed text-white font-semibold py-3 px-6 rounded-lg transition-colors flex items-center gap-2"
            >
              Continue to Seat Selection
              <ArrowRight size={20} />
            </button>
          </div>
          
          {!isFormValid && (
            <div className="mt-4 p-3 bg-amber-50 border border-amber-200 rounded-lg">
              <p className="text-sm text-amber-700">
                Please fill in all required fields for all passengers before continuing.
              </p>
            </div>
          )}
        </div>
      </div>
    </div>
  );
};