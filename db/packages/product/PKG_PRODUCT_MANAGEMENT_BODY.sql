CREATE OR REPLACE PACKAGE BODY PKG_PRODUCT_MANAGEMENT AS

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
  
    INSERT INTO TBL_PRODUCT_MANAGEMENT_LOG
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

  PROCEDURE PRC_ADD_PRODUCT(P_NAME          IN VARCHAR2,
                            P_DESCRIPTION   IN VARCHAR2,
                            P_PRICE         IN NUMBER,
                            P_PROD_SIZE     IN VARCHAR2,
                            P_PROD_LENGTH   IN NUMBER,
                            P_PROD_MATERIAL IN VARCHAR2,
                            P_PROD_COLOR    IN VARCHAR2,
                            P_IMAGE_SRC     IN VARCHAR2,
                            P_BARCODE       IN VARCHAR2,
                            P_CATEGORY_ID   IN VARCHAR2,
                            P_BRAND_ID      IN VARCHAR2,
                            P_RESPONSE_CODE OUT NUMBER,
                            P_RESPONSE_DESC OUT VARCHAR2) IS
  
    V_LOG_MESSAGE    VARCHAR2(4000);
    V_PARAMETERS     CLOB;
    V_START_TIME     TIMESTAMP;
    V_END_TIME       TIMESTAMP;
    V_EXECUTION_TIME VARCHAR2(400);
    V_METHODE_NAME   VARCHAR2(4000);
    V_RESPONSE_CODE  NUMBER;
    V_RESPONSE_DESC  VARCHAR2(4000);
    V_ERROR_CODE     VARCHAR2(400);
    V_ERROR_DESC     VARCHAR2(4000);
  BEGIN
  
    V_START_TIME := SYSTIMESTAMP;
  
    V_LOG_MESSAGE   := 'SUCCESS';
    V_RESPONSE_CODE := 0;
    V_RESPONSE_DESC := 'SUCCESS';
    V_ERROR_CODE    := 0;
    V_ERROR_DESC    := 'SUCCESS';
  
    V_PARAMETERS := JSON_OBJECT(KEY 'NAME' VALUE P_NAME,
                                KEY 'DESCRIPTION' VALUE P_DESCRIPTION,
                                KEY 'PRICE' VALUE P_PRICE,
                                KEY 'PROD_SIZE' VALUE P_PROD_SIZE,
                                KEY 'PROD_LENGTH' VALUE P_PROD_LENGTH,
                                KEY 'PROD_MATERIAL' VALUE P_PROD_MATERIAL,
                                KEY 'PROD_COLOR' VALUE P_PROD_COLOR,
                                KEY 'IMAGE_SRC' VALUE P_IMAGE_SRC,
                                KEY 'BARCODE' VALUE P_BARCODE,
                                KEY 'CATEGORY_ID' VALUE P_CATEGORY_ID,
                                KEY 'BRAND_ID' VALUE P_BRAND_ID);
  
    BEGIN
    
      INSERT INTO N_PROD_PRODUCT
        (ROW_ID,
         NAME,
         DESCRIPTION,
         PRICE,
         PROD_SIZE,
         PROD_LENGTH,
         PROD_MATERIAL,
         PROD_COLOR,
         IMAGE_SRC,
         BARCODE,
         CATEGORY_ID,
         BRAND_ID)
      VALUES
        (FNC_GENERATE_PRODUCT_MANAGEMENT_ROW_ID(),
         P_NAME,
         P_DESCRIPTION,
         P_PRICE,
         P_PROD_SIZE,
         P_PROD_LENGTH,
         P_PROD_MATERIAL,
         P_PROD_COLOR,
         P_IMAGE_SRC,
         P_BARCODE,
         P_CATEGORY_ID,
         P_BRAND_ID);
    
      COMMIT;
    
    EXCEPTION
      WHEN OTHERS THEN
        V_LOG_MESSAGE   := 'Error occurred: ' || SQLERRM;
        V_ERROR_CODE    := SQLCODE;
        V_ERROR_DESC    := SQLERRM;
        V_RESPONSE_CODE := -100;
        V_RESPONSE_DESC := 'Faild';
        ROLLBACK;
      
    END;
  
    V_END_TIME := SYSDATE;
    SELECT TO_CHAR((V_END_TIME - V_START_TIME))
      INTO V_EXECUTION_TIME
      FROM DUAL;
    P_RESPONSE_CODE := V_RESPONSE_CODE;
    P_RESPONSE_DESC := V_RESPONSE_DESC;
    PRC_INSERT_LOG(P_LOG_MESSAGE    => V_LOG_MESSAGE,
                   P_METHODE_NAME   => V_METHODE_NAME,
                   P_PARAMETERS     => V_PARAMETERS,
                   P_REQUEST_ID     => NULL,
                   P_EXECUTION_TIME => V_EXECUTION_TIME,
                   P_RESPONSE_CODE  => V_RESPONSE_CODE,
                   P_RESPONSE_DESC  => V_RESPONSE_DESC,
                   P_ERROR_CODE     => V_ERROR_CODE,
                   P_ERROR_DESC     => V_ERROR_DESC);
  END;

  ---------------------------------------------------------------------------------------------------    

  ---------------------------------------------------------------------------------------------------      
  PROCEDURE PRC_DELETE_PRODUCT(P_PRODUCT_ID    IN VARCHAR2,
                               P_RESPONSE_CODE OUT NUMBER,
                               P_RESPONSE_DESC OUT VARCHAR2) IS
  
    V_LOG_MESSAGE    VARCHAR2(4000);
    V_PARAMETERS     CLOB;
    V_START_TIME     TIMESTAMP;
    V_END_TIME       TIMESTAMP;
    V_EXECUTION_TIME VARCHAR2(400);
    V_METHODE_NAME   VARCHAR2(4000);
    V_RESPONSE_CODE  NUMBER;
    V_RESPONSE_DESC  VARCHAR2(4000);
    V_ERROR_CODE     VARCHAR2(400);
    V_ERROR_DESC     VARCHAR2(4000);
  BEGIN
  
    V_START_TIME := SYSTIMESTAMP;
  
    V_LOG_MESSAGE   := 'SUCCESS';
    V_RESPONSE_CODE := 0;
    V_RESPONSE_DESC := 'SUCCESS';
    V_ERROR_CODE    := 0;
    V_ERROR_DESC    := 'SUCCESS';
  
    V_PARAMETERS := JSON_OBJECT(KEY 'PRODUCT_ID' VALUE P_PRODUCT_ID);
  
    BEGIN
    
      DELETE N_PROD_PRODUCT P WHERE P.ROW_ID = P_PRODUCT_ID;
    
      COMMIT;
    
    EXCEPTION
      WHEN OTHERS THEN
        V_LOG_MESSAGE   := 'Error occurred: ' || SQLERRM;
        V_ERROR_CODE    := SQLCODE;
        V_ERROR_DESC    := SQLERRM;
        V_RESPONSE_CODE := -100;
        V_RESPONSE_DESC := 'Faild';
        ROLLBACK;
      
    END;
  
    V_END_TIME := SYSDATE;
    SELECT TO_CHAR((V_END_TIME - V_START_TIME))
      INTO V_EXECUTION_TIME
      FROM DUAL;
    P_RESPONSE_CODE := V_RESPONSE_CODE;
    P_RESPONSE_DESC := V_RESPONSE_DESC;
    PRC_INSERT_LOG(P_LOG_MESSAGE    => V_LOG_MESSAGE,
                   P_METHODE_NAME   => V_METHODE_NAME,
                   P_PARAMETERS     => V_PARAMETERS,
                   P_REQUEST_ID     => NULL,
                   P_EXECUTION_TIME => V_EXECUTION_TIME,
                   P_RESPONSE_CODE  => V_RESPONSE_CODE,
                   P_RESPONSE_DESC  => V_RESPONSE_DESC,
                   P_ERROR_CODE     => V_ERROR_CODE,
                   P_ERROR_DESC     => V_ERROR_DESC);
  END;

  ---------------------------------------------------------------------------------------------------    

  ---------------------------------------------------------------------------------------------------   

  PROCEDURE PRC_ADD_CATEGORY(P_NAME          IN VARCHAR2,
                             P_DESCRIPTION   IN VARCHAR2,
                             P_PARENT_ID     IN VARCHAR2,
                             P_RESPONSE_CODE OUT NUMBER,
                             P_RESPONSE_DESC OUT VARCHAR2) IS
  
    V_LOG_MESSAGE    VARCHAR2(4000);
    V_PARAMETERS     CLOB;
    V_START_TIME     TIMESTAMP;
    V_END_TIME       TIMESTAMP;
    V_EXECUTION_TIME VARCHAR2(400);
    V_METHODE_NAME   VARCHAR2(4000);
    V_RESPONSE_CODE  NUMBER;
    V_RESPONSE_DESC  VARCHAR2(4000);
    V_ERROR_CODE     VARCHAR2(400);
    V_ERROR_DESC     VARCHAR2(4000);
  BEGIN
  
    V_START_TIME := SYSTIMESTAMP;
  
    V_LOG_MESSAGE   := 'SUCCESS';
    V_RESPONSE_CODE := 0;
    V_RESPONSE_DESC := 'SUCCESS';
    V_ERROR_CODE    := 0;
    V_ERROR_DESC    := 'SUCCESS';
  
    V_PARAMETERS := JSON_OBJECT(KEY 'NAME' VALUE P_NAME,
                                KEY 'DESCRIPTION' VALUE P_DESCRIPTION,
                                KEY 'PARENT_ID' VALUE P_PARENT_ID);
  
    BEGIN
    
      INSERT INTO N_PROD_CATEGORY
        (ROW_ID, NAME, DESCRIPTION, PARENT_ID)
      VALUES
        (FNC_GENERATE_PRODUCT_MANAGEMENT_ROW_ID(),
         P_NAME,
         P_DESCRIPTION,
         P_PARENT_ID);
    
      COMMIT;
    
    EXCEPTION
      WHEN OTHERS THEN
        V_LOG_MESSAGE   := 'Error occurred: ' || SQLERRM;
        V_ERROR_CODE    := SQLCODE;
        V_ERROR_DESC    := SQLERRM;
        V_RESPONSE_CODE := -100;
        V_RESPONSE_DESC := 'Faild';
        ROLLBACK;
      
    END;
  
    V_END_TIME := SYSDATE;
    SELECT TO_CHAR((V_END_TIME - V_START_TIME))
      INTO V_EXECUTION_TIME
      FROM DUAL;
    P_RESPONSE_CODE := V_RESPONSE_CODE;
    P_RESPONSE_DESC := V_RESPONSE_DESC;
    PRC_INSERT_LOG(P_LOG_MESSAGE    => V_LOG_MESSAGE,
                   P_METHODE_NAME   => V_METHODE_NAME,
                   P_PARAMETERS     => V_PARAMETERS,
                   P_REQUEST_ID     => NULL,
                   P_EXECUTION_TIME => V_EXECUTION_TIME,
                   P_RESPONSE_CODE  => V_RESPONSE_CODE,
                   P_RESPONSE_DESC  => V_RESPONSE_DESC,
                   P_ERROR_CODE     => V_ERROR_CODE,
                   P_ERROR_DESC     => V_ERROR_DESC);
  END;

  ---------------------------------------------------------------------------------------------------    

  ---------------------------------------------------------------------------------------------------   

  PROCEDURE PRC_DELETE_CATEGORY(P_CATEGORY_ID   IN VARCHAR2,
                                P_RESPONSE_CODE OUT NUMBER,
                                P_RESPONSE_DESC OUT VARCHAR2) IS
  
    V_LOG_MESSAGE    VARCHAR2(4000);
    V_PARAMETERS     CLOB;
    V_START_TIME     TIMESTAMP;
    V_END_TIME       TIMESTAMP;
    V_EXECUTION_TIME VARCHAR2(400);
    V_METHODE_NAME   VARCHAR2(4000);
    V_RESPONSE_CODE  NUMBER;
    V_RESPONSE_DESC  VARCHAR2(4000);
    V_ERROR_CODE     VARCHAR2(400);
    V_ERROR_DESC     VARCHAR2(4000);
  BEGIN
  
    V_START_TIME := SYSTIMESTAMP;
  
    V_LOG_MESSAGE   := 'SUCCESS';
    V_RESPONSE_CODE := 0;
    V_RESPONSE_DESC := 'SUCCESS';
    V_ERROR_CODE    := 0;
    V_ERROR_DESC    := 'SUCCESS';
  
    V_PARAMETERS := JSON_OBJECT(KEY 'CATEGORY_ID' VALUE P_CATEGORY_ID);
  
    BEGIN
    
      DELETE N_PROD_CATEGORY P WHERE P.ROW_ID = P_CATEGORY_ID;
    
      COMMIT;
    
    EXCEPTION
      WHEN OTHERS THEN
        V_LOG_MESSAGE   := 'Error occurred: ' || SQLERRM;
        V_ERROR_CODE    := SQLCODE;
        V_ERROR_DESC    := SQLERRM;
        V_RESPONSE_CODE := -100;
        V_RESPONSE_DESC := 'Faild';
        ROLLBACK;
      
    END;
  
    V_END_TIME := SYSDATE;
    SELECT TO_CHAR((V_END_TIME - V_START_TIME))
      INTO V_EXECUTION_TIME
      FROM DUAL;
    P_RESPONSE_CODE := V_RESPONSE_CODE;
    P_RESPONSE_DESC := V_RESPONSE_DESC;
    PRC_INSERT_LOG(P_LOG_MESSAGE    => V_LOG_MESSAGE,
                   P_METHODE_NAME   => V_METHODE_NAME,
                   P_PARAMETERS     => V_PARAMETERS,
                   P_REQUEST_ID     => NULL,
                   P_EXECUTION_TIME => V_EXECUTION_TIME,
                   P_RESPONSE_CODE  => V_RESPONSE_CODE,
                   P_RESPONSE_DESC  => V_RESPONSE_DESC,
                   P_ERROR_CODE     => V_ERROR_CODE,
                   P_ERROR_DESC     => V_ERROR_DESC);
  END;

---------------------------------------------------------------------------------------------------    

---------------------------------------------------------------------------------------------------    

END;
