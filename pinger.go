package main

import (
	"flag"
	"fmt"
	"runtime"
	"time"

	"github.com/sparrc/go-ping"
)

//Mainloop - will make sure main loop runs as long as this is true
var Mainloop bool

func main() {
	Mainloop = true
	var hostStatus int // 0 - unknown, 1 up, -1 down
	hostStatus = 0
	host := flag.String("s", "8.8.8.8", "IP host to track.")
	dTime := flag.Int("d", 3, "How many pings get lost when the host declared as down.")
	downInterval := flag.Int("o", 1, "Interval between pings when host is down.")
	upInterval := flag.Int("u", 1, "Interval between pings when host is up.")
	verbose := flag.Bool("v", false, "Verbosity")
	flag.Parse()

	var ipHost string
	var downTime int
	var verbosity bool
	verbosity = *verbose
	downTime = *dTime
	ipHost = *host
	fmt.Printf("Starting pinger. Monitoring host %s\n", ipHost)

	if verbosity {
		fmt.Printf("Host %s\n", ipHost)
		fmt.Printf("DownTime %d\n", downTime)
		fmt.Printf("Up interval %d\n", *upInterval)
		fmt.Printf("Down interval %d\n", *downInterval)
	}

	for Mainloop {
		pinger, err := ping.NewPinger(ipHost)
		if err != nil {
			panic(err)
		}
		pinger.Count = 1
		pinger.Timeout = 1
		if runtime.GOOS == "windows" {
			pinger.SetPrivileged(true)
		}
		pinger.Run()
		stats := pinger.Statistics()
		if stats.PacketsRecv == 1 {
			if hostStatus <= 0 {
				fmt.Printf("Host %s is up at %s\n", ipHost, time.Now().Format(time.RFC1123))
				hostStatus = 1
				downTime = *dTime
			}
			time.Sleep(time.Second * time.Duration(*upInterval))
		} else {
			if hostStatus >= 0 {
				if downTime == 0 {
					fmt.Printf("Host %s is down at %s\n", ipHost, time.Now().Format(time.RFC1123))
					hostStatus = -1
				} else {
					downTime--
				}

			}
			time.Sleep(time.Second * time.Duration(*downInterval))
		}
		if verbosity {
			fmt.Printf("Packets rect/send: %d/%d\n", stats.PacketsRecv, stats.PacketsSent)
		}
	}

}
