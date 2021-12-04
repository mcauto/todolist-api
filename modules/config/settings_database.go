package config

import "fmt"

// DatabaseCmd represents the database command
func (s Settings) DatabaseDSN() string {
	return fmt.Sprintf(
		"%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local",
		s.Database.User,
		s.Database.Password,
		s.Database.Host,
		s.Database.Port,
		s.Database.Name,
	)
}
