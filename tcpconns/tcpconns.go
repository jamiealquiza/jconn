// The MIT License (MIT)
//
// Copyright (c) 2015 Jamie Alquiza
//
// http://knowyourmeme.com/memes/deal-with-it.
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
package tcpconns

import (
	"bufio"
	"encoding/hex"
	"os"
	"strconv"
	"strings"
)

// Get fetches the current tcp connection data from /proc/net/tcp and
// returns it as a human readable, multidimensional slice of strings.
// Format: [[local_ip, local_port, remote_ip, remote_port, state], ...]
func Get() ([][]string, error) {
	file, err := os.Open("/proc/net/tcp")
	if err != nil {
		return nil, err
	}

	// [][]string that will hold [localIp:localPort, remoteIp:remotePort] pairs.
	connections := make([][]string, 0)

	// Fetch the IP:port pairs and append into connections.
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		fields := strings.Fields(line)
		conn := fields[1:4]
		if len(fields) > 12 {
			conn = append(conn, fields[12])
		}
		connections = append(connections, conn)
	}
	file.Close()

	// For each [local, remote] ip:port pairs, we convert the captured
	// hex value to human readable strings and append into the [][]string formatted.
	// We use a new slice so we can do sanity checks against our return data rather than
	// just returning whatever was captured from proc.
	formatted := make([][]string, 0)
	for _, conn := range connections[1:] {
		// For each connection set, handle the local / remote ip:port pair.
		pair := []string{}
		for _, c := range conn {
			if len(conn) < 3 {
				break
			}
			ipPort := strings.Split(c, ":")
			// Lazy check. In case we just got something other than ip:port.
			if len(ipPort) == 2 {
				ipHex, _ := hex.DecodeString(ipPort[0])
				ip, port := hexToString(ipHex, ipPort[1])
				pair = append(pair, ip)
				pair = append(pair, port)
			} else {
				break
			}
		}
		state := hexStateToString(conn[2])
		pair = append(pair, state)
		if len(conn) == 4 {
			pair = append(pair, conn[3])
		}
		formatted = append(formatted, pair)
	}
	return formatted, nil

}

// hexToString does the conversion to a human readable string.
// For instance, inserting periods in between each octet.
// Also, we get back a bunch of uint8's anyway and there's
// no quick conversions in the std lib.
func hexToString(ipHex []byte, portHex string) (string, string) {
	var ip, port string

	// Handle the IP address.
	for i := len(ipHex) - 1; i > -1; i-- {
		ip += strconv.Itoa(int(ipHex[i]))
		if i > 0 {
			ip += "."
		}
	}

	// Handle the port.
	portInt64, _ := strconv.ParseInt(portHex, 16, 32)
	port = strconv.Itoa(int(portInt64))
	return ip, port
}

func hexStateToString(stateHex string) string {
	switch stateHex {
	case "01":
		return "ESTABLISHED"
	case "02":
		return "SYN_SENT"
	case "03":
		return "SYN_RECV"
	case "04":
		return "FIN_WAIT1"
	case "05":
		return "FIN_WAIT2"
	case "06":
		return "TIME_WAIT"
	case "07":
		return "CLOSE"
	case "08":
		return "CLOSE_WAIT"
	case "09":
		return "LAST_ACK"
	case "0A":
		return "LISTEN"
	case "0B":
		return "CLOSING"
	}
	return ""
}
