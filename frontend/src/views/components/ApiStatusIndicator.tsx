import React from 'react';
import { Wifi, WifiOff, Database, RefreshCw } from 'lucide-react';

interface ApiStatusIndicatorProps {
  apiStatus: {
    isConnected: boolean;
    dataSource: 'mock' | 'api';
    apiConfigured: boolean;
  } | null;
  onRefresh: () => void;
  loading: boolean;
}

export const ApiStatusIndicator: React.FC<ApiStatusIndicatorProps> = ({
  apiStatus,
  onRefresh,
  loading,
}) => {
  if (!apiStatus) return null;

  const getStatusColor = () => {
    if (apiStatus.dataSource === 'mock') return 'text-amber-600 bg-amber-50 border-amber-200';
    if (apiStatus.isConnected) return 'text-green-600 bg-green-50 border-green-200';
    return 'text-red-600 bg-red-50 border-red-200';
  };

  const getStatusIcon = () => {
    if (apiStatus.dataSource === 'mock') return <Database size={16} />;
    if (apiStatus.isConnected) return <Wifi size={16} />;
    return <WifiOff size={16} />;
  };

  const getStatusText = () => {
    if (apiStatus.dataSource === 'mock') return 'Using Mock Data';
    if (apiStatus.isConnected) return 'API Connected';
    if (!apiStatus.apiConfigured) return 'API Not Configured';
    return 'API Disconnected';
  };

  return (
    <div className={`flex items-center gap-2 px-3 py-2 rounded-lg border text-sm ${getStatusColor()}`}>
      {getStatusIcon()}
      <span className="font-medium">{getStatusText()}</span>
      <button
        onClick={onRefresh}
        disabled={loading}
        className="ml-2 p-1 rounded hover:bg-black/10 transition-colors disabled:opacity-50"
        title="Refresh data"
      >
        <RefreshCw size={14} className={loading ? 'animate-spin' : ''} />
      </button>
    </div>
  );
};