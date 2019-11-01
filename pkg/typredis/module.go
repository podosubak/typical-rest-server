package typredis

import (
	"fmt"
	"os"
	"os/exec"
	"time"

	"github.com/go-redis/redis"
	"github.com/typical-go/typical-rest-server/EXPERIMENTAL/typiobj"
	"github.com/typical-go/typical-rest-server/pkg/utility/envkit"
	"github.com/urfave/cli"
)

// Config is Redis Configuration
type Config struct {
	Host     string `required:"true" default:"localhost"`
	Port     string `required:"true" default:"6379"`
	Password string `default:"redispass"`
	DB       int    `default:"0"`

	PoolSize           int           `envconfig:"POOL_SIZE"  default:"20" required:"true"`
	DialTimeout        time.Duration `envconfig:"DIAL_TIMEOUT" default:"5s" required:"true"`
	ReadWriteTimeout   time.Duration `envconfig:"READ_WRITE_TIMEOUT" default:"3s" required:"true"`
	IdleTimeout        time.Duration `envconfig:"IDLE_TIMEOUT" default:"5m" required:"true"`
	IdleCheckFrequency time.Duration `envconfig:"IDLE_CHECK_FREQUENCY" default:"1m" required:"true"`
	MaxConnAge         time.Duration `envconfig:"MAX_CONN_AGE" default:"30m" required:"true"`
}

// Module of redis
func Module() interface{} {
	return &redisModule{
		Name: "Redis",
		Configuration: typiobj.Configuration{
			Prefix: "REDIS",
			Spec:   &Config{},
		},
	}
}

type redisModule struct {
	typiobj.Configuration
	Name string
}

// Provide dependencies
func (r redisModule) Provide() []interface{} {
	return []interface{}{
		r.loadConfig,
		r.connect,
	}
}

// Destroy dependencies
func (r redisModule) Destroy() []interface{} {
	return []interface{}{
		r.disconnect,
	}
}

// CommandLine return command
func (r redisModule) CommandLine() cli.Command {
	return cli.Command{
		Name:   "redis",
		Usage:  "Redis Utility Tool",
		Before: envkit.CliLoadEnvFile,
		Subcommands: []cli.Command{
			{Name: "console", ShortName: "c", Action: typiobj.Action(r, r.console)},
		},
	}
}

func (r redisModule) loadConfig() (cfg *Config, err error) {
	err = r.Configuration.Load()
	cfg = r.Configuration.Spec.(*Config)
	return
}

func (redisModule) connect(cfg *Config) (client *redis.Client, err error) {
	client = redis.NewClient(&redis.Options{
		Addr:               fmt.Sprintf("%s:%s", cfg.Host, cfg.Port),
		Password:           cfg.Password,
		DB:                 cfg.DB,
		PoolSize:           cfg.PoolSize,
		DialTimeout:        cfg.DialTimeout,
		ReadTimeout:        cfg.ReadWriteTimeout,
		WriteTimeout:       cfg.ReadWriteTimeout,
		IdleTimeout:        cfg.IdleTimeout,
		IdleCheckFrequency: cfg.IdleCheckFrequency,
		MaxConnAge:         cfg.MaxConnAge,
	})
	err = client.Ping().Err()
	return
}

func (redisModule) disconnect(client *redis.Client) (err error) {
	fmt.Println("Redis Client close")
	return client.Close()
}

func (redisModule) console(config *Config) (err error) {
	args := []string{
		"-h", config.Host,
		"-p", config.Port,
	}
	if config.Password != "" {
		args = append(args, "-a", config.Password)
	}
	// TODO: using docker -it
	cmd := exec.Command("redis-cli", args...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stdout
	cmd.Stdin = os.Stdin
	return cmd.Run()
}