package testEndpoint

type TestEndpoint interface {
	ResponseTestService(name string) (TestResponse, error)
}

type TestResponse struct {
	Name     string `json:"name"`
	LastName string `json:"lastName"`
}
