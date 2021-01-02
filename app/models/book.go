package models

type Book struct {
	BookId int
	Price  float64
	Title  string
	// CoverImage base64.Encoding
	CoverImage string
}
