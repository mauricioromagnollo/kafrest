package config

// Config holds all the application environments.
type Config struct {
	App   AppEnvironments
	API   APIEnvironments
	Kafka KafkaEnvironments
}
