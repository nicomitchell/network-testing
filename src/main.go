package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"

	"github.com/nicomitchell/network-testing/src/node"
)

func main() {
	if len(os.Args) != 4 {
		fmt.Println("required args: ip list filepath, server port, gossip network ports")
		os.Exit(0)
	}
	ipList := os.Args[1]
	httpServerPort := os.Args[2]
	portRange := os.Args[3]
	f, err := os.Open(ipList)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	var lines []string
	lines, err = getLines(f)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	// fmt.Println(httpServerPort)
	openPorts, err := parsePortRange(portRange)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	// fmt.Println(openPorts)
	// fmt.Println(lines)
	node, err := node.NewNode(httpServerPort, openPorts)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	err = node.Start()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println(lines)

}

func getLines(r io.ReadCloser) ([]string, error) {
	br := bufio.NewReader(r)
	l, err := br.ReadBytes('\n')
	lines := make([]string, 0)
	if err != nil {
		if err != io.EOF {
			return nil, err
		}
		lines = append(lines, strings.Trim(string(l), "\n"))
		return lines, nil
	}
	for err == nil {
		lines = append(lines, strings.Trim(string(l), "\n"))
		l, err = br.ReadBytes('\n')
	}
	if err != io.EOF {
		return nil, err
	}
	r.Close()
	return lines, nil
}

func parsePortRange(portRange string) ([]string, error) {
	startAndEnd := strings.Split(portRange, ":")
	if len(startAndEnd) != 2 {
		return nil, fmt.Errorf("Invalid formatting for port range")
	}
	start, err := strconv.Atoi(startAndEnd[0])
	if err != nil {
		return nil, err
	}
	end, err := strconv.Atoi(startAndEnd[1])
	if err != nil {
		return nil, err
	}
	out := []string{}
	for i := start; i <= end; i++ {
		port := strconv.Itoa(i)
		out = append(out, port)
	}
	return out, nil
}
