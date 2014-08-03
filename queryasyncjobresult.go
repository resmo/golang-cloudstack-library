package cloudstack

import (
	"encoding/json"
	"fmt"
	"log"
)

type QueryAsyncJobResultParameter struct {
	Jobid ID
}

func (p *QueryAsyncJobResultParameter) SetJobid(s string) {
	p.Jobid.String = s
	p.Jobid.Valid = true
}
func (p *QueryAsyncJobResultParameter) ToMap() map[string]string {
	m := map[string]string{}
	if p.Jobid.Valid {
		m["jobid"] = fmt.Sprint(p.Jobid.String)
	}
	return m
}

type QueryAsyncJobResultResponse struct {
	Accountid       ID              `json:"accountid"`
	Cmd             NullString      `json:"cmd"`
	Created         NullString      `json:"created"`
	Jobid           ID              `json:"jobid"`
	Jobinstanceid   ID              `json:"jobinstanceid"`
	Jobinstancetype NullString      `json:"jobinstancetype"`
	Jobprocstatus   NullInt64       `json:"jobprocstatus"`
	Jobresult       json.RawMessage `json:"jobresult"`
	Jobresultcode   NullInt64       `json:"jobresultcode"`
	Jobresulttype   NullString      `json:"jobresulttype"`
	Jobstatus       NullInt64       `json:"jobstatus"`
	Userid          ID              `json:"userid"`
}

func (c *Client) QueryAsyncJobResult(p QueryAsyncJobResultParameter) (*QueryAsyncJobResultResponse, error) {
	resp := new(QueryAsyncJobResultResponse)
	b, err := c.Request("queryAsyncJobResult", p.ToMap())
	if err != nil {
		log.Println("Request failed:", err)
		return nil, err
	}
	err = json.Unmarshal(b, resp)
	if err != nil {
		log.Println("json.Unmarshal failed:", err)
	}
	return resp, err
}
