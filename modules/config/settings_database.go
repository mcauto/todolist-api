package config

// DatabaseCmd represents the database command
func (s Settings) DatabaseDSN() string {
	return s.Database.DSN
}
