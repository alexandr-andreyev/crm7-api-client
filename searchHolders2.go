package crm7client

func (c Crm7Client) SearchHolders2(phoneNumber string) (*Holders, error) {
	body := Message{
		Action:       "Search holders 2",
		TerminalType: c.crmTerminalType,
		UnitID:       "1",
		UserID:       "1",
		Include:      "Holder_Contact, Holder_Address, Holder_Card",
		Count:        1,
		Items: []Item{
			{
				Mode: "Clear",
			},
			{
				Mode: "Add",
				Contacts: &Contacts{
					Phone: Phone{
						Value:    phoneNumber,
						IsNumber: true,
					},
				},
			},
			{Mode: "Clear"},
		},
	}
	req, err := c.newRequest("POST", body)
	if err != nil {
		return nil, err
	}
	result := Holders{}
	_, err = c.do(req, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}
