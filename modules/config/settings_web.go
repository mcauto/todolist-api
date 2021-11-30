package config

import "fmt"

// BindAddress is the address to bind to
func (s Settings) BindAddress() string {
	return fmt.Sprintf(":%d", s.App.Port)
}
