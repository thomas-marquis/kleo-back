package transactions

type Factory struct {
}

func NewFactory() (*Factory, error) {
	return &Factory{}, nil
}

func (f *Factory) Transaction() (*Transaction, error) {
	return &Transaction{}, nil
}
