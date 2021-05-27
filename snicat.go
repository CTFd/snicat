package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/therootcompany/sclient"
)

var insecure = flag.Bool("insecure", false, "Disable verification of server SSL certificate")
var servername = flag.String("servername", "", "Server Name Indication (SNI) to provide to the server (e.g. ssl.example.com)")
var bind = flag.String("bind", "", "Tunnel connection to a local unencrypted port (e.g. localhost:3000, 3000)")
var (
	commit  = "unknown"
	version = "v0.0.0"
)

func init() {
	flag.BoolVar(insecure, "k", false, "Shorthand for -insecure")
	flag.StringVar(servername, "sni", "", "Shorthand for -servername")
	flag.StringVar(bind, "b", "", "Shorthand for -bind")
}

func Usage() {
	fmt.Println(os.Args[0], "<hostname>", "<port>")
	fmt.Println("version:", version)
	fmt.Println("commit:", commit)
	fmt.Println()
	flag.Usage()
}

func main() {
	flag.Parse()

	if flag.NArg() < 1 {
		Usage()
		os.Exit(1)
	}

	hostname := flag.Arg(0)
	var port int
	var err error
	if strings.Contains(hostname, ":") {
		split := strings.Split(hostname, ":")
		hostname = split[0]
		port, err = strconv.Atoi(split[1])
		if err != nil {
			fmt.Println(flag.Arg(1), "is not a valid remote port")
			os.Exit(1)
		}
	} else {
		portArg := flag.Arg(1)
		if len(portArg) != 0 {
			port, err = strconv.Atoi(portArg)
			if err != nil {
				fmt.Println(flag.Arg(1), "is not a valid remote port")
				os.Exit(1)
			}
		} else {
			port = 443
		}
	}

	localAddress := "-"
	var localPort int
	if len(*bind) > 0 {
		if strings.Contains(*bind, ":") {
			split := strings.Split(*bind, ":")
			localAddress = split[0]
			localPort, err = strconv.Atoi(split[1])
			if err != nil {
				fmt.Println(*bind, "is not a valid local port")
				os.Exit(1)
			}
		} else {
			localAddress = "localhost"
			localPort, err = strconv.Atoi(*bind)
			if err != nil {
				fmt.Println(*bind, "is not a valid local port")
				os.Exit(1)
			}
		}
	}

	if len(*servername) == 0 {
		servername = &hostname
	}
	sclient := &sclient.Tunnel{
		ServerName:         *servername,
		RemoteAddress:      hostname,
		RemotePort:         port,
		LocalAddress:       localAddress,
		LocalPort:          localPort,
		InsecureSkipVerify: *insecure,
	}
	sclient.DialAndListen()
}
