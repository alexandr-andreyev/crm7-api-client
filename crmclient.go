package crm7client

import (
	"fmt"
	"net/http"
	"time"
)

type Crm7Client struct {
	crmUrl          string
	crmTerminalType string
	hTTPClient      *http.Client
}

func NewClient(CrmServerIP string, CrmPort int, CrmTerminalType string) *Crm7Client {
	url := fmt.Sprintf("http://%s:%d", CrmServerIP, CrmPort)
	return &Crm7Client{
		crmUrl:          url,
		crmTerminalType: CrmTerminalType,
		hTTPClient: &http.Client{
			Timeout: time.Minute,
			//Transport: &http.Transport{TLSClientConfig: &tls.Config{InsecureSkipVerify: true}},
		},
	}
}
