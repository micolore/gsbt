package main

import (
	"context"
	"fmt"
	"log"

	"github.com/rsocket/rsocket-go"
	"github.com/rsocket/rsocket-go/payload"
	"github.com/rsocket/rsocket-go/rx/mono"
)

func main() {
	err := rsocket.Receive().
		Acceptor(func(ctx context.Context, setup payload.SetupPayload, sendingSocket rsocket.CloseableRSocket) (rsocket.RSocket, error) {
			// bind responder
			return rsocket.NewAbstractSocket(
				rsocket.RequestResponse(func(msg payload.Payload) mono.Mono {
					fmt.Sprintf("%s\n", msg.DataUTF8())
					return mono.Just(msg)
				}),
			), nil
		}).
		Transport(rsocket.TCPServer().SetAddr(":7878").Build()).
		Serve(context.Background())
	log.Fatalln(err)
}
