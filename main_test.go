package main

import (
	"Project/sha1func"
	"crypto/sha1"
	"testing"
)

func TestHash(t *testing.T) {
	message := "Hello, world!"

	mySha1 := sha1func.Sha1hashing(message)

	libSha1 := sha1.New()
	libSha1.Write([]byte(message))
	libRealisation := libSha1.Sum(nil)

	// Перетворення []byte в [20]byte
	var libRealisationFixed [20]byte
	copy(libRealisationFixed[:], libRealisation)

	if mySha1 != libRealisationFixed {
		t.Errorf("Array mySha1: %x, array from library: %x", mySha1, libRealisationFixed)
	}
}
