package main

import (
	"fmt"

	"github.com/khurlbut/fakehttp"
	conf "github.com/khurlbut/fakeserverconf"
	// "github.com/tkanos/gonfig"
	"net"
	"os"
	"os/signal"
	"syscall"
)

type Configuration struct {
	path   string
	body   string
	status int
}

func main() {
	fmt.Println("Version 0.1.4")
	server := fakehttp.Server()

	// Set up capture of <Ctrl-C> for server shutdown
	c := make(chan os.Signal)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-c
		server.Close()
		os.Exit(1)
	}()

	config := conf.ReadJSONFile("./config.json")

	for _, c := range config {
		server.NewHandler().Get(c.Path).Reply(c.Status).BodyString(c.Body)
	}

	fmt.Printf("resolveHostIp(): %s\n", resolveHostIp())
	server.Start(resolveHostIp(), "8181")
	fmt.Printf("Server Running at: %s\n", server.URL())

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
