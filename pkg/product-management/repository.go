package product_management

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"nextgen/internals/db/ora"
)

func PrcAddProduct(product Product) (ResponseCode int, ResponseDesc string) {
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

func PrcGetAllProducts() (products []Product, err error) {
	var p_response_code int
	var p_response_desc string
	var cursor *sql.Rows
	var prods []Product
	ctx := context.Background()
	query := "BEGIN PKG_PRODUCT_MANAGEMENT.PRC_GET_ALL_PRODUCTS(:1, :2, :3); END;"

	_, err = ora.OraDB.ExecContext(ctx,
		query,
		sql.Out{Dest: &cursor},
		sql.Out{Dest: &p_response_code},
		sql.Out{Dest: &p_response_desc})

	defer cursor.Close()

	if err != nil {
		log.Println("Error executing stored procedure:", err)
		return nil, err
	}
	for cursor.Next() {
		var product Product
		err = cursor.Scan(&product.Name, product.Description)
		if err != nil {
			log.Println("Error scanning product:", err)
			return nil, err
		}
		prods = append(prods, product)
	}

	return prods, nil
}
