package domain


type CustomerRepositoryStub struct {
	customers []Customer
}

func (s CustomerRepositoryStub) FindAll() ([]Customer, error) {
	return s.customers, nil
}

func NewCustomerRepositoryStub() CustomerRepositoryStub {
	customers := []Customer {
		{Id: "1001", Name: "Víctor", City: "New Delhi", Zipcode: "28212", DateOfBirth: "21/05/1992", Status: "1"},
		{Id: "1002", Name: "Carlos", City: "Cancún", Zipcode: "43444", DateOfBirth: "25/05/1992", Status: "1"},
	}

	return CustomerRepositoryStub{customers}
}