package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
)

func main() {
	http.HandleFunc("/", foo)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)
}

func foo(w http.ResponseWriter, req *http.Request) {
	var s string
	fmt.Println(req.Method)
	if req.Method == http.MethodPost {

		//open
		f, h, err := req.FormFile("q")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		defer f.Close()

		// for your information
		fmt.Println("\nfile:", f, "\nheader:", h, "\nerr", err)

		//read
		bs, err := ioutil.ReadAll(f)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		s = string(bs)
	}

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	io.WriteString(w, `
		<form method="POST" enctype="multipart/form-data">
		<input type="file" name="q">
		<input type="submit">
		</form>
		<br>`+s)
}

/*
Run Result:
POST
GET
POST

file: {0xc0420c45a0}
header: &{keys to success.txt map[Content-Disposition:[form-data; name="q"; filename="keys to success.txt"] Content-Type:[text/plain]] [84 104 101 32 115 117 99 99 101 115 115 102 117 108 32 119 97 114 114 105 111
114 32 105 115 32 116 104 101 32 97 118 101 114 97 103 101 32 109 97 110 44 32 119 105 116 104 32 108 97 115 101 114 45 108 105 107 101 32 102 111 99 117 115 46 13 10 13 10 77 121 32 115 117 99 99 101 115 115 44 32 112 97 114 116 32 111 102 32 105 116 32 99 101 114 116 97 105 110 108 121 44 32 105 115 32 116 104 97 116 32 73 32 104 97 118 101 32 102 111 99 117 115 101 100 32 105 110 32 111 110 32 97 32 102 101 119 32 116 104 105 110 103 115 46] }
err <nil>
*/