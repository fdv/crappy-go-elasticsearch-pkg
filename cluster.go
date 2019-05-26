package elasticsearch

import (
	"github.com/tidwall/gjson"
)

// Getsettings returns the cluster global settings
func (c *Cluster) Getsettings() string {
	return (c.buildgetquery("_cluster", "settings"))
}

// Getsetting returns a defined setting
func (c *Cluster) Getsetting(setting string) string {
	settings := c.Getsettings()

	result := gjson.Get(settings, setting)
	return (result.String())
}
