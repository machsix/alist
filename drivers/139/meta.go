package _139

import (
	"github.com/machsix/alist/v3/internal/driver"
	"github.com/machsix/alist/v3/internal/op"
)

type Addition struct {
	//Account       string `json:"account" required:"true"`
	Authorization string `json:"authorization" type:"text" required:"true"`
	driver.RootID
	Type    string `json:"type" type:"select" options:"personal,family,personal_new" default:"personal"`
	CloudID string `json:"cloud_id"`
}

var config = driver.Config{
	Name:             "139Yun",
	LocalSort:        true,
	ProxyRangeOption: true,
}

func init() {
	op.RegisterDriver(func() driver.Driver {
		d := &Yun139{}
		d.ProxyRange = true
		return d
	})
}
