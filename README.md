# print-timezone

Print the time in multiple time zone.

## USAGE

```
# now
$ print-timezone
2017-01-15 09:43 UTC +0000
2017-01-15 01:43 PST -0800
2017-01-15 09:43 GMT +0000
2017-01-15 18:43 JST +0900
2017-01-15 20:43 AEDT +1100

# specify time
$ print-timezone 2017-12-31 23:59 JST
2017-12-31 14:59 UTC +0000
2017-12-31 06:59 PST -0800
2017-12-31 14:59 GMT +0000
2017-12-31 23:59 JST +0900
2018-01-01 01:59 AEDT +1100
```

## Configuration

Load `~/.print-timezone.yml` by default.

```
# Example
locations:
  - UTC
  - America/Los_Angeles
  - Europe/London
  - Europe/Berlin
  - Asia/Tokyo
  - Australia/Sydney
```

You can specify configuration file by `-f` switch.

```
$ print-timezone -f /path/to/another-timezone.yml 2017-12-31 23:59 JST
2017-12-31 14:59 UTC +0000
2017-12-31 06:59 PST -0800
2017-12-31 14:59 GMT +0000
2017-12-31 15:59 CET +0100
2017-12-31 23:59 JST +0900
2018-01-01 01:59 AEDT +1100
```
