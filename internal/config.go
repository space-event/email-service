package internal

type Config struct {
	Addr     string     `toml:"addr"`
	Smtp     SMTPConfig `toml:"smtp"`
	LogLevel string     `toml:"loglevel"`
}

type SMTPConfig struct {
	HostSMTP     string `toml:"host_smtp"`
	PortSMTP     string `toml:"port_smtp"`
	EmailSMTP    string `toml:"email_smtp"`
	PasswordSMTP string `toml:"password_smtp"`
}
