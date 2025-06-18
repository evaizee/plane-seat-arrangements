import React from 'react';
import { Seat, SeatModel } from '../../models/SeatModel';

interface SeatButtonProps {
  seat: Seat;
  rowNumber: number;
  isSelected: boolean;
  isHovered: boolean;
  onSelect: (seat: Seat, rowNumber: number) => void;
  onHover: (seatCode: string | null) => void;
  canSelectMore?: boolean;
}

export const SeatButton: React.FC<SeatButtonProps> = ({
  seat,
  rowNumber,
  isSelected,
  isHovered,
  onSelect,
  onHover,
  canSelectMore = true,
}) => {
  if (seat.storefrontSlotCode === 'BLANK') {
    return <div className="w-10 h-10 md:w-12 md:h-12" />;
  }

  if (seat.storefrontSlotCode === 'AISLE') {
    return (
      <div className="w-6 md:w-8 h-10 md:h-12 flex items-center justify-center">
        <div className="w-1 h-6 bg-gray-300 rounded-full" />
      </div>
    );
  }

  if (seat.storefrontSlotCode === 'BULKHEAD') {
    return (
      <div className="w-10 h-10 md:w-12 md:h-12 flex items-center justify-center">
        <div className="w-8 h-2 bg-gray-400 rounded-sm" />
      </div>
    );
  }

  const seatType = SeatModel.getSeatType(seat);
  const price = SeatModel.getSeatPrice(seat);
  const isSelectable = SeatModel.isSeatSelectable(seat) && (canSelectMore || isSelected);
  const isFree = seat.freeOfCharge;
  
  const getButtonStyles = () => {
    const baseStyles = "w-10 h-10 md:w-12 md:h-12 rounded-lg border-2 flex items-center justify-center text-xs md:text-sm font-semibold transition-all duration-200 relative";
    
    if (!isSelectable) {
      const disabledReason = !SeatModel.isSeatSelectable(seat) ? 'occupied' : 'limit-reached';
      return `${baseStyles} ${
        disabledReason === 'occupied' 
          ? 'bg-gray-300 border-gray-400 text-gray-500 cursor-not-allowed'
          : 'bg-orange-100 border-orange-300 text-orange-600 cursor-not-allowed opacity-60'
      }`;
    }
    
    if (isSelected) {
      return `${baseStyles} bg-blue-500 border-blue-600 text-white shadow-lg transform scale-105`;
    }
    
    if (isHovered) {
      return `${baseStyles} bg-blue-100 border-blue-300 text-blue-800 shadow-md transform scale-105`;
    }
    
    // Free seats get a special green styling
    if (isFree) {
      return `${baseStyles} bg-emerald-50 border-emerald-300 text-emerald-800 hover:bg-emerald-100 hover:border-emerald-400 cursor-pointer`;
    }
    
    switch (seatType) {
      case 'window':
        return `${baseStyles} bg-green-50 border-green-300 text-green-800 hover:bg-green-100 hover:border-green-400 cursor-pointer`;
      case 'aisle':
        return `${baseStyles} bg-blue-50 border-blue-300 text-blue-800 hover:bg-blue-100 hover:border-blue-400 cursor-pointer`;
      default:
        return `${baseStyles} bg-gray-50 border-gray-300 text-gray-800 hover:bg-gray-100 hover:border-gray-400 cursor-pointer`;
    }
  };

  const getSeatIcon = () => {
    if (isFree) return '🆓';
    
    switch (seatType) {
      case 'window':
        return '🪟';
      case 'aisle':
        return '🚶';
      default:
        return '💺';
    }
  };

  const getTooltipText = () => {
    const baseText = `Seat ${seat.code} - ${SeatModel.getSeatLabel(seat)}`;
    const priceText = isFree ? ' (Free)' : price > 0 ? ` (${SeatModel.getSeatCurrency(seat)} ${price})` : '';
    
    if (!SeatModel.isSeatSelectable(seat)) {
      return `${baseText} - Occupied`;
    }
    
    if (!canSelectMore && !isSelected) {
      return `${baseText} - Maximum seats selected`;
    }
    
    return `${baseText}${priceText}`;
  };

  return (
    <div className="relative">
      <button
        className={getButtonStyles()}
        onClick={() => onSelect(seat, rowNumber)}
        onMouseEnter={() => seat.code && onHover(seat.code)}
        onMouseLeave={() => onHover(null)}
        disabled={!isSelectable}
        title={getTooltipText()}
      >
        <span className="hidden md:block text-xs">{getSeatIcon()}</span>
        <span className="md:hidden text-xs font-bold">{seat.code?.slice(-1)}</span>
      </button>
      
      {(isHovered || isSelected) && (
        <div className="absolute -top-8 left-1/2 transform -translate-x-1/2 bg-gray-800 text-white text-xs px-2 py-1 rounded whitespace-nowrap z-10">
          {isFree ? 'Free' : price > 0 ? `${SeatModel.getSeatCurrency(seat)} ${price}` : 'Included'}
        </div>
      )}
    </div>
  );
};