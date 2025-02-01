CREATE OR REPLACE TRIGGER trg_n_loy_txn_insert
BEFORE INSERT ON N_LOY_TXN
FOR EACH ROW
BEGIN
    :NEW.CREATED_BY := USER; -- Automatically populates with the database user executing the update
END;
/


