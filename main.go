package main

import (
	"context"
	"log"

	"github.com/vroomy/vroomy"

	_ "github.com/vroomy/hello-world/plugins/companies"
)

func main() {
	var (
		svc *vroomy.Vroomy
		err error
	)

	if svc, err = vroomy.New("./config.toml"); err != nil {
		log.Fatal(err)
	}

	if err = svc.ListenUntilSignal(context.Background()); err != nil {
		log.Fatal(err)
	}
}
