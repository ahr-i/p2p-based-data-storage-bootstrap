package nodeHandler

import (
	"fmt"
	"net/http"
)

func GetNodeCntHandler(w http.ResponseWriter, r *http.Request) {
	mutex.Lock()
	defer mutex.Unlock()

	fmt.Fprintf(w, "%d", len(node_list))
}
