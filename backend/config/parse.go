package config

import (
	"os"
	"runtime"
	"runtime/pprof"
	"sync"
	"time"

	log "github.com/h2san/h2sanlog"

	"github.com/BurntSushi/toml"
)

var config Config
var once sync.Once

//GetConfig ...
func GetConfig() *Config {
	once.Do(func() {
		err := parseConfig(&config)
		if err != nil {
			panic(err)
		}
		log.Debug("[parse config] => %+v\n", config)

		//start log
		logInfo := config.Log
		logWriter, err := log.NewFileWriter(logInfo.Name, logInfo.MaxSize, logInfo.MaxNum)
		if err != nil {
			panic(err)
		}
		log.SetOutput(logWriter)
		config.Log.Writer = logWriter

		


		//start profile
		if config.Profile.CPU != "" {
			go func() {
				f, e := os.Create(config.Profile.CPU)
				if e != nil {
					panic(e)
				}
				if e = pprof.StartCPUProfile(f); e != nil {
					panic(e)
				}
				time.Sleep(time.Duration(config.Profile.Duration) * time.Minute)
				pprof.StopCPUProfile()
				f.Close()
			}()
		}
		if config.Profile.Mem != "" {
			go func() {
				f, e := os.Create(config.Profile.Mem)
				if e != nil {
					panic(e)
				}
				time.Sleep(time.Duration(config.Profile.Duration) * time.Minute)
				runtime.GC()
				if e = pprof.WriteHeapProfile(f); e != nil {
					panic(e)
				}
				f.Close()
			}()
		}
		if config.Profile.Stack != "" {
			go func() {
				f, e := os.Create(config.Profile.Stack)
				if e != nil {
					panic(e)
				}
				time.Sleep(time.Duration(config.Profile.Duration) * time.Minute)
				buf := make([]byte, 16*1024*1024)
				buf = buf[:runtime.Stack(buf, true)]
				f.WriteString("\n\n")
				f.WriteString(time.Now().String())
				f.WriteString("\n")
				f.Write(buf)
				f.Close()
			}()
		}
	})
	return &config
}

//ConfPath config.toml path
var ConfPath = "config/config.toml"

func parseConfig(config interface{}) error {
	if _, err := toml.DecodeFile(ConfPath, config); err != nil {
		return err
	}
	return nil
}
