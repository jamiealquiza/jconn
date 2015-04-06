package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"os"
	"time"

	"github.com/jamiealquiza/tcpconns"
)

var (
	hostname     string
	ts string
	tagHostname  bool
	tagTimestamp bool
)

func init() {
	flag.BoolVar(&tagHostname, "tag-hostname", false, "Add @hostname: hostname field-value")
	flag.BoolVar(&tagTimestamp, "tag-timestamp", false, "Add @timestamp: RFC3339 field-value")
	flag.Parse()
}

type connection struct {
	LocalIp    string `json:"local_ip"`
	LocalPort  string `json:"local_port"`
	RemoteIp   string `json:"remote_ip"`
	RemotePort string `json:"remote_port"`
	State      string `json:"state"`
	Timestamp  string `json:"@timestamp"`
	Hostname   string `json:"@hostname"`
}

func main() {
	if tagHostname {
		hostname, _ = os.Hostname()
	}

	if tagTimestamp {
		ts = time.Now().Format(time.RFC3339)
	}

	conns, _ := tcpconns.Get()
	for _, c := range conns {

		m := &connection{LocalIp: c[0],
			LocalPort:  c[1],
			RemoteIp:   c[2],
			RemotePort: c[3],
			State:      c[4],
			Hostname:   hostname,
			Timestamp:  ts,
		}

		msg, err := json.Marshal(m)
		if err != nil {
			fmt.Println(err)
		}

		fmt.Println(string(msg))
	}
}
