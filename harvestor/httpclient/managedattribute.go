package httpclient

import (
	_ "bytes"
	//c "github.com/hashicorp/go-retryablehttp"
	_ "github.com/liamylian/jsontime"
	_ "harvestor/config"
	"harvestor/db"
	//l "harvestor/logger"
	_ "io"
	_ "net/http"
)

// post Sidecard data as Managed Metadata
func postManagedMeta(file *db.File) (db.Sidecard, error) {
	// custom json for all time formats
	//var json = jsontime.ConfigWithCustomTimeFormat
	// init logger
	//var logger = l.NewLogger()
	// init conf
	//conf := config.GetConf()
	// init sidecard struct
	var sidecard db.Sidecard

	return sidecard, nil
}
