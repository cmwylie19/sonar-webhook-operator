package helper

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetMongoClient(t *testing.T) {
	os.Setenv("MONGO_URL", "mongodb://localhost:27017")
	_, err := GetMongoClient()
	assert.NoError(t, err)
}
