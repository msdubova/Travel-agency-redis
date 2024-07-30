CREATE TABLE bookings (
    reservation_number VARCHAR PRIMARY KEY,
    name VARCHAR,
    email VARCHAR,
    price FLOAT,
    status VARCHAR,
    tour_id VARCHAR
);

CREATE TABLE tours (
    id VARCHAR PRIMARY KEY,
    price FLOAT,
    start_date TIMESTAMP,
    end_date TIMESTAMP,
    destination VARCHAR,
    description TEXT
);

INSERT INTO tours (id, price, start_date, end_date, destination, description) VALUES 
('1', 299.99, NOW() + INTERVAL '2 day', NOW() + INTERVAL '7 day', 'Spain, Alicante', 'Enjoy a sunny beach'),
('2', 399.99, NOW() + INTERVAL '4 day', NOW() + INTERVAL '10 day', 'Austria, Courshavel', 'Explore the mountains');
