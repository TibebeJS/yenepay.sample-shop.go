package models

import "github.com/revel/revel"

type Book struct {
	BookId     int
	Price      float64
	Title      string
	CoverImage string
}

func (book *Book) Validate(v *revel.Validation) {
	v.Check(book.Title,
		revel.Required{},
		revel.MaxSize{
			Max: 50,
		},
		revel.MinSize{
			Min: 4,
		},
	)

	v.Check(book.CoverImage,
		revel.Required{},
		revel.MaxSize{
			Max: 50,
		},
		revel.MinSize{
			Min: 4,
		},
	)

}
