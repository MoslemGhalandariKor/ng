CREATE TABLE N_LOY_TIER (
    ROW_ID              VARCHAR2(15 CHAR) NOT NULL, -- Primary Key
    CREATED             TIMESTAMP DEFAULT CURRENT_TIMESTAMP, -- Creation timestamp
    LAST_UPD            TIMESTAMP DEFAULT CURRENT_TIMESTAMP, -- Last update timestamp
    MODIFICATION_NUM    NUMBER DEFAULT 0, -- Modification count
    CONFLICT_ID         VARCHAR2(15 CHAR), -- Conflict resolution ID
    PROGRAM_ID         VARCHAR2(15 CHAR) NOT NULL, -- FK to loyalty program
    TIER_NUM           VARCHAR2(50 CHAR) UNIQUE, -- Unique tier number
    NAME               VARCHAR2(255 CHAR) NOT NULL, -- Tier name (e.g., "Gold", "Platinum")
    DESCRIPTION        VARCHAR2(2000 CHAR), -- Description of the tier
    STATUS_CD          VARCHAR2(30 CHAR), -- Status of the tier (e.g., "Active", "Inactive")
    MIN_POINTS_REQ     NUMBER DEFAULT 0, -- Minimum points required for this tier
    MAX_POINTS_REQ     NUMBER, -- Maximum points allowed for this tier (if applicable)
    QUALIFICATION_CRITERIA VARCHAR2(4000 CHAR), -- JSON containing qualification rules
    BENEFITS           VARCHAR2(4000 CHAR), -- JSON containing benefits for this tier
    EXPIRATION_DAYS    NUMBER DEFAULT 365, -- Number of days before tier expires
    RENEWAL_POLICY     VARCHAR2(4000 CHAR), -- JSON defining tier renewal rules
    AUTO_UPGRADE_FLG   CHAR(1) DEFAULT 'Y', -- Auto-upgrade based on points ('Y' or 'N')
    AUTO_DOWNGRADE_FLG CHAR(1) DEFAULT 'Y', -- Auto-downgrade based on inactivity ('Y' or 'N')
    PROMOTION_ELIGIBLE_FLG CHAR(1) DEFAULT 'Y', -- If tier members can receive exclusive promotions
    CREATED_BY         VARCHAR2(100 CHAR), -- User who created the record
    LAST_UPD_BY        VARCHAR2(100 CHAR), -- User who last updated the record
    PRIMARY KEY (ROW_ID),
    FOREIGN KEY (PROGRAM_ID) REFERENCES N_LOY_PROGRAM(ROW_ID)
);