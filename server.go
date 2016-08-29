package goresto

import "time"

// One server can cook a meal for one - and only one - client
type Server struct {
	Busy    bool      // True: Server is being cooked meal
	Client  *Client   // Pointer to a client object
	Off     bool      // True: Server has stopped his work and gone at home
	suspend chan bool // Stop server waiting to write to Billables
}

// Create a new Server, free to cook meals
func NewServer(in <-chan *Client, out chan<- *Client) *Server {
	// Suspend chan is bufferized chan
	server := Server{suspend: make(chan bool)}
	// Wait clients and return the server pointer
	go server.waitClient(in, out)
	return &server
}

// (private) Server waits a client to serve
func (srv *Server) waitClient(in <-chan *Client, out chan<- *Client) {
	for clt := range in {
		srv.Busy = true
		srv.Client = clt
		// Hard work simulation
		time.Sleep(100 * time.Millisecond)
		// Work is done, Server will wait another client
		srv.Busy = false
		srv.Client = nil
		select {
		case <-srv.suspend: // If restaurant is being closed, Clients have a free meal :)
		case out <- clt: // Or return result if chan is ready to get data
		}
	}
	// When the channel "in" is closed, escape the for loop.
	// Set the current server "Off", he return to his home.
	srv.Off = true
}

// Suspend the server
func (srv *Server) SuspendMe() {
	close(srv.suspend)
}
