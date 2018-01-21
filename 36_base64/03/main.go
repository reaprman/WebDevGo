package main

import (
	"encoding/base64"
	"fmt"
	"log"
)

func main() {
	s := "People who succeed have momentum. The more they succeed, the more they want to succeed,	and the more they find a way to succeed. Similarly, when someone is failing, the tendency is	to get on a downward spiral that can even become a self-fulfilling prophecy"

	s64 := base64.StdEncoding.EncodeToString([]byte(s))
	fmt.Println(s64)

	bs, err := base64.StdEncoding.DecodeString(s64 + "s")
	if err != nil {
		log.Fatalln("I'm giiving her all she's got Captain!", err)
	}

	fmt.Println(string(bs))
}

/*
Example run twice with matching data and without
Run Result:
match:
UGVvcGxlIHdobyBzdWNjZWVkIGhhdmUgbW9tZW50dW0uIFRoZSBtb3JlIHRoZXkgc3VjY2VlZCwgdGhlIG1vcmUgdGhleSB3YW50IHRvIHN1Y2NlZWQsCWFuZCB0aGUgbW9yZSB0aGV5IGZpbmQgYSB3YXkgdG8gc3VjY2VlZC4gU2ltaWxhcmx5LCB3aGVuIHNvbWVvbmUgaXMgZmFpbGluZywgdGhlIHRlbmRlbmN5IGlzCXRvIGdldCBvbiBhIGRvd253YXJkIHNwaXJhbCB0aGF0IGNhbiBldmVuIGJlY29tZSBhIHNlbGYtZnVsZmlsbGluZyBwcm9waGVjeQ==
People who succeed have momentum. The more they succeed, the more they want to succeed,   and the more they find a way to succeed. Similarly, when someone is failing, the tendency is      to get on a downward spiral that can even become a self-fulfilling prophecy

not match:
UGVvcGxlIHdobyBzdWNjZWVkIGhhdmUgbW9tZW50dW0uIFRoZSBtb3JlIHRoZXkgc3VjY2VlZCwgdGhlIG1vcmUgdGhleSB3YW50IHRvIHN1Y2NlZWQsCWFuZCB0aGUgbW9yZSB0aGV5IGZpbmQgYSB3YXkgdG8gc3VjY2VlZC4gU2ltaWxhcmx5LCB3aGVuIHNvbWVvbmUgaXMgZmFpbGluZywgdGhlIHRlbmRlbmN5IGlzCXRvIGdldCBvbiBhIGRvd253YXJkIHNwaXJhbCB0aGF0IGNhbiBldmVuIGJlY29tZSBhIHNlbGYtZnVsZmlsbGluZyBwcm9waGVjeQ==
2018/01/20 23:59:40 I'm giiving her all she's got Captain! illegal base64
data at input byte 344
exit status 1

*/
