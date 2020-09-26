package util

import (
	"math/rand"
	"time"
	"log"

	"github.com/matoous/go-nanoid"
)

const charset = "_-0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"

var seededRand *rand.Rand = rand.New(rand.NewSource(time.Now().UnixNano()))

func StringWithCharset(length int, charset string) string {
	b := make([]byte, length)
	for i := range b {
		b[i] = charset[seededRand.Intn(len(charset))]
	}

	return string(b)
}

func GenerateTag() string {
	tag, err := gonanoid.Generate(charset, 6)
	if err != nil {
		log.Print(err)
	}
	return tag
	// return StringWithCharset(3, charset)
}
