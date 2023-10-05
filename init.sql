-- Create the schema
CREATE SCHEMA IF NOT EXISTS product_services COLLATE utf8mb4_0900_ai_ci;

-- Create the table
CREATE TABLE IF NOT EXISTS product_services.product_list (
                                                             id               INT AUTO_INCREMENT PRIMARY KEY,
                                                             product_name     VARCHAR(200) NULL,
                                                             product_desc     TEXT NULL,
                                                             product_price    INT NULL,
                                                             product_quantity INT NULL,
                                                             date_created     TIMESTAMP NULL
);