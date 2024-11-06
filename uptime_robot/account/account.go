package account

type Account struct {
	Email      string `json:"email"            xml:"email,attr"`
	UserID     int64  `json:"user_id"`
	FirstName  string `json:"firstname"`
	SMSCredits int64  `json:"sms_credits"`
	// TODO PaymentProcessor ???
	// TODO PaymentPeriod ???
	// TODO SubscriptionExpiryDate ???
	MonitorLimit       int64 `json:"monitor_limit"    xml:"monitor_limit,attr"`
	MonitorInterval    int64 `json:"monitor_interval" xml:"monitor_interval,attr"` // TODO would this be useful as a time.Duration?
	UpMonitors         int64 `json:"up_monitors"      xml:"up_monitors,attr"`
	DownMonitors       int64 `json:"down_monitors"    xml:"down_monitors,attr"`
	PausedMonitors     int64 `json:"paused_monitors"  xml:"paused_monitors,attr"`
	TotalMonitorsCount int64 `json:"total_monitors_count"`
	// TODO RegisteredAt time.Time `json:"registered_at"`
	// TODO ActiveSubscription ??? `json:"active_subscriptions"`
	// TODO Organizations []???
}
