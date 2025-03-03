package product_management

import "time"

// Product represents the n_prod_product table
type Product struct {
	RowID       string    // Primary Key
	Created     time.Time // Default: CURRENT_TIMESTAMP
	CreatedBy   string    // User who created the record
	LastUpd     time.Time // Default: CURRENT_TIMESTAMP
	LastUpdBy   string    // User who last updated the record
	Name        string    // Name of the product (NOT NULL)
	Description string    // Detailed description (CLOB equivalent)
	SKU         string    // Unique Stock Keeping Unit (NOT NULL)
	Barcode     string    // Unique barcode identifier
	CategoryID  string    // Foreign key to product category
	BrandID     string    // Foreign key to product brand
	Weight      float64   // Product weight in kg (DECIMAL(10,2))
	Dimensions  string    // Product dimensions (LxWxH)
	Status      string    // Status (CHECK: 'Active', 'Inactive', 'Discontinued')
}
