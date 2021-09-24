package hmac

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
)

// CheckMAC verifies hash checksum
func CheckMAC(message, messageMAC, key []byte) bool {
	mac := hmac.New(sha256.New, key)
	mac.Write(message)
	expectedMAC := mac.Sum(nil)

	return hmac.Equal(messageMAC, expectedMAC)
}

// Sign a message with the key and return bytes.
// Note: for human readable output see encoding/hex and
// encode string functions.
func Sign(message, key []byte) []byte {
	mac := hmac.New(sha256.New, key)
	mac.Write(message)
	signed := mac.Sum(nil)
	return signed
}

// Validate validate an encodedHash taken
// from GitHub via X-Hub-Signature HTTP Header.
func Validate(bytesIn []byte, encodedHash string, secretKey string) error {
	var validated error

	if len(encodedHash) > 5 {
		messageMACBuf, _ := hex.DecodeString(encodedHash)
		res := CheckMAC(bytesIn, []byte(messageMACBuf), []byte(secretKey))
		if !res {
			validated = fmt.Errorf("invalid message digest or secret")
		}
	} else {
		return fmt.Errorf("invalid encodedHash, should have at least 5 characters")
	}

	return validated
}

func init() {}
