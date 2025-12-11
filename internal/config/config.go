// Package config loads YAML configuration and defines relay settings for
// logging, sinks, and runtime options.
package config

import (
	"fmt"
	"os"
	"time"

	"gopkg.in/yaml.v3"
)

// Config represents the application configuration
type Config struct {
	Sinks       SinksConfig   `yaml:"sinks"`
	Logging     LoggingConfig `yaml:"logging"`
	Debug       bool
	TLSCertPath string
	TLSKeyPath  string
	GrpcPort    int
	BufferSize  int
}

// SinksConfig contains configuration for all data sinks
type SinksConfig struct {
	S3            *S3Config            `yaml:"s3,omitempty"`
	GCS           *GCSConfig           `yaml:"gcs,omitempty"`
	BigQuery      *BigQueryConfig      `yaml:"bigquery,omitempty"`
	Timestream    *TimestreamConfig    `yaml:"timestream,omitempty"`
	InfluxDB      *InfluxDBConfig      `yaml:"influxdb,omitempty"`
	Prometheus    *PrometheusConfig    `yaml:"prometheus,omitempty"`
	Elasticsearch *ElasticsearchConfig `yaml:"elasticsearch,omitempty"`
	Kafka         *KafkaConfig         `yaml:"kafka,omitempty"`
	File          *FileConfig          `yaml:"file,omitempty"`
	NATS          *NATSConfig          `yaml:"nats,omitempty"`
}

// S3Config contains S3 sink configuration
type S3Config struct {
	Bucket             string        `yaml:"bucket"`
	Region             string        `yaml:"region"`
	AccessKey          string        `yaml:"access_key"`
	SecretKey          string        `yaml:"secret_key"`
	Prefix             string        `yaml:"prefix"`
	FlushInterval      time.Duration `yaml:"flush_interval"`
	QueueSize          int           `yaml:"queue_size"`
	BackpressurePolicy string        `yaml:"backpressure_policy"`
}

// GCSConfig contains Google Cloud Storage sink configuration
type GCSConfig struct {
	Bucket             string        `yaml:"bucket"`
	ProjectID          string        `yaml:"project_id"`
	Credentials        string        `yaml:"credentials"` // Path to service account JSON file
	Prefix             string        `yaml:"prefix"`
	FlushInterval      time.Duration `yaml:"flush_interval"` // How often to flush buffered data (e.g., "30s")
	QueueSize          int           `yaml:"queue_size"`
	BackpressurePolicy string        `yaml:"backpressure_policy"`
}

// BigQueryConfig contains BigQuery sink configuration
type BigQueryConfig struct {
	ProjectID          string `yaml:"project_id"`
	Dataset            string `yaml:"dataset"`
	Table              string `yaml:"table"`
	Credentials        string `yaml:"credentials"`    // Path to service account JSON file
	BatchSize          int    `yaml:"batch_size"`     // Number of messages to batch before insert
	FlushInterval      string `yaml:"flush_interval"` // How often to flush (e.g., "30s", "1m")
	QueueSize          int    `yaml:"queue_size"`
	BackpressurePolicy string `yaml:"backpressure_policy"`
}

// TimestreamConfig contains AWS Timestream sink configuration
type TimestreamConfig struct {
	Database           string `yaml:"database"`
	Table              string `yaml:"table"`
	Region             string `yaml:"region"`
	AccessKey          string `yaml:"access_key"`
	SecretKey          string `yaml:"secret_key"`
	SessionToken       string `yaml:"session_token,omitempty"` // For temporary credentials
	BatchSize          int    `yaml:"batch_size"`              // Number of records to batch
	FlushInterval      string `yaml:"flush_interval"`          // How often to flush (e.g., "30s", "1m")
	QueueSize          int    `yaml:"queue_size"`
	BackpressurePolicy string `yaml:"backpressure_policy"`
}

// InfluxDBConfig contains InfluxDB sink configuration
type InfluxDBConfig struct {
	URL                string `yaml:"url"`
	Database           string `yaml:"database"`
	Username           string `yaml:"username"`
	Password           string `yaml:"password"`
	Token              string `yaml:"token"`        // For InfluxDB 2.x
	Organization       string `yaml:"organization"` // For InfluxDB 2.x
	Bucket             string `yaml:"bucket"`       // For InfluxDB 2.x
	BatchSize          int    `yaml:"batch_size"`
	FlushInterval      string `yaml:"flush_interval"`
	QueueSize          int    `yaml:"queue_size"`
	BackpressurePolicy string `yaml:"backpressure_policy"`
}

// PrometheusConfig contains Prometheus sink configuration
type PrometheusConfig struct {
	URL                string `yaml:"url"`
	Job                string `yaml:"job"`
	Instance           string `yaml:"instance"`
	BatchSize          int    `yaml:"batch_size"`
	FlushInterval      string `yaml:"flush_interval"`
	QueueSize          int    `yaml:"queue_size"`
	BackpressurePolicy string `yaml:"backpressure_policy"`
}

