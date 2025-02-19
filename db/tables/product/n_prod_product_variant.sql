-- Create Product Variant Table
CREATE TABLE n_prod_product_variant (
    row_id VARCHAR2(15) PRIMARY KEY, -- Unique identifier for the variant
    CREATED                TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    CREATED_BY             VARCHAR2(400)    , -- User who created the record
    LAST_UPD               TIMESTAMP DEFAULT CURRENT_TIMESTAMP, -- Last update timestamp
    LAST_UPD_BY            VARCHAR2(400), -- User who last updated the record
    product_id VARCHAR2(15) NOT NULL, -- Reference to the main product
    name VARCHAR2(255) NOT NULL, -- Name of the variant
    sku VARCHAR2(50) UNIQUE NOT NULL, -- Unique SKU for the variant
    barcode VARCHAR2(50) UNIQUE, -- Unique barcode for the variant
    price DECIMAL(10,2) NOT NULL, -- Selling price of the variant
    stock INTEGER DEFAULT 0, -- Available stock for the variant
    CONSTRAINT fk_variant_product FOREIGN KEY (product_id) REFERENCES n_prod_product(row_id)
);