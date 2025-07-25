package config

// Config holds all the application environments.
type Config struct {
	App   appEnvironments
	API   apiEnvironments
	Kafka kafkaEnvironments
}
