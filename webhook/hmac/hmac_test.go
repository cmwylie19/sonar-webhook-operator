package hmac

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// CheckMAC verifies hash checksum
func TestCheckMAC(t *testing.T) {
	secret := []byte("secret")
	msg := []byte("Test CheckMAC")
	mac := hmac.New(sha256.New, secret)
	mac.Write(msg)
	signature := mac.Sum(nil)

	shouldBeTrue := CheckMAC(msg, signature, secret)

	// Incorrect signature
	shouldBeFalse := CheckMAC(msg, signature, []byte("wrong"))

	assert.Equal(t, shouldBeTrue, true, "should be equal")
	assert.Equal(t, shouldBeFalse, false, "should be equal")
}

func TestSign(t *testing.T) {
	secret := []byte("secret")
	msg := []byte("Test CheckMAC")
	mac := hmac.New(sha256.New, secret)
	mac.Write(msg)
	signature := mac.Sum(nil)

	resultSignature := Sign(msg, secret)
	assert.Equal(t, signature, resultSignature, "should be equal")
}

type SonarResult struct {
	Status string `json:"status"`
}

func TestValidate(t *testing.T) {
	// Request Body
	status := SonarResult{
		Status: "Success",
	}

	// Secret Key
	secret := []byte("secret")

	// Marshal the status
	json_data, err := json.Marshal(status)
	if err != nil {
		fmt.Println("could not unmarshal data")
	}
	encodedHash := "a00b02bfa6721f4218fd60aac2d2741df08abce8f562963da65121d6d39d69f5"
	shouldBeNil := Validate(json_data, encodedHash, string(secret))
	shouldNotBeNil := Validate(json_data, encodedHash, "testing")
	assert.Nil(t, shouldBeNil)
	assert.NotNil(t, shouldNotBeNil)
}
