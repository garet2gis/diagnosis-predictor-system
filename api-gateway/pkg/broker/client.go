package broker

import (
	"fmt"
	"github.com/nats-io/nats.go"
	"time"
)

type Config struct {
	NatsURI               string
	MaxConnectionAttempts int
}

func NewNATSClient(cfg Config) (*NatsBroker, error) {
	var connection *nats.Conn
	var err error

	natsURL := cfg.NatsURI

	err = DoWithTries(func() error {
		connection, err = nats.Connect(natsURL)
		if err != nil {
			return fmt.Errorf("can't connect to NATS (no credentials) (%s) due to error: %w", natsURL, err)
		}

		return nil
	}, cfg.MaxConnectionAttempts, 2*time.Second)

	if err != nil {
		return nil, err
	}

	return newNatsBroker(connection), nil
}

func DoWithTries(fn func() error, attempts int, delay time.Duration) (err error) {
	for attempts > 0 {
		if err = fn(); err != nil {
			time.Sleep(delay)
			attempts--
			continue
		}
		return nil
	}
	return
}
