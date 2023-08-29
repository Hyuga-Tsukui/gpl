package ex02

import (
	"crypto/sha256"
	"crypto/sha512"
	"flag"
	"fmt"
)

const (
	SHA256 = "SHA256"
	SHA384 = "SHA384"
	SHA512 = "SHA512"
)

func Hash(d, alg string) (string, error) {
	if d == "" {
		return "", fmt.Errorf("empty string")
	}

	switch alg {
	case SHA256:
		return fmt.Sprintf("%x", sha256.Sum256([]byte(d))), nil
	case SHA384:
		return fmt.Sprintf("%x", sha512.Sum384([]byte(d))), nil
	case SHA512:
		return fmt.Sprintf("%x", sha512.Sum512([]byte(d))), nil
	default:
		return "", fmt.Errorf("unknown algorithm: %s", alg)
	}
}

func main() {
	var (
		alg = flag.String("alg", SHA256, "hash algorithm (SHA256, SHA384, SHA512) default: SHA256")
		d   = flag.String("d", "", "string to hash")
	)

	flag.Parse()
	h, err := Hash(*d, *alg)
	if err != nil {
		fmt.Printf("error: %v\n", err)
		return
	}
	fmt.Printf("%s\n", h)
}
