CREATE OR REPLACE TRIGGER trg_n_loy_member_update_last_upd
BEFORE UPDATE ON N_LOY_MEMBER
FOR EACH ROW
BEGIN
    :NEW.LAST_UPD := CURRENT_TIMESTAMP;
    :NEW.LAST_UPD_BY := USER; -- Automatically populates with the database user executing the update
END;
/