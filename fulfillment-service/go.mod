module github.com/Emoto13/sort-system/fulfillment-service

go 1.16

replace github.com/Emoto13/sort-system/gen => ../gen

require (
	github.com/Emoto13/sort-system/gen v0.0.0-20210623104657-36fa702e85f3
	go.uber.org/atomic v1.9.0
	google.golang.org/grpc v1.38.0
)
