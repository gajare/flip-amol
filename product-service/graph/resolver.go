package graph

import "product-service/internal/product/service"

type Resolver struct {
	ProductService service.ProductService
}
