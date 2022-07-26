package parser

import (
    "github.com/stretchr/testify/assert"
    "testing"
    "time"
)

func TestParser(t *testing.T) {

    t.Run("test - 1", func(t *testing.T) {
        tl, err := time.Parse(timeLayout, "01/Jan/2020:14:17:02 +0000")
        if err != nil {
            return
        }

        test(
            t,
            "123.456.189.000 - - [01/Jan/2020:14:17:02 +0000] \"GET / HTTP/1.1\" 200 11111 \"https://unit-test.com/\" \"Unit test agent\"",
            ResolvedLine{
                ConnectingIp: "123.456.189.000",
                TimeLocal:    tl,
                Request:      "GET",
                BodyBytes:    11111,
            },
        )
    })

    t.Run("test - 2", func(t *testing.T) {
        tl, err := time.Parse(timeLayout, "01/Jan/2020:14:17:02 +0000")
        if err != nil {
            return
        }

        test(
            t,
            "2001:8a0:ffde:4e00:f463:3069:5ab4:2ae5 - - [01/Jan/2020:14:17:02 +0000] \"GET / HTTP/1.1\" 200 11111 \"https://unit-test.com/\" \"Unit test agent\"\n",
            ResolvedLine{
                ConnectingIp: "2001:8a0:ffde:4e00:f463:3069:5ab4:2ae5",
                TimeLocal:    tl,
                Request:      "GET",
                BodyBytes:    11111,
            },
        )
    })

    t.Run("test - 3", func(t *testing.T) {
        tl, err := time.Parse(timeLayout, "01/Jan/2020:14:17:02 +0000")
        if err != nil {
            return
        }

        test(
            t,
            "- - - [01/Jan/2020:14:17:02 +0000] \"POST / HTTP/1.1\" 404 0 \"-\" \"-\"",
            ResolvedLine{
                ConnectingIp: "-",
                TimeLocal:    tl,
                Request:      "POST",
                BodyBytes:    0,
            },
        )
    })

}

func test(t *testing.T, line string, expected ResolvedLine) {
    result, err := Parse(line)
    if err != nil {
        assert.Error(t, err)
    }

    assert.Equal(t, result, expected)
}
