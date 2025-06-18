-- Remove mock data in reverse order of dependencies

-- Remove booking seats
DELETE FROM booking_seats WHERE booking_id = '12345678-1234-1234-1234-123456789012';

-- Remove bookings
DELETE FROM bookings WHERE id = '12345678-1234-1234-1234-123456789012';

-- Remove seat prices
DELETE FROM seat_prices WHERE seat_id IN (
  'aaaaaaaa-aaaa-aaaa-aaaa-aaaaaaaaaaaa',
  'bbbbbbbb-bbbb-bbbb-bbbb-bbbbbbbbbbbb',
  'cccccccc-cccc-cccc-cccc-cccccccccccc',
  'dddddddd-dddd-dddd-dddd-dddddddddddd',
  'eeeeeeee-eeee-eeee-eeee-eeeeeeeeeeee',
  'ffffffff-ffff-ffff-ffff-ffffffffffff'
);

-- Remove seats
DELETE FROM seats WHERE cabin_id = '77777777-7777-7777-7777-777777777777';

-- Remove seat rows
DELETE FROM seat_rows WHERE cabin_id = '77777777-7777-7777-7777-777777777777';

-- Remove cabins
DELETE FROM cabins WHERE id = '77777777-7777-7777-7777-777777777777';

-- Remove aircraft
DELETE FROM aircraft WHERE id = '66666666-6666-6666-6666-666666666666';

-- Remove segments
DELETE FROM segments WHERE id = '55555555-5555-5555-5555-555555555555';

-- Remove itineraries
DELETE FROM itineraries WHERE id = '44444444-4444-4444-4444-444444444444';

-- Remove users
DELETE FROM users WHERE id IN (
  '11111111-1111-1111-1111-111111111111',
  '22222222-2222-2222-2222-222222222222',
  '33333333-3333-3333-3333-333333333333'
);

DELETE FROM frequent_flyers;
DELETE FROM passengers;