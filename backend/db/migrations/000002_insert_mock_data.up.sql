-- Mock data for the seat arrangements application
-- Based on the SeatMapResponse.json file

-- Insert mock users
INSERT INTO users (id, email, password, first_name, last_name, created_at, updated_at)
VALUES 
  ('11111111-1111-1111-1111-111111111111', 'john.doe@example.com', '$2a$10$N9qo8uLOickgx2ZMRZoMyeIjZAgcfl7p92ldGxad68LJZdL17lhWy', 'John', 'Doe', NOW(), NOW()),
  ('22222222-2222-2222-2222-222222222222', 'jane.doe@example.com', '$2a$10$N9qo8uLOickgx2ZMRZoMyeIjZAgcfl7p92ldGxad68LJZdL17lhWy', 'Jane', 'Doe', NOW(), NOW()),
  ('33333333-3333-3333-3333-333333333333', 'rutwik.sabre@example.com', '$2a$10$N9qo8uLOickgx2ZMRZoMyeIjZAgcfl7p92ldGxad68LJZdL17lhWy', 'Rutwik', 'Sabre', NOW(), NOW());

-- Insert mock itinerary
INSERT INTO itineraries (id, created_at, updated_at)
VALUES ('44444444-4444-4444-4444-444444444444', NOW(), NOW());

-- Insert mock segment
INSERT INTO segments (id, itinerary_id, origin, destination, departure, arrival, equipment, flight_number, airline_code, operating_flight_number, operating_airline_code, departure_terminal, arrival_terminal, duration, booking_class, cabin_class, created_at, updated_at)
VALUES ('55555555-5555-5555-5555-555555555555', '44444444-4444-4444-4444-444444444444', 'KUL', 'CGK', '2025-08-27 17:55:00', '2025-08-27 19:10:00', '738', '312', 'OD', '312', 'OD', 'TERMINAL 1', 'TERMINAL 2', 2, 'T', 'Economy', NOW(), NOW());

-- Insert mock aircraft
INSERT INTO aircraft (id, code, name, created_at, updated_at)
VALUES ('66666666-6666-6666-6666-666666666666', '738', 'Boeing 737-800', NOW(), NOW());

-- Insert mock cabin
INSERT INTO cabins (id, aircraft_id, segment_id, deck, first_row, last_row, seat_columns, created_at, updated_at)
VALUES ('77777777-7777-7777-7777-777777777777', '66666666-6666-6666-6666-666666666666', '55555555-5555-5555-5555-555555555555', 'MAIN', 4, 6, 'LEFT_SIDE,A,B,C,AISLE,D,E,F,RIGHT_SIDE', NOW(), NOW());

-- Insert mock seat row
INSERT INTO seat_rows (id, cabin_id, row_number, seat_codes, created_at, updated_at)
VALUES 
  ('88888888-8888-8888-8888-888888888888', '77777777-7777-7777-7777-777777777777', 0, 'BULKHEAD,BLANK', NOW(), NOW()),
  ('99999999-9999-9999-9999-999999999999', '77777777-7777-7777-7777-777777777777', 4, 'BLANK,AISLE,SEAT', NOW(), NOW());

