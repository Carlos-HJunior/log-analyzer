package datastore

import (
    "database/sql"
    _ "github.com/go-sql-driver/mysql"
    "log-analyzer/cmd/dumper/parser"
)

func NewDatastore() (*Mysql, error) {
    db, err := sql.Open("mysql", "root:Root@123@tcp(localhost:3306)/dblog?parseTime=true")
    if err != nil {
        return nil, err
    }

    return &Mysql{
        db: db,
    }, nil
}

type Mysql struct {
    db *sql.DB
}

func (s Mysql) SaveLogs(list []parser.ResolvedLine) error {

    batch, err := s.db.Begin()
    if err != nil {
        return err
    }

    stmt, err := batch.Prepare(`
        INSERT IGNORE INTO log (http_cf_connecting_ip, time_local, method, body_bytes_sent)
                VALUES (?, ?, ?, ?);
    `)
    if err != nil {
        return err
    }

    for _, it := range list {
        _, err = stmt.Exec(it.ConnectingIp, it.TimeLocal, it.Request, it.BodyBytes)
        if err != nil {
            return err
        }
    }

    stmt.Close()
    err = batch.Commit()
    if err != nil {
        return err
    }
    return nil
}

func (s Mysql) GetLogsByMethod(method string, qty uint) (err error, items []parser.ResolvedLine) {
    query, err := s.db.Query(`
        SELECT http_cf_connecting_ip,
               time_local,
               method,
               body_bytes_sent 
        FROM log 
        WHERE method = ?
        ORDER BY time_local DESC
        LIMIT ?`,
        method, qty,
    )

    if err != nil {
        return
    }

    for query.Next() {
        var it parser.ResolvedLine

        err = query.Scan(
            &it.ConnectingIp,
            &it.TimeLocal,
            &it.Request,
            &it.BodyBytes,
        )

        if err != nil {
            return
        }

        items = append(items, it)
    }

    err = query.Close()
    return
}

func (s Mysql) Dispose() error {
    return s.db.Close()
}
