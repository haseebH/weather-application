package config

import (
	"fmt"
	"log"
	"os"

	"github.com/spf13/viper"
	"github.com/urfave/cli/v2"
)

var cfg *Config

type Config struct {
	ServerPort   string        `mapstructure:"port"`
	Database     *database     `mapstructure:"database"`
	MessageQueue *MessageQueue `mapstructure:"message_queue"`
}
type MessageQueue struct {
	RabbitMQURI  string `mapstructure:"rabbitmq_uri"`
	ExchangeName string `mapstructure:"exchange_name"`
	RoutingKey   string `mapstructure:"routing_key"`
}
type database struct {
	Host        string `mapstructure:"host"`
	Port        int64  `mapstructure:"port"`
	User        string `mapstructure:"user"`
	Password    string `mapstructure:"password"`
	DBName      string `mapstructure:"db_name"`
	MinPoolSize uint64 `mapstructure:"min_pool"`
	MaxPoolSize uint64 `mapstructure:"max_pool"`
}

func LoadConfig() *Config {
	if cfg != nil {
		return cfg
	}
	cfg = new(Config)
	fmt.Println(os.Getenv("ENV"))
	if os.Getenv("ENV") == "prod" {
		cfg.Database = new(database)
		cfg.MessageQueue = new(MessageQueue)
		loadENV(cfg)
		return cfg
	}

	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error reading config file, %s", err)
	}
	err := viper.Unmarshal(cfg)
	if err != nil {
		log.Fatalf("Unable to decode into struct, %v", err)
	}

	return cfg
}

func loadENV(cfg *Config) {
	app := &cli.App{
		Name:  "RBAC Application",
		Usage: "authentication",
		Action: func(*cli.Context) error {
			return nil
		},
	}
	app.Flags = []cli.Flag{
		&cli.StringFlag{
			Name:        "port",
			Usage:       "port for the application",
			Destination: &cfg.ServerPort,
			EnvVars:     []string{"PORT"},
		},
		&cli.StringFlag{
			Name:        "database-host",
			Usage:       "rbac manager url  http://rbac-manager:6971",
			Destination: &cfg.Database.Host,
			EnvVars:     []string{"DATABASE_HOST"},
			Required:    true,
		},
		&cli.Int64Flag{
			Name:        "database-port",
			Usage:       "root token",
			Destination: &cfg.Database.Port,
			EnvVars:     []string{"DATABASE_PORT"},
			Required:    true,
		},
		&cli.StringFlag{
			Name:        "database-user",
			Usage:       "audit manager path, http://audit-url:8973",
			Destination: &cfg.Database.User,
			EnvVars:     []string{"DATABASE_USER"},
			Required:    true,
		},
		&cli.StringFlag{
			Name:        "database-password",
			Usage:       "audit manager path, http://audit-url:8973",
			Destination: &cfg.Database.Password,
			EnvVars:     []string{"DATABASE_PASSWORD"},
			Required:    true,
		},
		&cli.StringFlag{
			Name:        "database-db",
			Usage:       "audit manager path, http://audit-url:8973",
			Destination: &cfg.Database.DBName,
			EnvVars:     []string{"DATABASE_DB"},
			Required:    true,
		},
		&cli.Uint64Flag{
			Name:        "database-min-pool",
			Usage:       "audit manager path, http://audit-url:8973",
			Destination: &cfg.Database.MinPoolSize,
			EnvVars:     []string{"DATABASE_MIN_POOL"},
			Required:    true,
		},
		&cli.Uint64Flag{
			Name:        "database-max-pool",
			Usage:       "audit manager path, http://audit-url:8973",
			Destination: &cfg.Database.MaxPoolSize,
			EnvVars:     []string{"DATABASE_MAX_POOL"},
			Required:    true,
		},
		&cli.StringFlag{
			Name:        "message-queue-uri",
			Usage:       "audit manager path, http://audit-url:8973",
			Destination: &cfg.MessageQueue.RabbitMQURI,
			EnvVars:     []string{"MESSAGE_QUEUE_URI"},
			Required:    true,
		},
		&cli.StringFlag{
			Name:        "message-queue-exchange-name",
			Usage:       "audit manager path, http://audit-url:8973",
			Destination: &cfg.MessageQueue.ExchangeName,
			EnvVars:     []string{"MESSAGE_QUEUE_EXCHANGE_NAME"},
			Required:    true,
		},
		&cli.StringFlag{
			Name:        "message-queue-routing-key",
			Usage:       "audit manager path, http://audit-url:8973",
			Destination: &cfg.MessageQueue.RoutingKey,
			EnvVars:     []string{"MESSAGE_QUEUE_ROUTING_KEY"},
			Required:    true,
		},
	}
	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
