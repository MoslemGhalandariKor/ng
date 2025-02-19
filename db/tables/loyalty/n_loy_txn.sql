CREATE TABLE N_LOY_TXN (
    ROW_ID               VARCHAR2(15 CHAR) NOT NULL, -- Primary Key
    CREATED              TIMESTAMP DEFAULT CURRENT_TIMESTAMP, -- Creation timestamp
    CREATED_BY           VARCHAR2(100 CHAR), -- User who created the record
    LAST_UPD             TIMESTAMP DEFAULT CURRENT_TIMESTAMP, -- Last update timestamp
    LAST_UPD_BY          VARCHAR2(100 CHAR), -- User who last updated the record
    MODIFICATION_NUM     NUMBER DEFAULT 0, -- Modification count
    CONFLICT_ID          VARCHAR2(15 CHAR), -- Conflict resolution ID
    TXN_NUM              VARCHAR2(50 CHAR) UNIQUE, -- Unique transaction number
    MEMBER_ID            VARCHAR2(15 CHAR) NOT NULL, -- FK to N_LOY_MEMBER
    PROGRAM_ID           VARCHAR2(15 CHAR) NOT NULL, -- FK to N_LOY_PROGRAM
    TXN_TYPE_CD          VARCHAR2(30 CHAR) NOT NULL, -- Transaction type code (e.g., "Accrual", "Redemption")
    TXN_DATE             TIMESTAMP DEFAULT CURRENT_TIMESTAMP, -- Date of the transaction
    TXN_STATUS_CD        VARCHAR2(30 CHAR), -- Status of the transaction (e.g., "Pending", "Completed")
    CHANNEL_CD           VARCHAR2(30 CHAR), -- Channel code (e.g., "Online", "In-Store")
    POINTS_EARNED        NUMBER DEFAULT 0, -- Points earned in the transaction
    POINTS_REDEEMED      NUMBER DEFAULT 0, -- Points redeemed in the transaction
    POINTS_ADJUSTED      NUMBER DEFAULT 0, -- Points adjusted in the transaction
    POINTS_BALANCE       NUMBER DEFAULT 0, -- Points balance after the transaction
    CURRENCY_CD          VARCHAR2(10 CHAR), -- Currency code for the transaction
    AMOUNT               NUMBER(15, 2), -- Monetary amount associated with the transaction
    BASE_POINTS          NUMBER DEFAULT 0, -- Base points earned before bonuses
    BONUS_POINTS         NUMBER DEFAULT 0, -- Bonus points earned
    MULTIPLIER           NUMBER DEFAULT 1, -- Multiplier applied to base points
    REF_TXN_ID           VARCHAR2(15 CHAR), -- Reference transaction ID for reversals or linked transactions
    HOUSEHOLD_ID         VARCHAR2(15 CHAR), -- FK to N_HOUSEHOLD, if linked to a household
    PARTNER_ID           VARCHAR2(15 CHAR), -- FK to N_PARTNER, if partner-related
    OFFER_ID             VARCHAR2(15 CHAR), -- FK to N_OFFER, for promotions or special offers
    PROMOTION_ID         VARCHAR2(15 CHAR), -- FK to N_PROMOTION, if tied to a promotion
    PRODUCT_ID           VARCHAR2(15 CHAR), -- FK to N_PRODUCT, if tied to a product
    STORE_ID             VARCHAR2(15 CHAR), -- Store or location ID
    PAYMENT_METHOD_CD    VARCHAR2(30 CHAR), -- Payment method code (e.g., "Credit Card", "Cash")
    SALES_PERSON_ID      VARCHAR2(15 CHAR), -- FK to N_CONTACT, if tied to a salesperson
    SEGMENT_ID           VARCHAR2(15 CHAR), -- FK to N_SEGMENT, if applicable
    TIER_ID              VARCHAR2(15 CHAR), -- FK to N_LOY_TIER, member tier at the time of the transaction
    ACTIVITY_CD          VARCHAR2(30 CHAR), -- Activity type code (e.g., "Purchase", "Referral")
    TXN_DESC             VARCHAR2(255 CHAR), -- Description of the transaction
    IS_REVERSAL_FLG      CHAR(1) DEFAULT 'N', -- Indicates if this transaction is a reversal ('Y' or 'N')
    REVERSAL_REASON_CD   VARCHAR2(30 CHAR), -- Reason code for reversal
    PRIMARY KEY (ROW_ID),
    FOREIGN KEY (MEMBER_ID) REFERENCES N_LOY_MEMBER(ROW_ID)
);