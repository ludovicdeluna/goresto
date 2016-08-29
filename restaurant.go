package goresto

import "errors"

// Represent a restaurant
type Resto struct {
	Servers   []*Server    // Slide of Server's pointers
	Name      string       // The name of restaurant
	Billables chan *Client // Clients ready to pay
	waiters   chan *Client // (private) Clients waiting
	open      bool         // (private) Is Open ?
}

// Create a new restaurant with empty servers
func New() *Resto {
	return &Resto{
		Name:      "The Restaurant",
		waiters:   make(chan *Client),
		Billables: make(chan *Client),
		open:      true,
	}
}

// Add a server to the restaurant, ready to serve gentles clients
func (r *Resto) AddServer() {
	if r.open == false {
		// If restaurant is closed, do not accept new server
		return
	}
	r.Servers = append(r.Servers, NewServer(r.waiters, r.Billables))
}

// Add a client to the Resto
func (r *Resto) GetClient(clt *Client) error {
	switch {
	case r.open == false:
		// If restaurant is closed, do not accept new client
		return errors.New("Restaurant is closed")
	case len(r.Servers) == 0:
		// If there is no server into the restaurant, can't accept client
		return errors.New("No server into the restaurant")
	}
	go func() {
		// Do not block the process here if all servers are busy.
		// With goroutine, simply add the client when chan is read (in background)
		r.waiters <- clt
	}()
	return nil
}

// Close restaurant
func (r *Resto) CloseMe() {
	r.open = false
	close(r.waiters) // Other go-routines using this channel will exit
}
