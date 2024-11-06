package public_status_page

import "github.com/WileESpaghetti/go-uptimerobot-v2/uptime_robot/api"

// GetPublicStatusPages is details about your account and API response metadata
type GetPublicStatusPages struct {
	api.Envelope
	PublicStatusPages []PublicStatusPage `json:"psps,omitempty"`
}

type GetPublicStatusPagesRequest struct {
	PublicStatusPages PublicStatusPages `form:"psps,omitempty"`
}

type PublicStatusPageOptions func(*GetPublicStatusPagesRequest)

// WithOptions allows setting all options for [getPSPs] request at once.
//
// This is a good option if you are setting many query parameters and do not
// want to pass several AlertContactOptions.
//
// [getPSPs]: https://uptimerobot.com/api/
func WithOptions(getPublicStatusPages *GetPublicStatusPagesRequest) PublicStatusPageOptions {
	return func(options *GetPublicStatusPagesRequest) {
		*options = *getPublicStatusPages
	}
}

func WithPublicStatusPages(psps PublicStatusPages) PublicStatusPageOptions {
	return func(options *GetPublicStatusPagesRequest) {
		options.PublicStatusPages = psps
	}
}
