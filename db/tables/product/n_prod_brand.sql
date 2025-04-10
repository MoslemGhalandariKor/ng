CREATE TABLE n_prod_brand (
    row_id                 VARCHAR2(15) PRIMARY KEY, -- Unique identifier for the category
    CREATED                TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    CREATED_BY             VARCHAR2(400)    , -- User who created the record
    LAST_UPD               TIMESTAMP DEFAULT CURRENT_TIMESTAMP, -- Last update timestamp
    LAST_UPD_BY            VARCHAR2(400), -- User who last updated the record
    BRAND_NAME             VARCHAR2(255) NOT NULL, -- Name of the category
    BRAND_COUNTRY          VARCHAR2(255) NOT NULL, -- Name of the category country
    FULL_DESCRIPTION       CLOB, -- Detailed description of the category
    SHORT_DESCRIPTION      CLOB, -- Detailed description of the category
    P_BRAND_LOGO           VARCHAR2(400)   -- Url of Brand Logo
);