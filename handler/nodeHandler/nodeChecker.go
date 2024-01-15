package nodeHandler

import "time"

func NodeChecker() {
	for {
		CheckNodes()

		time.Sleep(5 * time.Second) // 5초 간격으로 Check
	}
}
