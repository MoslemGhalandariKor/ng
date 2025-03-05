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
		    PKG_PRODUCT_MANAGEMENT.PRC_ADD_PRODUCT(:1, :2, :3, :4, :5, :6); 
		  END;`

	_, err := ora.OraDB.ExecContext(ctx,
		q,
		product.Name,
		product.Description,
		product.CategoryID,
		'2',
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

	q := "SELECT * FROM N_PROD_PRODUCT"

	rows, err := ora.OraDB.Query(q)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	for rows.Next() {
		var product Product
		if err := rows.Scan(&product.RowID, &product.Name, &product.Price, &product.Name); err != nil {
			log.Fatal(err)
		}
		products = append(products, product)
	}
	return products, nil
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
