CREATE OR REPLACE TRIGGER trg_n_loy_acrl_update
BEFORE UPDATE ON N_LOY_ACRL
FOR EACH ROW
BEGIN
    :NEW.LAST_UPD := CURRENT_TIMESTAMP;
    :NEW.LAST_UPD_BY := USER; -- Automatically populates with the database user executing the update
END;
/