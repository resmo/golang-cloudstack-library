package cloudstack

type Result struct {
	// any text associated with the success or failure
	DisplayText NullString `json:"displaytext"`
	// true if operation is executed successfully
	Success NullBool `json:"success"`
	// CloudStack API Client
	Client *Client
}
