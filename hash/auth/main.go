package main

import (
	"fmt"
	"hash/crc32"
)

func main() {
	hash := Hash("tencent")
	fmt.Println(hash)
}

func Hash(password string) uint64 {
	v := int64(crc32.ChecksumIEEE([]byte(password)))
	if v >= 0 {
		return uint64(v)
	}
	if -v >= 0 {
		return uint64(-v)
	}
	// v == MinInt
	return 0
}
