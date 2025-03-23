package nats

import (
	"fmt"
	"log"
	"time"

	"github.com/nats-io/nats.go"
)

// NATSClient represents a NATS client connection
type NATSClient struct {
	conn *nats.Conn
}

// NewNATSClient creates and returns a new NATS client
func NewNATSClient(url string) (*NATSClient, error) {
	// Connect to NATS server with options
	opts := []nats.Option{
		nats.Name("Medusa Service"),
		nats.Timeout(10 * time.Second),
		nats.ReconnectWait(5 * time.Second),
		nats.MaxReconnects(-1), // Unlimited reconnects
		nats.DisconnectErrHandler(func(nc *nats.Conn, err error) {
			log.Printf("NATS disconnected: %v", err)
		}),
		nats.ReconnectHandler(func(nc *nats.Conn) {
			log.Printf("NATS reconnected to %s", nc.ConnectedUrl())
		}),
		nats.ErrorHandler(func(nc *nats.Conn, sub *nats.Subscription, err error) {
			log.Printf("NATS error: %v", err)
		}),
	}

	conn, err := nats.Connect(url, opts...)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to NATS: %w", err)
	}

	log.Println("Connected to NATS successfully")

	return &NATSClient{conn: conn}, nil
}

// Close disconnects from NATS
func (c *NATSClient) Close() error {
	c.conn.Close()
	log.Println("Disconnected from NATS")
	return nil
}

// Publish publishes a message to a subject
func (c *NATSClient) Publish(subject string, data []byte) error {
	return c.conn.Publish(subject, data)
}

// Subscribe subscribes to a subject
func (c *NATSClient) Subscribe(subject string, callback nats.MsgHandler) (*nats.Subscription, error) {
	return c.conn.Subscribe(subject, callback)
}

// QueueSubscribe subscribes to a subject as part of a queue group
func (c *NATSClient) QueueSubscribe(subject, queue string, callback nats.MsgHandler) (*nats.Subscription, error) {
	return c.conn.QueueSubscribe(subject, queue, callback)
}

// Request sends a request and waits for a response
func (c *NATSClient) Request(subject string, data []byte, timeout time.Duration) (*nats.Msg, error) {
	return c.conn.Request(subject, data, timeout)
}

// GetConn returns the underlying NATS connection
func (c *NATSClient) GetConn() *nats.Conn {
	return c.conn
}

// JetStream returns a JetStream context for the connection
func (c *NATSClient) JetStream(opts ...nats.JSOpt) (nats.JetStreamContext, error) {
	return c.conn.JetStream(opts...)
}
