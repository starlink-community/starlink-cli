This is a basic command line interface that demonstrates using the [golang GRPC API](starlink-community/starlink-grpc-go).

The API is reused between dishy and the wifi router, however both servers do not implement all the APIs. The `basic-cli` demonstrates that the `status` API works on dishy, but not on the router. Conversely, `ping` works on the router, but not dishy. 

```
$ go get github.com/starlink-community/starlink-cli/cmd/basic-cli
$ basic-cli -h
Usage of basic-cli:
  -addr string
    	grpc addr (dishy is at 192.168.100.1:9200, wifi is at 192.168.1.1:9000 (default "192.168.100.1:9200")
  -req string
    	status or ping (default "status")
$ basic-cli 
dish_get_status: <
  device_info: <
    id: "ut01000000-00000000-00001234"
    hardware_version: "rev1_pre_production"
    software_version: "e68dfc80-fa1a-4fa4-9b21-d7ee2a918496.release"
    country_code: "US"
  >
  device_state: <
    uptime_s: 312929
  >
  state: CONNECTED
  alerts: <
  >
  snr: 9
  downlink_throughput_bps: 5737.808
  uplink_throughput_bps: 1006.63293
  pop_ping_latency_ms: 29.75
  obstruction_stats: <
    last_24h_obstructed_s: 4
    valid_s: 61838.32
    wedge_fraction_obstructed: 0
    wedge_fraction_obstructed: 0
    wedge_fraction_obstructed: 0
    wedge_fraction_obstructed: 0
    wedge_fraction_obstructed: 0
    wedge_fraction_obstructed: 0
    wedge_fraction_obstructed: 0
    wedge_fraction_obstructed: 0
    wedge_fraction_obstructed: 0
    wedge_fraction_obstructed: 0
    wedge_fraction_obstructed: 0
    wedge_fraction_obstructed: 0
    wedge_abs_fraction_obstructed: 0
    wedge_abs_fraction_obstructed: 0
    wedge_abs_fraction_obstructed: 0
    wedge_abs_fraction_obstructed: 0
    wedge_abs_fraction_obstructed: 0
    wedge_abs_fraction_obstructed: 0
    wedge_abs_fraction_obstructed: 0
    wedge_abs_fraction_obstructed: 0
    wedge_abs_fraction_obstructed: 0
    wedge_abs_fraction_obstructed: 0
    wedge_abs_fraction_obstructed: 0
    wedge_abs_fraction_obstructed: 0
  >
>

$ basic-cli -req ping -addr 192.168.1.1:9000
get_ping: <
  results: <
    key: "103.10.124.1"
    value: <
      target: <
        service: "Counter-Strike: GO"
        location: "Singapore"
        address: "103.10.124.1"
      >
      dropRate: 0.4
      latencyMs: 217.1374
    >
  >
  results: <
    key: "103.10.125.1"
    value: <
      target: <
        service: "Counter-Strike: GO"
        location: "Sydney"
        address: "103.10.125.1"
      >
      dropRate: 1
    >
  >
  results: <
    key: "104.160.131.3"
    value: <
      target: <
        service: "League of Legends"
        location: "Chicago"
        address: "104.160.131.3"
      >
      dropRate: 1
    >
  >
  results: <
    key: "104.160.136.3"
    value: <
      target: <
        service: "League of Legends"
        location: "Los Angeles"
        address: "104.160.136.3"
      >
      dropRate: 1
    >
  >
  results: <
    key: "104.160.141.3"
    value: <
      target: <
        service: "League of Legends"
        location: "Amsterdam"
        address: "104.160.141.3"
      >
      dropRate: 1
    >
  >
  results: <
    key: "104.160.142.3"
    value: <
      target: <
        service: "League of Legends"
        location: "Frankfurt"
        address: "104.160.142.3"
      >
      dropRate: 1
    >
  >
  results: <
    key: "104.160.156.1"
    value: <
      target: <
        service: "League of Legends"
        location: "Sydney"
        address: "104.160.156.1"
      >
      dropRate: 1
    >
  >
  results: <
    key: "146.66.152.1"
    value: <
      target: <
        service: "Counter-Strike: GO"
        location: "Luxemborg"
        address: "146.66.152.1"
      >
      dropRate: 1
    >
  >
  results: <
    key: "146.66.155.1"
    value: <
      target: <
        service: "Counter-Strike: GO"
        location: "Vienna"
        address: "146.66.155.1"
      >
      dropRate: 0.4
      latencyMs: 210.87733
    >
  >
  results: <
    key: "146.66.156.1"
    value: <
      target: <
        service: "Counter-Strike: GO"
        location: "Stockholm"
        address: "146.66.156.1"
      >
      dropRate: 1
    >
  >
  results: <
    key: "185.25.183.1"
    value: <
      target: <
        service: "Counter-Strike: GO"
        location: "Dubai"
        address: "185.25.183.1"
      >
      dropRate: 0.6
      latencyMs: 335.9776
    >
  >
  results: <
    key: "192.168.100.1"
    value: <
      target: <
        service: "Your Starlink"
        address: "192.168.100.1"
      >
      latencyMs: 31.793804
    >
  >
  results: <
    key: "192.69.96.1"
    value: <
      target: <
        service: "Counter-Strike: GO"
        location: "Seattle"
        address: "192.69.96.1"
      >
      latencyMs: 61.9395
    >
  >
  results: <
    key: "197.80.200.1"
    value: <
      target: <
        service: "Counter-Strike: GO"
        location: "Johannesburg"
        address: "197.80.200.1"
      >
      dropRate: 0.6
      latencyMs: 371.28906
    >
  >
  results: <
    key: "208.78.164.1"
    value: <
      target: <
        service: "Counter-Strike: GO"
        location: "Virginia"
        address: "208.78.164.1"
      >
      dropRate: 0.2
      latencyMs: 149.59204
    >
  >
  results: <
    key: "209.197.29.1"
    value: <
      target: <
        service: "Counter-Strike: GO"
        location: "S\303\243o Paulo"
        address: "209.197.29.1"
      >
      dropRate: 1
    >
  >
  results: <
    key: "dynamodb.ap-east-1.amazonaws.com"
    value: <
      target: <
        service: "Fortnite"
        location: "Hong Kong"
        address: "dynamodb.ap-east-1.amazonaws.com"
      >
      dropRate: 0.2
      latencyMs: 189.16917
    >
  >
  results: <
    key: "dynamodb.ap-northeast-1.amazonaws.com"
    value: <
      target: <
        service: "Fortnite"
        location: "Tokyo"
        address: "dynamodb.ap-northeast-1.amazonaws.com"
      >
      dropRate: 0.2
      latencyMs: 169.83228
    >
  >
  results: <
    key: "dynamodb.ap-northeast-2.amazonaws.com"
    value: <
      target: <
        service: "Fortnite"
        location: "Seoul"
        address: "dynamodb.ap-northeast-2.amazonaws.com"
      >
      dropRate: 0.4
      latencyMs: 229.78513
    >
  >
  results: <
    key: "dynamodb.ap-northeast-3.amazonaws.com"
    value: <
      target: <
        service: "Fortnite"
        location: "Osaka"
        address: "dynamodb.ap-northeast-3.amazonaws.com"
      >
      dropRate: 0.4
      latencyMs: 171.29086
    >
  >
  results: <
    key: "dynamodb.ap-south-1.amazonaws.com"
    value: <
      target: <
        service: "Fortnite"
        location: "Mumbai"
        address: "dynamodb.ap-south-1.amazonaws.com"
      >
      dropRate: 0.6
      latencyMs: 306.35556
    >
  >
  results: <
    key: "dynamodb.ap-southeast-1.amazonaws.com"
    value: <
      target: <
        service: "Fortnite"
        location: "Singapore"
        address: "dynamodb.ap-southeast-1.amazonaws.com"
      >
      dropRate: 0.4
      latencyMs: 229.85423
    >
  >
  results: <
    key: "dynamodb.ap-southeast-2.amazonaws.com"
    value: <
      target: <
        service: "Fortnite"
        location: "Sydney"
        address: "dynamodb.ap-southeast-2.amazonaws.com"
      >
      dropRate: 0.6
      latencyMs: 248.15205
    >
  >
  results: <
    key: "dynamodb.ca-central-1.amazonaws.com"
    value: <
      target: <
        service: "Fortnite"
        location: "Montreal"
        address: "dynamodb.ca-central-1.amazonaws.com"
      >
      dropRate: 0.2
      latencyMs: 154.19655
    >
  >
  results: <
    key: "dynamodb.cn-north-1.amazonaws.com.cn"
    value: <
      target: <
        service: "Fortnite"
        location: "Beijing"
        address: "dynamodb.cn-north-1.amazonaws.com.cn"
      >
      dropRate: 0.6
      latencyMs: 332.88885
    >
  >
  results: <
    key: "dynamodb.cn-northwest-1.amazonaws.com.cn"
    value: <
      target: <
        service: "Fortnite"
        location: "Ningxia"
        address: "dynamodb.cn-northwest-1.amazonaws.com.cn"
      >
      dropRate: 0.4
      latencyMs: 307.33124
    >
  >
  results: <
    key: "dynamodb.eu-central-1.amazonaws.com"
    value: <
      target: <
        service: "Fortnite"
        location: "Frankfurt"
        address: "dynamodb.eu-central-1.amazonaws.com"
      >
      dropRate: 0.4
      latencyMs: 223.75505
    >
  >
  results: <
    key: "dynamodb.eu-north-1.amazonaws.com"
    value: <
      target: <
        service: "Fortnite"
        location: "Stockholm"
        address: "dynamodb.eu-north-1.amazonaws.com"
      >
      dropRate: 0.4
      latencyMs: 222.32887
    >
  >
  results: <
    key: "dynamodb.eu-west-1.amazonaws.com"
    value: <
      target: <
        service: "Fortnite"
        location: "Ireland"
        address: "dynamodb.eu-west-1.amazonaws.com"
      >
      dropRate: 0.4
      latencyMs: 252.07127
    >
  >
  results: <
    key: "dynamodb.eu-west-2.amazonaws.com"
    value: <
      target: <
        service: "Fortnite"
        location: "London"
        address: "dynamodb.eu-west-2.amazonaws.com"
      >
      dropRate: 0.2
      latencyMs: 183.14899
    >
  >
  results: <
    key: "dynamodb.eu-west-3.amazonaws.com"
    value: <
      target: <
        service: "Fortnite"
        location: "Paris"
        address: "dynamodb.eu-west-3.amazonaws.com"
      >
      dropRate: 0.4
      latencyMs: 243.78366
    >
  >
  results: <
    key: "dynamodb.me-south-1.amazonaws.com"
    value: <
      target: <
        service: "Fortnite"
        location: "Bahrain"
        address: "dynamodb.me-south-1.amazonaws.com"
      >
      dropRate: 0.4
      latencyMs: 322.5198
    >
  >
  results: <
    key: "dynamodb.sa-east-1.amazonaws.com"
    value: <
      target: <
        service: "Fortnite"
        location: "S\303\243o Paulo"
        address: "dynamodb.sa-east-1.amazonaws.com"
      >
      dropRate: 0.4
      latencyMs: 262.59018
    >
  >
  results: <
    key: "dynamodb.us-east-1.amazonaws.com"
    value: <
      target: <
        service: "Fortnite"
        location: "Virginia"
        address: "dynamodb.us-east-1.amazonaws.com"
      >
      dropRate: 0.2
      latencyMs: 146.55888
    >
  >
  results: <
    key: "dynamodb.us-east-2.amazonaws.com"
    value: <
      target: <
        service: "Fortnite"
        location: "Ohio"
        address: "dynamodb.us-east-2.amazonaws.com"
      >
      dropRate: 0.2
      latencyMs: 120.64577
    >
  >
  results: <
    key: "dynamodb.us-west-1.amazonaws.com"
    value: <
      target: <
        service: "Fortnite"
        location: "California"
        address: "dynamodb.us-west-1.amazonaws.com"
      >
      latencyMs: 84.610985
    >
  >
  results: <
    key: "dynamodb.us-west-2.amazonaws.com"
    value: <
      target: <
        service: "Fortnite"
        location: "Oregon"
        address: "dynamodb.us-west-2.amazonaws.com"
      >
      latencyMs: 55.702183
    >
  >
  results: <
    key: "google.com"
    value: <
      target: <
        service: "Google"
        address: "google.com"
      >
      dropRate: 0.2
      latencyMs: 51.706455
    >
  >
>
```

