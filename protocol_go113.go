//go:build go1.13
// +build go1.13

package mssql

import (
	"fmt"

	"github.com/io1o/go-mssqldb/msdsn"
)

func wrapConnErr(p *msdsn.Config, err error) error {
	f := "unable to open tcp connection with host '%v:%v': %w"
	return fmt.Errorf(f, p.Host, resolveServerPort(p.Port), err)
}
