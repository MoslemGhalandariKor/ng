CREATE TABLE N_LOY_PROD_INT (
    ROW_ID              VARCHAR2(15 CHAR) NOT NULL, -- Primary Key
    CREATED             TIMESTAMP DEFAULT CURRENT_TIMESTAMP, -- Creation timestamp
    LAST_UPD            TIMESTAMP DEFAULT CURRENT_TIMESTAMP, -- Last update timestamp
    MODIFICATION_NUM    NUMBER DEFAULT 0, -- Modification count
    CONFLICT_ID         VARCHAR2(15 CHAR), -- Conflict resolution ID
    PROD_NUM            VARCHAR2(50 CHAR) UNIQUE, -- Unique product number
    NAME                VARCHAR2(255 CHAR) NOT NULL, -- Product name
    DESCRIPTION         VARCHAR2(2000 CHAR), -- Product description
    LONG_DESCRIPTION    CLOB, -- Extended product details
    STATUS_CD           VARCHAR2(30 CHAR), -- Product status (e.g., "Active", "Inactive")
    CATEGORY_ID         VARCHAR2(15 CHAR), -- FK to product category
    TYPE_CD             VARCHAR2(30 CHAR), -- Product type (e.g., "Physical", "Digital")
    SUB_TYPE_CD         VARCHAR2(30 CHAR), -- Sub-type within product type
    PRICE               NUMBER(15, 2), -- Base price
    DISCOUNT_PRICE      NUMBER(15, 2), -- Discounted price (if applicable)
    CURRENCY_CD         VARCHAR2(10 CHAR), -- Currency code
    TAX_CODE            VARCHAR2(30 CHAR), -- Applicable tax code
    IS_TAXABLE_FLG      CHAR(1) DEFAULT 'Y', -- Taxable flag ('Y' or 'N')
    STOCK_QTY           NUMBER DEFAULT 0, -- Current stock quantity
    STOCK_STATUS_CD     VARCHAR2(30 CHAR), -- Stock status (e.g., "In Stock", "Out of Stock")
    MIN_ORDER_QTY       NUMBER DEFAULT 1, -- Minimum order quantity
    MAX_ORDER_QTY       NUMBER, -- Maximum order quantity
    MANUFACTURER        VARCHAR2(255 CHAR), -- Manufacturer name
    BRAND               VARCHAR2(255 CHAR), -- Brand name
    MODEL               VARCHAR2(255 CHAR), -- Model number or name
    WARRANTY_PERIOD     NUMBER, -- Warranty period in months
    DIMENSIONS          VARCHAR2(255 CHAR), -- Product dimensions (e.g., "10x20x30 cm")
    WEIGHT              NUMBER, -- Product weight
    COLOR               VARCHAR2(50 CHAR), -- Product color
    PROD_SIZE           VARCHAR2(50 CHAR), -- Product size
    MATERIAL            VARCHAR2(255 CHAR), -- Material composition
    COUNTRY_OF_ORIGIN   VARCHAR2(100 CHAR), -- Country of manufacture
    IMAGE_URL           VARCHAR2(2000 CHAR), -- Main product image
    ADDITIONAL_IMAGES   CLOB, -- JSON array of additional image URLs
    VIDEO_URL           VARCHAR2(2000 CHAR), -- Product video URL
    SEO_NAME            VARCHAR2(255 CHAR), -- SEO-friendly product name
    SEO_DESCRIPTION     VARCHAR2(2000 CHAR), -- SEO description
    SEO_KEYWORDS        VARCHAR2(1000 CHAR), -- SEO keywords
    URL_SLUG            VARCHAR2(255 CHAR), -- Product URL slug
    FEATURED_FLG        CHAR(1) DEFAULT 'N', -- Indicates if the product is featured
    BESTSELLER_FLG      CHAR(1) DEFAULT 'N', -- Indicates if the product is a bestseller
    PROMOTION_ID        VARCHAR2(15 CHAR), -- FK to promotion table (if product is on sale)
    CREATED_BY          VARCHAR2(100 CHAR), -- User who created the record
    LAST_UPD_BY         VARCHAR2(100 CHAR), -- User who last updated the record
    PRIMARY KEY (ROW_ID),
    FOREIGN KEY (CATEGORY_ID) REFERENCES N_LOY_PROD_CATEGORY(ROW_ID),
    FOREIGN KEY (PROMOTION_ID) REFERENCES N_LOY_PROMOTION(ROW_ID)
);

