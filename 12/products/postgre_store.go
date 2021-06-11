package products

type postgreProductStore struct {
	//element from postgr (CRUD)
}

func NewPostgreProductStore() (ProductStore, error) {
	//connection to postgre
	return &postgreProductStore{}, nil
}

func (p *postgreProductStore) Create(product *Product) (*Product, error) {
	return nil, nil
}

func (p *postgreProductStore) List() ([]Product, error) {
	return nil, nil
}

func (p *postgreProductStore) GetById(id string) (*Product, error) {
	return nil, nil
}
