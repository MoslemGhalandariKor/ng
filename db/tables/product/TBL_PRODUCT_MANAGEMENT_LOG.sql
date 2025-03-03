CREATE TABLE TBL_PRODUCT_MANAGEMENT_LOG
(
    ROW_ID         NUMBER,
    LOG_DATE       TIMESTAMP     DEFAULT CURRENT_TIMESTAMP,
    LOG_LEVEL      VARCHAR2(20),
    LOG_MESSAGE    VARCHAR2(4000),
    METHODE_NAME   VARCHAR2(400),
    PARAMETERS     JSON,
    REQUEST_ID     VARCHAR2(100),                                                -- Unique ID for request tracking
    SESSION_ID     VARCHAR2(50)  DEFAULT SYS_CONTEXT('USERENV', 'SESSIONID'),    -- Session ID
    USER_NAME      VARCHAR2(100) DEFAULT SYS_CONTEXT('USERENV', 'SESSION_USER'), -- DB user
    CLIENT_IP      VARCHAR2(50)  DEFAULT SYS_CONTEXT('USERENV', 'IP_ADDRESS'),   -- Client IP
    MODULE_NAME    VARCHAR2(200) DEFAULT SYS_CONTEXT('USERENV', 'MODULE'),       -- Application module
    ACTION_NAME    VARCHAR2(200) DEFAULT SYS_CONTEXT('USERENV', 'ACTION'),
    EXECUTION_TIME NUMBER,
    RESPONSE_CODE  NUMBER,
    RESPONSE_DESC  VARCHAR2(4000),
    ERROR_CODE     VARCHAR2(400),
    ERROR_DESC     VARCHAR2(4000)
);

SELECT *
FROM TBL_PRODUCT_MANAGEMENT_LOG;

DROP TABLE TBL_PRODUCT_MANAGEMENT_LOG;
CREATE SEQUENCE SEQ_PRODUCT_MANAGEMENT_LOG START WITH 1 INCREMENT BY 1 CACHE 25;

SELECT SEQ_PRODUCT_MANAGEMENT_LOG.nextval
FROM DUAL;

SELECT SYS_CONTEXT('USERENV', 'ACTION')
FROM DUAL;
