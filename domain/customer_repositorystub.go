package domain

type CustomerRepositoryStub struct {
	customers []Customer
}

func (s CustomerRepositoryStub) FindAll() ([]Customer, error) {
	return s.customers, nil
}

func NewCustomerRepositoryStub() CustomerRepositoryStub {
	customer := []Customer {
		{ID: "1", Name: "Wander", City: "BH", Zipcode: "123", DateofBirth: "1990-09-19", Status: "active"},
		{ID: "2", Name: "Erika", City: "BH", Zipcode: "123", DateofBirth: "1989-07-17", Status: "active"},
	}

	return CustomerRepositoryStub{customers: customer}
}
