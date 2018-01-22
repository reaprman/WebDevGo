package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/bar", bar)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)
}

func index(w http.ResponseWriter, req *http.Request) {
	ctx := req.Context()

	ctx = context.WithValue(ctx, "userID", 777)
	ctx = context.WithValue(ctx, "fname", "Bond")

	results := dbAccess(ctx)
	fmt.Println(w, results)
}

func dbAccess(ctx context.Context) int {
	uid := ctx.Value("userID").(int)
	return uid
}

func bar(w http.ResponseWriter, req *http.Request) {
	ctx := req.Context()
	log.Println(ctx)
	fmt.Fprintln(w, ctx)
}

/*
per request variables
good candidate for putting into context

Run Result:
&{0xc04203a6e0 0xc042030400 {} 0x4b2140 false false false false 0xc04200a280 {0xc0420d6000 map[] false false} map[] false 0 -1 0 false false [] 0 [0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0] [0 0
0 0 0 0 0 0 0 0] 0xc04200e1c0 0} 777

2018/01/22 01:20:32 context.Background.WithValue(&http.contextKey{name:"http-server"}, &http.Server{Addr:":8080", Handler:http.Handler(nil), TLSConfig:(*tls.Config)(0xc042054180), ReadTimeout:0, ReadHeaderTimeout:0, WriteTimeout:0, IdleTimeout:0, MaxHeaderBytes:0, TLSNextProto:map[string]func(*http.Server, *tls.Conn, http.Handler){"h2":(func(*http.Server, *tls.Conn, http.Handler))(0x5ed4b0)}, ConnState:(func(net.Conn,
http.ConnState))(nil), ErrorLog:(*log.Logger)(nil), disableKeepAlives:0, inShutdown:0, nextProtoOnce:sync.Once{m:sync.Mutex{state:0, sema:0x0}, done:0x1}, nextProtoErr:error(nil), mu:sync.Mutex{state:0, sema:0x0}, listeners:map[net.Listener]struct {}{http.tcpKeepAliveListener{TCPListener:(*net.TCPListener)(0xc04204e030)}:struct {}{}}, activeConn:map[*http.conn]struct {}{(*http.conn)(0xc04203a6e0):struct {}{}}, doneChan:(chan struct {})(nil)}).WithValue(&http.contextKey{name:"local-addr"}, &net.TCPAddr{IP:net.IP{0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0}, Port:8080, Zone:""}).WithCancel.WithCancel
*/
