package main

import (
	"flag"
	"fmt"
	"github.com/khurlbut/fakehttp"
	"github.com/khurlbut/fakeserverconf"
	"log"
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
	server := fakehttp.Server()

	shutdownServerOnCntrlC(server)

	config := fakeserverconf.DefaultConfig()
	configfile := flag.String("config-file", "", "JSON configuration file")
	flag.Parse()

	if len(*configfile) > 0 {
		config = fakeserverconf.ReadJSONFile(*configfile)
	}

	for _, c := range config {
		server.NewHandler().Get(c.Path).Reply(c.Status).BodyString(c.Body)
	}

	server.Start(resolveHostIp(), "8181")
	fmt.Printf("Server Running at: %s\n", server.URL())

	// Don't exit
	for {
	}
}

func resolveHostIp() string {

	netInterfaceAddresses, err := net.InterfaceAddrs()

	if err != nil {
		log.Fatal(err)
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

func shutdownServerOnCntrlC(server *fakehttp.HTTPFake) {
	c := make(chan os.Signal)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-c
		server.Close()
		os.Exit(1)
	}()
}
