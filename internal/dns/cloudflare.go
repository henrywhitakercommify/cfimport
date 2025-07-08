// Package dns
// An abstraction over the cloudflare api client to retrieve dns records for a zone
package dns

import (
	"context"
	"fmt"

	"github.com/cloudflare/cloudflare-go/v4"
	"github.com/cloudflare/cloudflare-go/v4/dns"
	"github.com/cloudflare/cloudflare-go/v4/option"
	"github.com/henrywhitakercommify/cfimport/internal/slice"
)

type Client struct {
	client *cloudflare.Client
}

func New(apiToken string) *Client {
	api := cloudflare.NewClient(
		option.WithAPIToken(apiToken),
	)
	return &Client{
		client: api,
	}
}

type Record struct {
	ID      string
	Type    string
	Value   string
	Proxied bool
}

func (c *Client) Records(ctx context.Context, zone string) ([]Record, error) {
	records, err := c.client.DNS.Records.List(ctx, dns.RecordListParams{
		ZoneID:  cloudflare.String(zone),
		PerPage: cloudflare.Float(200),
	})
	if err != nil {
		return nil, fmt.Errorf("list zone records: %w", err)
	}

	return slice.Map(records.Result, func(record dns.RecordResponse) Record {
		return Record{
			ID:      record.ID,
			Type:    string(record.Type),
			Value:   record.Content,
			Proxied: record.Proxied,
		}
	}), nil
}
