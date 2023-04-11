package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	http.HandleFunc("/sse", func(w http.ResponseWriter, r *http.Request) {
		// 设置响应头
		w.Header().Set("Content-Type", "text/event-stream")
		w.Header().Set("Cache-Control", "no-cache")
		w.Header().Set("Connection", "keep-alive")
		w.Header().Set("Access-Control-Allow-Origin", "*")

		// 向客户端发送数据
		fmt.Fprintf(w, "event: message\n")
		fmt.Fprintf(w, "data: Hello, world!\n\n")
		flusher := w.(http.Flusher)
		flusher.Flush()

		// 定期发送无意义的数据
		ticker := time.NewTicker(10 * time.Second)
		defer ticker.Stop()
		for {
			select {
			case <-ticker.C:
				fmt.Fprintf(w, "ping: \n\n")
				flusher.Flush()
			case <-r.Context().Done():
				return
			}
		}
	})

	http.ListenAndServe(":3000", nil)
}
