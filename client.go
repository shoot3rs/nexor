package nexor

import (
	"github.com/nats-io/nats.go"
	"log"
	"os"
	"strconv"
	"time"
)

type Nexor struct {
	conn *nats.Conn
	js   nats.JetStreamContext
}

type nexorConfig struct {
	ClientName string
	Debug      bool
	MaxConn    int
	MaxRecon   int
	ReconWait  int
}

func getConfig() *nexorConfig {
	var debugMode = false
	var clientName = "Nexor"
	var maxConn, maxWait = 5, 5
	debugModeValue, found := os.LookupEnv("NEXOR.DEBUG")
	if found {
		debugMode = debugModeValue == "true"
	}

	clientNameValue, found := os.LookupEnv("NEXOR.CLIENT")
	if found {
		clientName = clientNameValue
	}

	maxConnValue := os.Getenv("NEXOR.MAX_CONNECTIONS")
	if maxConnValue != "" {
		maxConn, _ = strconv.Atoi(maxConnValue)
	}

	maxWaitValue := os.Getenv("NEXOR.MAX_RECONNECT_WAIT")
	if maxWaitValue != "" {
		maxWait, _ = strconv.Atoi(maxWaitValue)
	}

	return &nexorConfig{
		ClientName: clientName,
		Debug:      debugMode,
		MaxConn:    maxConn,
		ReconWait:  maxWait,
	}
}

// New initializes a new Nexor client with optional NATS options.
// Usage: nexor.New("nats://localhost:4222")
func New(url string, opts ...nats.Option) (*Nexor, error) {
	cfg := getConfig()
	if len(opts) == 0 {
		opts = []nats.Option{
			nats.Name(cfg.ClientName),
			nats.MaxReconnects(cfg.MaxRecon),
			nats.ReconnectWait(time.Duration(cfg.ReconWait) * time.Second),
		}
	}

	conn, err := nats.Connect(url, opts...)
	if err != nil {
		if cfg.Debug {
			log.Printf("🔌 Failed to connect to NATS: %v 🔌\n\n", err)
			os.Exit(1)
		}

		return nil, err
	}

	js, err := conn.JetStream()
	if err != nil {
		if cfg.Debug {
			log.Printf("🔌 Failed to connect to Jetstream: %v 🔌", err)
			os.Exit(1)
		}

		conn.Close()
		return nil, err
	}

	if cfg.Debug {
		log.Printf("🚀 Connected to NATS server successful 🚀\n")
	}

	return &Nexor{conn: conn, js: js}, nil
}

// Close safely closes the NATS connection.
func (n *Nexor) Close() {
	if n.conn != nil && !n.conn.IsClosed() {
		n.conn.Close()
	}
}

// JetStream exposes the underlying JetStream context
// so that microservices can create/manage streams and consumers.
func (n *Nexor) JetStream() nats.JetStreamContext {
	return n.js
}
