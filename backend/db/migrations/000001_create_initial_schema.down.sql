-- Drop indexes
DROP INDEX IF EXISTS idx_booking_seats_seat_id;
DROP INDEX IF EXISTS idx_booking_seats_booking_id;
DROP INDEX IF EXISTS idx_bookings_flight_id;
DROP INDEX IF EXISTS idx_bookings_user_id;
DROP INDEX IF EXISTS idx_seat_prices_seat_id;
DROP INDEX IF EXISTS idx_seats_segment_id;
DROP INDEX IF EXISTS idx_seats_cabin_id;
DROP INDEX IF EXISTS idx_seat_rows_cabin_id;
DROP INDEX IF EXISTS idx_cabins_segment_id;
DROP INDEX IF EXISTS idx_cabins_aircraft_id;
DROP INDEX IF EXISTS idx_segments_itinerary_id;
DROP INDEX IF EXISTS idx_users_email;
DROP INDEX IF EXISTS idx_frequent_flyers_passenger_id;
DROP INDEX IF EXISTS idx_passengers_segment_id;

-- Drop tables
DROP TABLE IF EXISTS frequent_flyers;
DROP TABLE IF EXISTS passengers;
DROP TABLE IF EXISTS booking_seats;
DROP TABLE IF EXISTS bookings;
DROP TABLE IF EXISTS seat_prices;
DROP TABLE IF EXISTS seats;
DROP TABLE IF EXISTS seat_rows;
DROP TABLE IF EXISTS cabins;
DROP TABLE IF EXISTS aircraft;
DROP TABLE IF EXISTS segments;
DROP TABLE IF EXISTS itineraries;
DROP TABLE IF EXISTS users;
