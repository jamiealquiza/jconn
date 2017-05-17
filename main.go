package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"os"
	"time"

	"github.com/jamiealquiza/jconn/tcpconns"
)

var (
	hostname     string
	ts           string
	tagHostname  bool
	tagTimestamp bool
	tagRto       bool
)

func init() {
	flag.BoolVar(&tagHostname, "tag-hostname", false, "Add @hostname: hostname field-value")
	flag.BoolVar(&tagTimestamp, "tag-timestamp", false, "Add @timestamp: RFC3339 field-value")
	flag.BoolVar(&tagRto, "tag-rto", true, "Add RTO field if applicable")
	flag.Parse()
}

type connection struct {
	LocalIP    string `json:"local_ip"`
	LocalPort  string `json:"local_port"`
	RemoteIP   string `json:"remote_ip"`
	RemotePort string `json:"remote_port"`
	State      string `json:"state"`
	RTO        string `json:"rto,omitempty"`
	Timestamp  string `json:"@timestamp,omitempty"`
	Hostname   string `json:"@hostname,omitempty"`
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
		m := &connection{
			LocalIP:    c[0],
			LocalPort:  c[1],
			RemoteIP:   c[2],
			RemotePort: c[3],
			State:      c[4],
			Hostname:   hostname,
			Timestamp:  ts,
		}

		if tagRto && len(c) == 6 {
			m.RTO = c[5]
		}

		msg, err := json.Marshal(m)
		if err != nil {
			fmt.Println(err)
		}

		fmt.Println(string(msg))
	}
}
