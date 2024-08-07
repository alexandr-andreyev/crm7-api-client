package crm7client

import "encoding/xml"

type Message struct {
	XMLName      xml.Name `xml:"Message"`
	Action       string   `xml:"Action,attr"`
	TerminalType string   `xml:"Terminal_Type,attr"`
	GlobalType   string   `xml:"Global_Type,attr"`
	UnitID       string   `xml:"Unit_ID,attr"`
	UserID       string   `xml:"User_ID,attr"`
}

type Version struct {
	XMLName    xml.Name `xml:"Version"`
	MessageID  string   `xml:"Message_ID,attr"`
	ServerInfo string   `xml:"Server_Info"`
	Major      string   `xml:"Major"`
	Minor      string   `xml:"Minor"`
	Release    string   `xml:"Release"`
	Build      string   `xml:"Build"`
}
