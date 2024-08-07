package crm7client

import (
	"fmt"
	"net/http"
	"time"
)

type Crm7Client struct {
	CrmUrl          string
	CrmTerminalType string
	HTTPClient      *http.Client
}

func NewClient(CrmServerIP string, CrmPort int, CrmTerminalType string) *Crm7Client {
	url := fmt.Sprintf("http://%s:%d", CrmServerIP, CrmPort)
	return &Crm7Client{
		CrmUrl:          url,
		CrmTerminalType: CrmTerminalType,
		HTTPClient: &http.Client{
			Timeout: time.Minute,
			//Transport: &http.Transport{TLSClientConfig: &tls.Config{InsecureSkipVerify: true}},
		},
	}
}
