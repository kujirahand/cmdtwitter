package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"net/url"

	"github.com/ChimeraCoder/anaconda"
)

func main() {
	// get parameters
	key := flag.String("key", "", "API Key")
	sec := flag.String("sec", "", "API Secret Key")
	accTok := flag.String("acc_tok", "", "Access Token")
	accSec := flag.String("acc_sec", "", "Access Token Secret")
	cmd := flag.String("cmd", "tweet", "Command, cmd=tweek,hometimeline")
	text := flag.String("text", "", "Text")
	cnt := flag.String("count", "10", "Count")
	flag.Parse()

	// fmt.Println(flag.Args())
	if key == nil || sec == nil {
		fmt.Println("error: no key, sec parameters")
		return
	}
	if accTok == nil || accSec == nil {
		fmt.Println("error: no acc_tok, acc_sec parameters")
		return
	}

	// fmt.Println(flag.Args())
	// fmt.Printf("key=[%s]\nsec=[%s]\n", *key, *sec)
	// fmt.Printf("acc_tok=[%s]\nacc_sec=[%s]\n", *accTok, *accSec)

	api := anaconda.NewTwitterApiWithCredentials(*accTok, *accSec, *key, *sec)
	if *cmd == "tweet" && text == nil {
		fmt.Println("error: no text parameter")
		return
	}

	if *cmd == "tweet" {
		_, err := api.PostTweet(*text, nil)
		if err != nil {
			fmt.Printf("error: %s\n", err.Error())
			return
		}
		fmt.Println("ok")
		return
	}
	if *cmd == "hometimeline" {
		v := url.Values{}
		v.Set("count", *cnt)
		tweets, err := api.GetHomeTimeline(v)
		if err != nil {
			fmt.Println("error: %w", err)
			return
		}
		jsonData, _ := json.MarshalIndent(tweets, "", "  ")
		fmt.Println(string(jsonData))
		return
	}
	if *cmd == "usertimeline" {
		v := url.Values{}
		v.Set("count", *cnt)
		tweets, err := api.GetUserTimeline(v)
		if err != nil {
			fmt.Println("error: %w", err)
			return
		}
		jsonData, _ := json.MarshalIndent(tweets, "", "  ")
		fmt.Println(string(jsonData))
		return
	}

	fmt.Println("error: cmd parameter unsupported command")
}
