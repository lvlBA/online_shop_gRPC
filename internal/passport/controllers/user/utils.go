package user

import "crypto/sha512"

func toHash(data string) string {
	hash := sha512.Sum512([]byte(data))

	return string(hash[:])
}
