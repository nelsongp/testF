package getTestEndpoint

import (
	kitlog "github.com/go-kit/kit/log"
	"github.com/nelsongp/testF/internal/testEndpoint"
)

type testEndpointService struct {
	log kitlog.Logger
}

func NewTestEndpointService(log kitlog.Logger) *testEndpointService {
	return &testEndpointService{log: log}
}

func (t *testEndpointService) ResponseTestService(name string) (testEndpoint.TestResponse, error) {
	return testEndpoint.TestResponse{
		Name:     "Test",
		LastName: "Prueba",
	}, nil
}
