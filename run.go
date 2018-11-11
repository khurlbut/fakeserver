package main

import (
	"fmt"
	"github.com/khurlbut/mockhttp"
	"net"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	server := mockhttp.Server()

	// Set up capture of <Ctrl-C> for server shutdown
	c := make(chan os.Signal)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-c
		server.Close()
		os.Exit(1)
	}()

	server.NewHandler().Get("/").Reply(200).BodyString("Content Service Upstream")
	server.NewHandler().Get("/browse/").Reply(200).BodyString("Browse at Content Service Upstream")
	server.NewHandler().Get("/browse/catalog").Reply(200).BodyString("Browse Catalog at Content Service Upstream")
	server.NewHandler().Get("/oldpage").Reply(302).BodyString("Redirect")

	fmt.Printf("resolveHostIp(): %s\n", resolveHostIp())
	server.Start()
	fmt.Printf("Server Running at: %s\n", server.Server.URL)

	// Don't exit
	for {
	}
}

func resolveHostIp() string {

	netInterfaceAddresses, err := net.InterfaceAddrs()

	if err != nil {
		return ""
	}

	for _, netInterfaceAddress := range netInterfaceAddresses {

		networkIp, ok := netInterfaceAddress.(*net.IPNet)

		if ok && !networkIp.IP.IsLoopback() && networkIp.IP.To4() != nil {

			ip := networkIp.IP.String()

			fmt.Println("Resolved Host IP: " + ip)

			return ip
		}
	}
	return ""
}
