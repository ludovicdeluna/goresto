package goresto

// A client, want eat something
// For this example, what is exactly a Client is not interesting.
// It's simply "an object"
type Client int

// Build a new client
func NewClient() *Client {
	var c Client
	return &c
}
