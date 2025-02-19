-- Create Product Table
CREATE TABLE n_prod_product (
    row_id VARCHAR2(15) PRIMARY KEY, -- Unique identifier for the product
    CREATED                TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    CREATED_BY             VARCHAR2(400)    , -- User who created the record
    LAST_UPD               TIMESTAMP DEFAULT CURRENT_TIMESTAMP, -- Last update timestamp
    LAST_UPD_BY            VARCHAR2(400), -- User who last updated the record
    name VARCHAR2(255) NOT NULL, -- Name of the product
    description CLOB, -- Detailed description of the product
    sku VARCHAR2(50) UNIQUE NOT NULL, -- Unique Stock Keeping Unit
    barcode VARCHAR2(50) UNIQUE, -- Unique barcode identifier
    category_id VARCHAR2(15), -- Reference to product category
    brand_id VARCHAR2(15), -- Reference to product brand
    weight DECIMAL(10,2), -- Product weight in kg
    dimensions VARCHAR2(50), -- Product dimensions (LxWxH)
    status VARCHAR2(20) CHECK (status IN ('Active', 'Inactive', 'Discontinued')), -- Status of the product
    CONSTRAINT fk_product_category FOREIGN KEY (category_id) REFERENCES n_prod_category(row_id),
    CONSTRAINT fk_product_brand FOREIGN KEY (brand_id) REFERENCES n_prod_brand(row_id)
);