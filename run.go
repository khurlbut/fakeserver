package main

import (
	"flag"
	"log"
	"net"
	"os"
	"os/signal"
	"strings"
	"syscall"

	"github.com/khurlbut/fakehttp"
	"github.com/khurlbut/fakeserverconf"
)

const defaultPort = "8181"

func main() {
	server := fakehttp.Server()

	shutdownServerOnCntrlC(server)

	config := readConfiguration()
	for _, p := range config.Pages {
		rh := server.NewHandler()
		rh.Get(p.Path).Reply(p.Status).BodyString(p.Body)
		for _, h := range p.Headers {
			s := strings.Split(h, ":")
			k, v := s[0], s[1]
			rh.AddHeader(k, v)
		}
		rh.CustomHandle = fakehttp.RequireHeadersResponder
	}

	server.Start(config.IPAddress, config.Port)

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

	if config.IPAddress == "host" {
		config.IPAddress = resolveHostIP()
	}

	if len(config.Port) == 0 {
		config.Port = defaultPort
	}

	return config
}

func resolveHostIP() string {

	netInterfaceAddresses, err := net.InterfaceAddrs()

	if err != nil {
		log.Fatal(err)
	}

	for _, netInterfaceAddress := range netInterfaceAddresses {
		networkIP, ok := netInterfaceAddress.(*net.IPNet)
		if ok && !networkIP.IP.IsLoopback() && networkIP.IP.To4() != nil {
			ip := networkIP.String()
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
