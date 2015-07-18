package dnsimple

import (
	"fmt"
	"log"

	"github.com/pearkes/dnsimple"
)

type Config struct {
	Email       string
	Token       string
	DomainToken string
}

// Client() returns a new client for accessing dnsimple.
func (c *Config) Client() (*dnsimple.Client, error) {
	// client, err := dnsimple.NewClient(c.Email, c.Token)
	// client, err := dnsimple.NewClientWithDomainToken(c.DomainToken)

	var client *dnsimple.Client
	var err error

	if c.DomainToken != "" {
		client, err = dnsimple.NewClientWithDomainToken(c.DomainToken)
	} else {
		client, err = dnsimple.NewClient(c.Email, c.Token)
	}

	if err != nil {
		return nil, fmt.Errorf("Error setting up client: %s", err)
	}

	if c.DomainToken != "" {
		log.Print("[INFO] DNSimple Client configured with domain token")
	} else {
		log.Printf("[INFO] DNSimple Client configured for user: %s", client.Email)
	}

	return client, nil
}
