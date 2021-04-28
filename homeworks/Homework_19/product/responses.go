package product

type TemplateResponse struct {
	Title       string
	Description string
	Price       int
	Error       string
}

type ListProductResponse struct {
	ProductField Product
}
