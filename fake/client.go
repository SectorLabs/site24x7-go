package fake

import (
	"github.com/Bonial-International-GmbH/site24x7-go"
	"github.com/Bonial-International-GmbH/site24x7-go/api/endpoints"
	"github.com/Bonial-International-GmbH/site24x7-go/api/endpoints/fake"
)

var _ site24x7.Client = &Client{}

// Client is an implementation of site24x7.Client that stubs out all endpoints
// with mocks. In can be used in unit tests.
type Client struct {
	FakeCurrentStatus        *fake.CurrentStatus
	FakeITAutomations        *fake.ITAutomations
	FakeLocationTemplate     *fake.LocationTemplate
	FakeLocationProfiles     *fake.LocationProfiles
	FakeMonitorGroups        *fake.MonitorGroups
	FakeMonitors             *fake.Monitors
	FakeNotificationProfiles *fake.NotificationProfiles
	FakeThresholdProfiles    *fake.ThresholdProfiles
	FakeUserGroups           *fake.UserGroups
}

// NewClient creates a new fake site24x7 API client.
func NewClient() *Client {
	return &Client{
		FakeCurrentStatus:        &fake.CurrentStatus{},
		FakeITAutomations:        &fake.ITAutomations{},
		FakeLocationProfiles:     &fake.LocationProfiles{},
		FakeLocationTemplate:     &fake.LocationTemplate{},
		FakeMonitorGroups:        &fake.MonitorGroups{},
		FakeMonitors:             &fake.Monitors{},
		FakeNotificationProfiles: &fake.NotificationProfiles{},
		FakeThresholdProfiles:    &fake.ThresholdProfiles{},
		FakeUserGroups:           &fake.UserGroups{},
	}
}

// CurrentStatus implements Client.
func (c *Client) CurrentStatus() endpoints.CurrentStatus {
	return c.FakeCurrentStatus
}

// ItAutomations implements Client.
func (c *Client) ITAutomations() endpoints.ITAutomations {
	return c.FakeITAutomations
}

// LocationProfiles implements Client.
func (c *Client) LocationProfiles() endpoints.LocationProfiles {
	return c.FakeLocationProfiles
}

// LocationTemplate implements Client.
func (c *Client) LocationTemplate() endpoints.LocationTemplate {
	return c.FakeLocationTemplate
}

// Monitors implements Client.
func (c *Client) Monitors() endpoints.Monitors {
	return c.FakeMonitors
}

// MonitorGroups implements Client.
func (c *Client) MonitorGroups() endpoints.MonitorGroups {
	return c.FakeMonitorGroups
}

// NotificationProfiles implements Client.
func (c *Client) NotificationProfiles() endpoints.NotificationProfiles {
	return c.FakeNotificationProfiles
}

// ThresholdProfiles implements Client.
func (c *Client) ThresholdProfiles() endpoints.ThresholdProfiles {
	return c.FakeThresholdProfiles
}

// UserGroups implements Client.
func (c *Client) UserGroups() endpoints.UserGroups {
	return c.FakeUserGroups
}
