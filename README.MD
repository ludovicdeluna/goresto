Restaurant exercice with Go
---------------------------

This is a simple exercice where the objective is to use collaborative
goroutines using channels and simulate the situation of a restaurant.

In this case, a restaurant will engage servers to cook meals and
will wait for clients. When a client is coming, a Server cook a meal. When
it's done (we don't wait for customer to eat), the client goes to pay.
