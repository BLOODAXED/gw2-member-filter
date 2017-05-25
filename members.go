package main

import "net/http"
import "fmt"
import "io/ioutil"
import "encoding/json"

import "os"
import str "strings"
import "time"
import "flag"

type Member struct {
	Name   string
	Rank   string
	Joined string
}
type Conf struct {
	AccessToken string
	GuildID     string
}

func main() {

	flag.Parse()
	file, _ := os.Open("config.json")
	decoder := json.NewDecoder(file)
	conf := Conf{}
	err := decoder.Decode(&conf)
	if err != nil {
		fmt.Println("error:", err)
	}

	var when string
	whenParsed := time.Now()

	flag.StringVar(&when, "date", "default", "a date in YYYY-MM-DD format")

	//get the rank argument
	rank := string(flag.Arg(0))

	if when == "default" {
		whenParsed = time.Now()
	} else {
		whenParsed, _ = time.Parse("2006-01-02", when)
	}

	var m []Member

	resp, err := http.Get("https://api.guildwars2.com/v2/guild/" + conf.GuildID + "/members?access_token=" + conf.AccessToken)
	if err != nil {

	}

	defer resp.Body.Close()
	if resp.StatusCode == 200 {
		body, _ := ioutil.ReadAll(resp.Body)
		err := json.Unmarshal(body, &m)
		if err != nil {
			fmt.Println("error", err)
		}
		for _, x := range m {
			tmp, _ := time.Parse(time.RFC3339, x.Joined)
			if str.ToLower(x.Rank) == str.ToLower(rank) {
				if tmp.Before(whenParsed) {
					fmt.Printf("%+v\n", x)
				}
			}
		}
	}
}
