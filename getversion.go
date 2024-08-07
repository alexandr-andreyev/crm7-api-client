package crm7client

func (c Crm7Client) GetVersion() (*Version, error) {
	body := Message{
		Action:       "Get version",
		TerminalType: c.CrmTerminalType,
		UnitID:       "1",
		UserID:       "1",
	}
	req, err := c.newRequest("POST", body)
	if err != nil {
		return nil, err
	}
	result := Version{}
	_, err = c.do(req, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}
