package client

const (
	testEndpoint = "http://localhost:8080"
	testVersion  = "v1"
)

// newTestClient create a new tokend client for testing
func newTestClient() (*Client, error) {
	return Dial(testEndpoint, testVersion)
}
