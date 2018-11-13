package main

import (
	"flag"
	"github.com/khurlbut/fakehttp"
	"github.com/khurlbut/fakeserverconf"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	server := fakehttp.Server()

	shutdownServerOnCntrlC(server)

	config := readConfiguration()

	for _, p := range config.Pages {
		server.NewHandler().Get(p.Path).Reply(p.Status).BodyString(p.Body)
	}

	server.Start(resolveHostIp(), "8181")

	log.Print("Server Running at --> " + server.URL())
	for {
	}
}

func readConfiguration() fakeserverconf.Configuration {
	config := fakeserverconf.DefaultConfig()
	configfile := flag.String("config-file", "", "JSON configuration file")
	flag.Parse()

	if len(*configfile) > 0 {
		config = fakeserverconf.ReadJSONFile(*configfile)
	}
	return config
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
