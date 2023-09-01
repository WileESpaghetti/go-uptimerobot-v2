package models

type Monitor struct {
	Id           int                `schema:"id,omitempty"  json:"id,omitempty"`
	FriendlyName string             `schema:"friendly_name" json:"friendly_name"`
	Type         MonitorType        `schema:"type"          json:"type"`
	Status       MonitorStatus      `schema:"status"        json:"status"`
	SubType      MonitorSubType     `schema:"sub_type"      json:"sub_type"`
	KeywordType  MonitorKeywordType `schema:"keyword_type"      json:"keyword_type"`
}
