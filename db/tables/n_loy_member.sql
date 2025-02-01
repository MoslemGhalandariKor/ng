CREATE TABLE N_LOY_MEMBER (
    ROW_ID                 VARCHAR2(15 CHAR) NOT NULL, -- Primary key, Siebel-style ROW_ID
    CREATED                TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    CREATED_BY             VARCHAR2(400), -- User who created the record
    LAST_UPD               TIMESTAMP DEFAULT CURRENT_TIMESTAMP, -- Last update timestamp
    LAST_UPD_BY            VARCHAR2(400), -- User who last updated the record
    MODIFICATION_NUM       NUMBER DEFAULT 0, -- Modification count for record tracking
    CONFLICT_ID            VARCHAR2(15 CHAR), -- Conflict resolution ID
    PARTY_ID               VARCHAR2(15 CHAR), -- Foreign key to N_PARTY
    CONTACT_ID             VARCHAR2(15 CHAR), -- Foreign key to N_CONTACT
    ORG_ID                 VARCHAR2(15 CHAR), -- Foreign key to N_ORG_ID
    FIRST_NAME             VARCHAR2(400 CHAR), -- First name of member
    LAST_NAME              VARCHAR2(400 CHAR), -- Last name of member
    PROGRAM_ID             VARCHAR2(15 CHAR), -- Foreign key to N_LOY_PROGRAM
    MEMBER_NUM             VARCHAR2(50 CHAR), -- Unique loyalty member number
    STATUS_CD              VARCHAR2(30 CHAR), -- Status code for the member
    SUB_STATUS_CD          VARCHAR2(30 CHAR), -- Sub Status code for the member
    TIER_ID                VARCHAR2(15 CHAR), -- Foreign key to N_LOY_TIER
    START_DATE             DATE, -- Membership start date
    EXPIRATION_DT          DATE, -- Membership expiration date
    POINT_TYPE_A_VAL       NUMBER DEFAULT 0, -- Total points
    POINT_TYPE_B_VAL       NUMBER DEFAULT 0, -- Total points
    POINT_TYPE_C_VAL       NUMBER DEFAULT 0, -- Total points
    SEGMENT_ID             VARCHAR2(15 CHAR), -- Foreign key to member's segment
    ACCOUNT_ID             VARCHAR2(15 CHAR), -- Foreign key to associated account
    HOUSEHOLD_ID           VARCHAR2(15 CHAR), -- Foreign key to associated household
    MEMBER_TYPE_CD         VARCHAR2(30 CHAR), -- Type of member (individual, household, etc.)
    ENROLLMENT_CHANNEL_CD  VARCHAR2(30 CHAR), -- Enrollment channel code
    PREFERRED_LANG_CD      VARCHAR2(30 CHAR), -- Preferred language code
    REFERRED_COMM_CD       VARCHAR2(30 CHAR), -- Preferred communication method
    TOTAL_VISITS           NUMBER DEFAULT 0, -- Total number of visits or transactions
    LAST_VISIT_DT          DATE, -- Last visit or transaction date
    REFERRAL_SRC_CD        VARCHAR2(30 CHAR), -- Referral source code
    REFERRAL_ID            VARCHAR2(15 CHAR), -- Referral entity ID
    MEMBER_DESC            VARCHAR2(255 CHAR), -- Description or notes about the member
    PRIMARY KEY (ROW_ID) -- Enforce unique primary key constraint
);

-- Create indexes for faster lookups
CREATE INDEX IDX_N_LOY_MEMBER_PARTY ON N_LOY_MEMBER (PARTY_ID);
CREATE INDEX IDX_N_LOY_MEMBER_PROGRAM ON N_LOY_MEMBER (PROGRAM_ID);
CREATE INDEX IDX_N_LOY_MEMBER_TIER ON N_LOY_MEMBER (TIER_ID);
CREATE INDEX IDX_N_LOY_MEMBER_ACCOUNT ON N_LOY_MEMBER (ACCOUNT_ID);
CREATE INDEX IDX_N_LOY_MEMBER_HOUSEHOLD ON N_LOY_MEMBER (HOUSEHOLD_ID);
CREATE INDEX IDX_N_LOY_MEMBER_STATUS ON N_LOY_MEMBER (STATUS_CD);
CREATE INDEX IDX_N_LOY_MEMBER_SUB_STATUS ON N_LOY_MEMBER (SUB_STATUS_CD);
CREATE INDEX IDX_N_LOY_MEMBER_CON_ID ON N_LOY_MEMBER (CONTACT_ID);


