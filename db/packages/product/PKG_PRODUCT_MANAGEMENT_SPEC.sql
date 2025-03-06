CREATE OR REPLACE PACKAGE PKG_PRODUCT_MANAGEMENT AS
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
                            P_RESPONSE_DESC OUT VARCHAR2);

  PROCEDURE PRC_DELETE_PRODUCT(P_PRODUCT_ID    IN VARCHAR2,
                               P_RESPONSE_CODE OUT NUMBER,
                               P_RESPONSE_DESC OUT VARCHAR2);

  PROCEDURE PRC_ADD_CATEGORY(P_NAME          IN VARCHAR2,
                             P_DESCRIPTION   IN VARCHAR2,
                             P_PARENT_ID     IN VARCHAR2,
                             P_RESPONSE_CODE OUT NUMBER,
                             P_RESPONSE_DESC OUT VARCHAR2);

  PROCEDURE PRC_DELETE_CATEGORY(P_CATEGORY_ID   IN VARCHAR2,
                                P_RESPONSE_CODE OUT NUMBER,
                                P_RESPONSE_DESC OUT VARCHAR2);
END;
