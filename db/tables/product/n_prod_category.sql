CREATE TABLE n_prod_category (
    row_id VARCHAR2(15) PRIMARY KEY, -- Unique identifier for the category
    CREATED                TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    CREATED_BY             VARCHAR2(400)    , -- User who created the record
    LAST_UPD               TIMESTAMP DEFAULT CURRENT_TIMESTAMP, -- Last update timestamp
    LAST_UPD_BY            VARCHAR2(400), -- User who last updated the record
    name VARCHAR2(255) NOT NULL, -- Name of the category
    parent_id VARCHAR2(15) NULL, -- Parent category for hierarchical structuring
    description CLOB, -- Detailed description of the category
    CONSTRAINT fk_category_parent FOREIGN KEY (parent_id) REFERENCES n_prod_category(row_id)
);
