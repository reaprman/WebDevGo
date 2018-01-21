package main

import (
	"crypto/hmac"
	"crypto/sha256"
	"fmt"
	"io"
)

func main() {
	c := getCode("test@example.com")
	fmt.Println(c)
	c = getCode("test@exampl.com")
	fmt.Println(c)
}

func getCode(s string) string {
	h := hmac.New(sha256.New, []byte("ourkey"))
	io.WriteString(h, s)
	return fmt.Sprintf("%x", h.Sum(nil))
}

/*
Run Result:
453159a62580804892cc90a27dfb8bfdf2309107336445dcfd0186674111ee71
7e35ec99f8dfdd96d54c5185a81b25f662e2f786c36e737d4b65f903fe4862bc
*/