-- Insert mock seats
INSERT INTO seats (id, row_id, segment_id, storefront_slot_code, code, available, entitled, fee_waived, free_of_charge, originally_selected, entitled_rule_id, fee_waived_rule_id, refund_indicator, seat_characteristics, raw_characteristics, created_at, updated_at)
VALUES 
  -- Row 4 seats
  ('aaaaaaaa-aaaa-aaaa-aaaa-aaaaaaaaaaaa', '99999999-9999-9999-9999-999999999999', '55555555-5555-5555-5555-555555555555', 'SEAT', '4A', true, true, false, false, false, '', '', 'R', 'CH,W', 'K,W,LS,L,CH', NOW(), NOW()),
  ('bbbbbbbb-bbbb-bbbb-bbbb-bbbbbbbbbbbb', '99999999-9999-9999-9999-999999999999', '55555555-5555-5555-5555-555555555555', 'SEAT', '4B', true, true, false, false, false, '', '', 'R', 'CH,9', 'K,LS,L,CH,9', NOW(), NOW()),
  ('cccccccc-cccc-cccc-cccc-cccccccccccc', '99999999-9999-9999-9999-999999999999', '55555555-5555-5555-5555-555555555555', 'SEAT', '4C', true, true, false, false, false, '', '', 'R', 'A,CH', 'A,K,LS,L,CH', NOW(), NOW()),
  ('dddddddd-dddd-dddd-dddd-dddddddddddd', '99999999-9999-9999-9999-999999999999', '55555555-5555-5555-5555-555555555555', 'SEAT', '4D', true, true, false, false, false, '', '', 'R', 'A,CH', 'A,K,RS,L,CH', NOW(), NOW()),
  ('eeeeeeee-eeee-eeee-eeee-eeeeeeeeeeee', '99999999-9999-9999-9999-999999999999', '55555555-5555-5555-5555-555555555555', 'SEAT', '4E', true, true, false, false, false, '', '', 'R', 'CH,9', 'K,RS,L,CH,9', NOW(), NOW()),
  ('ffffffff-ffff-ffff-ffff-ffffffffffff', '99999999-9999-9999-9999-999999999999', '55555555-5555-5555-5555-555555555555', 'SEAT', '4F', true, true, false, false, false, '', '', 'R', 'CH,W', 'K,W,RS,L,CH', NOW(), NOW()),
  -- Blank and aisle seats
  ('11111111-aaaa-bbbb-cccc-dddddddddddd', '88888888-8888-8888-8888-888888888888', '55555555-5555-5555-5555-555555555555', 'BLANK', NULL, false, false, false, true, false, NULL, NULL, NULL, 'LEFT_SIDE', 'LEFT_SIDE', NOW(), NOW()),
  ('22222222-aaaa-bbbb-cccc-dddddddddddd', '88888888-8888-8888-8888-888888888888', '55555555-5555-5555-5555-555555555555', 'AISLE', NULL, false, false, false, true, false, NULL, NULL, NULL, NULL, NULL, NOW(), NOW()),
  ('33333333-aaaa-bbbb-cccc-dddddddddddd', '88888888-8888-8888-8888-888888888888', '55555555-5555-5555-5555-555555555555', 'BLANK', NULL, false, false, false, true, false, NULL, NULL, NULL, 'RIGHT_SIDE', 'RIGHT_SIDE', NOW(), NOW());

-- Insert mock seat prices
INSERT INTO seat_prices (id, seat_id, type, amount, currency)
VALUES 
  ('11111111-2222-3333-4444-555555555555', 'aaaaaaaa-aaaa-aaaa-aaaa-aaaaaaaaaaaa', 'standard', 65.00, 'MYR'),
  ('22222222-3333-4444-5555-666666666666', 'bbbbbbbb-bbbb-bbbb-bbbb-bbbbbbbbbbbb', 'standard', 65.00, 'MYR'),
  ('33333333-4444-5555-6666-777777777777', 'cccccccc-cccc-cccc-cccc-cccccccccccc', 'standard', 65.00, 'MYR'),
  ('44444444-5555-6666-7777-888888888888', 'dddddddd-dddd-dddd-dddd-dddddddddddd', 'standard', 65.00, 'MYR'),
  ('55555555-6666-7777-8888-999999999999', 'eeeeeeee-eeee-eeee-eeee-eeeeeeeeeeee', 'standard', 65.00, 'MYR'),
  ('66666666-7777-8888-9999-aaaaaaaaaaaa', 'ffffffff-ffff-ffff-ffff-ffffffffffff', 'standard', 65.00, 'MYR');

-- Insert mock booking
INSERT INTO bookings (id, user_id, flight_id, status, created_at, updated_at)
VALUES ('12345678-1234-1234-1234-123456789012', '33333333-3333-3333-3333-333333333333', '55555555-5555-5555-5555-555555555555', 'confirmed', NOW(), NOW());

