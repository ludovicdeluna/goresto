package goresto

import "errors"

// Represent a restaurant
type Resto struct {
	Servers   []*Server // Slide of Server's pointers
	Name      string    // The name of restaurant
	Waiters   Waiters   // Clients waiting
	Billables Waiters   // Clients ready to pay
	open      bool      // (private) Is Open ?
}

// Create a new restaurant with empty servers
func New() *Resto {
	return &Resto{
		Name:      "The Restaurant",
		Waiters:   NewWaiters(),
		Billables: NewWaiters(),
		open:      true,
	}
}

// Add a server to the restaurant, ready to serve gentles clients
func (r *Resto) AddServer() {
	if r.open == false {
		// If restaurant is closed, do not accept new server
		return
	}
	r.Servers = append(r.Servers, NewServer(r))
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
		r.Waiters <- clt
	}()
	return nil
}

// Close restaurant
func (r *Resto) CloseMe() {
	r.open = false
	close(r.Waiters) // Other go-routines using this channel will exit
}

// Waiters, a channel who accept Client'pointers
type Waiters chan *Client

// Buile a Waiters channel
func NewWaiters() Waiters {
	return make(Waiters)
}
