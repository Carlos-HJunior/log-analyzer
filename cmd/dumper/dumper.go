package main

import (
    "bufio"
    "log"
    "log-analyzer/cmd/dumper/parser"
    "log-analyzer/datastore"
    "os"
    "sync"
)

const maxChunkSize = 3000

func main() {
    log.Println("start")

    db, err := datastore.NewDatastore()
    if err != nil {
        log.Fatalln(err)
    }

    file, err := os.Open("golang.log")
    if err != nil {
        log.Fatalln(err)
    }

    scanner := bufio.NewScanner(file)

    var wg sync.WaitGroup

    var chunks []string

    for scanner.Scan() {
        chunks = append(chunks, scanner.Text())

        if len(chunks) < maxChunkSize {
            continue
        }

        wg.Add(1)
        go routine(chunks, &wg, db)

        chunks = []string{}
    }

    wg.Add(1)
    go routine(chunks, &wg, db)

    wg.Wait()
    log.Println("end")
}

func routine(chunk []string, wg *sync.WaitGroup, db *datastore.Mysql) {
    defer wg.Done()
    var items []parser.ResolvedLine

    for _, line := range chunk {

        item, err := parser.Parse(line)
        if err != nil {
            log.Println(err)
            continue
        }

        if item.Request != "GET" && item.Request != "POST" {
            continue
        }

        items = append(items, item)
    }

    err := db.SaveLogs(items)
    if err != nil {
        log.Fatalln(err)
    }
}
