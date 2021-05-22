package autovpn

import (
	"errors"
	"fmt"
	"reflect"
	"strings"

	"github.com/mysteriumnetwork/go-openvpn/openvpn3"
)

type callbacks interface {
	openvpn3.Logger
	openvpn3.EventConsumer
	openvpn3.StatsConsumer
}

type Config struct {
	openvpn3.UserCredentials
	openvpn3.Logger
	configDir string
	callbacks
}

type defaultCallbacks struct {
}

func (lc *defaultCallbacks) Log(text string) {
	lines := strings.Split(text, "\n")
	for _, line := range lines {
		fmt.Println("Openvpn log >>", line)
	}
}

func (lc *defaultCallbacks) OnEvent(event openvpn3.Event) {
	fmt.Printf("Openvpn event >> %+v\n", event)
}

func (lc *defaultCallbacks) OnStats(stats openvpn3.Statistics) {
	fmt.Printf("Openvpn stats >> %+v\n", stats)
}

func NewConfig(username, password, vpnConfigDir string, logger openvpn3.Logger, openVpnCallbacks interface{}) (*Config, error) {
	if username == "" || password == "" {
		return nil, errors.New("Invalid credentials")
	} else if vpnConfigDir == "" {
		return nil, errors.New("No Openvpn config directory provided, please specify a valid directory with your VPNs configuration")
	}
	cb, ok := openVpnCallbacks.(callbacks)
	if !ok && reflect.TypeOf(openVpnCallbacks) == nil {
		cb = &defaultCallbacks{}
	} else if !ok {
		return nil, errors.New("Invalid callbacks implementation")
	}
	return &Config{
		openvpn3.UserCredentials{username, password},
		logger,
		vpnConfigDir,
		cb,
	}, nil
}
