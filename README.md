# SHA-1

## Short information
Secure Hash Algorithm 1 is a cryptographic hashing algorithm for messages of arbitrary length; it generates a 160-bit value (equal to 20 bytes). Previously, it was recommended as the main one for the state in the United States

## Structure description
- `forTesting/checkTimeFunc.go` - is used to compare the library implementation of the hashing algorithm and my implementation
- `operations/operations.go` - intermediate operations used in the sha-1 algorithm
- `sha1func/sha1.go` - implementation of the hashing algorithm
- `main.go` - in the main, the CheckTime function from checkTimeFunc is called 5 times to compare its own implementation with the library version
- `main_test.go` - main_test.go checks the correctness of the algorithm by comparing hash values with the finished implementation

## Example of use
```
package main

import (
	"Project/sha1func"
	"fmt"
)

func main() {
	message := "Hello, github!"
	hash := sha1func.Sha1hashing(message)
	fmt.Printf("Hash for string %s is: %x\n", message, hash)
}
```
## Note
<i>The U.S. National Institute of Standards and Technology, Inc. (NIST) stated that the SHA-1 hash algorithm is obsolete, insecure, and its use is not recommended; the code was written for research purposes only<i>
