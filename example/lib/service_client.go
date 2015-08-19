package lib

type ServiceClient interface {
	DoFancyStuff() string
}

type MockServiceClient struct {
	private string
}

func NewMockServiceClient(s string) *MockServiceClient {
	return &MockServiceClient{s}
}

func (c *MockServiceClient) DoFancyStuff() string {
	return c.private
}
