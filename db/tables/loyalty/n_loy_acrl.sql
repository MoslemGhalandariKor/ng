CREATE TABLE N_LOY_ACRL (
    ROW_ID              VARCHAR2(15 CHAR) NOT NULL, -- Primary Key
    CREATED             TIMESTAMP DEFAULT CURRENT_TIMESTAMP, -- Creation timestamp
    LAST_UPD            TIMESTAMP DEFAULT CURRENT_TIMESTAMP, -- Last update timestamp
    MODIFICATION_NUM    NUMBER DEFAULT 0, -- Modification count
    CONFLICT_ID         VARCHAR2(15 CHAR), -- Conflict resolution ID
    TXN_ID              VARCHAR2(15 CHAR) NOT NULL, -- FK to N_LOY_TXN
    MEMBER_ID           VARCHAR2(15 CHAR) NOT NULL, -- FK to N_LOY_MEMBERS
    PROGRAM_ID          VARCHAR2(15 CHAR) NOT NULL, -- FK to N_LOY_PROGRAM
    TIER_ID             VARCHAR2(15 CHAR), -- FK to N_LOY_TIER
    POINT_TYPE_CD       VARCHAR2(30 CHAR), -- Type of points (e.g., "Standard", "Bonus")
    POINTS_EARNED       NUMBER DEFAULT 0, -- Points earned in the accrual
    MULTIPLIER          NUMBER DEFAULT 1, -- Multiplier applied to base points
    BASE_POINTS         NUMBER DEFAULT 0, -- Base points before any multipliers
    BONUS_POINTS        NUMBER DEFAULT 0, -- Additional bonus points
    EXPIRATION_DT       DATE, -- Expiration date of the points earned
    PARTNER_ID          VARCHAR2(15 CHAR), -- FK to N_PARTNER, if earned through a partner
    PRODUCT_ID          VARCHAR2(15 CHAR), -- FK to N_PRODUCT, if linked to a product
    STORE_ID            VARCHAR2(15 CHAR), -- Store or location where points were accrued
    OFFER_ID            VARCHAR2(15 CHAR), -- FK to N_OFFER, if linked to an offer
    PROMOTION_ID        VARCHAR2(15 CHAR), -- FK to N_PROMOTION, if linked to a promotion
    ACTIVITY_CD         VARCHAR2(30 CHAR), -- Activity code (e.g., "Purchase", "Referral")
    STATUS_CD           VARCHAR2(30 CHAR), -- Status of the accrual (e.g., "Pending", "Completed")
    CREATED_BY          VARCHAR2(100 CHAR), -- User who created the record
    LAST_UPD_BY         VARCHAR2(100 CHAR), -- User who last updated the record
    ACRL_DESC           VARCHAR2(255 CHAR), -- Description or comments for the accrual
    PRIMARY KEY (ROW_ID),
    FOREIGN KEY (TXN_ID) REFERENCES N_LOY_TXN(ROW_ID)
);