package main

import (
	"fmt"
	httpTestWrapper "gihub.com/khurlbut/fakeserver"
	"net"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	fake := httpTestWrapper.New()

	// Set up capture of <Ctrl-C> for server shutdown
	c := make(chan os.Signal)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-c
		fake.Close()
		os.Exit(1)
	}()

	fake.NewHandler().Get("/").Reply(200).BodyString("Content Service Upstream")
	fake.NewHandler().Get("/browse/").Reply(200).BodyString("Browse at Content Service Upstream")
	fake.NewHandler().Get("/browse/catalog").Reply(200).BodyString("Browse Catalog at Content Service Upstream")
	fake.NewHandler().Get("/oldpage").Reply(302).BodyString("Redirect")

	fmt.Printf("resolveHostIp(): %s\n", resolveHostIp())
	fake.Start()
	fmt.Printf("Server Running at: %s\n", fake.Server.URL)

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
