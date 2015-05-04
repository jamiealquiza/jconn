// The MIT License (MIT)
//
// Copyright (c) 2015 Jamie Alquiza
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.
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
	LocalIp    string `json:"local_ip"`
	LocalPort  string `json:"local_port"`
	RemoteIp   string `json:"remote_ip"`
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
			LocalIp:    c[0],
			LocalPort:  c[1],
			RemoteIp:   c[2],
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
