export interface Passenger {
  id: string;
  name: string;
  identityNumber: string;
  identityType: 'passport' | 'national_id' | 'driving_license';
}

export interface PassengerFormData {
  numberOfPassengers: number;
  passengers: Passenger[];
}

export class PassengerModel {
  static createEmptyPassenger(id: string): Passenger {
    return {
      id,
      name: '',
      identityNumber: '',
      identityType: 'passport',
    };
  }

  static validatePassenger(passenger: Passenger): string[] {
    const errors: string[] = [];
    
    if (!passenger.name.trim()) {
      errors.push('Name is required');
    } else if (passenger.name.trim().length < 2) {
      errors.push('Name must be at least 2 characters');
    }
    
    if (!passenger.identityNumber.trim()) {
      errors.push('Identity number is required');
    } else if (passenger.identityNumber.trim().length < 5) {
      errors.push('Identity number must be at least 5 characters');
    }
    
    return errors;
  }

  static validateAllPassengers(passengers: Passenger[]): Record<string, string[]> {
    const errors: Record<string, string[]> = {};
    
    passengers.forEach(passenger => {
      const passengerErrors = this.validatePassenger(passenger);
      if (passengerErrors.length > 0) {
        errors[passenger.id] = passengerErrors;
      }
    });
    
    return errors;
  }

  static getIdentityTypeLabel(type: Passenger['identityType']): string {
    switch (type) {
      case 'passport':
        return 'Passport';
      case 'national_id':
        return 'National ID';
      case 'driving_license':
        return 'Driving License';
      default:
        return 'Unknown';
    }
  }
}