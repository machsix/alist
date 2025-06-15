package onedrive

import (
	"github.com/machsix/alist/v3/internal/driver"
	"github.com/machsix/alist/v3/internal/op"
)

type Addition struct {
	driver.RootPath
	Region       string `json:"region" type:"select" required:"true" options:"global,cn,us,de" default:"global"`
	IsSharepoint bool   `json:"is_sharepoint"`
  ClientID     string `json:"client_id" required:"true" default:"b15665d9-eda6-4092-8539-0eec376afd59"`
  ClientSecret string `json:"client_secret" required:"true" default:"qtyfaBBYA403=unZUP40~_#"`
	RedirectUri  string `json:"redirect_uri" required:"true" default:"http://localhost:53682/"`
	RefreshToken string `json:"refresh_token" required:"true"`
	SiteId       string `json:"site_id"`
	ChunkSize    int64  `json:"chunk_size" type:"number" default:"5"`
	CustomHost   string `json:"custom_host" help:"Custom host for onedrive download link"`
}

var config = driver.Config{
	Name:        "Onedrive",
	LocalSort:   true,
	DefaultRoot: "/",
}

func init() {
	op.RegisterDriver(func() driver.Driver {
		return &Onedrive{}
	})
}
