package nodeHandler

import (
	"net/http"
	"sync"

	"github.com/gorilla/mux"
	"github.com/unrolled/render"
)

var rend *render.Render = render.New()
var node_list []NodeInfo
var mutex sync.Mutex

type NodeInfo struct {
	IP   string
	Port string
}

func CreateHandler() http.Handler {
	mux := mux.NewRouter()

	mux.HandleFunc("/register", RegisterNodeHandler).Methods("GET")     // Node 등록 URL
	mux.HandleFunc("/nodes", GetNodeListHandler).Methods("GET")         // Node List Return URL
	mux.HandleFunc("/alive_node_cnt", GetNodeCntHandler).Methods("GET") // 살아있는 Node의 개수 Return URL

	mux.HandleFunc("/log_get_node_list", LogGetNodeList).Methods("GET") // Log용 Node List 출력 URL

	return mux
}
