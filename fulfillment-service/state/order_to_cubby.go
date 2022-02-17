package state

import (
	"hash/fnv"
	"strconv"
)

func MapOrderToCubby(orderID string, times, cubbiesCount uint32) string {
	result := hash(orderID)
	for times > 1 {
		result = hash(strconv.Itoa(int(result)))
		times--
	}

	return strconv.Itoa(int(result%cubbiesCount) + 1)
}

func hash(s string) uint32 {
	h := fnv.New32a()
	h.Write([]byte(s))
	return h.Sum32()
}
