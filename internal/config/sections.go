package config

type logs struct{}

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

var (
	Logs = logs{}
)
