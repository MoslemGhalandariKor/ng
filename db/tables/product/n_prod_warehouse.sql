CREATE TABLE n_prod_warehouse (
    row_id VARCHAR2(15) PRIMARY KEY, -- Unique identifier for the warehouse
    CREATED                TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    CREATED_BY             VARCHAR2(400)    , -- User who created the record
    LAST_UPD               TIMESTAMP DEFAULT CURRENT_TIMESTAMP, -- Last update timestamp
    LAST_UPD_BY            VARCHAR2(400), -- User who last updated the record
    name VARCHAR2(255) NOT NULL, -- Name of the warehouse
    location VARCHAR2(255) NOT NULL, -- Location/address of the warehouse
    capacity INTEGER NOT NULL, -- Maximum storage capacity
    manager VARCHAR2(255) -- Warehouse manager name

);