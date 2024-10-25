DROP INDEX IF EXISTS "idx_order_customer_id_order_date_desc";

DROP TABLE IF EXISTS "order_item";
DROP TABLE IF EXISTS "order";
DROP TABLE IF EXISTS "user";
DROP TABLE IF EXISTS "book_author";
DROP TABLE IF EXISTS "book";
DROP TABLE IF EXISTS "author";

DROP EXTENSION IF EXISTS "uuid-ossp";