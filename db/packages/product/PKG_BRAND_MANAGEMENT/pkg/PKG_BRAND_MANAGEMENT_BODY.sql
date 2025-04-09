CREATE OR REPLACE PACKAGE BODY PKG_BRAND_MANAGEMENT IS


  PROCEDURE PRC_INSERT_LOG(P_LOG_MESSAGE    VARCHAR2,
                           P_METHODE_NAME   VARCHAR2,
                           P_PARAMETERS     CLOB,
                           P_REQUEST_ID     VARCHAR2,
                           P_EXECUTION_TIME VARCHAR2,
                           P_RESPONSE_CODE  NUMBER,
                           P_RESPONSE_DESC  VARCHAR2,
                           P_ERROR_CODE     VARCHAR2,
                           P_ERROR_DESC     VARCHAR2) IS
  BEGIN

    INSERT INTO TBL_BRAND_MANAGEMENT_LOG
      (ROW_ID,
       LOG_MESSAGE,
       METHODE_NAME,
       PARAMETERS,
       REQUEST_ID,
       EXECUTION_TIME,
       RESPONSE_CODE,
       RESPONSE_DESC,
       ERROR_CODE,
       ERROR_DESC)
    VALUES
      (SEQ_PRODUCT_MANAGEMENT_LOG.NEXTVAL,
       P_LOG_MESSAGE,
       P_METHODE_NAME,
       P_PARAMETERS,
       P_REQUEST_ID,
       P_EXECUTION_TIME,
       P_RESPONSE_CODE,
       P_RESPONSE_DESC,
       P_ERROR_CODE,
       P_ERROR_DESC);
    COMMIT;
  EXCEPTION
    WHEN OTHERS THEN
      NULL;
  END;

  ---------------------------------------------------------------------------------------------------

  ---------------------------------------------------------------------------------------------------



END;
