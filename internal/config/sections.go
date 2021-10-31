package config

type logs struct{}
type database struct{}
type server struct{}

// logs section
func (l logs) LogLevel() int {
	val, _ := cfg.GetInt("logs", "log_level")
	return val
}

func (l logs) ConsoleLogging() bool {
	val, _ := cfg.GetBool("logs", "console_logging")
	return val
}

func (l logs) FileLogging() bool {
	val, _ := cfg.GetBool("logs", "file_logging")
	return val
}

func (l logs) LogPath() string {
	val, _ := cfg.GetString("logs", "log_path")
	return val
}

func (l logs) MaxSize() int {
	val, _ := cfg.GetInt("logs", "max_age")
	return val
}

func (l logs) MaxAge() int {
	val, _ := cfg.GetInt("logs", "max_age")
	return val
}

func (l logs) MaxBackups() int {
	val, _ := cfg.GetInt("logs", "max_age")
	return val
}

// database
func (d database) Host() string {
	val, _ := cfg.GetString("database", "host")
	return val
}

func (d database) Port() int {
	val, _ := cfg.GetInt("database", "port")
	return val
}

func (d database) User() string {
	val, _ := cfg.GetString("database", "user")
	return val
}

func (d database) Password() string {
	val, _ := cfg.GetString("database", "password")
	return val
}

func (d database) DbName() string {
	val, _ := cfg.GetString("database", "db_name")
	return val
}

func (d database) SslMode() string {
	val, _ := cfg.GetString("database", "ssl_mode")
	return val
}

// server
func (s server) ListenAddress() string {
	val, _ := cfg.GetString("server", "listen_address")
	return val
}

func (s server) Port() int {
	val, _ := cfg.GetInt("server", "port")
	return val
}

func (s server) BaseAddress() string {
	val, _ := cfg.GetString("server", "base_address")
	return val
}

var (
	Logs     = logs{}
	Server   = server{}
	Database = database{}
	Login    = login{}
)
