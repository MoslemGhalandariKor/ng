CREATE OR REPLACE FUNCTION FNC_GENERATE_PRODUCT_MANAGEMENT_ROW_ID(partition_prefix IN VARCHAR2 DEFAULT '1-') RETURN VARCHAR2 IS
    next_val  NUMBER;
    base36_id VARCHAR2(100);
    row_id VARCHAR2(100);
BEGIN
    -- Get the next value from the sequence
    SELECT seq_loy_row_id.NEXTVAL INTO next_val FROM DUAL;
    -- Convert the sequence value to Base-36
    base36_id := FNC_TO_BASE36(next_val);

    -- Prepend the partition prefix to create the ROW_ID
    row_id := partition_prefix || base36_id;

    RETURN row_id;
END FNC_GENERATE_PRODUCT_MANAGEMENT_ROW_ID;
/
