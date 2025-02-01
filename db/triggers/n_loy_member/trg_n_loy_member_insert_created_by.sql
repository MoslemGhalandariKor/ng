CREATE OR REPLACE TRIGGER trg_n_loy_member_insert_created_by
BEFORE INSERT ON N_LOY_MEMBER
FOR EACH ROW
BEGIN
    :NEW.CREATED_BY := USER; -- Automatically populates with the database user executing the update
END;
/


