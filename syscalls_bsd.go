// +build dragonfly freebsd netbsd openbsd

package water

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"syscall"
)

func openDev(config Config) (*Interface, error) {
	var file *os.File
	var err error
	var ifaceName = config.Name

	prefix := "tun"
	if config.DeviceType == TAP {
		prefix = "tap"
	}

	if ifaceName != "" {
		n, err := strconv.Atoi(strings.TrimPrefix(config.Name, prefix))
		if err != nil || n > 255 {
			return nil, fmt.Errorf("Interface name must be '%s[0-9]+' (till 255)", prefix)
		}

		file, err = openDevice(config.Name)
		if err != nil {
			return nil, err
		}
	} else {
		// When a default device is not provided, we try to open device named /dev/tun0
		// till tun255 (same for tap) and return the file handle which is successful.
		// from OpenVPN source:
		//     https://github.com/OpenVPN/openvpn/blob/1d86fae87/src/openvpn/tun.c#L1680-L1692
		for i := 0; i < 255; i++ {
			ifaceName = fmt.Sprintf("%s%d", prefix, i)
			file, err = openDevice("/dev/" + ifaceName)
			if err == nil {
				break
			}
		}

		// if all fails there will be err.
		if err != nil {
			return nil, err
		}
	}

	return &Interface{
		isTAP:           config.DeviceType == TAP,
		ReadWriteCloser: file,
		name:            ifaceName,
	}, nil
}

// openDeviceFile opens a tun/tap device with the
// provided device name and returns error if any
func openDevice(ifaceName string) (*os.File, error) {
	fdInt, err := syscall.Open(ifaceName, syscall.O_RDWR|syscall.O_NONBLOCK, 0)
	if err != nil {
		return nil, err
	}
	return os.NewFile(uintptr(fdInt), ifaceName), nil
}
