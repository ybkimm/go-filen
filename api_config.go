package filen

import (
	"fmt"
	"net/http"
	"time"

	"github.com/ybkimm/go-filen/internal/httpx"
)

type Config struct {
	Maintenance   bool                 `json:"maintenance"`
	ReadOnly      bool                 `json:"readOnly"`
	Announcements []ConfigAnnouncement `json:"announcements"`
	Pricing       struct {
		LifetimeEnabled bool          `json:"lifetimeEnabled"`
		SaleEnabled     bool          `json:"saleEnabled"`
		Plans           []ConfigPlans `json:"plans"`
	} `json:"pricing"`
}

type ConfigAnnouncement struct {
	UUID      string `json:"uuid"`
	Title     string `json:"title"`
	Message   string `json:"message"`
	Active    bool   `json:"active"`
	Timestamp int    `json:"timestamp"`
}

type ConfigPlans struct {
	TermType int    `json:"termType"`
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Cost     int    `json:"cost"`
	Sale     int    `json:"sale"`
	Storage  int    `json:"storage"`
	Popular  bool   `json:"popular"`
	Term     string `json:"term"`
}

type ConfigPlanTermType int

const (
	ConfigPlanTermTypeStarter  ConfigPlanTermType = 1
	ConfigPlanTermTypeMonthly  ConfigPlanTermType = 2
	ConfigPlanTermTypeAnnually ConfigPlanTermType = 3
	ConfigPlanTermTypeLifetime ConfigPlanTermType = 4
)

type ConfigPlanTerm string

const (
	ConfigPlanTermMonthly  ConfigPlanTerm = "monthly"
	ConfigPlanTermAnnually ConfigPlanTerm = "annually"
	ConfigPlanTermLifetime ConfigPlanTerm = "lifetime"
)

func (c *APIClient) GetConfig() (*Config, error) {
	var config Config

	// ConfigEndpoint is a CDN endpoint. Since it's a CDN endpoint,
	// it does not affected by the endpoints list.
	_, err := httpx.DoJsonRequest(
		c.httpClient,
		http.MethodGet,
		fmt.Sprintf("%s?noCache=%d", ConfigEndpoint, time.Now().UnixNano()),
		nil,
		&config,
	)
	if err != nil {
		return nil, err
	}

	return &config, nil
}
