# ReportHealCheck
This project is program to check website and response time when given the CSV list, calls the Healthcheck Report API to send the statistic of each website.

#### Setup project env

```sh
PORT=8000
CHANELID=1234
CHANELSECRET=123456
REDIRECTURL=https://www.oommi04.com/callback
REPORTHEALCHECLURL=https://www.oommi04.com/healcheck/report
```

#### How to run

```bash
$ go run main.go emp.csv
or
$ go run main.go -accessToken="token" emp.csv
```

> NOTE: If you run the program with the access token, the login process would be skipped. But if your access token is denied, the program will verify the login process again.

#### How to build

```bash
$ go build
```

#### How to test

```bash
$ go test ./... -v
```

#### Folder Architecture

```sh
├── usecase
│   └── usecase1
│       └── usecase1Interface.go
├── domain
│   └── domain1
│       ├── errors.go
│       └── entity.go
├── drivers
│   └── driver1
│      └──driver.go
├── external
│   └── external1
│      └──external1.go
└── README.md
```

#### Layer


| Layer | Folder |
| ------ | ------ |
| Entities | domain |
| Use Case | usecase |
| Interface Adapters | external |
| Frameworks and Drivers | drivers |


#### FlowChart
![alt text](https://github.com/oommi04/ReportHealcheck/blob/master/flowChart.png?raw=true)
