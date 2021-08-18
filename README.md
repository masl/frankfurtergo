[![Go Report Card](https://goreportcard.com/badge/github.com/masl/frankfurtergo)](https://goreportcard.com/report/github.com/masl/frankfurtergo)
[![Go Reference](https://pkg.go.dev/badge/github.com/masl/frankfurtergo.svg)](https://pkg.go.dev/github.com/masl/frankfurtergo)

# frankfurtergo
A neat Go wrapper for the [frankfurter.app](https://www.frankfurter.app/docs/) API

## Installation
```bash
go get -u github.com/masl/frankfurtergo
```

## Usage

### Intialize
```go
client := frankfurtergo.New()
```

### Latest Data
```go
latest, err := client.FetchLatest()
if err != nil {
    return err
}
```

### Historical Data
```go
t := time.Date(
    2010, 8, 18, 16, 0, 0, 0, time.UTC)

hist, err := client.FetchHistorical(t)
if err != nil {
    return err
}
```

### Time Series Data
```go
f := time.Date(
    2010, 8, 18, 16, 0, 0, 0, time.UTC)

t := time.Date(
    2012, 4, 20, 16, 0, 0, 0, time.UTC)

series, err := client.FetchSeries(f, t)
if err != nil {
    return err
}
```

## Contribution
Feel free to help me out as I'm learning Go with this project and would appreciate any kind of help
