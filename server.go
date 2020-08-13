package main

import (
	"net"
	"net/http"
	"os"

	"github.com/sirupsen/logrus"
)

var (
	DefaultSocketLocation = "/tmp/log.sock"
)

// Server structure is used to the store backend information
type Server struct {
	SocketLocation string
	Debug          bool
}

// StartServerWithDefaults starts the server with default values
func StartServerWithDefaults() {
	s := Server{
		SocketLocation: DefaultSocketLocation,
	}
	s.Start()
}

// Start the server
func (s *Server) Start() {
	os.Remove(s.SocketLocation)
	s.ListenAndServe()
}

// ListenAndServe is used to setup handlers and
// start listening on the specified location
func (s *Server) ListenAndServe() error {
	logrus.Infof("Listening on %s", s.SocketLocation)
	server := http.Server{}
	http.HandleFunc("/alloc", s.alloc)
	http.HandleFunc("/status", s.stats)
	socketListener, err := net.Listen("unix", s.SocketLocation)
	if err != nil {
		return err
	}
	return server.Serve(socketListener)
}

func (s *Server) alloc(rw http.ResponseWriter, req *http.Request) {
	logrus.Debugf("Received alloc request")
	if req.Method == http.MethodGet {
		allcateMemory()
		rw.Write([]byte("OK\n"))
	}
}

func (s *Server) stats(rw http.ResponseWriter, req *http.Request) {
	logrus.Debugf("Received stats request")
	if req.Method == http.MethodGet {
		printMemoryStats()
		rw.Write([]byte("OK\n"))
	}
}
