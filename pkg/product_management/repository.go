package product_management

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"nextgen/internals/db/ora"
)

func AddProduct(product Product) (ResponseCode int, ResponseDesc string) {
	ctx := context.Background()

	var p_response_code int
	var p_response_desc string

	q := `BEGIN 
		    PKG_PRODUCT_MANAGEMENT.PRC_ADD_PRODUCT(:1, :2, :3, :4, :5, :6, :7, :8, :9, :10, :11, :12, :13); 
		  END;`

	_, err := ora.OraDB.ExecContext(ctx,
		q,
		product.Name,
		product.Description,
		product.Price,
		// product.Amount,
		product.ProdSize,
		product.ProdLength,
		product.ProdMaterial,
		product.ProdColor,
		product.ImageSrc,
		product.Barcode,
		product.CategoryID,
		product.BrandID,
		sql.Out{Dest: &p_response_code},
		sql.Out{Dest: &p_response_desc},
	)

	if err != nil {
		log.Fatalf("Failed to execute procedure: %v", err)
	}
	fmt.Printf("Mapped Output P_RESPONSE_CODE: %d\n", p_response_code)
	fmt.Printf("Mapped Output P_RESPONSE_DESC: %s\n", p_response_desc)

	return p_response_code, p_response_desc
}

func GetAllProducts() (products []Product, err error) {

	q := `
		SELECT 	P.ROW_ID,
      			P.NAME,
       			P.DESCRIPTION,
				P.PROD_SIZE,
				P.PROD_LENGTH,
				P.PROD_MATERIAL,
				P.PROD_COLOR,
				P.IMAGE_SRC,
				
				P.BARCODE,
				P.CATEGORY_ID,
				P.BRAND_ID,
				P.STATUS,
				P.PRICE
  		FROM 	N_PROD_PRODUCT P	
`

	rows, err := ora.OraDB.Query(q)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	for rows.Next() {
		var product Product
		if err := rows.Scan(&product.RowID, &product.Name, &product.Description, &product.ProdSize, &product.ProdLength, &product.ProdMaterial, &product.ProdColor, &product.ImageSrc,  &product.Barcode, &product.CategoryID, &product.BrandID, &product.Status, &product.Price); err != nil {
			log.Fatal(err)
		}
		products = append(products, product)
	}
	return products, nil
}

func DeleteProductById(Product_id string) (ResponseCode int, ResponseDesc string) {
	ctx := context.Background()

	var p_response_code int
	var p_response_desc string

	q := `BEGIN 
		    PKG_PRODUCT_MANAGEMENT.PRC_DELETE_Product(:1, :2, :3); 
		  END;`

	_, err := ora.OraDB.ExecContext(ctx,
		q,
		Product_id,
		sql.Out{Dest: &p_response_code},
		sql.Out{Dest: &p_response_desc},
	)

	if err != nil {
		log.Fatalf("Failed to execute procedure: %v", err)
	}
	fmt.Printf("Mapped Output P_RESPONSE_CODE: %d\n", p_response_code)
	fmt.Printf("Mapped Output P_RESPONSE_DESC: %s\n", p_response_desc)

	return p_response_code, p_response_desc
}

func AddCategory(category Category) (ResponseCode int, ResponseDesc string) {
	ctx := context.Background()

	var p_response_code int
	var p_response_desc string

	q := `BEGIN 
		    PKG_PRODUCT_MANAGEMENT.PRC_ADD_CATEGORY(:1, :2, :3, :4, :5); 
		  END;`

	_, err := ora.OraDB.ExecContext(ctx,
		q,
		category.Name,
		category.Description,
		category.ParentId,
		sql.Out{Dest: &p_response_code},
		sql.Out{Dest: &p_response_desc},
	)

	if err != nil {
		log.Fatalf("Failed to execute procedure: %v", err)
	}
	fmt.Printf("Mapped Output P_RESPONSE_CODE: %d\n", p_response_code)
	fmt.Printf("Mapped Output P_RESPONSE_DESC: %s\n", p_response_desc)

	return p_response_code, p_response_desc
}

func GetAllCategories() (categories []Category, err error) {

	q := `
		SELECT  	C.ROW_ID,
      				C.NAME,
       				C.DESCRIPTION,
       				C1.NAME 
  		FROM 		N_PROD_CATEGORY C
  		LEFT JOIN 	N_PROD_CATEGORY C1
  		ON C1.ROW_ID = C.PARENT_ID
`

	rows, err := ora.OraDB.Query(q)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	for rows.Next() {
		var category Category
		if err := rows.Scan(&category.RowID, &category.Name, &category.Description, &category.ParentId); err != nil {
			log.Fatal(err)
		}
		categories = append(categories, category)
	}
	return categories, nil
}

func DeleteCategoryById(category_id string) (ResponseCode int, ResponseDesc string) {
	ctx := context.Background()

	var p_response_code int
	var p_response_desc string

	q := `BEGIN 
		    PKG_PRODUCT_MANAGEMENT.PRC_DELETE_CATEGORY(:1, :2, :3); 
		  END;`

	_, err := ora.OraDB.ExecContext(ctx,
		q,
		category_id,
		sql.Out{Dest: &p_response_code},
		sql.Out{Dest: &p_response_desc},
	)

	if err != nil {
		log.Fatalf("Failed to execute procedure: %v", err)
	}
	fmt.Printf("Mapped Output P_RESPONSE_CODE: %d\n", p_response_code)
	fmt.Printf("Mapped Output P_RESPONSE_DESC: %s\n", p_response_desc)

	return p_response_code, p_response_desc
}