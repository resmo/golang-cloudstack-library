package cloudstack

type AffinityGroupType struct {
	// the type of the affinity group
	Type NullString `json:"type"`
	// CloudStack API Client
	Client *Client
}
