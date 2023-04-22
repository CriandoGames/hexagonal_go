package application

type ProductService struct {
	ProductRepository ProductPersistenceInterface
}

func (p *ProductService) Get(id string) (ProductInterface, error) {

	product, err := p.ProductRepository.Get(id)
	if err != nil {
		return nil, err
	}

	return product, nil

}
