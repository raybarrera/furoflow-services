package crypto

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
)

//SecretHash takes parameters first and second, combines them, then hashes them using the secret and returns the resulting base64-encoded string.
func SecretHash(first, second, secret string) string {
	mac := hmac.New(sha256.New, []byte(secret))
	mac.Write([]byte(first + second))
	return base64.StdEncoding.EncodeToString(mac.Sum(nil))
}
