package monitor

// CustomHTTPStatus shows which HTTP statuses are considered "up" by the UptimeRobot API. By default success responses
// (200-299) and redirect (300-399) are treated as being "up". The ability to customize this is only allowed under
// specific types of UptimeRobot accounts. For flexibility, statuses are kept as their integer value and can be passed
// to http.StatusText in cases where you want the human friendly version of the status.
//
// Up defaults to the 2xx and 3xx HTTP statuses. For monitoring purposes this might include all statuses between
// 200 and 399, but the API only includes the standardized/common HTTP statuses within that range.
//
// Down defaults to 4xx and 5xx HTTP statuses and returns the list of standardized/common HTTP statuses within that
// range.
type CustomHTTPStatus struct {
	Up   []int64 `json:"up"`
	Down []int64 `json:"down"`
}
