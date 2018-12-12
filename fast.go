package fast

// Config fast ftp config.
type Config struct {
	Address  string
	User     string
	BuffSize string
}

// Fast fast ftp struct.
type Fast struct {
}

// New new a fast ftp.
func New(c *Config) *Fast {
	return &Fast{}
}
