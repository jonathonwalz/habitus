package configuration

import (
	"errors"
	"strings"

	"github.com/op/go-logging"
)

type TupleItem struct {
	Key   string
	Value string
}

type TupleArray []TupleItem

// Recognized OS type for --os switch.
// Needs to agree with builder.go
var OsTypes = []string{
	"debian",
	"redhat",
	"busybox",
	"alpine",
}

type StringArray []string

// Config stores application configurations
type Config struct {
	Buildfile                         string
	Workdir                           string
	NoCache                           bool
	SuppressOutput                    bool
	RmTmpContainers                   bool
	ForceRmTmpContainer               bool
	UniqueID                          string
	Logger                            logging.Logger
	DockerHost                        string
	DockerCert                        string
	EnvVars                           TupleArray
	BuildArgs                         TupleArray
	KeepSteps                         bool
	KeepArtifacts                     bool
	Network                           string
	NoSquash                          bool
	NoPruneRmImages                   bool
	UseTLS                            bool
	UseStatForPermissions             bool
	FroceRmImages                     bool
	OsType                            string
	ApiPort                           int
	ApiBinding                        string
	SecretService                     bool
	AllowAfterBuildCommands           bool
	SecretProviders                   string
	DockerMemory                      string
	DockerCPUSetCPUs                  string
	DockerCPUShares                   int
	UseAuthenticatedSecretServer      bool
	AuthenticatedSecretServerPassword string
	AuthenticatedSecretServerUser     string
	CacheFrom                         StringArray
}

func (i *TupleArray) String() string {
	return ""
}

func (i *TupleArray) Set(value string) error {
	parts := strings.Split(value, "=")

	if len(parts) != 2 {
		return errors.New("invalid key/value format (key=value)")
	}

	item := TupleItem{Key: parts[0], Value: parts[1]}
	*i = append(*i, item)
	return nil
}

func (i *TupleArray) Find(key string) string {
	for _, item := range *i {
		if item.Key == key {
			return item.Value
		}
	}

	return ""
}

func (i *StringArray) String() string {
    return ""
}

func (i *StringArray) Set(value string) error {
    *i = append(*i, value)
    return nil
}

// CreateConfig creates a new configuration object
func CreateConfig() Config {
	return Config{}
}

func (c *Config) ValidateOsType() bool {
	for _, os := range OsTypes {
		if c.OsType == os {
			return true
		}
	}
	return false
}
