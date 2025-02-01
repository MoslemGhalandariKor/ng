CREATE OR REPLACE FUNCTION format_uuid(raw_uuid RAW) RETURN VARCHAR2 IS
BEGIN
    RETURN REGEXP_REPLACE(
        RAWTOHEX(raw_uuid),
        '(.{8})(.{4})(.{4})(.{4})(.{12})',
        '\1-\2-\3-\4-\5'
    );
END;
/
