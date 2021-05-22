package autovpn

import (
	"io/ioutil"
	"math/rand"
	"time"

	"github.com/mysteriumnetwork/go-openvpn/openvpn3"
)

func (c *Config) Start() error {
	openvpn3.SelfCheck(c.Logger)
	files, err := filePathWalkDir(c.configDir)
	if err != nil {
		panic(err)
	}
	rand.Seed(time.Now().UnixNano())
	bytes, err := ioutil.ReadFile(files[rand.Intn(len(files))])
	if err != nil {
		return err
	}
	vpnConfig := openvpn3.NewConfig(string(bytes))

	session := openvpn3.NewSession(vpnConfig, c.UserCredentials, c.callbacks)
	session.Start()
	err = session.Wait()
	return err
}
