package crm7client

import "encoding/xml"

type Message struct {
	XMLName      xml.Name `xml:"Message"`
	Action       string   `xml:"Action,attr"`
	TerminalType string   `xml:"Terminal_Type,attr"`
	GlobalType   string   `xml:"Global_Type,attr"`
	UnitID       string   `xml:"Unit_ID,attr"`
	UserID       string   `xml:"User_ID,attr"`
	Include      string   `xml:"Include"`
	Count        int      `xml:"Count"`
	Items        []Item   `xml:"Item"`
}

type Item struct {
	Mode     string    `xml:"Mode,attr"`
	Contacts *Contacts `xml:"Contacts,omitempty"`
}

type Contacts struct {
	Phone Phone `xml:"Phone"`
}

// Phone представляет телефон с атрибутами
type Phone struct {
	Value    string `xml:"Value,attr"`
	IsNumber bool   `xml:"IsNumber,attr"`
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

type SystemTime struct {
	XMLName   xml.Name `xml:"SystemTime"`
	MessageID string   `xml:"Message_ID,attr"`
	Date      string   `xml:"Date"`
	Time      string   `xml:"Time"`
	DateTime  string   `xml:"DateTime"`
}

type Holders struct {
	XMLName     xml.Name `xml:"Holders"`
	MessageID   string   `xml:"Message_ID,attr"`
	IndexFrom   int      `xml:"IndexFrom,attr"`
	IndexTo     int      `xml:"IndexTo,attr"`
	Count       int      `xml:"Count,attr"`
	SearchGUID  string   `xml:"Search_GUID,attr"`
	HoldersList []Holder `xml:"Holder"`
}

// Holder представляет элемент Holder
type Holder struct {
	Deleted          string           `xml:"Deleted"`
	GroupID          int              `xml:"Group_ID"`
	GroupName        string           `xml:"Group_Name"`
	DivisionID       int              `xml:"Division_ID"`
	DivisionName     string           `xml:"Division_Name"`
	HolderID         string           `xml:"Holder_ID"`
	INN              string           `xml:"INN"`
	ExternalCode     string           `xml:"External_Code"`
	UnpayTypeID      string           `xml:"Unpay_Type_ID"`
	UnpayTypeName    string           `xml:"Unpay_Type_Name"`
	LName            string           `xml:"L_Name"`
	FName            string           `xml:"F_Name"`
	MName            string           `xml:"M_Name"`
	FullName         string           `xml:"Full_Name"`
	Birth            string           `xml:"Birth"`
	Gender           string           `xml:"Gender"`
	Marrital         string           `xml:"Marrital"`
	Smoke            string           `xml:"Smoke"`
	Verification     string           `xml:"Verification"`
	Image            string           `xml:"Image"`
	LanguageID       string           `xml:"Language_ID"`
	LanguageName     string           `xml:"Language_Name"`
	Source           string           `xml:"Source"`
	Remarks          string           `xml:"Remarks"`
	HoldersContacts  HoldersContacts  `xml:"Holders_Contacts"`
	HoldersAddresses HoldersAddresses `xml:"Holders_Addresses"`
	HoldersCards     HoldersCards     `xml:"Holders_Cards"`
}

// HoldersContacts представляет элемент Holders_Contacts
type HoldersContacts struct {
	HolderContacts []HolderContact `xml:"Holder_Contact"`
}

// HolderContact представляет элемент Holder_Contact
type HolderContact struct {
	HolderID string  `xml:"Holder_ID"`
	Contact  Contact `xml:"Contact"`
}

// Contact представляет контакт
type Contact struct {
	Deleted       string `xml:"Deleted"`
	ContactID     string `xml:"Contact_ID"`
	IsVirtualCard string `xml:"Is_Virtual_Card"`
	Dispatch      string `xml:"Dispatch"`
	TypeID        int    `xml:"Type_ID"`
	TypeCode      string `xml:"Type_Code"`
	TypeName      string `xml:"Type_Name"`
	Value         string `xml:"Value"`
	TypeClass     string `xml:"Type_Class"`
	Notes         string `xml:"Notes"`
}

// HoldersAddresses представляет элемент Holders_Addresses
type HoldersAddresses struct {
	HolderAddresses []HolderAddress `xml:"Holder_Address"`
}

// HolderAddress представляет элемент Holder_Address
type HolderAddress struct {
	HolderID string `xml:"Holder_ID"`
}

// HoldersCards представляет элемент Holders_Cards
type HoldersCards struct {
	HolderCards []HolderCard `xml:"Holder_Card"`
}

// HolderCard представляет элемент Holder_Card
type HolderCard struct {
	HolderID string `xml:"Holder_ID"`
	Card     Card   `xml:"Card"`
}

// Card представляет элемент Card
type Card struct {
	CardCode         string `xml:"Card_Code"`
	IsVirtualCard    string `xml:"Is_Virtual_Card"`
	IsConfirmManager string `xml:"Is_Confirm_Manager"`
	Status           string `xml:"Status"`
	CarrierData      string `xml:"Carrier_Data"`
	Offered          string `xml:"Offered"`
	Expired          string `xml:"Expired"`
	GroupID          int    `xml:"Group_ID"`
	GroupName        string `xml:"Group_Name"`
	OwnerID          string `xml:"Owner_ID"`
	Verification     string `xml:"Verification"`
}
