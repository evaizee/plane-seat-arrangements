import { useState, useCallback, useEffect } from 'react';
import { Seat, SeatRow, SeatSelection, SeatModel, BackendSeatResponse, BackendPassenger, Segment } from '../models/SeatModel';
import { SeatDataService } from '../services/SeatDataService';

export const useSeatController = (flightId?: string) => {
  const [backendData, setBackendData] = useState<BackendSeatResponse | null>(null);
  const [seatData, setSeatData] = useState<SeatRow[]>([]);
  const [seatColumns, setSeatColumns] = useState<string[]>([]);
  const [passengers, setPassengers] = useState<BackendPassenger[]>([]);
  const [flightInfo, setFlightInfo] = useState<Segment | null>(null);
  const [selectedSeats, setSelectedSeats] = useState<SeatSelection[]>([]);
  const [hoveredSeat, setHoveredSeat] = useState<string | null>(null);
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState<string | null>(null);
  const [apiStatus, setApiStatus] = useState<{
    isConnected: boolean;
    dataSource: 'mock' | 'api';
    apiConfigured: boolean;
  } | null>(null);

  // Load seat data
  const loadSeatData = useCallback(async () => {
    try {
      setLoading(true);
      setError(null);
      
      const [data, status] = await Promise.all([
        SeatDataService.getSeatData(flightId),
        SeatDataService.checkApiStatus()
      ]);
      
      setBackendData(data);
      
      // Extract seat rows and columns
      const extractedSeatRows = SeatDataService.extractSeatRows(data);
      const extractedSeatColumns = SeatDataService.extractSeatColumns(data);
      
      setSeatData(extractedSeatRows);
      setSeatColumns(extractedSeatColumns);
      
      // Extract passengers and flight info
      const firstItinerary = data.seatsItineraryParts[0];
      const firstSegment = firstItinerary?.segmentSeatMaps[0];
      
      if (firstSegment) {
        const extractedPassengers = firstSegment.passengerSeatMaps.map(psm => psm.passenger);
        setPassengers(extractedPassengers);
        setFlightInfo(firstSegment.segment);
      }
      
      setApiStatus(status);
    } catch (err) {
      setError(err instanceof Error ? err.message : 'Failed to load seat data');
      console.error('Error loading seat data:', err);
    } finally {
      setLoading(false);
    }
  }, [flightId]);

  // Load data on mount
  useEffect(() => {
    loadSeatData();
  }, [loadSeatData]);

  const selectSeat = useCallback((seat: Seat, rowNumber: number) => {
    if (!SeatModel.isSeatSelectable(seat) || !seat.code) return;

    setSelectedSeats(prev => {
      const existing = prev.find(s => s.seat.code === seat.code);
      if (existing) {
        // Deselect if already selected
        return prev.filter(s => s.seat.code !== seat.code);
      } else {
        // Check if we've reached the maximum number of seats (based on passengers)
        const maxSeats = passengers.length;
        if (maxSeats && prev.length >= maxSeats) {
          alert(`You can only select up to ${maxSeats} seat${maxSeats > 1 ? 's' : ''} based on the number of passengers.`);
          return prev;
        }
        // Select new seat
        return [...prev, { seat, rowNumber }];
      }
    });
  }, [passengers.length]);

  const clearSelection = useCallback(() => {
    setSelectedSeats([]);
  }, []);

  const getTotalPrice = useCallback(() => {
    return selectedSeats.reduce((total, selection) => {
      return total + SeatModel.getSeatPrice(selection.seat);
    }, 0);
  }, [selectedSeats]);

  const getCurrency = useCallback(() => {
    return selectedSeats.length > 0 
      ? SeatModel.getSeatCurrency(selectedSeats[0].seat)
      : 'MYR';
  }, [selectedSeats]);

  const isSeatSelected = useCallback((seatCode: string) => {
    return selectedSeats.some(s => s.seat.code === seatCode);
  }, [selectedSeats]);

  const handleSeatHover = useCallback((seatCode: string | null) => {
    setHoveredSeat(seatCode);
  }, []);

  const refreshData = useCallback(() => {
    loadSeatData();
  }, [loadSeatData]);

  const canSelectMoreSeats = useCallback(() => {
    const maxSeats = passengers.length;
    if (!maxSeats) return true;
    return selectedSeats.length < maxSeats;
  }, [selectedSeats.length, passengers.length]);

  const getRemainingSeats = useCallback(() => {
    const maxSeats = passengers.length;
    if (!maxSeats) return null;
    return maxSeats - selectedSeats.length;
  }, [selectedSeats.length, passengers.length]);

  return {
    backendData,
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
    maxSeats: passengers.length,
  };
};