// ElasticsearchConfig contains Elasticsearch sink configuration
type ElasticsearchConfig struct {
	URLs               []string `yaml:"urls"`
	Index              string   `yaml:"index"`
	Username           string   `yaml:"username"`
	Password           string   `yaml:"password"`
	APIKey             string   `yaml:"api_key"`
	BatchSize          int      `yaml:"batch_size"`
	FlushInterval      string   `yaml:"flush_interval"`
	QueueSize          int      `yaml:"queue_size"`
	BackpressurePolicy string   `yaml:"backpressure_policy"`
}

// KafkaConfig contains Kafka sink configuration
type KafkaConfig struct {
	Brokers            []string `yaml:"brokers"`
	Topic              string   `yaml:"topic"`
	QueueSize          int      `yaml:"queue_size"`
	BackpressurePolicy string   `yaml:"backpressure_policy"`
}

// FileConfig contains file-based sink configuration
type FileConfig struct {
	Path               string        `yaml:"path"`              // Path to the file, without the filename
	Prefix             string        `yaml:"prefix"`            // Prefix for the filename, will be appended to the path
	Format             string        `yaml:"format"`            // json, csv, binary
	RotationInterval   time.Duration `yaml:"rotation_interval"` // 24h, 1h, 10m, etc.
	QueueSize          int           `yaml:"queue_size"`
	BackpressurePolicy string        `yaml:"backpressure_policy"`
}

// NATSConfig contains NATS JetStream sink configuration
type NATSConfig struct {
	URL                string        `yaml:"url"`
	Subject            string        `yaml:"subject"`              // Template: "{entity_id}.mavlink" or static "mavlink.telemetry"
	Token              string        `yaml:"token,omitempty"`      // JWT token for auth
	CredsFile          string        `yaml:"creds_file,omitempty"` // Path to credentials file
	QueueSize          int           `yaml:"queue_size"`
	BackpressurePolicy string        `yaml:"backpressure_policy"`
	Stream             *StreamConfig `yaml:"stream,omitempty"` // JetStream configuration
	KV                 *KVConfig     `yaml:"kv,omitempty"`     // KeyValue store configuration
}

// StreamConfig contains NATS JetStream stream configuration
type StreamConfig struct {
	Name        string   `yaml:"name"`                  // Stream name
	Subjects    []string `yaml:"subjects"`              // Subject patterns for the stream
	Storage     string   `yaml:"storage,omitempty"`     // "memory" or "file"
	Replicas    int      `yaml:"replicas,omitempty"`    // Number of replicas
	MaxAge      string   `yaml:"max_age,omitempty"`     // Message retention period
	MaxBytes    int64    `yaml:"max_bytes,omitempty"`   // Max bytes stored
	MaxMsgs     int64    `yaml:"max_msgs,omitempty"`    // Max messages stored
	Compression bool     `yaml:"compression,omitempty"` // Enable compression
}

// KVConfig contains NATS KeyValue store configuration for device state
type KVConfig struct {
	Bucket      string   `yaml:"bucket"`                // KV bucket name
	KeyPattern  string   `yaml:"key_pattern"`           // Key pattern: "{entity_id}.mavlink"
	TTL         string   `yaml:"ttl,omitempty"`         // Value TTL (e.g., "1h", "24h")
	MaxBytes    int64    `yaml:"max_bytes,omitempty"`   // Max bytes for bucket
	Replicas    int      `yaml:"replicas,omitempty"`    // Number of replicas
	Storage     string   `yaml:"storage,omitempty"`     // "memory" or "file"
	Description string   `yaml:"description,omitempty"` // Bucket description
	MessageTypes []string `yaml:"message_types,omitempty"` // Message types to track (e.g., ["Heartbeat", "GlobalPositionInt"])
}

// LoggingConfig contains logging configuration
type LoggingConfig struct {
	Level  string `yaml:"level"`
	Format string `yaml:"format"` // json, text
	Output string `yaml:"output"` // stdout, file
	File   string `yaml:"file,omitempty"`
}

// Load loads configuration from a YAML file
func Load(path string) (*Config, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("%w: %w", ErrFailedToReadConfigFile, err)
	}

	dataStr := os.ExpandEnv(string(data))

	var config Config
	if err := yaml.Unmarshal([]byte(dataStr), &config); err != nil {
		return nil, fmt.Errorf("%w: %w", ErrFailedToParseConfigFile, err)
	}

	if config.Logging.Level == "" {
		config.Logging.Level = "info"
	}
	if config.Logging.Format == "" {
		config.Logging.Format = "text"
	}
	if config.Logging.Output == "" {
		config.Logging.Output = "stdout"
	}

	return &config, nil
}
