package main

import (
	"encoding/base64"
	"fmt"
)

func main() {
	s := "People who succeed have momentum. The more they succeed, the more they want to succeed,	and the more they find a way to succeed. Similarly, when someone is failing, the tendency is	to get on a downward spiral that can even become a self-fulfilling prophecy"

	encodeStd :=
		"ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789+/"
	s64 := base64.NewEncoding(encodeStd).EncodeToString([]byte(s))

	fmt.Println(len(s))
	fmt.Println(len(s64))
	fmt.Println(s)
	fmt.Println(s64)
}

/*
Run Result:
256
344
People who succeed have momentum. The more they succeed, the more they want to succeed,   and the more they find a way to succeed. Similarly, when someone is failing, the tendency is      to get on a downward spiral that can even become a self-fulfilling prophecy
UGVvcGxlIHdobyBzdWNjZWVkIGhhdmUgbW9tZW50dW0uIFRoZSBtb3JlIHRoZXkgc3VjY2VlZCwgdGhlIG1vcmUgdGhleSB3YW50IHRvIHN1Y2NlZWQsCWFuZCB0aGUgbW9yZSB0aGV5IGZpbmQgYSB3YXkgdG8gc3VjY2VlZC4gU2ltaWxhcmx5LCB3aGVuIHNvbWVvbmUgaXMgZmFpbGluZywgdGhlIHRlbmRlbmN5IGlzCXRvIGdldCBvbiBhIGRvd253YXJkIHNwaXJhbCB0aGF0IGNhbiBldmVuIGJlY29tZSBhIHNlbGYtZnVsZmlsbGluZyBwcm9waGVjeQ==
*/
