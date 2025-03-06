create table N_PROD_PRODUCT
(
  row_id        VARCHAR2(15) not null,
  created       TIMESTAMP(6),
  created_by    VARCHAR2(400),
  last_upd       TIMESTAMP(6),
  last_upd_by    VARCHAR2(400),
  name          VARCHAR2(400),
  description   VARCHAR2(4000),
  price         NUMBER,
  amount        NUMBER,
  prod_size     VARCHAR2(40),
  prod_length   NUMBER,
  prod_material VARCHAR2(400),
  prod_color    VARCHAR2(400),
  image_src     VARCHAR2(400),
  barcode       VARCHAR2(50),
  category_id   VARCHAR2(15),
  brand_id      VARCHAR2(15),
  status        VARCHAR2(20) default 'Active' -- Default value for Status
);
ALTER TABLE N_PROD_PRODUCT DROP COLUMN sku;
ALTER TABLE N_PROD_PRODUCT ADD sku VARCHAR2(50);
ALTER TABLE n_prod_product MODIFY (description VARCHAR2(4000));
DROP TABLE N_PROD_PRODUCT
ALTER TABLE n_prod_product DROP CONSTRAINT fk_product_brand;
SELECT * FROM n_prod_product;


SELECT FNC_GENERATE_ROW_ID() FROM DUAL;

SELECT * FROM TBL_TEST