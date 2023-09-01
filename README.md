# go-uptimemonitor-v2
Uptime Robot APIv2 command-line client and library

[UptimeRobot API Documentation](https://uptimerobot.com/api)


## Build
```shell
go build -o uptimerobot cmd/uptimerobot/main.go
```

## Usage

```shell
./uptimerobot GROUP [COMMAND] --api-key=$KEY
```

### Implemented groups
* account - displays account information
* help
* other cobra built-ins