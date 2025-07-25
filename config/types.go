package config

// appEnvironments represents the configuration settings for the application environments.
type appEnvironments struct {
	// Name is the name of the application
	Name string `envconfig:"APP_NAME" required:"true" default:"kafrest-api"`
	// Env is the environment in which the application is running (e.g., development, production)
	Env string `envconfig:"APP_ENV" required:"true" default:"development"`
}

// apiEnvironments represents the configuration settings for the API environments.
type apiEnvironments struct {
	// Port is the port on which the API server will listen for incoming requests.
	Port string `envconfig:"API_PORT" required:"true" default:"8080"`
}

// kafkaEnvironments represents the configuration settings for connecting to a Kafka broker.
type kafkaEnvironments struct {
	// BrokerConnect is the connection string for the Kafka broker.
	BrokerConnect string `envconfig:"KAFKA_BROKER_CONNECT" required:"true" default:"localhost:9092"`
}
