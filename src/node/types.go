package node

import (
	"encoding/json"
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
	"sort"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

type NodeInfo struct {
	HostName  string
	IPAddress net.IP
	//tbd
}

type Node struct {
	ServerPort string
	Ports      map[string]*net.Conn
	Info       NodeInfo
	Router     *mux.Router
}

func getNodeInfo() NodeInfo {
	hostName := os.Getenv("HOSTNAME")
	conn, err := net.Dial("udp", "8.8.8.8:80")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	localAddr := conn.LocalAddr().(*net.UDPAddr)
	ip := localAddr.IP.String()
	return NodeInfo{
		HostName:  hostName,
		IPAddress: net.ParseIP(ip),
	}
}

func NewNode(serverPort string, ports []string) (*Node, error) {
	portsMap := make(map[string]*net.Conn)
	for _, port := range ports {
		portsMap[port] = nil
	}
	n := new(Node)
	n.ServerPort = serverPort
	n.Ports = portsMap
	n.Info = getNodeInfo()
	n.Router = mux.NewRouter()
	n.configureRouter()
	return n, nil
}

func (n *Node) configureRouter() {
	n.Router.HandleFunc("/ports", n.getOpenPorts).Methods("GET")
}

func (n *Node) Start() error {
	return http.ListenAndServe(
		fmt.Sprintf(":%s", n.ServerPort),
		handlers.CORS(
			handlers.AllowedOrigins([]string{"*"}),
		)(n.Router),
	)
}

func (n *Node) getOpenPorts(w http.ResponseWriter, r *http.Request) {
	type openNodePortsResponse struct {
		Ports []string `json:"ports"`
	}
	openPorts := make([]string, 0)
	for k, v := range n.Ports {
		if v == nil {
			openPorts = append(openPorts, k)
		}
	}
	sort.Strings(openPorts)
	resp := openNodePortsResponse{Ports: openPorts}
	json.NewEncoder(w).Encode(resp)
}
