package utils

import (
	"github.com/brenordv/go-monitor-internet-connection/internal/core"
	"strings"
)

func GetDefaultUrls() []string {
	return strings.Split(core.Urls, ",")
}
