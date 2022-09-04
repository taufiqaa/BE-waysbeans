package productdto

type ProductRequest struct {
	Title       string `json:"title" form:"name" gorm:"type: varchar(255)"`
	Price       int    `json:"price" form:"price" gorm:"type: int"`
	Stock       int    `json:"stock" form:"stock" gorm:"type: int"`
	Description string `json:"description" form:"description" gorm:"type: varchar(255)"`
	Image       string `json:"image" form:"image" gorm:"type: varchar(255)"`
	// Qty   int    `json:"qty" form:"qty" gorm:"type: int"`
	// CategoryID int    `json:"category_id" form:"category_id" gorm:"type: int"`
}

type UpdateProductRequest struct {
	Title       string `json:"title" form:"name" gorm:"type: varchar(255)"`
	Price       int    `json:"price" form:"price" gorm:"type: int"`
	Stock       int    `json:"stock" form:"stock" gorm:"type: int"`
	Description string `json:"description" form:"description" gorm:"type: varchar(255)"`
	Image       string `json:"image" form:"image" gorm:"type: varchar(255)"`
	// Qty   int    `json:"qty" form:"qty" gorm:"type: int"`
	// CategoryID int    `json:"category_id" form:"category_id" gorm:"type: int"`
}

type UpdateProductStock struct {
	Stock int `json:"stock" gorm:"type: int"`
}

type ProductResponse struct {
	Title       string `json:"title" form:"title"`
	Price       int    `json:"price" form:"price"`
	Stock       int    `json:"stock" form:"stock" gorm:"type: int"`
	Description string `json:"description" form:"description" gorm:"type: varchar(255)"`
	Image       string `json:"image" form:"image"`
}
