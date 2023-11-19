package utils

import (
	"fmt"
	"net"
	"time"
)

func CheckPort(host string, ports []string) {
	for _, port := range ports {
		timeout := time.Second
		conn, err := net.DialTimeout("tcp", net.JoinHostPort(host, port), timeout)
		if err != nil {
			fmt.Printf("Porta %s está fechada.\n", port)
		}

		if conn != nil {
			fmt.Printf("Porta %s está aberta. \n", port)
			defer conn.Close()
		}
	}

}