// +build !windows
// Copyright 2012-2017 Apcera Inc. All rights reserved.

package server

// Run starts the NATS server. This wrapper function allows Windows to add a
// hook for running NATS as a service.
func Run(server *Server) {
	server.Start()
}
