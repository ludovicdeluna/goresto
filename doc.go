// License MIT

/*
	This is a simple example of using Go routines to make a similar approach to
	collaborative process seen in an Erlang program (internal challange into
	our company). This function in the same way but using Go.

	The test file is here to see how it's work, but it's more interesting to
	dive into the code.

	For this simple code, I choose to use pointers for all objects and open
	all methods (no private, to be clear to understand)
*/
package goresto
