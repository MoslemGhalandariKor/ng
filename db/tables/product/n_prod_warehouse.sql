CREATE TABLE n_prod_warehouse (
    row_id              VARCHAR2(15) PRIMARY KEY,
    -- Unique identifier for the warehose
    CREATED             TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    CREATED_BY          VARCHAR2(400),
    -- User who created the record
    LAST_UPD            TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    -- Last update timestamp
    LAST_UPD_BY         VARCHAR2(400),
    -- User who last updated the record
    WAREHOUSE_NAME      VARCHAR2(255),
    -- Name of the warehose
    WAREHOUSE_MANAGER   VARCHAR2(255),
    -- Name of the warehose manager
    WAREHOUSE_ADDRESS   VARCHAR2(400),
    -- Detailed address of the warehose
    NUMBER_WORKERS      NUMBER,
    -- Detailed number of workers of the warehose
    NUMBER_PRODUCTS     NUMBER, 
    -- Detailed number of products of the warehose
    WAREHOUSE_IMG      VARCHAR2(400),
    status              VARCHAR2(20) default 'Active' 
    -- Default value for Status
);