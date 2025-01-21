# crm7-api-client

package main

import (
"crm-client/crm7client"
"log"
)

func main() {
client := crm7client.NewClient("127.0.0.1", 9192, "dlv7")

    resp, err := client.GetVersion()
    if err != nil {
    	log.Println(err)
    }

    log.Printf("%+v", resp)

}
