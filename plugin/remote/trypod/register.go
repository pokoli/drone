package trypod

import (
	"github.com/drone/config"
	"github.com/drone/drone/plugin/remote"
)

var (
	trypodURL   = config.String("trypod-url", "")
	trypodOwner = config.String("trypod-owner", "")
	trypodOpen  = config.Bool("trypod-open", false)
)

func Register() {
	if len(*trypodURL) == 0 {
		return
	}
	remote.Register(
		New(
			*trypodURL,
			*trypodOwner,
			*trypodOpen,
		),
	)
}
