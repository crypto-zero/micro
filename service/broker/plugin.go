package broker

import (
	"fmt"

	"github.com/crypto-zero/micro/v2/plugin"
)

var (
	defaultManager = plugin.NewManager()
)

// Plugins lists the broker plugins
func Plugins() []plugin.Plugin {
	return defaultManager.Plugins()
}

// Register registers an broker plugin
func Register(pl plugin.Plugin) error {
	if plugin.IsRegistered(pl) {
		return fmt.Errorf("%s registered globally", pl.String())
	}
	return defaultManager.Register(pl)
}
