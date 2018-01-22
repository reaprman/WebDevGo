package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", index)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)
}

func index(w http.ResponseWriter, req *http.Request) {
	ctx := req.Context()
	log.Println(ctx)
	fmt.Fprintln(w, ctx)
}

/*
Run Result:
2018/01/22 01:12:42 context.Background.WithValue(&http.contextKey{name:"http-server"}, &http.Server{Addr:":8080", Handler:http.Handler(nil), TLSConfig:(*tls.Config)(0xc042054180), ReadTimeout:0, ReadHeaderTimeout:0, WriteTimeout:0, IdleTimeout:0, MaxHeaderBytes:0, TLSNextProto:map[string]func(*http.Server, *tls.Conn, http.Handler){"h2":(func(*http.Server, *tls.Conn, http.Handler))(0x5ed420)}, ConnState:(func(net.Conn,
http.ConnState))(nil), ErrorLog:(*log.Logger)(nil), disableKeepAlives:0, inShutdown:0, nextProtoOnce:sync.Once{m:sync.Mutex{state:0, sema:0x0}, done:0x1}, nextProtoErr:error(nil), mu:sync.Mutex{state:0, sema:0x0}, listeners:map[net.Listener]struct {}{http.tcpKeepAliveListener{TCPListener:(*net.TCPListener)(0xc04204e030)}:struct {}{}}, activeConn:map[*http.conn]struct {}{(*http.conn)(0xc04203a6e0):struct {}{}},
doneChan:(chan struct {})(nil)}).WithValue(&http.contextKey{name:"local-addr"}, &net.TCPAddr{IP:net.IP{0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0}, Port:8080, Zone:""}).WithCancel.WithCancel
*/
