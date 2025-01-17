//go:build windows
// +build windows

package winsspi

import "github.com/io1o/go-mssqldb/integratedauth"

// AuthProvider handles SSPI Windows Authentication via secur32.dll functions
var AuthProvider integratedauth.Provider = integratedauth.ProviderFunc(getAuth)

func init() {
	err := integratedauth.SetIntegratedAuthenticationProvider("winsspi", AuthProvider)
	if err != nil {
		panic(err)
	}
}
