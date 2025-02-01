CREATE TABLE N_LOY_RDM (
    ROW_ID              VARCHAR2(15 CHAR) NOT NULL, -- Primary Key
    CREATED             TIMESTAMP DEFAULT CURRENT_TIMESTAMP, -- Creation timestamp
    LAST_UPD            TIMESTAMP DEFAULT CURRENT_TIMESTAMP, -- Last update timestamp
    MODIFICATION_NUM    NUMBER DEFAULT 0, -- Modification count
    CONFLICT_ID         VARCHAR2(15 CHAR), -- Conflict resolution ID
    TXN_ID              VARCHAR2(15 CHAR) NOT NULL, -- FK to N_LOY_TXN
    MEMBER_ID           VARCHAR2(15 CHAR) NOT NULL, -- FK to N_LOY_MEMBERS
    PROGRAM_ID          VARCHAR2(15 CHAR) NOT NULL, -- FK to N_LOY_PROGRAM
    TIER_ID             VARCHAR2(15 CHAR), -- FK to N_LOY_TIER
    POINT_TYPE_CD       VARCHAR2(30 CHAR), -- Type of points redeemed (e.g., "Standard", "Bonus")
    POINTS_REDEEMED     NUMBER DEFAULT 0, -- Total points redeemed
    REWARD_ID           VARCHAR2(15 CHAR), -- FK to N_LOY_REWARDS, reward being redeemed
    PARTNER_ID          VARCHAR2(15 CHAR), -- FK to N_PARTNER, if redeemed through a partner
    PRODUCT_ID          VARCHAR2(15 CHAR), -- FK to N_PRODUCT, if linked to a product
    STORE_ID            VARCHAR2(15 CHAR), -- Store or location where redemption occurred
    PROMOTION_ID        VARCHAR2(15 CHAR), -- FK to N_PROMOTION, if tied to a promotion
    EXPIRATION_DT       DATE, -- Expiration date of the reward redeemed
    STATUS_CD           VARCHAR2(30 CHAR), -- Status of the redemption (e.g., "Pending", "Completed")
    CREATED_BY          VARCHAR2(100 CHAR), -- User who created the record
    LAST_UPD_BY         VARCHAR2(100 CHAR), -- User who last updated the record
    RDM_DESC            VARCHAR2(255 CHAR), -- Description or comments for the redemption
    PRIMARY KEY (ROW_ID),
    FOREIGN KEY (TXN_ID) REFERENCES N_LOY_TXN(ROW_ID)
);

