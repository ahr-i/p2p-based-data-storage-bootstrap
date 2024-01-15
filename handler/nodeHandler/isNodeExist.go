package nodeHandler

func IsNodeExist(ip string, port string) bool {
	for _, node := range node_list {
		/* Node 존재 시 */
		if node.IP == ip && node.Port == port {
			return true
		}
	}

	return false
}
