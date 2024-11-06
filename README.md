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
* monitor
* contact
* psp
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
- [X] all_time_uptime_ratio
- [X] all_time_uptime_durations
- [X] logs
  - [ ] logs_start_date [Requires Pro Plan]
  - [ ] logs_end_date [Requires Pro Plan]
  - [X] log_types
  - [X] logs_limit
- [X] response_times
  - [X] response_times_limit
  - [X] response_times_average
  - [ ] response_times_start_date
  - [ ] response_times_end_date
- [X] alert_contacts
- [ ] mwindows
- [ ] ssl
- [ ] custom_http_headers
- [X] custom_http_statuses
- [ ] http_request_details
- [X] auth_type
- [ ] timezone
  - not really a property of a monitor and is returned as a property of the main response object
- [ ] offset
- [ ] limit
- [ ] search

#### `getAlertContacts`
- [X] api_key
- [X] alert_contacts
- [ ] offset
- [ ] limit
 
#### `getPSPs`
- [X] api_key
- [X] psps
- [ ] offset
- [ ] limit

#### `getMWindows`
- [ ] api_key
- [ ] mwindows
- [ ] offset
- [ ] limit

## FIXME
* library
  * [types] use unsigned ints when possible
  * [types] use best practices for int vs int8 vs int16...
  * [types] more consistent usage of exported error messages vs Error structs
  * [types] more consistent usage of string vs url.URL usage
  * [types] should optional number be an *int instead of int?
  * [types] double check time.Duration conversions they might need to be in int32 instead of int64
  * [types] some of the type stuff don't need to be exported and should only expose New* functions to create them
  * [types] maybe add typed versions of alert contacts where Value is not a string
  * [types] account monitor interval -> time.duration
  * [types] maybe add error/warning returns for options if you are assigning conflicting options (ex. response times limit and response times start/end)
  * [types] use generics instead of custom response types?
  * [types] double check struct tags match api docs (plural vs singular)
  * [slices] best practices for custom slice types
  * [consistency] ensure we conform to common api client common usage patterns
  * [consistency] use of uptime_robot vs uptimerobot
  * [consistency] some stuff uses 0/1 for boolean, but others (ex. auth_type) uses true/false (good for fuzzing/unit test consideration)
  * [consistency] make sure function args match across interfaces (ex UnmarshalJSON(b []byte) vs UnmarshalJSON(data []byte))
  * [consistency] consistent handling of ID's. Alert contacts are strings and can have leading zeros so needs to remain a string, but monitors are returned as numbers
  * [consistency] ensure best practices for implementing MarshalJSON and UnMarshalJSON seems like MarshalJSON - pass pointer, use json.* functions
  * [testing] test with different key types
  * [enums] monitors.Status: does it make more sense to have this as a struct {ID uint, Value string} and return error from NewStatus() if not enum? (all enums)
  * There's some New*() functions that take a parameter, might make sense to separate them into New*FromT(T) functions
  * [api] consider removing some of the string stuff. it's mostly to make the command output easier, but doesn't really make sense as a part of a library
  * [api] WithMonitors: is it more ergonomic to have a secondary function that just takes variadic ID's instead of a list or structs
  * [api] WithMonitors: is it more ergonomic as a variadic instead of accepting a slice. I think the biggest reason I do this is because of serialization
  * [pkg] move stuff that is only needed for command line opts/args to the package and use aliases instead of cluttering up the main module with stuff that is only getters/setters
  * [pkg] api.Envelop: does this need to be exported, or is having it as a part of the endpoint type sufficient?
  * [pkg] api.Envelop.Stat: should the Stat and related constants be a custom type?
  * [pkg] move all of the stuff out of the api package right now api.Error/Envelop causes circular references
  * when checking errors use the new Error.Is (100 go mistakes)
  * go:generate stringer - https://last9.io/blog/golang-stringer-tool/
  * [tests] use strings.folds when comparing strings in unit tests
  * [types] should monitor.KeywordCaseType be a boolean IsCaseSensitive instead?
  * [godoc] check for consistent formatting of fields in godoc
  * [godoc] add godoc for Request options
  * ensure zero values match what api zero values/defaults are
  * User friendly error messages: does it make more sense to move these to the command line client or some other higher level package?
  * uptimerobot python client uses the term "List separator" for the dashed numbers. should probably use that in code comments instead of other stuff (ex. listSeparatedValueFlag)
  * rename some of the GetMonitorRequest fields to make more semantic sense (ex. ResponseTimesLimit -> MaxResponseTimes, ResponseTimesAverage -> ResponseTimeAverageInterval?)
  * [Account] double check for missing attributes
  * [Monitor] add missing form structs
  * I don't like the CustomHttpStatus struct name
  * just use lowercase version of structs instead of unencodable*
  * there might need to be some distinction/cleanup between api json serialization vs normal json serialization (ex. UptimeDuration 123-234-234 vs {"up": 123, "down": ...})
  * consistent naming in tests (ex. want vs expected)
  * definitely need to make sure we use pointers for some stuff that's optional especially for dates (ex. maintenance_window.status)
  * maybe make a function for Maintenance window to get the next time.Time that it will be running
  * handle retry error headers
  * GetMonitorsRequest Response times should accept time.Time and convert to unix time
  * add options to hide sensitive info like psp>password
  * should url.URL be *url.URL?
* commands
  * separate user agent for the command line client than the library usage
  * ensure we conform to common api client common usage patterns
  * find out how other command line clients handle pagination
  * zero values vs nils on flags
  * NO_COLOR for commands - http://no-color.org/
  * show error if `monitors list --logs-limit -1`
  * dashboard with bubbletea components
  * fix generic type information in command help
  * maybe move flag stuff to a shared package to make building other commands easier
  * fix command help documentation
  * time.Time flags should also support a date format string to allow more flexibility
  * command output formatters like docker and kubectl
  * remove Get* prefix when not needed (ex. GetSlice)
  * verify all premium features and creating/editing them
  * use consistent capitalization of HTTP in field/struct/function names
  * if `monitor list` args are a string then use search. if ids conflict then split into 2 requests and merge
* uncategorized
  * does it cause bad ergonomics to strict type ResponseTimeEntry instead of keeping them as unix timestamp and ms
  * maybe group stuff like all of the ResponseTime stuff
  * add golangci-lint to github actions
  * see if there is anything worth borrowing from https://github.com/bitfield/uptimerobot