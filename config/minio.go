package config

type Minio struct {
	Location        string `mapstructure:"location"  json:"location" yaml:"location"`
	BucketName      string `mapstructure:"bucket_name"  json:"bucket_name" yaml:"bucket_name"`
	Endpoint        string `mapstructure:"endpoint" json:"endpoint" yaml:"endpoint"`
	AccessKeyID     string `mapstructure:"accessKeyID" json:"accessKeyID" yaml:"accessKeyID"`
	SecretAccessKey string `mapstructure:"secretAccessKey" json:"secretAccessKey" yaml:"secretAccessKey"`
	Secure          bool   `mapstructure:"secure" json:"secure" yaml:"secure"`
}
