CREATE OR REPLACE TRIGGER trg_n_prod_product_insert
BEFORE INSERT ON N_PROD_PRODUCT
FOR EACH ROW
BEGIN
    :NEW.CREATED := CURRENT_TIMESTAMP;
    :NEW.CREATED_BY := USER; -- Automatically populates with the database user executing the update
END;
/

