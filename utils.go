package autovpn

import (
	"errors"
	"fmt"
	"net"
	"os"
	"path/filepath"
	"regexp"
	"sort"
	"time"

	"github.com/go-ping/ping"
)

const (
	pingTimeout = 10
)

func filePathWalkDir(root string) ([]string, error) {
	var files []string
	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if !info.IsDir() {
			files = append(files, path)
		}
		return nil
	})
	return files, err
}

func testInternetConnection(ip, port string) (bool, error) {
	timeOut := time.Duration(pingTimeout) * time.Second
	conn, err := net.DialTimeout("tcp", ip+":"+port, timeOut)
	if err != nil {
		return false, err
	}
	fmt.Printf("Remote Address : %s \n", conn.RemoteAddr().String())
	return true, nil
}

func pingAddress(ip string) (*ping.Statistics, error) {
	pinger, err := ping.NewPinger("www.google.com")
	if err != nil {
		return nil, err
	}
	pinger.Count = 5
	err = pinger.Run() // Blocks until finished.
	if err != nil {
		return nil, err
	}
	return pinger.Statistics(), nil
}

func getRemoteAddress(config string) (string, error) {
	//example= Remote Address : 142.250.185.14:80
	//re := regexp.MustCompile(`Remote.Address.\:.*$`)
	//re := regexp.MustCompile(`^(((25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)(\.|$)){4})`)
	//re := regexp.MustCompile(`(25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)(\.(25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)){3}`)
	re := regexp.MustCompile(`\w+\.\w+\.com`)
	submatchall := re.FindAllString(config, -1)
	if len(submatchall) < 1 {
		return "", errors.New("No remote address in openvpn configuration file")
	}
	sort.Strings(submatchall)
	auxElem := submatchall[0]
	for _, element := range submatchall {
		if auxElem != element {
			return "", errors.New("More than one remote address was found in the configuration file")
		}
	}
	return submatchall[0], nil
}
