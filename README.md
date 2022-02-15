# sort-system-v1

This is the first version of the [SORT system](https://www.youtube.com/watch?v=BQDliV7w7_8).

## How to run the project
 * `make grpc-compile` to generate all grpc-related files in the `gen/` folder
 * Enter `sorting-service` and type `go run *.go`
 * Enter `fulfillment-service` and type `go run *.go`
 * use `sh scripts/seed-orders.sh` to send orders to send items to the sorting service and fulfill orders with the fulfillment service

Functionalities:
 * LoadItems - loads an input array of items in the service. E.g. `[{"code": "123", "label": "tomatoes"},{"code": "456", "label": "cucumber"}]`
 * SelectItem -> Choose an item at random from the remaining ones in the array. E.g. choose "tomatoes" at random && remove item from existing array
 * MoveItem -> Move the selected item in the input cubby.
 * LoadOrders -> Load orders and fulfill them
 * GetOrderFulfillmentStatusById -> Get Status of order by ID
 * GetAllOrdersFulfillmentStatus -> Get Status of all orders
 
Return an error in any of the following cases:
 * SelectItem is invokes but there are no items in input bin
 * MoveItem is invoked but no item is selected yet
 * SelectItem is invoked when an item is already selected
