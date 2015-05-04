# tcp-conns
Fetches TCP connection info in single-digit millisecond time.

Output format: [local addr, local port, remote addr, remote port, state]

In example.go:
```go
package main

import (
	"fmt"
	
	"github.com/jamiealquiza/jconn/tcpconns"
)

func main() {
	conns, _ := tcpconns.Get()
	for _, c := range conns {
		if c[4] == "ESTABLISHED" {
			fmt.Println(c)
		}
	}
}
```

<pre>
% ./example
[xxx.xxx.239.215 37625 xxx.xxx.250.120 18000 ESTABLISHED]
[127.0.0.1 2101 xxx.xxx.239.215 53789 ESTABLISHED]
[xxx.xxx.239.215 53789 127.0.0.1 2101 ESTABLISHED]
[xxx.xxx.239.215 50805 127.0.0.1 2102 ESTABLISHED]
[xxx.xxx.239.215 22 xxx.xxx.250.5 34094 ESTABLISHED]
[xxx.xxx.239.215 52846 xxx.xxx.250.76 1515 ESTABLISHED]
[xxx.xxx.239.215 40579 127.0.0.1 2103 ESTABLISHED]
[127.0.0.1 2102 xxx.xxx.239.215 50805 ESTABLISHED]
[127.0.0.1 2100 xxx.xxx.239.215 56339 ESTABLISHED]
[127.0.0.1 2103 xxx.xxx.239.215 40579 ESTABLISHED]
[xxx.xxx.239.215 58818 xxx.xxx.250.96 2003 ESTABLISHED]
[xxx.xxx.239.215 56339 127.0.0.1 2100 ESTABLISHED]
</pre>
