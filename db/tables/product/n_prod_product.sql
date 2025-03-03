create table N_PROD_PRODUCT
(
  row_id          VARCHAR2(15) not null,
  created         TIMESTAMP(6) default CURRENT_TIMESTAMP,
  created_by      VARCHAR2(400),
  last_upd        TIMESTAMP(6) default CURRENT_TIMESTAMP,
  last_upd_by     VARCHAR2(400),
  name            VARCHAR2(400) not null,
  description     VARCHAR2(4000),
  PRICE           NUMBER,
  PROD_SIZE       VARCHAR2(40),
  PROD_LENGTH     NUMBER,
  PROD_MATERIAL   VARCHAR2(400),
  PROD_COLOR      VARCHAR2(400),
  IMAGE_SRC       VARCHAR2(400),
  barcode         VARCHAR2(50) UNIQUE,
  category_id     VARCHAR2(15),
  brand_id        VARCHAR2(15),
  status          VARCHAR2(20) DEFAULT 'Active'
);
ALTER TABLE N_PROD_PRODUCT DROP COLUMN sku;
ALTER TABLE N_PROD_PRODUCT ADD sku VARCHAR2(50);
ALTER TABLE n_prod_product MODIFY (description VARCHAR2(4000));

ALTER TABLE n_prod_product DROP CONSTRAINT fk_product_brand;
SELECT * FROM n_prod_product;


SELECT FNC_GENERATE_ROW_ID() FROM DUAL;

SELECT * FROM TBL_TEST