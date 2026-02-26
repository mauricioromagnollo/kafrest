package config

// AppEnvironments represents the configuration settings for the application environments.
// Name is the name of the application
// Env is the environment in which the application is running (e.g., development, production)
type AppEnvironments struct {
	Name     string `envconfig:"APP_NAME" required:"true" default:"kafrest-api"`
	Env      string `envconfig:"APP_ENV" required:"true" default:"development"`
	LogLevel string `envconfig:"APP_LOG_LEVEL" default:"info" required:"false"`
}

// APIEnvironments represents the configuration settings for the API environments.
// Port is the port on which the API server will listen for incoming requests.
type APIEnvironments struct {
	Port string `envconfig:"API_PORT" required:"true" default:"8080"`
}

// KafkaEnvironments represents the configuration settings for connecting to a Kafka broker.
// BrokerConnect is the connection string for the Kafka broker.
type KafkaEnvironments struct {
	BrokersServers string `envconfig:"KAFKA_BROKERS_SERVERS" required:"true"`
	GroupID        string `envconfig:"KAFKA_GROUP_ID" required:"true" default:"kafrest"`
	Username       string `envconfig:"KAFKA_USERNAME" required:"false" default:""`
	Password       string `envconfig:"KAFKA_PASSWORD" required:"false" default:""`
	Topics         struct {
		Sample string `envconfig:"KAFKA_TOPIC_SAMPLE" required:"true" default:"my.topic_sample.0"`
	}
}
