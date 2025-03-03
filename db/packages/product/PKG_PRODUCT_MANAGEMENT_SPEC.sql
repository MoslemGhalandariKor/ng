CREATE OR REPLACE PACKAGE PKG_PRODUCT_MANAGEMENT AS
    PROCEDURE PRC_ADD_PRODUCT(P_NAME            IN VARCHAR2,
                              P_DESCRIPTION     IN VARCHAR2,
                              P_CATEGORY_ID     IN VARCHAR2,
                              P_SKU             IN VARCHAR2,
                              P_RESPONSE_CODE   OUT INT,
                              P_RESPONSE_DESC   OUT VARCHAR2);
END;
