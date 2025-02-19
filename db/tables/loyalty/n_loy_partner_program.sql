CREATE TABLE N_LOY_PARTNER_PROGRAM (
    ROW_ID               VARCHAR2(15 CHAR) NOT NULL, -- Primary Key
    CREATED              TIMESTAMP DEFAULT CURRENT_TIMESTAMP, -- Creation timestamp
    LAST_UPD             TIMESTAMP DEFAULT CURRENT_TIMESTAMP, -- Last update timestamp
    MODIFICATION_NUM     NUMBER DEFAULT 0, -- Modification count
    CONFLICT_ID          VARCHAR2(15 CHAR), -- Conflict resolution ID
    PROGRAM_ID           VARCHAR2(15 CHAR) NOT NULL, -- FK to N_LOY_PROGRAM
    PARTNER_ID           VARCHAR2(15 CHAR) NOT NULL, -- FK to N_PARTNER
    PARTNER_NAME         VARCHAR2(255 CHAR) NOT NULL, -- Partner company name
    STATUS_CD           VARCHAR2(30 CHAR) DEFAULT 'ACTIVE', -- Partner program status (Active, Inactive)
    AGREEMENT_START_DT   DATE NOT NULL, -- Agreement start date
    AGREEMENT_END_DT     DATE, -- Agreement end date (optional)
    REVENUE_SHARE_PCT    NUMBER(5,2) DEFAULT 0, -- Revenue share percentage with partner
    POINT_TRANSFER_FLG   CHAR(1) DEFAULT 'N', -- If points can be transferred between brands ('Y' or 'N')
    POINT_MULTIPLIER     NUMBER DEFAULT 1, -- Multiplier for points earned via partner purchases
    MAX_POINTS_EARN      NUMBER DEFAULT 0, -- Max points earned via partner (0 = no limit)
    REDEMPTION_ELIGIBLE_FLG CHAR(1) DEFAULT 'Y', -- If partner purchases can be redeemed with loyalty points
    EXCLUSIVE_OFFER_FLG  CHAR(1) DEFAULT 'N', -- If the partner provides exclusive deals for members
    PROMOTION_ELIGIBLE_FLG CHAR(1) DEFAULT 'Y', -- If members receive partner-specific promotions
    INTEGRATION_METHOD   VARCHAR2(100 CHAR), -- API, Manual, Batch Upload, etc.
    NOTES               VARCHAR2(2000 CHAR), -- Additional notes on the partnership
    CREATED_BY          VARCHAR2(100 CHAR), -- User who created the record
    LAST_UPD_BY         VARCHAR2(100 CHAR), -- User who last updated the record
    PRIMARY KEY (ROW_ID),
    FOREIGN KEY (PROGRAM_ID) REFERENCES N_LOY_PROGRAM(ROW_ID),
    FOREIGN KEY (PARTNER_ID) REFERENCES N_LOY_PARTNER(ROW_ID)
);