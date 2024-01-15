package nodeHandler

import (
	"net/http"
)

func RegisterNodeHandler(w http.ResponseWriter, r *http.Request) {
	ip := r.URL.Query().Get("ip")
	port := r.URL.Query().Get("port")

	if IsNodeExist(ip, port) {
		w.WriteHeader(http.StatusConflict)
		w.Write([]byte("Node already exists"))

		return
	}

	mutex.Lock()
	node_list = append(node_list, NodeInfo{IP: ip, Port: port})
	mutex.Unlock()

	w.WriteHeader(http.StatusOK)
}
