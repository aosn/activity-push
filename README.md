AOSN activity-push
==================

[![Build Status](http://jenkins.tasktoys.com/buildStatus/icon?job=aosn/activity-push/master)](http://jenkins.tasktoys.com/job/aosn/job/activity-push/job/master/)
[![Code Climate](https://codeclimate.com/github/aosn/activity-push/badges/gpa.svg)](https://codeclimate.com/github/aosn/activity-push)
[![Go Report Card](https://goreportcard.com/badge/github.com/aosn/activity-push)](https://goreportcard.com/report/github.com/aosn/activity-push)

## LT

データで見るAOSN読書会 (2016 合宿 LT)

http://www.slideshare.net/YutakaKato/aosn-2016-lt

## How it works

1. Parse markdown table of the AOSN workshop activities.
2. Generate JSON data.
3. Push to Elasticsearch.

## Usage

* `-h` - Host of your Elasticsearch (default: `localhost`)
* `-p` - Port of your Elasticsearch (default: `9200`)
* `-t` - Target file of [workshop record](https://github.com/aosn/aosn.github.io/tree/master/workshop) (e.g. 1-java8)
* `-stat` - Print statistics markdown (don't push to Elasticsearch)

Short example:

```bash
./activity-push -t 1-java8 
```
Complete example:

```bash
./activity-push -h localhost -p 9200 -t 1-java8 
```

Print statistics example:

```bash
./activity-push -stat -t 1-java8
```


## Author

mikan

## License

Apache License 2.0
