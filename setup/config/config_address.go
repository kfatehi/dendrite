package config

import (
	"io/fs"
	"net/url"
)

const (
	NetworkTCP  = "tcp"
	NetworkUnix = "unix"
)

type ServerAddress struct {
	Address              string
	Scheme               string
	UnixSocketPermission fs.FileMode
}

func (s ServerAddress) Enabled() bool {
	return s.Address != ""
}

func (s ServerAddress) IsUnixSocket() bool {
	return s.Scheme == NetworkUnix
}

func (s ServerAddress) Network() string {
	if s.Scheme == NetworkUnix {
		return NetworkUnix
	} else {
		return NetworkTCP
	}
}

func UnixSocketAddress(path string, perm fs.FileMode) ServerAddress {
	return ServerAddress{Address: path, Scheme: NetworkUnix, UnixSocketPermission: perm}
}

func HTTPAddress(urlAddress string) (ServerAddress, error) {
	parsedUrl, err := url.Parse(urlAddress)
	if err != nil {
		return ServerAddress{}, err
	}
	return ServerAddress{parsedUrl.Host, parsedUrl.Scheme, 0}, nil
}
