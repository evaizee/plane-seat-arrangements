export interface SeatPrice {
  amount: number;
  currency: string;
}

export interface SeatPricing {
  alternatives: SeatPrice[][];
}

export interface Seat {
  slotCharacteristics?: string[];
  storefrontSlotCode: 'BLANK' | 'AISLE' | 'SEAT' | 'BULKHEAD';
  available: boolean;
  entitled: boolean;
  feeWaived: boolean;
  freeOfCharge: boolean;
  originallySelected: boolean;
  code?: string;
  designations?: string[];
  entitledRuleId?: string;
  feeWaivedRuleId?: string;
  seatCharacteristics?: string[];
  limitations?: string[];
  refundIndicator?: string;
  prices?: SeatPricing;
  taxes?: SeatPricing;
  total?: SeatPricing;
  rawSeatCharacteristics?: string[];
}

export interface SeatRow {
  rowNumber: number;
  seatCodes: string[];
  seats: Seat[];
}

export interface SeatSelection {
  seat: Seat;
  rowNumber: number;
}

// New interfaces for the nested backend structure
export interface PassengerDetails {
  firstName: string;
  lastName: string;
}

export interface PassengerInfo {
  dateOfBirth: string;
  gender: string;
  type: string;
  emails: string[];
  phones: string[];
  address: {
    street1: string;
    street2: string;
    postcode: string;
    state: string;
    city: string;
    country: string;
    addressType: string;
  };
}

export interface DocumentInfo {
  issuingCountry: string;
  countryOfBirth: string;
  documentType: string;
  nationality: string;
}

export interface BackendPassenger {
  passengerIndex: number;
  passengerNameNumber: string;
  passengerDetails: PassengerDetails;
  passengerInfo: PassengerInfo;
  documentInfo: DocumentInfo;
  preferences: any;
}

export interface Flight {
  flightNumber: number;
  operatingFlightNumber: number;
  airlineCode: string;
  operatingAirlineCode: string;
  stopAirports: string[];
  departureTerminal: string;
  arrivalTerminal: string;
}

export interface Segment {
  duration: number;
  cabinClass: string;
  equipment: string;
  flight: Flight;
  origin: string;
  destination: string;
  departure: string;
  arrival: string;
  bookingClass: string;
  layoverDuration: number;
  fareBasis: string;
  subjectToGovernmentApproval: boolean;
  segmentRef: string;
}

export interface Cabin {
  deck: string;
  seatColumns: string[];
  seatRows: SeatRow[];
  firstRow: number;
  lastRow: number;
}

export interface SeatMap {
  rowsDisabledCauses: string[];
  aircraft: string;
  cabins: Cabin[];
}

export interface PassengerSeatMap {
  seatSelectionEnabledForPax: boolean;
  seatMap: SeatMap;
  passenger: BackendPassenger;
}

export interface SegmentSeatMap {
  passengerSeatMaps: PassengerSeatMap[];
  segment: Segment;
}

export interface SeatsItineraryPart {
  segmentSeatMaps: SegmentSeatMap[];
}

export interface BackendSeatResponse {
  seatsItineraryParts: SeatsItineraryPart[];
  selectedSeats: any[];
}

export class SeatModel {
  static getSeatType(seat: Seat): 'window' | 'middle' | 'aisle' | 'blank' | 'aisle-space' | 'bulkhead' {
    if (seat.storefrontSlotCode === 'BLANK') return 'blank';
    if (seat.storefrontSlotCode === 'AISLE') return 'aisle-space';
    if (seat.storefrontSlotCode === 'BULKHEAD') return 'bulkhead';
    
    const characteristics = seat.seatCharacteristics || [];
    if (characteristics.includes('W')) return 'window';
    if (characteristics.includes('A')) return 'aisle';
    return 'middle';
  }

  static getSeatPrice(seat: Seat): number {
    return seat.total?.alternatives?.[0]?.[0]?.amount || 0;
  }

  static getSeatCurrency(seat: Seat): string {
    return seat.total?.alternatives?.[0]?.[0]?.currency || 'MYR';
  }

  static isSeatSelectable(seat: Seat): boolean {
    return seat.storefrontSlotCode === 'SEAT' && seat.available;
  }

  static getSeatLabel(seat: Seat): string {
    const type = this.getSeatType(seat);
    
    if (type === 'window') return 'Window';
    if (type === 'aisle') return 'Aisle';
    if (type === 'middle') return 'Middle';
    if (type === 'bulkhead') return 'Bulkhead';
    return '';
  }

  static formatPassengerName(passenger: BackendPassenger): string {
    return `${passenger.passengerDetails.firstName} ${passenger.passengerDetails.lastName}`;
  }

  static formatFlightNumber(segment: Segment): string {
    return `${segment.flight.airlineCode} ${segment.flight.flightNumber}`;
  }

  static formatDateTime(dateTimeString: string): string {
    const date = new Date(dateTimeString);
    return date.toLocaleString('en-US', {
      year: 'numeric',
      month: 'short',
      day: 'numeric',
      hour: '2-digit',
      minute: '2-digit',
      hour12: true
    });
  }

  static formatTime(dateTimeString: string): string {
    const date = new Date(dateTimeString);
    return date.toLocaleString('en-US', {
      hour: '2-digit',
      minute: '2-digit',
      hour12: true
    });
  }

  static formatDate(dateTimeString: string): string {
    const date = new Date(dateTimeString);
    return date.toLocaleString('en-US', {
      year: 'numeric',
      month: 'short',
      day: 'numeric'
    });
  }
}