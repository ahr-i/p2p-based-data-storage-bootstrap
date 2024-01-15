package nodeHandler

import (
	"fmt"
	"net/http"
)

func LogGetNodeList(w http.ResponseWriter, r *http.Request) {
	mutex.Lock()
	defer mutex.Unlock()

	nodes := make([]NodeList, len(node_list))
	for i, node := range node_list {
		nodes[i] = NodeList{Address: fmt.Sprintf("%s:%s", node.IP, node.Port)}
	}

	rend.JSON(w, http.StatusOK, nodes)
}
