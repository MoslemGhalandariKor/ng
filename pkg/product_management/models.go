package product_management

import "time"

// Product represents the n_prod_product table
type Product struct {
	RowID        string    `form:"row_id"`        // VARCHAR2(15)
	Created      time.Time `form:"created"`       // TIMESTAMP(6)
	CreatedBy    string    `form:"created_by"`    // VARCHAR2(400)
	LastUpd      time.Time `form:"last_upd"`      // TIMESTAMP(6)
	LastUpdBy    string    `form:"last_upd_by"`   // VARCHAR2(400)
	Name         string    `form:"name"`          // VARCHAR2(400)
	Description  string    `form:"description"`   // VARCHAR2(4000)
	Price        string    `form:"price"`         // NUMBER
	Amount       string    `form:"amount"`        // NUMBER
	ProdSize     string    `form:"prod_size"`     // VARCHAR2(40)
	ProdLength   string    `form:"prod_length"`   // NUMBER
	ProdMaterial string    `form:"prod_material"` // VARCHAR2(400)
	ProdColor    string    `form:"prod_color"`    // VARCHAR2(400)
	ImageSrc     string    `form:"image_src"`     // VARCHAR2(400)
	Barcode      string    `form:"barcode"`       // VARCHAR2(50)
	CategoryID   string    `form:"category_id"`   // VARCHAR2(15)
	BrandID      string    `form:"brand_id"`      // VARCHAR2(15)
	Status       string    `form:"status"`        // VARCHAR2(20), default 'Active'
}

type ProductView struct {
	RowID            string
	Name             string
	Description      string
	Price            string
	Amount           string
	ProdSize         string
	ProdLength       string
	ProdMaterial     string
	ProdColor        string
	ImageSrc         string
	Barcode          string
	CategoryName     string
	BrandName        string
	Status           string
	DeleteProductUrl string
}

type Category struct {
	RowID       string    `form:"row_id"`      // VARCHAR2(15)
	Created     time.Time `form:"created"`     // TIMESTAMP(6)
	CreatedBy   string    `form:"created_by"`  // VARCHAR2(400)
	LastUpd     time.Time `form:"last_upd"`    // TIMESTAMP(6)
	LastUpdBy   string    `form:"last_upd_by"` // VARCHAR2(400)
	Name        string    `form:"name"`        // VARCHAR2(400)
	Description string    `form:"description"` // VARCHAR2(4000)
	ParentId    string    `form:"parent_id"`   // NUMBER

}

type CategoryView struct {
	RowID              string
	Name               string
	Description        string
	ParentCategoryName string
	DeleteCategoryUrl  string
}

type Brand struct {
	RowID            string    `form:"row_id"`           // VARCHAR2(15)
	Created          time.Time `form:"created"`          // TIMESTAMP(6)
	CreatedBy        string    `form:"created_by"`       // VARCHAR2(400)
	LastUpd          time.Time `form:"last_upd"`         // TIMESTAMP(6)
	LastUpdBy        string    `form:"last_upd_by"`      // VARCHAR2(400)
	BrandName        string    `form:"brandname"`        // VARCHAR2(400)
	BrandCountry     string    `form:"brandcountry"`     // VARCHAR2(400)
	FullDescription  string    `form:"fulldescription"`  // VARCHAR2(4000)
	ShortDescription string    `form:"shortdescription"` // VARCHAR2(4000)
	BrandLogo        string    `form:"brandlogo"`        // VARCHAR2(4000)
}

type BrandView struct {
	RowID            string
	BrandName        string
	BrandCountry     string
	FullDescription  string
	ShortDescription string
	BrandLogo        string
	DeleteBrandUrl   string
}

type Warehouse struct {
	RowID            string    `form:"row_id"`           // VARCHAR2(15)
	Created          time.Time `form:"created"`          // TIMESTAMP(6)
	CreatedBy        string    `form:"created_by"`       // VARCHAR2(400)
	LastUpd          time.Time `form:"last_upd"`         // TIMESTAMP(6)
	LastUpdBy        string    `form:"last_upd_by"`      // VARCHAR2(400)
	WarehouseName    string    `form:"warehouse_name"`    // VARCHAR2(400)
	WarehouseManager string    `form:"warehouse_manager"` // VARCHAR2(400)
	WarehouseAddress string    `form:"warehouse_address"` // VARCHAR2(4000)
	NumberWorkers    string    `form:"number_workers"`    // VARCHAR2(4000)
	NumberProducts   string    `form:"number_products"`   // VARCHAR2(4000)
	Status           string    `form:"status"`           // VARCHAR2(4000)
	WarehouseImg     string    `form:"warehouse_img"`     // VARCHAR2(4000)
}

type WarehouseView struct {
	RowID              string
	WarehouseName      string
	WarehouseManager   string
	WarehouseAddress   string
	NumberWorkers      string
	NumberProducts     string
	Status             string
	WarehouseImg       string
	DeleteWarehouseUrl string
}
