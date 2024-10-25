CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE IF NOT EXISTS "author" (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    name VARCHAR(255) NOT NULL
);

-- currently title is not indexed yet because there is no search by title feature
-- if the feature were to be implemented, the index should be on elastic search
-- or use a postgresql extension for text search index.
CREATE TABLE IF NOT EXISTS "book" (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    title VARCHAR(255) NOT NULL,
    price DECIMAL(10, 2) NOT NULL,
    image_url VARCHAR(100),
    isbn VARCHAR(13) UNIQUE NOT NULL,
    description TEXT
);

-- author_id can be indexed if need to search list of books created by certain authors
CREATE TABLE IF NOT EXISTS "book_author" (
    book_id UUID,
    author_id UUID,

    PRIMARY KEY (book_id, author_id),
    FOREIGN KEY (author_id) REFERENCES "author"(id) ON DELETE CASCADE,
    FOREIGN KEY (book_id) REFERENCES "book"(id) ON DELETE CASCADE
);

-- address related fields should be moved into another table if support multiple addresses
CREATE TABLE IF NOT EXISTS "user" (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    name VARCHAR(100) NOT NULL,
    email VARCHAR(255) UNIQUE NOT NULL,
    password VARCHAR(255) NOT NULL,
    phone_number VARCHAR(15),
    address VARCHAR(255),
    city VARCHAR(100),
    state VARCHAR(100),
    zip_code VARCHAR(20),
    country VARCHAR(100),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);


CREATE TABLE IF NOT EXISTS "order" (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    user_id UUID,
    total_amount DECIMAL(10, 2) NOT NULL,
    status VARCHAR(50) DEFAULT 'PENDING',
    shipping_address VARCHAR(255),
    city VARCHAR(100),
    state VARCHAR(100),
    zip_code VARCHAR(20),
    country VARCHAR(100),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,

    FOREIGN KEY (user_id) REFERENCES "user"(id) ON DELETE CASCADE
);

CREATE INDEX IF NOT EXISTS idx_order_customer_id_order_date_desc ON "order"(user_id, created_at DESC);


CREATE TABLE IF NOT EXISTS "order_item" (
    order_id UUID,
    book_id UUID,
    quantity INT NOT NULL,
    price DECIMAL(10, 2) NOT NULL,

    PRIMARY KEY (order_id, book_id),
    FOREIGN KEY (order_id) REFERENCES "order"(id) ON DELETE CASCADE,
    FOREIGN KEY (book_id) REFERENCES "book"(id) ON DELETE CASCADE
);