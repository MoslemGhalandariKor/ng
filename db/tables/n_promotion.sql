CREATE TABLE N_PROMOTION (
    ROW_ID              VARCHAR2(15 CHAR) NOT NULL, -- Primary Key
    CREATED             TIMESTAMP DEFAULT CURRENT_TIMESTAMP, -- Creation timestamp
    LAST_UPD            TIMESTAMP DEFAULT CURRENT_TIMESTAMP, -- Last update timestamp
    MODIFICATION_NUM    NUMBER DEFAULT 0, -- Modification count
    CONFLICT_ID         VARCHAR2(15 CHAR), -- Conflict resolution ID
    PROMOTION_NUM       VARCHAR2(50 CHAR) UNIQUE, -- Unique promotion code
    NAME                VARCHAR2(255 CHAR) NOT NULL, -- Promotion name
    DESCRIPTION         VARCHAR2(2000 CHAR), -- Promotion description
    STATUS_CD           VARCHAR2(30 CHAR), -- Promotion status (e.g., "Active", "Expired")
    TYPE_CD             VARCHAR2(30 CHAR), -- Type (e.g., "Discount", "Buy One Get One")
    START_DATE          DATE NOT NULL, -- Promotion start date
    END_DATE            DATE NOT NULL, -- Promotion end date
    DISCOUNT_TYPE_CD    VARCHAR2(30 CHAR), -- Discount type (e.g., "Percentage", "Fixed Amount")
    DISCOUNT_VALUE      NUMBER(15,2), -- Discount value (percentage or fixed amount)
    MIN_PURCHASE_AMT    NUMBER(15,2) DEFAULT 0, -- Minimum purchase amount required
    MAX_DISCOUNT_AMT    NUMBER(15,2), -- Maximum discount amount (if applicable)
    PROMO_CODE          VARCHAR2(100 CHAR), -- Promo code (if required for redemption)
    USAGE_LIMIT         NUMBER DEFAULT 0, -- Max times promo can be used (0 = unlimited)
    PER_USER_LIMIT      NUMBER DEFAULT 0, -- Max times a single user can use promo
    APPLICABLE_PROD     VARCHAR2(4000 CHAR), -- JSON list of applicable product IDs
    APPLICABLE_CATEGORIES VARCHAR2(4000 CHAR), -- JSON list of applicable category IDs
    LOYALTY_TIER_REQ    VARCHAR2(15 CHAR), -- Required loyalty tier for eligibility
    CHANNEL_CD          VARCHAR2(30 CHAR), -- Channel where promo is valid (e.g., "Online", "In-Store")
    STACKABLE_FLG       CHAR(1) DEFAULT 'N', -- Can be combined with other promotions ('Y' or 'N')
    AUTO_APPLY_FLG      CHAR(1) DEFAULT 'N', -- Automatically applied ('Y' or 'N')
    CREATED_BY          VARCHAR2(100 CHAR), -- User who created the record
    LAST_UPD_BY         VARCHAR2(100 CHAR), -- User who last updated the record
    PRIMARY KEY (ROW_ID),
    FOREIGN KEY (LOYALTY_TIER_REQ) REFERENCES N_LOY_TIER(ROW_ID)
);