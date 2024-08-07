package crm7client

func (c Crm7Client) GetSystemTyme() (*SystemTime, error) {
	body := Message{
		Action:       "Get system time",
		TerminalType: c.crmTerminalType,
		UnitID:       "1",
		UserID:       "1",
	}
	req, err := c.newRequest("POST", body)
	if err != nil {
		return nil, err
	}
	result := SystemTime{}
	_, err = c.do(req, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}
