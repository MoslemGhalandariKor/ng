CREATE TABLE n_prod_supplier (
    row_id VARCHAR2(15) PRIMARY KEY, -- Unique identifier for the supplier
    CREATED                TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    CREATED_BY             VARCHAR2(400)    , -- User who created the record
    LAST_UPD               TIMESTAMP DEFAULT CURRENT_TIMESTAMP, -- Last update timestamp
    LAST_UPD_BY            VARCHAR2(400), -- User who last updated the record
    name VARCHAR2(255) NOT NULL, -- Name of the supplier
    contact VARCHAR2(255), -- Contact details of the supplier
    email VARCHAR2(255), -- Email of the supplier
    phone VARCHAR2(20), -- Phone number
    address VARCHAR2(255), -- Address of the supplier
    country VARCHAR2(100) -- Country of the supplier

);