import React from 'react';
import { SeatMapView } from './views/SeatMapView';

function App() {
  // You can pass a flight ID here when available
  const flightId = 'FL123'; // Replace with actual flight ID or make it dynamic

  return <SeatMapView flightId={flightId} />;
}

export default App;