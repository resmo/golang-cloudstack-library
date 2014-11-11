package cloudstack

import (
	"encoding/json"
	"fmt"
	"log"
)

type CreateSecurityGroupParameter struct {
	Account     NullString
	Description NullString
	Domainid    ID
	Name        NullString
	Projectid   ID
}

func (p *CreateSecurityGroupParameter) SetAccount(s string) {
	p.Account.String = s
	p.Account.Valid = true
}
func (p *CreateSecurityGroupParameter) SetDescription(s string) {
	p.Description.String = s
	p.Description.Valid = true
}
func (p *CreateSecurityGroupParameter) SetDomainid(s string) {
	p.Domainid.String = s
	p.Domainid.Valid = true
}
func (p *CreateSecurityGroupParameter) SetName(s string) {
	p.Name.String = s
	p.Name.Valid = true
}
func (p *CreateSecurityGroupParameter) SetProjectid(s string) {
	p.Projectid.String = s
	p.Projectid.Valid = true
}
func (p *CreateSecurityGroupParameter) ToMap() map[string]string {
	m := map[string]string{}
	if p.Account.Valid {
		m["account"] = fmt.Sprint(p.Account.String)
	}
	if p.Description.Valid {
		m["description"] = fmt.Sprint(p.Description.String)
	}
	if p.Domainid.Valid {
		m["domainid"] = fmt.Sprint(p.Domainid.String)
	}
	if p.Name.Valid {
		m["name"] = fmt.Sprint(p.Name.String)
	}
	if p.Projectid.Valid {
		m["projectid"] = fmt.Sprint(p.Projectid.String)
	}
	return m
}
func (c *Client) CreateSecurityGroup(p CreateSecurityGroupParameter) (Securitygroup, error) {
	var v map[string]json.RawMessage
	var ret Securitygroup
	b, err := c.Request("createSecurityGroup", p.ToMap())
	if err != nil {
		log.Println("Request failed:", err)
		return ret, err
	}
	err = json.Unmarshal(b, &v)
	if err != nil {
		return ret, fmt.Errorf("Failed to unmarshal: %s", string(b))
	}
	content, ok := v["securitygroup"]
	if !ok {
		errortext, ok := v["errortext"]
		if ok {
			return ret, fmt.Errorf(string(errortext))
		} else {
			return ret, fmt.Errorf("Unexpected format")
		}
	}
	err = json.Unmarshal(content, &ret)
	if err != nil {
		return ret, fmt.Errorf("Failed to unmarshal: %s", string(content))
	}
	return ret, nil
}
