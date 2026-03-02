package modules

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"path/filepath"

	"github.com/sirupsen/logrus"
)

var versionParameter = "{{RELEASE_TAG}}"
var configDirectoryPath, _ = filepath.Abs("./config/")
var configFilePath = filepath.Join(configDirectoryPath, "config.json")
var ConfigFile = ConfigStruct{}

func LoadConfig() (err error) {
	ConfigFile = ConfigStruct{}

	// create config.json if it doesn't exist
	if _, err := os.Stat(configFilePath); errors.Is(err, os.ErrNotExist) {
		fmt.Println("config file does not exist. creating...")

		err := CreateConfigFile()
		if err != nil {
			return err
		}
	}

	// try to open file
	file, err := os.Open(configFilePath)
	if err != nil {
		fmt.Println("get config file threw error trying to open the file")
		return err
	}
	defer file.Close()
	decoder := json.NewDecoder(file)

	// decode file
	err = decoder.Decode(&ConfigFile)
	if err != nil {
		fmt.Println("get config file threw error trying to parse the file")
		return err
	}

	// verify values
	anythingChanged := false

	if ConfigFile.Version != versionParameter {
		ConfigFile.Version = versionParameter
		anythingChanged = true
	}

	if ConfigFile.Timezone == "" {
		ConfigFile.Timezone = "Europe/Paris"
	}

	if ConfigFile.Environment == "" {
		ConfigFile.Environment = "production"
	}

	if ConfigFile.PortFile == "" {
		ConfigFile.PortFile = "/tmp/gluetun/forwarded_port"
	}

	if ConfigFile.Interval == 0 {
		ConfigFile.Interval = 15
	}

	if ConfigFile.QBitTorrent.IP == "" {
		ConfigFile.QBitTorrent.IP = "localhost"
	}

	if ConfigFile.QBitTorrent.Port == 0 {
		ConfigFile.QBitTorrent.Port = 8080
	}

	if ConfigFile.QBitTorrent.Username == "" {
		ConfigFile.QBitTorrent.Username = "admin"
	}

	if anythingChanged {
		// save new version of config json
		fmt.Println("saving new config file version")
		err = SaveConfig()
		if err != nil {
			return err
		}
	}

	// return nil object
	return nil
}

// creates empty config.json
func CreateConfigFile() error {
	ConfigFile = ConfigStruct{}
	ConfigFile.QBitTorrent = QBitTorrentConfig{}
	level := logrus.InfoLevel
	ConfigFile.LogLevel = level.String()

	err := SaveConfig()
	if err != nil {
		fmt.Println("create config file threw error trying to save the file. error: " + err.Error())
		return errors.New("create config file threw error trying to save the file")
	}

	return nil
}

// saves the given config struct as config.json
func SaveConfig() error {
	err := os.MkdirAll(configDirectoryPath, os.ModePerm)
	if err != nil {
		fmt.Println("failed to create directory for config. error: " + err.Error())
		return errors.New("failed to create directory for config")
	}

	file, err := os.OpenFile(configFilePath, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0644)
	if err != nil {
		return err
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "\t")
	encoder.SetEscapeHTML(false) // disable &/< > escaping

	if Log != nil {
		newLogLevel, err := logrus.ParseLevel(ConfigFile.LogLevel)
		if err != nil {
			Log.Errorf("failed to parse loglevel '" + ConfigFile.LogLevel + "'")
			return err
		}

		if Log.Level != newLogLevel {
			Log.Level = newLogLevel
			Log.Infof("loglevel changed to '%s'", newLogLevel.String())
		}
	}

	return encoder.Encode(ConfigFile)
}
