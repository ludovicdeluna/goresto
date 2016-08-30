package goresto_test

import (
	"fmt"
	"github.com/ludovicdeluna/goresto"
)

// This simple sample introduce how use the package in general.
// We can see also that a NewResto return a pointer to the Resto object
func Example_createNewResto() {
	myRestaurant := goresto.New()
	fmt.Print(myRestaurant.Name)
	// Output: The Restaurant
}

// Show that a new Resto as empty servers
func ExampleNew() {
	myRestaurant := goresto.New()
	fmt.Print(myRestaurant.Name)
	// Output: The Restaurant
}

// Show the numbers of servers in the current restaurant. Here, because we
// add a Server (which wait Client in background), we need to defer CloseMe().
// CloseMe will stop any Server's go-routines and close channels.
func ExampleResto_AddServer() {
	myRestaurant := goresto.New()
	defer myRestaurant.CloseMe()
	myRestaurant.AddServer()
	fmt.Print(len(myRestaurant.Servers))
	// Output: 1
}

// We can only get a client if at least one server is in the restaurant
func ExampleResto_GetClient() {
	myRestaurant := goresto.New()
	client := goresto.NewClient()
	err := myRestaurant.GetClient(client)
	fmt.Println(err)
	// Output: No server into the restaurant
}
