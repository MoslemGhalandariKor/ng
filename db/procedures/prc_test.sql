CREATE PROCEDURE PRC_TEST(P_PARAM1 IN VARCHAR2,
                          P_PARAM2 IN VARCHAR2,
                          P_RESPONSE_CODE OUT NUMBER,
                          P_RESPONSE_DESC OUT VARCHAR2) IS

BEGIN
  P_RESPONSE_CODE := 0;
  P_RESPONSE_DESC := 'SUCCESS';
END ;
/

