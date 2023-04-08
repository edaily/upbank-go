package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/alecthomas/kong"
	"github.com/edaily/upbank-go/upsdk"
)

var cli struct {
	Token string `help:"Authentication token" required:"" env:"UP_TOKEN"`

	Accounts struct {
	} `cmd:"" help:"List accounts" type:"path"`

	Tags struct {
	} `cmd:"" help:"List tags" type:"path"`

	Transactions struct {
	} `cmd:"" help:"List transactions" type:"path"`

	Categories struct {
	} `cmd:"" help:"List categories" type:"path"`

	Ping struct {
	} `cmd:"" help:"Ping the upsdk" type:"path"`
}

func main() {
	ctx := kong.Parse(&cli)
	client := upsdk.NewClient(http.DefaultClient, cli.Token)

	switch ctx.Command() {
	case "tags":
		tags, _, err := client.Tags.List()
		if err != nil {
			ctx.Fatalf(err.Error())
		}
		fmt.Println(createJsonStringFromInterface(tags))

	case "accounts":
		accounts, _, err := client.Accounts.List()
		if err != nil {
			ctx.Fatalf(err.Error())
		}
		fmt.Println(createJsonStringFromInterface(accounts))

	case "transactions":
		transactions, _, err := client.Transactions.List()
		if err != nil {
			ctx.Fatalf(err.Error())
		}
		fmt.Println(createJsonStringFromInterface(transactions))

	case "categories":
		catagories, _, err := client.Categories.List()
		if err != nil {
			ctx.Fatalf(err.Error())
		}
		fmt.Println(createJsonStringFromInterface(catagories))

	case "ping":
		ping, _, err := client.Utils.Ping()
		if err != nil {
			ctx.Fatalf(err.Error())
		}
		fmt.Println(createJsonStringFromInterface(ping))

	default:
		panic(ctx.Command())
	}
}

func createJsonStringFromInterface(x interface{}) string {
	json, err := json.Marshal(x)
	if err != nil {
		log.Fatalf("unable to marshall %v, %v", json, err)
	}
	return string(json)
}
