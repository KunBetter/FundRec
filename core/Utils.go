package core

import "github.com/spaolacci/murmur3"

func Hash(cont string) uint64 {
	return murmur3.Sum64([]byte(cont))
}