-- Insert mock booking seat
INSERT INTO booking_seats (id, booking_id, seat_id, price, currency, created_at, updated_at)
VALUES ('87654321-4321-4321-4321-210987654321', '12345678-1234-1234-1234-123456789012', 'aaaaaaaa-aaaa-aaaa-aaaa-aaaaaaaaaaaa', 65.00, 'MYR', NOW(), NOW());

-- Insert mock data for passengers
INSERT INTO passengers (segment_id, passenger_index, passenger_name_number, first_name, last_name, date_of_birth, gender, type, email, phone, street, city, country, postcode, address_type, document_type, issuing_country, country_of_birth, nationality)
VALUES
    -- Passenger 1
    ('55555555-5555-5555-5555-555555555555', 1, '01.01', 'Rutwik', 'Sabre', '1970-08-17', 'MALE', 'ADT', 'johnsabre@domain.com', NULL, 'street1 street2', 'city', 'US', '75039', 'HOME', 'F', '', '', 'MY'),
    -- Passenger 2 (fictional data based on pattern)
    ('55555555-5555-5555-5555-555555555555', 2, '01.02', 'Jane', 'Smith', '1985-05-20', 'FEMALE', 'ADT', 'janesmith@domain.com', '+1234567890', '123 Main St', 'Boston', 'US', '02108', 'HOME', 'P', 'US', 'US', 'US'),
    -- Passenger 3 (fictional data based on pattern)
    ('55555555-5555-5555-5555-555555555555', 3, '02.01', 'Michael', 'Johnson', '1990-11-15', 'MALE', 'ADT', 'mjohnson@domain.com', '+9876543210', '456 Oak Ave', 'Chicago', 'US', '60601', 'HOME', 'P', 'US', 'US', 'US'),
    -- Passenger 4 (fictional data based on pattern)
    ('55555555-5555-5555-5555-555555555555', 4, '02.02', 'Emily', 'Brown', '1982-03-28', 'FEMALE', 'ADT', 'ebrown@domain.com', '+1122334455', '789 Pine St', 'Seattle', 'US', '98101', 'HOME', 'P', 'US', 'US', 'US'),
    -- Passenger 5 (fictional data based on pattern)
    ('55555555-5555-5555-5555-555555555555', 5, '03.01', 'David', 'Wilson', '1975-09-10', 'MALE', 'ADT', 'dwilson@domain.com', '+5544332211', '321 Elm St', 'Denver', 'US', '80202', 'HOME', 'P', 'US', 'US', 'US');

-- Insert mock data for frequent flyers
-- We need to use the actual passenger IDs that were auto-generated
-- For now, we'll use a different approach by selecting the IDs from the passengers table
INSERT INTO frequent_flyers (passenger_id, airline, number, tier_level)
SELECT id, 'OD', '88700194295', 'BASIC' FROM passengers WHERE first_name = 'Rutwik' AND last_name = 'Sabre';

INSERT INTO frequent_flyers (passenger_id, airline, number, tier_level)
SELECT id, 'AA', '123456789', 'GOLD' FROM passengers WHERE first_name = 'Jane' AND last_name = 'Smith';

INSERT INTO frequent_flyers (passenger_id, airline, number, tier_level)
SELECT id, 'UA', '987654321', 'SILVER' FROM passengers WHERE first_name = 'Michael' AND last_name = 'Johnson';

INSERT INTO frequent_flyers (passenger_id, airline, number, tier_level)
SELECT id, 'DL', '456789123', 'PLATINUM' FROM passengers WHERE first_name = 'Emily' AND last_name = 'Brown';

INSERT INTO frequent_flyers (passenger_id, airline, number, tier_level)
SELECT id, 'BA', '789123456', 'EXECUTIVE' FROM passengers WHERE first_name = 'David' AND last_name = 'Wilson';