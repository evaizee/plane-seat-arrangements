import React from 'react';
import { SeatRow as SeatRowType, Seat } from '../../models/SeatModel';
import { SeatButton } from './SeatButton';

interface SeatRowProps {
  row: SeatRowType;
  selectedSeats: string[];
  hoveredSeat: string | null;
  onSeatSelect: (seat: Seat, rowNumber: number) => void;
  onSeatHover: (seatCode: string | null) => void;
  canSelectMore?: boolean;
}

export const SeatRow: React.FC<SeatRowProps> = ({
  row,
  selectedSeats,
  hoveredSeat,
  onSeatSelect,
  onSeatHover,
  canSelectMore = true,
}) => {
  return (
    <div className="flex items-center justify-center gap-1 md:gap-2 py-2">
      {/* Row number */}
      <div className="w-8 text-center text-sm font-semibold text-gray-600">
        {row.rowNumber}
      </div>
      
      {/* Seats */}
      <div className="flex items-center gap-1 md:gap-2">
        {row.seats.map((seat, index) => (
          <SeatButton
            key={`${row.rowNumber}-${index}`}
            seat={seat}
            rowNumber={row.rowNumber}
            isSelected={seat.code ? selectedSeats.includes(seat.code) : false}
            isHovered={seat.code === hoveredSeat}
            onSelect={onSeatSelect}
            onHover={onSeatHover}
            canSelectMore={canSelectMore}
          />
        ))}
      </div>
      
      {/* Row number (right side) */}
      <div className="w-8 text-center text-sm font-semibold text-gray-600">
        {row.rowNumber}
      </div>
    </div>
  );
};