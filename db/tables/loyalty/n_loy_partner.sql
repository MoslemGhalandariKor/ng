CREATE TABLE N_LOY_PARTNER (
    ROW_ID              VARCHAR2(15 CHAR) NOT NULL, -- Primary Key
    CREATED             TIMESTAMP DEFAULT CURRENT_TIMESTAMP, -- Creation timestamp
    LAST_UPD            TIMESTAMP DEFAULT CURRENT_TIMESTAMP, -- Last update timestamp
    MODIFICATION_NUM    NUMBER DEFAULT 0, -- Modification count
    CONFLICT_ID         VARCHAR2(15 CHAR), -- Conflict resolution ID
    PARTNER_NUM         VARCHAR2(50 CHAR) UNIQUE, -- Unique partner identifier
    NAME                VARCHAR2(255 CHAR) NOT NULL, -- Partner name
    DESCRIPTION         VARCHAR2(2000 CHAR), -- Description of the partner
    STATUS_CD           VARCHAR2(30 CHAR) DEFAULT 'ACTIVE', -- Partner status (Active, Inactive)
    TYPE_CD             VARCHAR2(30 CHAR), -- Partner type (Merchant, Airline, Retail, Bank, etc.)
    INDUSTRY_CD         VARCHAR2(50 CHAR), -- Industry classification (e.g., Travel, Food, E-commerce)
    TAX_ID              VARCHAR2(50 CHAR), -- Partner's tax identification number
    CONTACT_ID          VARCHAR2(15 CHAR), -- FK to N_CONTACT (Primary contact person)
    ORG_ID              VARCHAR2(15 CHAR), -- FK to N_ORG (If applicable)
    ACCOUNT_ID          VARCHAR2(15 CHAR), -- FK to N_ACCOUNT (If linked to an account)
    ADDRESS             VARCHAR2(500 CHAR), -- Primary business address
    COUNTRY_CD          VARCHAR2(30 CHAR), -- Country code (ISO 3166-1)
    PHONE_NUM           VARCHAR2(50 CHAR), -- Primary contact number
    EMAIL              VARCHAR2(255 CHAR), -- Primary contact email
    WEBSITE_URL         VARCHAR2(255 CHAR), -- Partner's website URL
    AGREEMENT_START_DT  DATE NOT NULL, -- Partnership start date
    AGREEMENT_END_DT    DATE, -- Optional end date for the partnership
    REVENUE_SHARE_PCT   NUMBER(5,2) DEFAULT 0, -- Revenue share percentage
    POINT_ISSUE_FLG     CHAR(1) DEFAULT 'Y', -- If the partner can issue loyalty points ('Y' or 'N')
    POINT_REDEEM_FLG    CHAR(1) DEFAULT 'Y', -- If the partner allows point redemption ('Y' or 'N')
    MAX_POINTS_ISSUABLE NUMBER DEFAULT 0, -- Max points issuable per transaction (0 = no limit)
    MAX_POINTS_REDEEMABLE NUMBER DEFAULT 0, -- Max points redeemable per transaction (0 = no limit)
    EXCLUSIVE_PARTNER_FLG CHAR(1) DEFAULT 'N', -- If this is an exclusive partner ('Y' or 'N')
    PARTNER_TIER        VARCHAR2(30 CHAR), -- Partner tier level (e.g., Gold, Silver, Platinum)
    INTEGRATION_METHOD  VARCHAR2(100 CHAR), -- API, Manual, Batch Upload, etc.
    CREATED_BY          VARCHAR2(100 CHAR), -- User who created the record
    LAST_UPD_BY         VARCHAR2(100 CHAR), -- User who last updated the record
    NOTES               VARCHAR2(2000 CHAR), -- Additional partnership details
    PRIMARY KEY (ROW_ID)
);