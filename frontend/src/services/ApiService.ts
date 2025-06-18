import { BackendSeatResponse } from '../models/SeatModel';

export interface ApiResponse<T> {
  data: T;
  success: boolean;
  message?: string;
}

export class ApiService {
  private static baseUrl = import.meta.env.VITE_API_BASE_URL;
  private static apiKey = import.meta.env.VITE_API_KEY;

  private static getHeaders(): HeadersInit {
    return {
      'Content-Type': 'application/json',
      'Authorization': `Bearer ${this.apiKey}`,
      'X-API-Key': this.apiKey,
    };
  }

  static async getSeatMap(flightId: string): Promise<BackendSeatResponse> {
    try {
      const response = await fetch(
        `${this.baseUrl}/v1/flights/${flightId}/seat-map`,
        {
          method: 'GET',
          headers: this.getHeaders(),
        }
      );

      if (!response.ok) {
        throw new Error(`HTTP error! status: ${response.status}`);
      }

      const data = await response.json();
      
      // Handle different response formats
      if (data.seatsItineraryParts) {
        return data as BackendSeatResponse;
      } else if (data.data && data.data.seatsItineraryParts) {
        return data.data as BackendSeatResponse;
      }
      
      throw new Error('Invalid response format - missing seatsItineraryParts');
    } catch (error) {
      console.error('Error fetching seat map:', error);
      throw error;
    }
  }

  static async validateApiConnection(): Promise<boolean> {
    try {
      const response = await fetch(`${this.baseUrl}/health`, {
        method: 'GET',
        headers: this.getHeaders(),
      });
      return response.ok;
    } catch {
      return false;
    }
  }
}