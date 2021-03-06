package conf

import (
	"flag"
	"fmt"
	"os"

	"github.com/pion/ion/pkg/node/avp"
	"github.com/spf13/viper"
)

var (
	cfg     = Config{}
	Global  = &cfg.Global
	Etcd    = &cfg.Etcd
	Nats    = &cfg.Nats
	Avp     = &cfg.Avp
	CfgFile = &cfg.CfgFile
)

func init() {
	if !cfg.parse() {
		showHelp()
		os.Exit(-1)
	}
}

type global struct {
	Addr  string `mapstructure:"addr"`
	Pprof string `mapstructure:"pprof"`
	Dc    string `mapstructure:"dc"`
}

type log struct {
	Level string `mapstructure:"level"`
}

type etcd struct {
	Addrs []string `mapstructure:"addrs"`
}

type nats struct {
	URL string `mapstructure:"url"`
}

// Config for base AVP
type Config struct {
	Global  global     `mapstructure:"global"`
	Etcd    etcd       `mapstructure:"etcd"`
	Nats    nats       `mapstructure:"nats"`
	Avp     avp.Config `mapstructure:"avp"`
	CfgFile string
}

func showHelp() {
	fmt.Printf("Usage:%s {params}\n", os.Args[0])
	fmt.Println("      -c {config file}")
	fmt.Println("      -h (show help info)")
}

func (c *Config) unmarshal(rawVal interface{}) bool {
	if err := viper.Unmarshal(rawVal); err != nil {
		fmt.Printf("config file %s loaded failed. %v\n", c.CfgFile, err)
		return false
	}
	return true
}

func (c *Config) load() bool {
	_, err := os.Stat(c.CfgFile)
	if err != nil {
		return false
	}

	viper.SetConfigFile(c.CfgFile)
	viper.SetConfigType("toml")

	err = viper.ReadInConfig()
	if err != nil {
		fmt.Printf("config file %s read failed. %v\n", c.CfgFile, err)
		return false
	}
	if !c.unmarshal(&c) || !c.unmarshal(&c.Avp) || !c.unmarshal(&c.Avp.Config) {
		return false
	}

	fmt.Printf("config %s load ok!\n", c.CfgFile)
	return true
}

func (c *Config) parse() bool {

	flag.StringVar(&c.CfgFile, "c", "conf/conf.toml", "config file")
	help := flag.Bool("h", false, "help info")
	flag.Parse()
	if !c.load() {
		return false
	}

	if *help {
		showHelp()
		return false
	}
	return true
}
