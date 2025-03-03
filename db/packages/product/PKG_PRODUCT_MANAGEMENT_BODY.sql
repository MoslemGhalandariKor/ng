CREATE OR REPLACE PACKAGE BODY PKG_PRODUCT_MANAGEMENT AS
    PROCEDURE PRC_INSERT_LOG(P_METHODE_NAME  VARCHAR2,
                             P_RESPONSE_CODE NUMBER,
                             P_RESPONSE_DESC VARCHAR2,
                             P_ERROR_CODE    VARCHAR2,
                             P_ERROR_DESC    VARCHAR2) IS
    BEGIN

        INSERT INTO TBL_PRODUCT_MANAGEMENT_LOG(ROW_ID,
                                               METHODE_NAME,
                                               RESPONSE_CODE,
                                               RESPONSE_DESC,
                                               ERROR_CODE,
                                               ERROR_DESC)
                                        VALUES(SEQ_PRODUCT_MANAGEMENT_LOG.NEXTVAL,
                                               P_METHODE_NAME,
                                               P_RESPONSE_CODE,
                                               P_RESPONSE_DESC,
                                               P_ERROR_CODE,
                                               P_ERROR_DESC

                                              );
        COMMIT;

    END;
    PROCEDURE PRC_ADD_PRODUCT(P_NAME            IN VARCHAR2,
                              P_DESCRIPTION     IN VARCHAR2,
                              P_CATEGORY_ID     IN VARCHAR2,
                              P_SKU             IN VARCHAR2,
                              P_RESPONSE_CODE   OUT INT,
                              P_RESPONSE_DESC   OUT VARCHAR2) IS
    V_RESPONSE_CODE NUMBER;
    V_RESPONSE_DESC VARCHAR2(4000);
    V_ERROR_CODE VARCHAR2(400);
    V_ERROR_DESC VARCHAR2(4000);

    V_METHODE_NAME VARCHAR2(400);

    BEGIN
        V_RESPONSE_CODE := 0;
        V_RESPONSE_DESC := 'SUCCESS';
        V_ERROR_CODE := 0;
        V_ERROR_DESC := 'SUCCESS';
        V_METHODE_NAME := 'PRC_ADD_PRODUCT';
        BEGIN

        INSERT INTO N_PROD_PRODUCT(ROW_ID,
                                   CREATED,
                                   CREATED_BY,
                                   LAST_UPD,
                                   LAST_UPD_BY,
                                   NAME,
                                   SKU,
                                   BARCODE,
                                   CATEGORY_ID,
                                   BRAND_ID,
                                   WEIGHT,
                                   DIMENSIONS,
                                   STATUS,
                                   DESCRIPTION)
        VALUES (FNC_GENERATE_ROW_ID(),
                NULL,
                NULL,
                NULL,
                NULL,
                P_NAME,
                P_SKU,
                '1',
                P_CATEGORY_ID,
                NULL,
                NULL,
                NULL,
                'Active',
                P_DESCRIPTION

               );
        COMMIT;


        EXCEPTION
          WHEN OTHERS THEN
            V_RESPONSE_CODE := -100;
            V_RESPONSE_DESC := 'An Error Occurred';
            V_ERROR_CODE := SQLCODE;
            V_ERROR_DESC := SQLERRM;
            ROLLBACK ;

        END;

        P_RESPONSE_CODE := V_RESPONSE_CODE;
        P_RESPONSE_DESC := V_RESPONSE_DESC;

        PRC_INSERT_LOG(P_METHODE_NAME => V_METHODE_NAME,
                       P_RESPONSE_CODE => P_RESPONSE_CODE,
                       P_RESPONSE_DESC => P_RESPONSE_DESC,
                       P_ERROR_CODE => V_ERROR_CODE,
                       P_ERROR_DESC => V_ERROR_DESC);

   END;


END;
