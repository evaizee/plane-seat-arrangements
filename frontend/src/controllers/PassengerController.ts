import { useState, useCallback } from 'react';
import { Passenger, PassengerFormData, PassengerModel } from '../models/PassengerModel';

export const usePassengerController = () => {
  const [numberOfPassengers, setNumberOfPassengers] = useState<number>(1);
  const [passengers, setPassengers] = useState<Passenger[]>([
    PassengerModel.createEmptyPassenger('passenger-1')
  ]);
  const [errors, setErrors] = useState<Record<string, string[]>>({});
  const [isFormValid, setIsFormValid] = useState(false);

  const updateNumberOfPassengers = useCallback((count: number) => {
    const validCount = Math.max(1, Math.min(10, count)); // Limit between 1-10 passengers
    setNumberOfPassengers(validCount);
    
    // Adjust passengers array
    const newPassengers = [...passengers];
    
    if (validCount > newPassengers.length) {
      // Add new passengers
      for (let i = newPassengers.length; i < validCount; i++) {
        newPassengers.push(PassengerModel.createEmptyPassenger(`passenger-${i + 1}`));
      }
    } else if (validCount < newPassengers.length) {
      // Remove excess passengers
      newPassengers.splice(validCount);
    }
    
    setPassengers(newPassengers);
    validateForm(newPassengers);
  }, [passengers]);

  const updatePassenger = useCallback((passengerId: string, updates: Partial<Passenger>) => {
    const updatedPassengers = passengers.map(passenger =>
      passenger.id === passengerId
        ? { ...passenger, ...updates }
        : passenger
    );
    
    setPassengers(updatedPassengers);
    validateForm(updatedPassengers);
  }, [passengers]);

  const validateForm = useCallback((passengersToValidate: Passenger[]) => {
    const validationErrors = PassengerModel.validateAllPassengers(passengersToValidate);
    setErrors(validationErrors);
    
    const hasErrors = Object.keys(validationErrors).length > 0;
    const allFieldsFilled = passengersToValidate.every(p => 
      p.name.trim() && p.identityNumber.trim()
    );
    
    setIsFormValid(!hasErrors && allFieldsFilled);
  }, []);

  const getPassengerFormData = useCallback((): PassengerFormData => {
    return {
      numberOfPassengers,
      passengers,
    };
  }, [numberOfPassengers, passengers]);

  const resetForm = useCallback(() => {
    setNumberOfPassengers(1);
    setPassengers([PassengerModel.createEmptyPassenger('passenger-1')]);
    setErrors({});
    setIsFormValid(false);
  }, []);

  return {
    numberOfPassengers,
    passengers,
    errors,
    isFormValid,
    updateNumberOfPassengers,
    updatePassenger,
    getPassengerFormData,
    resetForm,
  };
};