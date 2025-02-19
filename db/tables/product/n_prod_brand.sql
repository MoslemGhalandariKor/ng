CREATE TABLE n_prod_brand (
    row_id VARCHAR2(15) PRIMARY KEY, -- Unique identifier for the brand
    CREATED                TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    CREATED_BY             VARCHAR2(400)    , -- User who created the record
    LAST_UPD               TIMESTAMP DEFAULT CURRENT_TIMESTAMP, -- Last update timestamp
    LAST_UPD_BY            VARCHAR2(400), -- User who last updated the record
    name VARCHAR2(255) NOT NULL, -- Name of the brand
    country VARCHAR2(100) NOT NULL, -- Country of origin
    description CLOB, -- Description of the brand
    established_year NUMBER(4), -- Year the brand was established
    website VARCHAR2(255) -- Official website URL
);