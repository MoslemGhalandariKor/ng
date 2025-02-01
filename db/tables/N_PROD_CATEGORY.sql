CREATE TABLE N_PROD_CATEGORY1 (
    ROW_ID              VARCHAR2(15 CHAR) NOT NULL, -- Primary Key
    CREATED             TIMESTAMP DEFAULT CURRENT_TIMESTAMP, -- Creation timestamp
    LAST_UPD            TIMESTAMP DEFAULT CURRENT_TIMESTAMP, -- Last update timestamp
    MODIFICATION_NUM    NUMBER DEFAULT 0, -- Modification count
    CONFLICT_ID         VARCHAR2(15 CHAR), -- Conflict resolution ID
    CATEGORY_NUM        VARCHAR2(50 CHAR) UNIQUE, -- Unique category identifier
    NAME                VARCHAR2(255 CHAR) NOT NULL, -- Category name
    DESCRIPTION         VARCHAR2(2000 CHAR), -- Category description
    PARENT_CATEGORY_ID  VARCHAR2(15 CHAR), -- FK to parent category (for hierarchy)
    STATUS_CD           VARCHAR2(30 CHAR), -- Category status (e.g., "Active", "Inactive")
    DISPLAY_SEQ         NUMBER DEFAULT 0, -- Display sequence order
    IMAGE_URL           VARCHAR2(2000 CHAR), -- Category image
    SEO_NAME            VARCHAR2(255 CHAR), -- SEO-friendly name
    SEO_DESCRIPTION     VARCHAR2(2000 CHAR), -- SEO meta description
    SEO_KEYWORDS        VARCHAR2(1000 CHAR), -- SEO keywords for search ranking
    URL_SLUG            VARCHAR2(255 CHAR), -- URL slug for category pages
    CATEGORY_TYPE_CD    VARCHAR2(30 CHAR), -- Type of category (e.g., "Retail", "Wholesale")
    FEATURED_FLG        CHAR(1) DEFAULT 'N', -- Indicates if the category is featured ('Y' or 'N')
    PROMOTION_ID        VARCHAR2(15 CHAR), -- FK to promotion table (if category is on sale)
    CREATED_BY          VARCHAR2(100 CHAR), -- User who created the record
    LAST_UPD_BY         VARCHAR2(100 CHAR), -- User who last updated the record
    PRIMARY KEY (ROW_ID),
    FOREIGN KEY (PARENT_CATEGORY_ID) REFERENCES N_PROD_CATEGORY(ROW_ID),
    FOREIGN KEY (PROMOTION_ID) REFERENCES N_PROMOTION(ROW_ID)
);
