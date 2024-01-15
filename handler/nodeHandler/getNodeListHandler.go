package nodeHandler

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type NodeList struct {
	Address string `json:"address"`
}

func GetNodeListHandler(w http.ResponseWriter, r *http.Request) {
	mutex.Lock()
	defer mutex.Unlock()

	nodes := make([]NodeList, len(node_list))
	for i, node := range node_list {
		nodes[i] = NodeList{Address: fmt.Sprintf("%s:%s", node.IP, node.Port)}
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(nodes)
}
