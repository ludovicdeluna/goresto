package goresto

import "time"

// One server can cook a meal for one - and only one - client
type Server struct {
	Busy   bool    // True: Server is being cooked meal
	Client *Client // Pointer to a client object
	Off    bool    // True: Server has stopped his work and gone at home
}

// Create a new Server, free to cook meals
func NewServer(in, out Waiters) *Server {
	server := Server{}
	go server.WaitClient(in, out)
	return &server
}

// Server waits a client to serve
func (srv *Server) WaitClient(in, out Waiters) {
	for clt := range in {
		srv.Busy = true
		srv.Client = clt
		// Hard work simulation
		time.Sleep(100 * time.Millisecond)
		// Work is done, Server will wait another client
		srv.Client = nil
		srv.Busy = false
		out <- clt
	}
	// When the channel "in" is closed, escape the for loop.
	// Set the current server "Off", he return to his home.
	srv.Off = true
}
