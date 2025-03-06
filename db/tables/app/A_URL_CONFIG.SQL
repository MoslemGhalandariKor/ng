CREATE TABLE A_URL_CONFIG(
      row_id        VARCHAR2(15) not null,
      created       TIMESTAMP(6),
      created_by    VARCHAR2(400),
      last_upd       TIMESTAMP(6),
      last_upd_by    VARCHAR2(400),
      METHODE        VARCHAR2(4000),
      URL           VARCHAR2(4000),
      METHODE_TYPE  VARCHAR2(4000)
);
ALTER TABLE A_URL_CONFIG RENAME COLUMN METHOD TO METHODE;


SELECT * FROM A_URL_CONFIG;

