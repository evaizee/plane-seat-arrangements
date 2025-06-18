-- Create users table
CREATE TABLE IF NOT EXISTS users (
    id UUID PRIMARY KEY,
    email VARCHAR(255) NOT NULL UNIQUE,
    password VARCHAR(255) NOT NULL,
    first_name VARCHAR(255) NOT NULL,
    last_name VARCHAR(255) NOT NULL,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL
);

-- Create itineraries table
CREATE TABLE IF NOT EXISTS itineraries (
    id UUID PRIMARY KEY,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL
);

-- Create segments table
CREATE TABLE IF NOT EXISTS segments (
    id UUID PRIMARY KEY,
    itinerary_id UUID REFERENCES itineraries(id),
    origin VARCHAR(3) NOT NULL,
    destination VARCHAR(3) NOT NULL,
    departure TIMESTAMP NOT NULL,
    arrival TIMESTAMP NOT NULL,
    equipment VARCHAR(10),
    flight_number VARCHAR(10) NOT NULL,
    airline_code VARCHAR(3) NOT NULL,
    operating_flight_number VARCHAR(10),
    operating_airline_code VARCHAR(3),
    departure_terminal VARCHAR(10),
    arrival_terminal VARCHAR(10),
    duration INTEGER NOT NULL,
    booking_class VARCHAR(10) NOT NULL,
    cabin_class VARCHAR(10) NOT NULL,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL
);

-- Create aircraft table
CREATE TABLE IF NOT EXISTS aircraft (
    id UUID PRIMARY KEY,
    code VARCHAR(10) NOT NULL,
    name VARCHAR(255) NOT NULL,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL
);

-- Create cabins table
CREATE TABLE IF NOT EXISTS cabins (
    id UUID PRIMARY KEY,
    aircraft_id UUID REFERENCES aircraft(id),
    segment_id UUID REFERENCES segments(id),
    deck VARCHAR(10) NOT NULL,
    first_row INTEGER NOT NULL,
    last_row INTEGER NOT NULL,
    seat_columns VARCHAR(255) NOT NULL,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL
);

-- Create seat_rows table
CREATE TABLE IF NOT EXISTS seat_rows (
    id UUID PRIMARY KEY,
    cabin_id UUID REFERENCES cabins(id),
    row_number INTEGER NOT NULL,
    seat_codes VARCHAR(255) NOT NULL,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL
);

-- Create seats table
CREATE TABLE IF NOT EXISTS seats (
    id UUID PRIMARY KEY,
    row_id UUID REFERENCES seat_rows(id),
    segment_id UUID REFERENCES segments(id),
    storefront_slot_code VARCHAR(50) NOT NULL,
    code VARCHAR(10),
    available BOOLEAN NOT NULL,
    entitled BOOLEAN NOT NULL,
    fee_waived BOOLEAN NOT NULL,
    free_of_charge BOOLEAN NOT NULL,
    originally_selected BOOLEAN NOT NULL,
    entitled_rule_id VARCHAR(50),
    fee_waived_rule_id VARCHAR(50),
    refund_indicator VARCHAR(10),
    seat_characteristics VARCHAR(255),
    raw_characteristics VARCHAR(255),
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL
);

-- Create seat_prices table
CREATE TABLE IF NOT EXISTS seat_prices (
    id UUID PRIMARY KEY,
    seat_id UUID REFERENCES seats(id),
    type VARCHAR(50) NOT NULL,
    amount DECIMAL(10, 2) NOT NULL,
    currency VARCHAR(3) NOT NULL
);

-- Create bookings table
CREATE TABLE IF NOT EXISTS bookings (
    id UUID PRIMARY KEY,
    user_id UUID REFERENCES users(id),
    flight_id UUID REFERENCES segments(id),
    status VARCHAR(20) NOT NULL,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL
);

-- Create booking_seats table
CREATE TABLE IF NOT EXISTS booking_seats (
    id UUID PRIMARY KEY,
    booking_id UUID REFERENCES bookings(id),
    seat_id UUID REFERENCES seats(id),
    price DECIMAL(10, 2) NOT NULL,
    currency VARCHAR(3) NOT NULL,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL
);

-- Create passengers table
CREATE TABLE IF NOT EXISTS passengers (
    id SERIAL PRIMARY KEY,
    segment_id UUID,
    passenger_index INTEGER NOT NULL,
    passenger_name_number VARCHAR(255) NOT NULL,
    first_name VARCHAR(255) NOT NULL,
    last_name VARCHAR(255) NOT NULL,
    date_of_birth DATE,
    gender VARCHAR(50),
    type VARCHAR(50),
    email VARCHAR(255),
    phone VARCHAR(50),
    street VARCHAR(255),
    city VARCHAR(255),
    country VARCHAR(255),
    postcode VARCHAR(50),
    address_type VARCHAR(50),
    document_type VARCHAR(50),
    issuing_country VARCHAR(50),
    country_of_birth VARCHAR(50),
    nationality VARCHAR(50),
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

-- Create frequent_flyers table
CREATE TABLE IF NOT EXISTS frequent_flyers (
    id SERIAL PRIMARY KEY,
    passenger_id SERIAL NOT NULL REFERENCES passengers(id) ON DELETE CASCADE,
    airline VARCHAR(255) NOT NULL,
    number VARCHAR(255) NOT NULL,
    tier_level VARCHAR(50),
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

-- Create indexes
CREATE INDEX idx_users_email ON users(email);
CREATE INDEX idx_segments_itinerary_id ON segments(itinerary_id);
CREATE INDEX idx_cabins_aircraft_id ON cabins(aircraft_id);
CREATE INDEX idx_cabins_segment_id ON cabins(segment_id);
CREATE INDEX idx_seat_rows_cabin_id ON seat_rows(cabin_id);
CREATE INDEX idx_seats_row_id ON seats(row_id);
CREATE INDEX idx_seats_segment_id ON seats(segment_id);
CREATE INDEX idx_seat_prices_seat_id ON seat_prices(seat_id);
CREATE INDEX idx_bookings_user_id ON bookings(user_id);
CREATE INDEX idx_bookings_flight_id ON bookings(flight_id);
CREATE INDEX idx_booking_seats_booking_id ON booking_seats(booking_id);
CREATE INDEX idx_booking_seats_seat_id ON booking_seats(seat_id);
CREATE INDEX idx_frequent_flyers_passenger_id ON frequent_flyers(passenger_id);