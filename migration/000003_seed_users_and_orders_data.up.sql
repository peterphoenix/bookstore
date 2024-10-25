INSERT INTO "user" (id, name, email, password, phone_number, address, city, state, zip_code, country, created_at, updated_at)
VALUES
    ('1ccc5f59-c962-41b4-b296-bda9dcbf0077', 'John Doe', 'john.doe@example.com', 'hashedpassword1', '1234567890', '123 Main St', 'New York', 'NY', '10001', 'USA', '2024-01-01 12:00:00', '2024-01-01 12:00:00'),
    ('d23699f0-96f7-4474-a6ec-5c6971c31504', 'Jane Smith', 'jane.smith@example.com', 'hashedpassword2', '0987654321', '456 Elm St', 'Los Angeles', 'CA', '90001', 'USA', '2024-01-01 12:00:00', '2024-01-01 12:00:00');

INSERT INTO "order" (id, user_id, total_amount, status, shipping_address, city, state, zip_code, country, created_at, updated_at)
VALUES
    ('d34434fc-b380-47ee-8af6-2f87afa92968', '1ccc5f59-c962-41b4-b296-bda9dcbf0077', 824550.00, 'COMPLETED', '123 Main St', 'New York', 'NY', '10001', 'USA', '2024-01-01 12:00:00', '2024-01-01 12:00:00'),
    ('cf843510-38dc-42db-aeea-482a48b5ac27', '1ccc5f59-c962-41b4-b296-bda9dcbf0077', 374850.00, 'COMPLETED', '123 Main St', 'New York', 'NY', '10001', 'USA', '2024-10-10 12:00:00', '2024-10-10 12:00:00'),
    ('b27b75a4-3006-4046-b2c9-9ad127ff97ef', 'd23699f0-96f7-4474-a6ec-5c6971c31504', 299850.00, 'COMPLETED', '456 Elm St', 'Los Angeles', 'CA', '90001', 'USA', '2024-01-01 12:00:00', '2024-01-01 12:00:00')
RETURNING id;

-- Insert seed data into "order_item" table
INSERT INTO "order_item" (order_id, book_id, quantity, price)
VALUES
    ('d34434fc-b380-47ee-8af6-2f87afa92968', (SELECT id from book where title = 'Harry Potter and the Philosophers Stone'), 2, 299850.00),
    ('d34434fc-b380-47ee-8af6-2f87afa92968', (SELECT id from book where title = 'The Hobbit'), 1, 224850.00),
    ('cf843510-38dc-42db-aeea-482a48b5ac27', (SELECT id from book where title = 'A Game of Thrones'), 1, 374850.00),
    ('b27b75a4-3006-4046-b2c9-9ad127ff97ef', (SELECT id from book where title = 'Harry Potter and the Philosophers Stone'), 1, 299850.00);
