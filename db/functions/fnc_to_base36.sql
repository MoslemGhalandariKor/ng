CREATE OR REPLACE FUNCTION FNC_TO_BASE36(num IN NUMBER) RETURN VARCHAR2 IS
    base36_chars CONSTANT VARCHAR2(36) := '0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZ';
    result       VARCHAR2(100) := '';
    quotient     NUMBER := num;
    remainder    NUMBER;
BEGIN
    -- Convert the number to Base-36
    WHILE quotient > 0 LOOP
        remainder := MOD(quotient, 36);
        quotient := FLOOR(quotient / 36);
        result := SUBSTR(base36_chars, remainder + 1, 1) || result;
    END LOOP;

    -- Handle the case where input is 0
    IF result IS NULL THEN
        result := '0';
    END IF;

    RETURN result;
END FNC_TO_BASE36;

