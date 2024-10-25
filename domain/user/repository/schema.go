package repository

const (
	userTableName = "\"user\""
)

var (
	userTable = struct {
		ID          string
		Name        string
		Email       string
		Password    string
		PhoneNumber string
		Address     string
		City        string
		State       string
		ZipCode     string
		Country     string
		CreatedAt   string
		UpdatedAt   string
	}{
		ID:          "id",
		Name:        "name",
		Email:       "email",
		Password:    "password",
		PhoneNumber: "phone_number",
		Address:     "address",
		City:        "city",
		State:       "state",
		ZipCode:     "zip_code",
		Country:     "country",
		CreatedAt:   "created_at",
		UpdatedAt:   "updated_at",
	}
)
