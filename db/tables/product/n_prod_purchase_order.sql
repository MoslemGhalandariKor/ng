CREATE TABLE n_prod_purchase_order (
    row_id VARCHAR2(15) PRIMARY KEY, -- Unique identifier for the purchase order
    CREATED                TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    CREATED_BY             VARCHAR2(400)    , -- User who created the record
    LAST_UPD               TIMESTAMP DEFAULT CURRENT_TIMESTAMP, -- Last update timestamp
    LAST_UPD_BY            VARCHAR2(400), -- User who last updated the record 
    supplier_id VARCHAR2(15) NOT NULL, -- Reference to the supplier
    status VARCHAR2(20) CHECK (status IN ('Pending', 'Completed', 'Cancelled')), -- Status of the purchase order
    total_amount DECIMAL(12,2) NOT NULL, -- Total amount of the purchase order
    CONSTRAINT fk_purchase_order_supplier FOREIGN KEY (supplier_id) REFERENCES n_prod_supplier(row_id)
);

