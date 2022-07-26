package parser

import (
    "errors"
    "regexp"
    "strconv"
    "time"
)

const regex = `(?P<ip>^\d+.\d+.\d+.\d+|^[\d\w:]+|-) - - \[(?P<date>\d+/[a-zA-Z]+/\d+:\d+:\d+:\d+ [+-]\d+|-)] \"(?P<method>[A-Z]+|-) (?P<path>/[0-9a-zA-Z-_+?=+,%:;~{}()\[\]&\/\.]*|-) (?P<version>HTTP/\d+.\d+|-)\" (?P<status>\d+|-) (?P<lenght>\d+|-) \"(?P<referer>[\w\-_+?=+,%:;&\/\.]+|.*)\" \"(?P<agent>.*)\"`
const timeLayout = "02/Jan/2006:15:04:05 -0700"

type ResolvedLine struct {
    ConnectingIp string
    TimeLocal    time.Time
    Request      string
    BodyBytes    uint
}

func Parse(line string) (resolvedLine ResolvedLine, err error) {
    re := regexp.MustCompile(regex)
    var result = re.FindStringSubmatch(line)

    if len(result) < 10 {
        err = errors.New("log file does not match regex")
        return
    }

    tl, err := time.Parse(timeLayout, result[2])
    if err != nil {
        return
    }

    atoi, err := strconv.Atoi(result[7])
    if err != nil {
        return
    }

    resolvedLine = ResolvedLine{
        ConnectingIp: result[1],
        TimeLocal:    tl,
        Request:      result[3],
        BodyBytes:    uint(atoi),
    }

    return
}
