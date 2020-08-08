// +build !linux,!darwin,!windows,!dragonfly,!freebsd,!netbsd,!openbsd

package water

import "errors"

func openDev(config Config) (*Interface, error) {
	return nil, errors.New("not implemented on this platform")
}
