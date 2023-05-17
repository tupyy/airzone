# airzone

## Introduction
Use Airzone to control and get information from your Airzone VAF webserver. 
Airzone allows you to control VAF parameters for one or all zones. 
Also, Airzone starts a metric server which servers zones temperature and humidity to Prometheus.


## Build
### Requirements
- Go1.19

### Using shell
```shell
make build.vendor
make build.local
```

### Using Podman
```shell
make build.podman
```


## Usage
```shell
➜  airzone 
Control your Airzone VAF

Usage:
  airzone [command]

Available Commands:
  completion  Generate the autocompletion script for the specified shell
  help        Help about any command
  metrics     Start the metric server
  zone        Control one zone only.
  zones       Control all the zones all together.

Flags:
    -h, --help            help for airzone
    --host string     airzone host url. Example: 192.168.1.1:3000 (default "airzone:3000")
    --log-level string   Log level (default "debug")
    -o, --output string      output style.Acceptes json or table (default "json")
    --system-id int   system id (default 1)

Use "airzone [command] --help" for more information about a command.
```

### Metrics
`metrics` command starts a server which serves temperature and humidity metrics for all zones.

```shell
➜  airzone metrics --help
This command starts a http server having single endpoint /metrics.
These metrics can be scraped by prometheus.

Usage:
  airzone metrics [flags]

Flags:
  -h, --help         help for metrics
      --port int     http port of the metric server (default 8080)
      --ticker int   interval of request to airzone server (default 10)

Global Flags:
      --host string     airzone host url. Example: 192.168.1.1:3000 (default "airzone:3000")
      --system-id int   system id (default 1)
```
The default port is `8080` and the metrics are fetched from VAF webserver every `10s`. 
The following metrics are exported:
- temperature
- humidity
- state (0 or 1 depeding of the the hvac is turn on or not)
- mode (heating, cooling ...)

### Zones

`zones` display information about all zones.
> The information could be displayed as table or json.

```shell
airzone zones -otable

--------+--------+-----+---------+-----------------+-----------------+----------------+------------+
| ZONEID | NAME   |  ON | MODE    | COOLSETPOINT °C | HEATSETPOINT °C | TEMPERATURE °C | HUMIDITY % |
+--------+--------+-----+---------+-----------------+-----------------+----------------+------------+
|      1 | xxxxxx | off | heating |           20.50 |           20.50 |          20.50 |         56 |
|      2 | xxxxxx | off | heating |           20.50 |           20.50 |          19.30 |         56 |
|      5 | xxxxxx | off | heating |           21.00 |           21.00 |          20.70 |         51 |
|      6 | xxxxxx | off | heating |           20.50 |           20.50 |          20.10 |         60 |
+--------+--------+-----+---------+-----------------+-----------------+----------------+------------+%
```

To set the temperature for all zone, you can use `set` subcommand:
```shell
airzone zones set temperature <value>
```


To turn on hvac for all zone, you can use `turn` subcommand:
```shell
airzone zones turn <on|off>
```

### Zone
`zone` display the information of a single zone.
```shell
airzone zone <name|id>
```
The command accepts either the name of the zone or the id.


To set the temperature for a particular zone, you can use `set` subcommand:
```shell
airzone zone set <zone-name|zone-id> temperature <value>
```

To turn on hvac for a zone, you can use `turn` subcommand:
```shell
airzone zone turn <name|id> <on|off>
```
