package runner

import (
	"errors"
	"fmt"
	"github.com/azate/giex/internal/utils"
	"github.com/spf13/viper"
	"os"
	"time"
)

type config struct {
	Input      string
	Output     string
	MaxWorkers uint
	MaxTasks   uint
	Proxy      string
}

func LoadConfig() *config {
	return &config{
		Input:      viper.GetString("input"),
		Output:     viper.GetString("output"),
		MaxWorkers: viper.GetUint("max-workers"),
		MaxTasks:   viper.GetUint("max-tasks"),
		Proxy:      viper.GetString("proxy"),
	}
}

func (cfg *config) Check() error {
	if !utils.FileExists(cfg.Input) {
		return errors.New("file with domains not found")
	}

	if cfg.Output == "/tmp" {
		cfg.Output += fmt.Sprintf("/GIEX_%s", time.Now().Format("20060102_150405"))
	}

	if !utils.DirExists(cfg.Output) {
		if err := os.MkdirAll(cfg.Output, 0600); err != nil {
			return err
		}
	}

	return nil
}
