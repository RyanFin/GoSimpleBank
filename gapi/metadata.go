package gapi

import (
	"context"
	"log"

	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/peer"
)

const (
	grpcGatewayUserAgentHeader = "grpcgateway-user-agent"
	xForwardForHeader          = "x-forwarded-for"
	userAgentHeader            = "user-agent"
)

type Metadata struct {
	UserAgent, ClientIP string
}

func (server *Server) extractMetadata(ctx context.Context) *Metadata {
	mtdt := &Metadata{}

	if md, ok := metadata.FromIncomingContext(ctx); ok {
		log.Printf("md %v\n", md)

		if userAgents := md.Get(grpcGatewayUserAgentHeader); len(userAgents) > 0 {
			mtdt.UserAgent = userAgents[0]
		}

		// check if userAgents slice is not empty
		if userAgents := md.Get(userAgentHeader); len(userAgents) > 0 {
			mtdt.UserAgent = userAgents[0]
		}

		if clientIP := md.Get(xForwardForHeader); len(clientIP) > 0 {
			mtdt.ClientIP = clientIP[0]
		}
	}

	if p, ok := peer.FromContext(ctx); ok {
		mtdt.ClientIP = p.Addr.String()
	}

	return mtdt
}
