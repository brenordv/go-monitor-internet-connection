package utils

import (
	"bufio"
	"encoding/json"
	"github.com/brenordv/go-monitor-internet-connection/internal/core"
	"io"
	"os"
	"path"
	"strings"
)

func GetAppDir() (string, error) {
	execPath, err := os.Executable()
	return path.Dir(execPath), err
}

func GetDbDir() (string, error) {
	appDir, err := GetAppDir()
	if err != nil {
		return "", err
	}

	dbDir := path.Join(appDir, ".app-data")
	if _, e := os.Stat(dbDir); os.IsNotExist(e) {
		err = os.MkdirAll(dbDir, os.ModePerm)
	}

	return dbDir, err
}

func ReadLn(r *bufio.Reader) (string, error) {
	var (
		isPrefix       = true
		err      error = nil
		line, ln []byte
	)

	for isPrefix && err == nil {
		line, isPrefix, err = r.ReadLine()
		ln = append(ln, line...)
	}

	return string(ln), err
}

func LoadWatchlistUrls() ([]string, error) {
	appDir, err := GetAppDir()
	if err != nil {
		return nil, err
	}
	urlsFile := path.Join(appDir, core.DefaultUrlsFile)
	if _, err := os.Stat(urlsFile); os.IsNotExist(err) {
		return GetDefaultUrls(), nil
	}
	var f *os.File
	f, err = os.Open(urlsFile)
	if err != nil {
		return nil, err
	}

	var urls []string
	var line string
	r := bufio.NewReader(f)
	for {
		line, err = ReadLn(r)

		if err == nil && strings.HasPrefix(strings.ToLower(line), "http") {
			urls = append(urls, line)
			continue
		}

		if err == io.EOF {
			return urls, nil
		}

		if err != nil {
			return nil, err
		}
	}
}

func LoadRuntimeConfig() (*core.RuntimeConfig, error) {
	appDir, err := GetAppDir()
	if err != nil {
		return nil, err
	}
	cfgFile := path.Join(appDir, core.RuntimeConfigFile)
	if _, err := os.Stat(cfgFile); os.IsNotExist(err) {
		return &core.RuntimeConfig{
			Headers:              nil,
			DelayInSeconds:       core.DefaultDelayInSeconds,
		}, nil
	}
	var f *os.File
	var runtimeCfg core.RuntimeConfig
	f, err = os.Open(cfgFile)
	jsonParser := json.NewDecoder(f)
	if err = jsonParser.Decode(&runtimeCfg); err != nil {
		return nil, err
	}

	if runtimeCfg.NoConnInfoTTLinHours == 0 {
		runtimeCfg.NoConnInfoTTLinHours = core.DefaultNoConnInfoTTLinHours
	}

	if runtimeCfg.DelayInSeconds == 0 {
		runtimeCfg.DelayInSeconds = core.DefaultDelayInSeconds
	}

	return &runtimeCfg, nil
}
