package main

import (
	"net"
	"net/http"
	"os"
	"runtime"
	"runtime/debug"

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
	http.HandleFunc("/stats", s.stats)
	http.HandleFunc("/freegc", s.freegc)
	http.HandleFunc("/freeosm", s.freeosm)
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

func (s *Server) freegc(rw http.ResponseWriter, req *http.Request) {
	logrus.Debugf("Received freegc request")
	if req.Method == http.MethodGet {
		runtime.GC()
		rw.Write([]byte("OK\n"))
	}
}

func (s *Server) freeosm(rw http.ResponseWriter, req *http.Request) {
	logrus.Debugf("Received freeosm request")
	if req.Method == http.MethodGet {
		debug.FreeOSMemory()
		rw.Write([]byte("OK\n"))
	}
}
