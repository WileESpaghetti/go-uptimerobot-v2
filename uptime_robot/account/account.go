package account

import (
	"encoding/json"
	"time"
)

type Account struct {
	Email      string `json:"email"            xml:"email,attr"`
	UserID     int64  `json:"user_id"`
	FirstName  string `json:"firstname"`
	SMSCredits int64  `json:"sms_credits"`
	// TODO PaymentProcessor ???
	// TODO PaymentPeriod ???
	// TODO SubscriptionExpiryDate ???
	MonitorLimit       int64         `json:"monitor_limit"    xml:"monitor_limit,attr"`
	MonitorInterval    time.Duration `json:"-"                xml:"-"`
	UpMonitors         int64         `json:"up_monitors"      xml:"up_monitors,attr"`
	DownMonitors       int64         `json:"down_monitors"    xml:"down_monitors,attr"`
	PausedMonitors     int64         `json:"paused_monitors"  xml:"paused_monitors,attr"`
	TotalMonitorsCount int64         `json:"total_monitors_count"`
	// TODO RegisteredAt time.Time `json:"registered_at"`
	// TODO ActiveSubscription ??? `json:"active_subscriptions"`
	// TODO Organizations []???
}

type account Account

type jsonAccount struct {
	account
	MonitorInterval int64 `json:"monitor_interval" xml:"monitor_interval,attr"`
}

func (a *Account) UnmarshalJSON(data []byte) error {
	var ja jsonAccount

	if err := json.Unmarshal(data, &ja); err != nil {
		return err
	}

	*a = Account{
		Email:              ja.Email,
		UserID:             ja.UserID,
		FirstName:          ja.FirstName,
		SMSCredits:         ja.SMSCredits,
		MonitorLimit:       ja.MonitorLimit,
		MonitorInterval:    time.Duration(ja.MonitorInterval) * time.Minute,
		UpMonitors:         ja.UpMonitors,
		DownMonitors:       ja.DownMonitors,
		PausedMonitors:     ja.PausedMonitors,
		TotalMonitorsCount: ja.TotalMonitorsCount,
	}

	return nil
}

func (a *Account) MarshalJSON() ([]byte, error) {
	uac := account(*a)
	return json.Marshal(jsonAccount{
		account:         uac,
		MonitorInterval: int64(a.MonitorInterval / time.Minute),
	})
}
