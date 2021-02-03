// harvestor is an application for recursively traverse directories and find media files.
// send each file of our interest to to object-store-api and map it back to the metadata 
// For more information please refer to https://github.com/AAFC-BICoE/object-store-harvestor/blob/dev/doc/design.md
package main

import (
        "log"
        "os"
)

func getEnv(key, fallback string) string {
        if value, ok := os.LookupEnv(key); ok {
                return value
        }
        return fallback
}

func main() {
        ////////////////////////////////////////////////////////////////////////////
        // Log env and release version
        var env = getEnv("DEPLOYMENT_ENV", "Not Set")
        var release = getEnv("RELEASE_VERSION", "Not Set")
        log.Println("||| Deployment Env  : ", env)
        log.Println("||| Release Version : ", release)
        ////////////////////////////////////////////////////////////////////////////

}
