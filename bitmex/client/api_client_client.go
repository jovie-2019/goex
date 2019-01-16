// Code generated by go-swagger; DO NOT EDIT.

package client

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"
	httptransport "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/nntaoli-project/GoEx/bitmex/client/announcement"
	"github.com/nntaoli-project/GoEx/bitmex/client/api_key"
	"github.com/nntaoli-project/GoEx/bitmex/client/chat"
	"github.com/nntaoli-project/GoEx/bitmex/client/execution"
	"github.com/nntaoli-project/GoEx/bitmex/client/funding"
	"github.com/nntaoli-project/GoEx/bitmex/client/instrument"
	"github.com/nntaoli-project/GoEx/bitmex/client/insurance"
	"github.com/nntaoli-project/GoEx/bitmex/client/leaderboard"
	"github.com/nntaoli-project/GoEx/bitmex/client/liquidation"
	"github.com/nntaoli-project/GoEx/bitmex/client/notification"
	"github.com/nntaoli-project/GoEx/bitmex/client/order"
	"github.com/nntaoli-project/GoEx/bitmex/client/order_book"
	"github.com/nntaoli-project/GoEx/bitmex/client/position"
	"github.com/nntaoli-project/GoEx/bitmex/client/quote"
	"github.com/nntaoli-project/GoEx/bitmex/client/schema"
	"github.com/nntaoli-project/GoEx/bitmex/client/settlement"
	"github.com/nntaoli-project/GoEx/bitmex/client/stats"
	"github.com/nntaoli-project/GoEx/bitmex/client/trade"
	"github.com/nntaoli-project/GoEx/bitmex/client/user"
)

// Default API client HTTP client.
var Default = NewHTTPClient(nil)

const (
	// DefaultHost is the default Host
	// found in Meta (info) section of spec file
	DefaultHost string = "localhost"
	// DefaultBasePath is the default BasePath
	// found in Meta (info) section of spec file
	DefaultBasePath string = "/api/v1"
)

// DefaultSchemes are the default schemes found in Meta (info) section of spec file
var DefaultSchemes = []string{"https"}

// NewHTTPClient creates a new API client HTTP client.
func NewHTTPClient(formats strfmt.Registry) *APIClient {
	return NewHTTPClientWithConfig(formats, nil)
}

// NewHTTPClientWithConfig creates a new API client HTTP client,
// using a customizable transport config.
func NewHTTPClientWithConfig(formats strfmt.Registry, cfg *TransportConfig) *APIClient {
	// ensure nullable parameters have default
	if cfg == nil {
		cfg = DefaultTransportConfig()
	}

	// create transport and client
	transport := httptransport.New(cfg.Host, cfg.BasePath, cfg.Schemes)
	return New(transport, formats)
}

// New creates a new API client client
func New(transport runtime.ClientTransport, formats strfmt.Registry) *APIClient {
	// ensure nullable parameters have default
	if formats == nil {
		formats = strfmt.Default
	}

	cli := new(APIClient)
	cli.Transport = transport

	cli.Announcement = announcement.New(transport, formats)

	cli.APIKey = api_key.New(transport, formats)

	cli.Chat = chat.New(transport, formats)

	cli.Execution = execution.New(transport, formats)

	cli.Funding = funding.New(transport, formats)

	cli.Instrument = instrument.New(transport, formats)

	cli.Insurance = insurance.New(transport, formats)

	cli.Leaderboard = leaderboard.New(transport, formats)

	cli.Liquidation = liquidation.New(transport, formats)

	cli.Notification = notification.New(transport, formats)

	cli.Order = order.New(transport, formats)

	cli.OrderBook = order_book.New(transport, formats)

	cli.Position = position.New(transport, formats)

	cli.Quote = quote.New(transport, formats)

	cli.Schema = schema.New(transport, formats)

	cli.Settlement = settlement.New(transport, formats)

	cli.Stats = stats.New(transport, formats)

	cli.Trade = trade.New(transport, formats)

	cli.User = user.New(transport, formats)

	return cli
}

// DefaultTransportConfig creates a TransportConfig with the
// default settings taken from the meta section of the spec file.
func DefaultTransportConfig() *TransportConfig {
	return &TransportConfig{
		Host:     DefaultHost,
		BasePath: DefaultBasePath,
		Schemes:  DefaultSchemes,
	}
}

// TransportConfig contains the transport related info,
// found in the meta section of the spec file.
type TransportConfig struct {
	Host     string
	BasePath string
	Schemes  []string
}

// WithHost overrides the default host,
// provided by the meta section of the spec file.
func (cfg *TransportConfig) WithHost(host string) *TransportConfig {
	cfg.Host = host
	return cfg
}

// WithBasePath overrides the default basePath,
// provided by the meta section of the spec file.
func (cfg *TransportConfig) WithBasePath(basePath string) *TransportConfig {
	cfg.BasePath = basePath
	return cfg
}

// WithSchemes overrides the default schemes,
// provided by the meta section of the spec file.
func (cfg *TransportConfig) WithSchemes(schemes []string) *TransportConfig {
	cfg.Schemes = schemes
	return cfg
}

// APIClient is a client for API client
type APIClient struct {
	Announcement *announcement.Client

	APIKey *api_key.Client

	Chat *chat.Client

	Execution *execution.Client

	Funding *funding.Client

	Instrument *instrument.Client

	Insurance *insurance.Client

	Leaderboard *leaderboard.Client

	Liquidation *liquidation.Client

	Notification *notification.Client

	Order *order.Client

	OrderBook *order_book.Client

	Position *position.Client

	Quote *quote.Client

	Schema *schema.Client

	Settlement *settlement.Client

	Stats *stats.Client

	Trade *trade.Client

	User *user.Client

	Transport runtime.ClientTransport
}

// SetTransport changes the transport on the client and all its subresources
func (c *APIClient) SetTransport(transport runtime.ClientTransport) {
	c.Transport = transport

	c.Announcement.SetTransport(transport)

	c.APIKey.SetTransport(transport)

	c.Chat.SetTransport(transport)

	c.Execution.SetTransport(transport)

	c.Funding.SetTransport(transport)

	c.Instrument.SetTransport(transport)

	c.Insurance.SetTransport(transport)

	c.Leaderboard.SetTransport(transport)

	c.Liquidation.SetTransport(transport)

	c.Notification.SetTransport(transport)

	c.Order.SetTransport(transport)

	c.OrderBook.SetTransport(transport)

	c.Position.SetTransport(transport)

	c.Quote.SetTransport(transport)

	c.Schema.SetTransport(transport)

	c.Settlement.SetTransport(transport)

	c.Stats.SetTransport(transport)

	c.Trade.SetTransport(transport)

	c.User.SetTransport(transport)

}
