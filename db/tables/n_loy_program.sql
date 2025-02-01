CREATE TABLE N_LOY_PROGRAM (
    ROW_ID              VARCHAR2(15 CHAR) NOT NULL, -- Primary Key
    CREATED             TIMESTAMP DEFAULT CURRENT_TIMESTAMP, -- Creation timestamp
    LAST_UPD            TIMESTAMP DEFAULT CURRENT_TIMESTAMP, -- Last update timestamp
    MODIFICATION_NUM    NUMBER DEFAULT 0, -- Modification count
    CONFLICT_ID         VARCHAR2(15 CHAR), -- Conflict resolution ID
    PROGRAM_NUM         VARCHAR2(50 CHAR) UNIQUE, -- Unique program identifier
    NAME               VARCHAR2(255 CHAR) NOT NULL, -- Loyalty program name
    DESCRIPTION        VARCHAR2(2000 CHAR), -- Program description
    STATUS_CD          VARCHAR2(30 CHAR), -- Status of the program (e.g., "Active", "Inactive")
    TYPE_CD            VARCHAR2(30 CHAR), -- Type of program (e.g., "Points-Based", "Tier-Based", "Hybrid")
    START_DATE         DATE NOT NULL, -- Program start date
    END_DATE           DATE, -- Optional end date for temporary programs
    ENROLLMENT_FLG     CHAR(1) DEFAULT 'Y', -- If customers can enroll ('Y' or 'N')
    AUTO_ENROLL_FLG    CHAR(1) DEFAULT 'N', -- If auto-enrollment is allowed ('Y' or 'N')
    POINT_EXPIRY_DAYS  NUMBER DEFAULT 365, -- Days before points expire
    TIER_EXPIRY_DAYS   NUMBER DEFAULT 365, -- Days before tier status expires
    MIN_PURCHASE_AMT   NUMBER(15,2) DEFAULT 0, -- Min purchase required for enrollment
    QUALIFICATION_CRITERIA VARCHAR2(4000 CHAR), -- JSON containing qualification rules
    REWARD_RULES       VARCHAR2(4000 CHAR), -- JSON containing reward accumulation rules
    REDEMPTION_RULES   VARCHAR2(4000 CHAR), -- JSON containing redemption rules
    MAX_REDEMPTION_LIMIT NUMBER DEFAULT 0, -- Max points redeemable per transaction (0 = unlimited)
    STACKABLE_FLG      CHAR(1) DEFAULT 'Y', -- Whether rewards stack with other programs ('Y' or 'N')
    PARTNER_ELIGIBLE_FLG CHAR(1) DEFAULT 'N', -- If program includes partner rewards
    PROMOTION_ELIGIBLE_FLG CHAR(1) DEFAULT 'Y', -- If members get exclusive promotions
    CHANNEL_CD         VARCHAR2(30 CHAR), -- Valid sales channels (e.g., "Online", "In-Store")
    CREATED_BY         VARCHAR2(100 CHAR), -- User who created the record
    LAST_UPD_BY        VARCHAR2(100 CHAR), -- User who last updated the record
    PRIMARY KEY (ROW_ID)
);