package forTesting

import (
	"Project/sha1func"
	"crypto/sha1"
	"fmt"
	"time"
)

func CheckTime(a string) {

	start := time.Now()

	h := sha1.New()
	h.Write([]byte(a))
	hashed := h.Sum(nil)
	fmt.Print(hashed, "\n")

	end := time.Now()

	start1 := time.Now()

	hash := sha1func.Sha1hashing(a)
	fmt.Print(hash, "\n")

	end1 := time.Now()

	durationLib := end.Sub(start)
	durationMySha := end1.Sub(start1)

	fmt.Printf("Lib implementation: %v My implementation: %v\n", durationLib, durationMySha)
}
