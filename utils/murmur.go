package utils

import (
	"fmt"

	"github.com/spaolacci/murmur3"
)

func joinString(str string, did uint) string {
	return fmt.Sprintf("%s%d", str, did)
}

func GetMurmur32(str string, did uint) uint32 {
	return murmur3.Sum32([]byte(joinString(str, did)))
}

func GetMurmur64(str string, did uint) uint64 {
	return murmur3.Sum64([]byte(joinString(str, did)))
}
