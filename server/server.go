package server

import (
	"context"
	"fmt"
	"log"
	"sync"

	"github.com/gobench-io/gobench/ent"
	_ "github.com/mattn/go-sqlite3"
)

type app struct {
	id       int
	name     string
	scenario string
}

type Server struct {
	mu       sync.Mutex
	pendings []app // pending apps
	curr     *app  // current processing app

	// database
	dbFilename string
	db         *ent.Client
}

// singleton server instance
var server Server

func init() {
	server = Server{}
}

// New returns the singleton instance of the server
func New() (*Server, error) {
	return &server, nil
}

// Start begin a gobench server
func (s *Server) Start() error {
	// default db name
	dbFilename := "./gobench.sqlite3"
	if err := server.setupDb(dbFilename); err != nil {
		return err
	}

	return nil
}

// NewApp creates a new application
// provided name and scenario
// returns application id and error
// func (s *Server) NewApp(name string, scenario string) (int, error) {
// 	s.mu.Lock()
// 	defer s.mu.Unlock()

// 	if s.curr == nil {
// 	}
// }

func (s *Server) setupDb(filename string) error {
	log.Printf("gobench result will be save in %s\n", filename)

	client, err := ent.Open(
		"sqlite3",
		filename+"?mode=rwc&cache=shared&&_busy_timeout=9999999&_fk=1")

	if err != nil {
		return fmt.Errorf("failed openning sqlite3 connection: %v", err)
	}

	if err = client.Schema.Create(context.Background()); err != nil {
		return fmt.Errorf("failed creating schema resources: %v", err)
	}

	s.mu.Lock()
	s.dbFilename = filename
	s.db = client
	s.mu.Unlock()

	return nil
}

// DB retuns the db client
func (s *Server) DB() *ent.Client {
	return s.db
}

// NewApplication create a new application with a name and a scenario
// return the application id and error
// func (s *Server) NewApplication(name, scenario string) (int, err) {
// }
