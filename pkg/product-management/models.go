package product_management

import "time"

// Product represents the n_prod_product table
type Product struct {
	RowID        string    `json:"row_id"`        // VARCHAR2(15)
	Created      time.Time `json:"created"`       // TIMESTAMP(6)
	CreatedBy    string    `json:"created_by"`    // VARCHAR2(400)
	LastUpd      time.Time `json:"last_upd"`      // TIMESTAMP(6)
	LastUpdBy    string    `json:"last_upd_by"`   // VARCHAR2(400)
	Name         string    `json:"name"`          // VARCHAR2(400)
	Description  string    `json:"description"`   // VARCHAR2(4000)
	Price        float64   `json:"price"`         // NUMBER
	ProdSize     string    `json:"prod_size"`     // VARCHAR2(40)
	ProdLength   float64   `json:"prod_length"`   // NUMBER
	ProdMaterial string    `json:"prod_material"` // VARCHAR2(400)
	ProdColor    string    `json:"prod_color"`    // VARCHAR2(400)
	ImageSrc     string    `json:"image_src"`     // VARCHAR2(400)
	Barcode      string    `json:"barcode"`       // VARCHAR2(50)
	CategoryID   string    `json:"category_id"`   // VARCHAR2(15)
	BrandID      string    `json:"brand_id"`      // VARCHAR2(15)
	Status       string    `json:"status"`        // VARCHAR2(20), default 'Active'
}

type Category struct {
	RowID       string    `json:"row_id"`      // VARCHAR2(15)
	Created     time.Time `json:"created"`     // TIMESTAMP(6)
	CreatedBy   string    `json:"created_by"`  // VARCHAR2(400)
	LastUpd     time.Time `json:"last_upd"`    // TIMESTAMP(6)
	LastUpdBy   string    `json:"last_upd_by"` // VARCHAR2(400)
	Name        string    `json:"name"`        // VARCHAR2(400)
	Description string    `json:"description"` // VARCHAR2(4000)
	ParentId    string    `json:"price"`       // NUMBER

}
