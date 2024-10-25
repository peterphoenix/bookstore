package repository

const (
	bookTableName       = "book"
	authorTableName     = "author"
	bookAuthorTableName = "book_author"
)

var (
	bookTable = struct {
		ID          string
		Title       string
		Price       string
		ISBN        string
		ImageURL    string
		Description string
	}{
		ID:          "id",
		Title:       "title",
		Price:       "price",
		ISBN:        "isbn",
		ImageURL:    "image_url",
		Description: "description",
	}

	authorTable = struct {
		ID   string
		Name string
	}{
		ID:   "id",
		Name: "name",
	}

	bookAuthorTable = struct {
		BookID   string
		AuthorID string
	}{
		BookID:   "book_id",
		AuthorID: "author_id",
	}
)
