package slave

import (
	"context"
	"net"
	"os"
	"os/user"
	"runtime"
	"time"
)

type NetworkInterface struct {
	Name   string
	IPAddr []string
	MAC    string
}

func (s *Slave) GetHostname() string {
	hostname, err := os.Hostname()
	if err != nil {
		return "unkown"
	}
	return hostname
}

func (s *Slave) GetOS() string {
	return runtime.GOOS
}

func (s *Slave) GetNetworkInterfaces() []*NetworkInterface {
	nIfaceList := []*NetworkInterface{}
	interfaces, err := net.Interfaces()
	if err != nil {
		return []*NetworkInterface{}
	}

	for _, iface := range interfaces {
		currentNetworkInterface := new(NetworkInterface)
		currentNetworkInterface.Name = iface.Name
		currentNetworkInterface.MAC = iface.HardwareAddr.String()

		if addrs, err := iface.Addrs(); err == nil {
			for _, addr := range addrs {
				currentNetworkInterface.IPAddr = append(currentNetworkInterface.IPAddr, addr.String())
			}
		}

		nIfaceList = append(nIfaceList, currentNetworkInterface)

	}
	return nIfaceList
}

func (s *Slave) GetPublicIP() string {
	resolver := &net.Resolver{
		PreferGo: true,
		Dial: func(ctx context.Context, network, address string) (net.Conn, error) {
			d := net.Dialer{
				Timeout: time.Millisecond * time.Duration(10000),
			}
			return d.DialContext(ctx, network, "resolver1.opendns.com:53")
		},
	}

	ip, err := resolver.LookupHost(context.Background(), "myip.opendns.com")
	if err != nil {
		return "unknown"
	}
	return ip[0]

}

func (s *Slave) GetCurrentUser() *user.User {
	if username, err := user.Current(); err == nil {
		return username
	}
	return &user.User{
		Name:     "Unknown",
		Uid:      "Unknown",
		Gid:      "Unknown",
		Username: "Unknown",
		HomeDir:  "Unknown",
	}

}
