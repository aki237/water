// +build !linux,!darwin,!windows,!dragonfly,!freebsd,!netbsd,!openbsd

package water

// PlatformSpeficParams
type PlatformSpecificParams struct {
}

func defaultPlatformSpecificParams() PlatformSpecificParams {
	return PlatformSpecificParams{}
}
