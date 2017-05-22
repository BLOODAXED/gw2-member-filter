package main

import "net/http"
import "fmt"
import "io/ioutil"
import "encoding/json"

import "os"
import str "strings"
import "time"
import "flag"

func main() {
	type Member struct {
		Name   string
		Rank   string
		Joined string
	}

	type Conf struct {
		Token string
	}

	file, _ := os.Open("config.json")
	decoder := json.NewDecoder(file)
	conf := Conf{}
	err := decoder.Decode(&conf)
	if err != nil {
		fmt.Println("error:", err)
	}

	var when string
	whenParsed := time.Now()

	//get the rank argument
	flag.StringVar(&when, "date", "default", "a date in YYYY-MM-DD format")

	flag.Parse()
	rank := string(flag.Arg(0))
	//token := string(flag.Arg(1))

	fmt.Println(when)
	fmt.Println(rank)

	if when == "default" {
		whenParsed = time.Now()
	} else {
		whenParsed, _ = time.Parse("2006-01-02", when)
	}
	fmt.Println(whenParsed)

	var m []Member

	resp, err := http.Get("https://api.guildwars2.com/v2/guild/6E313C3A-02E3-4170-88D4-F1709BAE0F1A/members?access_token=" + conf.Token)
	if err != nil {

	}

	defer resp.Body.Close()
	if resp.StatusCode == 200 {
		body, _ := ioutil.ReadAll(resp.Body)
		//bodyString := string(body)
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
