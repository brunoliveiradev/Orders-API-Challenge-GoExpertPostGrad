package configs

import (
	"github.com/spf13/viper"
	"os"
	"strconv"
)

type Envs struct {
	DBDriver          string `mapstructure:"DB_DRIVER"`
	DBHost            string `mapstructure:"DB_HOST"`
	DBPort            string `mapstructure:"DB_PORT"`
	DBUser            string `mapstructure:"DB_USER"`
	DBPassword        string `mapstructure:"DB_PASSWORD"`
	DBName            string `mapstructure:"DB_NAME"`
	DBMaxOpenConns    int    `mapstructure:"DB_MAX_OPEN_CONNS"`
	DBMaxIdleConns    int    `mapstructure:"DB_MAX_IDLE_CONNS"`
	WebServerPort     string `mapstructure:"WEB_SERVER_PORT"`
	GRPCServerPort    string `mapstructure:"GRPC_SERVER_PORT"`
	GraphQLServerPort string `mapstructure:"GRAPHQL_SERVER_PORT"`
	RabbitMQHost      string `mapstructure:"RABBITMQ_HOST"`
	RabbitMQUser      string `mapstructure:"RABBITMQ_USER"`
	RabbitMQPassword  string `mapstructure:"RABBITMQ_PASSWORD"`
	RabbitMQPort      string `mapstructure:"RABBITMQ_PORT"`
}

func LoadConfig() (*Envs, error) {
	// Tente carregar o arquivo .env
	viper.SetConfigFile("./.env")
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err == nil {
		// Sucesso ao ler o arquivo .env
		var cfg Envs
		if err := viper.Unmarshal(&cfg); err != nil {
			return nil, err
		}
		return &cfg, nil
	}

	// Falha ao ler o arquivo .env, carregue as variáveis de ambiente
	return loadLocalConfig(), nil
}

func getEnv(key, fallback string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return fallback
}

func loadLocalConfig() *Envs {
	dbMaxOpenConns, _ := strconv.Atoi(getEnv("DB_MAX_OPEN_CONNS", "10"))

	dbMaxIdleConns, _ := strconv.Atoi(getEnv("DB_MAX_IDLE_CONNS", "10"))

	return &Envs{
		DBDriver:          getEnv("DB_DRIVER", "mysql"),
		DBHost:            getEnv("DB_HOST", "localhost"),
		DBPort:            getEnv("DB_PORT", "3306"),
		DBUser:            getEnv("DB_USER", "root"),
		DBPassword:        getEnv("DB_PASSWORD", "root"),
		DBName:            getEnv("DB_NAME", "orders"),
		DBMaxOpenConns:    dbMaxOpenConns,
		DBMaxIdleConns:    dbMaxIdleConns,
		WebServerPort:     getEnv("WEB_SERVER_PORT", ":8000"),
		GRPCServerPort:    getEnv("GRPC_SERVER_PORT", "50051"),
		GraphQLServerPort: getEnv("GRAPHQL_SERVER_PORT", "8080"),
		RabbitMQHost:      getEnv("RABBITMQ_HOST", "localhost"),
		RabbitMQUser:      getEnv("RABBITMQ_USER", "guest"),
		RabbitMQPassword:  getEnv("RABBITMQ_PASSWORD", "guest"),
		RabbitMQPort:      getEnv("RABBITMQ_PORT", "5672"),
	}
}
