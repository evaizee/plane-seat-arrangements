import { BackendSeatResponse, SeatRow, Cabin } from '../models/SeatModel';
import { ApiService } from './ApiService';
import { mockBackendResponse } from '../utils/mockData';

export type DataSource = 'mock' | 'api';

export class SeatDataService {
  private static getDataSource(): DataSource {
    const source = import.meta.env.VITE_DATA_SOURCE as DataSource;
    return source === 'api' ? 'api' : 'mock';
  }

  static async getSeatData(flightId?: string): Promise<BackendSeatResponse> {
    const dataSource = this.getDataSource();
    
    console.log(`Using ${dataSource} data source`);
    
    if (dataSource === 'mock') {
      // Simulate API delay for realistic testing
      await new Promise(resolve => setTimeout(resolve, 500));
      return mockBackendResponse;
    }
    
    if (!flightId) {
      throw new Error('Flight ID is required when using API data source');
    }
    
    try {
      return await ApiService.getSeatMap(flightId);
    } catch (error) {
      console.warn('API call failed, falling back to mock data:', error);
      return mockBackendResponse;
    }
  }

  static async checkApiStatus(): Promise<{
    isConnected: boolean;
    dataSource: DataSource;
    apiConfigured: boolean;
  }> {
    const dataSource = this.getDataSource();
    const apiConfigured = !!(import.meta.env.VITE_API_BASE_URL && import.meta.env.VITE_API_KEY);
    
    if (dataSource === 'mock') {
      return {
        isConnected: false,
        dataSource: 'mock',
        apiConfigured,
      };
    }
    
    const isConnected = await ApiService.validateApiConnection();
    
    return {
      isConnected,
      dataSource: 'api',
      apiConfigured,
    };
  }

  static extractSeatRows(backendResponse: BackendSeatResponse): SeatRow[] {
    try {
      const firstItinerary = backendResponse.seatsItineraryParts[0];
      const firstSegment = firstItinerary?.segmentSeatMaps[0];
      const firstPassengerMap = firstSegment?.passengerSeatMaps[0];
      const seatMap = firstPassengerMap?.seatMap;
      const firstCabin = seatMap?.cabins[0];
      
      return firstCabin?.seatRows || [];
    } catch (error) {
      console.error('Error extracting seat rows:', error);
      return [];
    }
  }

  static extractSeatColumns(backendResponse: BackendSeatResponse): string[] {
    try {
      const firstItinerary = backendResponse.seatsItineraryParts[0];
      const firstSegment = firstItinerary?.segmentSeatMaps[0];
      const firstPassengerMap = firstSegment?.passengerSeatMaps[0];
      const seatMap = firstPassengerMap?.seatMap;
      const firstCabin = seatMap?.cabins[0];
      
      return firstCabin?.seatColumns || [];
    } catch (error) {
      console.error('Error extracting seat columns:', error);
      return [];
    }
  }
}