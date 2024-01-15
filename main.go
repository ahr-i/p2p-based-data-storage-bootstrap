package main

import (
	"net/http"

	"github.com/ahr-i/p2p-based-data-storage-bootstrap/handler/nodeHandler"

	"github.com/rs/cors"
	"github.com/urfave/negroni"
)

func main() {
	/* Bootstrap Setting */
	const port = "3000" // Bootstrap Port Setting

	/* ----- */
	mux := nodeHandler.CreateHandler()
	handler := negroni.Classic()
	cors_ := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},                            // 모든 출처의 Server 접근 허용
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE"}, // 허용된 HTTP Methods
		AllowedHeaders:   []string{"*"},                            // 모든 Header 허용
		AllowCredentials: true,                                     // 자격 증명 허용
	})
	handler.Use(cors_)
	handler.UseHandler(mux)

	go nodeHandler.NodeChecker()
	http.ListenAndServe(":"+port, handler)
}
