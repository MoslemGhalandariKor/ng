

-- Create Purchase Order Item Table
CREATE TABLE n_prod_purchase_order_item (
    row_id VARCHAR2(15) PRIMARY KEY, -- Unique identifier for the order item
    CREATED                TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    CREATED_BY             VARCHAR2(400)    , -- User who created the record
    LAST_UPD               TIMESTAMP DEFAULT CURRENT_TIMESTAMP, -- Last update timestamp
    LAST_UPD_BY            VARCHAR2(400), -- User who last updated the record
    purchase_order_id VARCHAR2(15) NOT NULL, -- Reference to the purchase order
    product_id VARCHAR2(15) NOT NULL, -- Reference to the product
    variant_id VARCHAR2(15) NULL, -- Reference to the variant (if applicable)
    quantity INTEGER NOT NULL, -- Quantity ordered
    price DECIMAL(10,2) NOT NULL, -- Price per unit
    total_price DECIMAL(12,2) GENERATED ALWAYS AS (quantity * price) VIRTUAL, -- Auto-calculated total price
    CONSTRAINT fk_po_item_order FOREIGN KEY (purchase_order_id) REFERENCES n_prod_purchase_order(row_id),
    CONSTRAINT fk_po_item_product FOREIGN KEY (product_id) REFERENCES n_prod_product(row_id),
    CONSTRAINT fk_po_item_variant FOREIGN KEY (variant_id) REFERENCES n_prod_product_variant(row_id)
);
