package gapi

import (
	"context"
	"log"

	"google.golang.org/grpc/metadata"
)

const (
	grpcGatewayUserAgentHeader = "grpcgateway-user-agent"
	xForwardForHeader          = "x-forwarded-for"
)

type Metadata struct {
	UserAgent, ClientIP string
}

func (server *Server) extractMetadata(ctx context.Context) *Metadata {
	mtdt := &Metadata{}

	if md, ok := metadata.FromIncomingContext(ctx); ok {
		log.Printf("md %v\n", md)

		// check if userAgents slice is not empty
		if userAgents := md.Get(grpcGatewayUserAgentHeader); len(userAgents) > 0 {
			mtdt.UserAgent = userAgents[0]
		}

		if clientIP := md.Get(xForwardForHeader); len(clientIP) > 0 {
			mtdt.ClientIP = clientIP[0]
		}
	}

	return mtdt
}
