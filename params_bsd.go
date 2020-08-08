// +build dragonfly freebsd netbsd openbsd

package water

type PlatformSpecificParams struct {
	// Name is the name to be set for the interface to be created.
	// The name should be of the format tunN or tapN (N is a number < 256).
	// When a zero value is given (empty string), then water iterates from 0
	// to 255 and returns the device whichever succeeds.
	//
	// Ref: https://github.com/OpenVPN/openvpn/blob/1d86fae87/src/openvpn/tun.c#L1680-L1692
	Name string
}

func defaultPlatformSpecificParams() PlatformSpecificParams {
	return PlatformSpecificParams{}
}
