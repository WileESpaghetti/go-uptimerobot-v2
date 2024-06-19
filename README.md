# go-uptimerobot-v2

![example workflow](https://github.com/WileESpaghetti/go-uptimerobot-v2/actions/workflows/go.yml/badge.svg)

Uptime Robot APIv2 command-line client and library

[UptimeRobot API Documentation](https://uptimerobot.com/api)


## Build
```shell
go build -o uptimerobot cmd/uptimerobot/main.go
```

## Command Usage

```shell
./uptimerobot GROUP [COMMAND] --api-key=$KEY
```

## Library Usage

```go
package main

import (
	"fmt"
	"github.com/WileESpaghetti/go-uptimerobot-v2/uptime_robot"
)

func main() {
	// initialize the API client
	apiKey := "your-api-key"
	ur := uptime_robot.NewClient(apiKey)

	// get all monitors
	monitors, err := ur.GetMonitors()
	if err != nil {
		fmt.Println(err)
		return
	}
	
    for _, m := range monitors {
		fmt.Printf("%d: %s: %s\n", m.ID, m.Url, m.Status)
    }
	
    // get specific monitors
    
	
    
}
```

### Implemented groups
* account - displays account information
* monitors
* help
* other cobra built-ins

### API Support

#### `getAccountDetails`
- [X] api_key

#### `getMonitors`
- [X] api_key
- [X] monitors
- [X] types
- [X] statuses
- [X] custom_uptime_ratios
  - does not accept hyphen as a separator yet
- [ ] custom_down_durations
- [ ] custom_uptime_ranges
- [ ] all_time_uptime_ratio
- [ ] all_time_uptime_durations
- [ ] logs
- [ ] logs_start_date
- [ ] logs_end_date
- [ ] log_types
- [ ] logs_limit
- [ ] response_times
- [ ] response_times_limit
- [ ] response_times_average
- [ ] response_times_start_date
- [ ] response_times_end_date
- [ ] alert_contacts
- [ ] mwindows
- [ ] ssl
- [ ] custom_http_headers
- [ ] custom_http_statuses
- [ ] http_request_details
- [ ] auth_type
- [ ] timezone
- [ ] offset
- [ ] limit
- [ ] search
