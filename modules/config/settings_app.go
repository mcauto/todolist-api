package config

// Debug mode
func (s Settings) Debug() bool {
	return s.App.Environment == "development"
}
