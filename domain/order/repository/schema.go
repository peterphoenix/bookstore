package repository

const (
	orderTableName     = "\"order\""
	orderItemTableName = "order_item"
)

var (
	orderTable = struct {
		ID              string
		UserID          string
		TotalAmount     string
		Status          string
		ShippingAddress string
		City            string
		State           string
		ZipCode         string
		Country         string
		CreatedAt       string
		UpdatedAt       string
	}{
		ID:              "id",
		UserID:          "user_id",
		TotalAmount:     "total_amount",
		Status:          "status",
		ShippingAddress: "shipping_address",
		City:            "city",
		State:           "state",
		ZipCode:         "zip_code",
		Country:         "country",
		CreatedAt:       "created_at",
		UpdatedAt:       "updated_at",
	}

	orderItemTable = struct {
		OrderID  string
		BookID   string
		Quantity string
		Price    string
	}{
		OrderID:  "order_id",
		BookID:   "book_id",
		Quantity: "quantity",
		Price:    "price",
	}
)
