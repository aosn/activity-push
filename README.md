chartgen
========

[![Code Climate](https://codeclimate.com/github/aosn/chartgen/badges/gpa.svg)](https://codeclimate.com/github/aosn/chartgen)
[![Go Report Card](https://goreportcard.com/badge/github.com/aosn/chartgen)](https://goreportcard.com/report/github.com/aosn/chartgen)

## LT

データで見るAOSN読書会 (2016 合宿 LT)

https://go-talks.appspot.com/github.com/mikan/talks/aosn-lt-2016.slide

## How it works

1. Parse markdown table of the AOSN workshop activities.
2. Generate JSON data.
3. Push to Elasticsearch.

## Usage

### aosn

`aosn` generates summary description written by markdown.

* `-t` - Target file of the [workshop record](https://github.com/aosn/aosn.github.io/tree/master/workshop) (e.g. 1-java8)

```bash
./aosn -t 1-java8
```

### aosn2es

`aosn2es` generates Elasticsearch (Kibana) datasets.

* `-h` - Host of your Elasticsearch (default: `localhost`)
* `-p` - Port of your Elasticsearch (default: `9200`)
* `-t` - Target file of the [workshop record](https://github.com/aosn/aosn.github.io/tree/master/workshop) (e.g. 1-java8)

Short example:

```bash
./aosn2es -t 1-java8 
```
Complete example:

```bash
./aosn2es -h localhost -p 9200 -t 1-java8 
```

## Author

mikan

## License

Apache License 2.0
