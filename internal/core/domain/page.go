package domain

type Page struct {
	Page uint32 //skip
	Size uint32 //limit
}

func NewDefaultPage() Page {
	return Page{
		Page: 1,
		Size: 10,
	}
}

func NewPage(page, size uint32) Page {
	if page < 1 {
		page = 1
	}
	if size < 1 || size > 100 {
		size = 100
	}

	return Page{
		Page: page,
		Size: size,
	}
}
