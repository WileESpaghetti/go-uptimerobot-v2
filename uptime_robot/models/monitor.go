package models

import "encoding/json"

type Monitor struct {
	Id           int                `schema:"id,omitempty"  json:"id,omitempty"`
	FriendlyName string             `schema:"friendly_name" json:"friendly_name"`
	Type         MonitorType        `schema:"type"          json:"type"`
	Status       MonitorStatus      `schema:"status"        json:"status"`
	SubType      MonitorSubType     `schema:"sub_type"      json:"sub_type"`
	KeywordType  MonitorKeywordType `schema:"keyword_type"  json:"keyword_type"`
	KeywordValue string             `schema:"keyword_value" json:"keyword_value"`
	HttpUsername string             `schema:"http_username" json:"http_username"`
	HttpPassword string             `schema:"http_password" json:"http_password"`
	Port         OptionalNumber     `schema:"port"          json:"port"`
	Interval     json.Number        `schema:"interval"      json:"interval"`
}
