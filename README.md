# publicdns

A client for [Google Public DNS](https://dns.google.com/query) written in [Go](https://golang.org/)

## Installation

```
go get -u github.com/peterhellberg/publicdns
```

### Command Line Client

```
go get -u github.com/peterhellberg/publicdns/cmd/resolve
```

```json
resolve c7.se A
{
  "Status": 0,
  "TC": false,
  "RD": true,
  "RA": true,
  "AD": true,
  "CD": false,
  "Question": [
    {
      "name": "c7.se.",
      "type": 1
    }
  ],
  "Answer": [
    {
      "name": "c7.se.",
      "type": 1,
      "TTL": 599,
      "data": "109.74.7.83"
    }
  ],
  "Comment": "Response from 93.188.0.20"
}
```
