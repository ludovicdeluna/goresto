package goresto

import (
	"testing"
	"time"
)

var msg string = "%s :\nGot:  %v\nWant: %v"

func TestRestoAddServer(t *testing.T) {
	resto := New()
	defer resto.CloseMe()
	if got, want := len(resto.Servers), 0; got != want {
		t.Errorf(msg, "New resto has no servers", got, want)
	}
	resto.AddServer()
	if got, want := len(resto.Servers), 1; got != want {
		t.Errorf(msg, "Add a Server to the Resto.servers slide", got, want)
	}
	if got, want := resto.Servers[0].Busy, false; got != want {
		t.Errorf(msg, "A new Server is NOT busy (false)", got, want)
	}
}

func TestRestoAddClient(t *testing.T) {
	resto := New()
	defer resto.CloseMe()
	client := NewClient()
	client2 := NewClient()
	if got, want := resto.GetClient(client), "No server into the restaurant"; got == nil {
		t.Errorf(msg, "A client is comming. No server should be in the resto", got, want)
	}
	resto.AddServer()
	resto.GetClient(client) // First client
	time.Sleep(20 * time.Millisecond)
	if got, want := resto.Servers[0].Busy, true; got != want {
		t.Errorf(msg, "Server is cooking for the Client and should be busy", got, want)
	}
	if got, want := <-resto.Billables, client; got != want {
		t.Errorf(msg, "Server finished with Client. Client should be billable", got, want)
	}
	resto.GetClient(client2) // Second client
	time.Sleep(20 * time.Millisecond)
	if got, want := <-resto.Billables, client2; got != want {
		t.Errorf(msg, "Server finished with Client2. Client2 should be billable", got, want)
	}
}

func TestRestoClose(t *testing.T) {
	resto := New()
	defer resto.CloseMe()
	resto.AddServer()
	if got, want := resto.Servers[0].Off, false; got != want {
		t.Errorf(msg, "Server is not off until the chan is closed", got, want)
	}
	resto.CloseMe()
	time.Sleep(20 * time.Millisecond)
	if got, want := resto.Servers[0].Off, true; got != want {
		t.Errorf(msg, "Resto close chan, so all servers goes off", got, want)
	}
	client := NewClient()
	if got, want := resto.GetClient(client), "Restaurant is closed"; got == nil {
		t.Errorf(msg, "Restaurant do not accept more client if it's closed", got, want)
	}
}
