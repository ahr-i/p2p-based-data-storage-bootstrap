package nodeHandler

import (
	"log"
	"net/http"
	"time"
)

func CheckNodes() {
	mutex.Lock()
	defer mutex.Unlock()

	log.Println("* (Node Checker) Node Checker Is Working.")
	var alive_node_list []NodeInfo
	client := &http.Client{
		Timeout: 2 * time.Second,
	}

	for _, node := range node_list {
		node_url := "http://" + node.IP + ":" + node.Port + "/ping"

		resp, err := client.Get(node_url)
		if err == nil && resp.StatusCode == http.StatusOK {
			log.Println("*(Node Checker)", node.IP+":"+node.Port, "\tIs Alive.")
			alive_node_list = append(alive_node_list, node)

			resp.Body.Close()
		} else {
			log.Println("*(Node Checker)", node.IP+":"+node.Port, "\tDelete.")
		}
	}

	node_list = alive_node_list
}
