package autovpn

import (
	"errors"
	"fmt"
	"io/ioutil"
	"math/rand"
	"sync"
	"time"

	"github.com/mysteriumnetwork/go-openvpn/openvpn3"
)

func Statistics(configDir string) error {
	files, err := filePathWalkDir(configDir)
	if err != nil {
		return err
	}
	var wg sync.WaitGroup
	for _, file := range files {
		wg.Add(1)
		go func(fileName string) {
			defer wg.Done()
			addr, err := getRemoteAddress(fileName)
			if err != nil {
				fmt.Println(err)
				return
			}
			stats, err := pingAddress(addr)
			if err != nil {
				fmt.Println(err)
				return
			}
			fmt.Printf("ConnFile: %s, Addr: %s, Average RTT: %v\n", fileName, addr, stats.AvgRtt)
		}(file)
	}
	wg.Wait()
	return nil
}

func (c *Config) Start() error {
	ok, err := testInternetConnection("google.com", "80")
	if !ok || err != nil {
		errors.New("Could not connect to Internet (google.com)")
	}
	openvpn3.SelfCheck(c.Logger)
	files, err := filePathWalkDir(c.configDir)
	if err != nil {
		return err
	}
	rand.Seed(time.Now().UnixNano())
	bytes, err := ioutil.ReadFile(files[rand.Intn(len(files))])
	if err != nil {
		return err
	}
	fmt.Println(string(bytes))
	vpnConfig := openvpn3.NewConfig(string(bytes))

	session := openvpn3.NewSession(vpnConfig, c.UserCredentials, c.callbacks)
	session.Start()
	err = session.Wait()
	return err
}
