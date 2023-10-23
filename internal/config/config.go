package config

// Config global config instance.
var Config *config

// ConfigInit init config.
func ConfigInit(configPath string) {
	// init viper
	initViper(configPath)

	// init Configuration
	Config = &config{
		App: App{
			Version: GetString("app.version"),
			Debug:   GetBool("app.debug"),
			LogFile: GetString("app.log_file"),
		},
		Server: Server{
			IP:   GetString("server.ip"),
			Port: GetInt("server.port"),
		},
		Agent: Agent{
			UUID:   GetString("agent.uuid"),
			Period: GetInt("agent.period"),
			GZip:   GetBool("agent.gzip"),
		},
		Redis: Redis{
			Ip:       GetString("redis.ip"),
			Port:     GetInt("redis.port"),
			Password: GetString("redis.password"),
			Database: GetInt("redis.database"),
		},
		DM: DM{
			Ip:       GetString("dm.ip"),
			Port:     GetInt("dm.port"),
			Username: GetString("dm.username"),
			Password: GetString("dm.password"),
		},
	}
}

type config struct {
	App    App
	Server Server
	Agent  Agent
	Redis  Redis
	DM     DM
}

// App config.
type App struct {
	Version string `mapstructure:"version"`
	Debug   bool   `mapstructure:"debug"`
	LogFile string `mapstructure:"log_file"`
}

// Server config.
type Server struct {
	IP   string `mapstructure:"ip"`
	Port int    `mapstructure:"port"`
}

// Agent config.
type Agent struct {
	UUID   string `mapstructure:"uuid"`
	Period int    `mapstructure:"period"`
	GZip   bool   `mapstructure:"gzip"`
}

// Redis config.
type Redis struct {
	Ip       string `mapstructure:"ip"`
	Port     int    `mapstructure:"port"`
	Password string `mapstructure:"password"`
	Database int    `mapstructure:"database"`
}

// DM config.
type DM struct {
	Ip       string `mapstructure:"ip"`
	Port     int    `mapstructure:"port"`
	Username string `mapstructure:"username"`
	Password string `mapstructure:"password"`
}
