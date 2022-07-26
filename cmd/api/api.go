package main

import (
    "github.com/labstack/echo/v4"
    "log"
    "log-analyzer/datastore"
    "net/http"
    "strconv"
)

func main() {
    db, err := datastore.NewDatastore()
    if err != nil {
        log.Fatalln(err)
    }

    e := echo.New()
    e.GET("/last/:qty/:method", func(ctx echo.Context) error {
        method := ctx.Param("method")
        qty := ctx.Param("qty")

        atoi, err := strconv.Atoi(qty)
        if err != nil {
            return err
        }

        err, result := db.GetLogsByMethod(method, uint(atoi))
        if err != nil {
            return err
        }

        return ctx.JSON(http.StatusOK, result)
    })

    e.Logger.Fatal(e.Start(":8080"))
}
