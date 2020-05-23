package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"

	"k8s.io/apimachinery/pkg/util/wait"
)

func main() {

	token := "xoxp-2253118358-891650867716-1125531635331-f2eec34c6ebd6065f85aeaba913d4163"
	startURL := fmt.Sprintf("https://slack.com/api/conversations.list?token=%s", token)
	id, name := getDataFromRestAPI(startURL)
	hist := fmt.Sprintf("https://slack.com/api/conversations.history?token=%s&channel=%s", token, id)
	//id, name := getDataFromRestAPI(startURL)
	fmt.Print("hello, ", hist)
	fmt.Print("hello, ", name)

}

func getDataFromRestAPI(startURL string) (string, string) {
	var id, name string

	err := wait.PollImmediate(5*time.Second, 1*time.Minute, func() (bool, error) {
		resp, err := http.Get(startURL)
		testerr(err)
		defer resp.Body.Close()

		decoder := json.NewDecoder(resp.Body)
		var data SlackResponse

		err = decoder.Decode(&data)
		testerr(err)
		for i, rack := range data.Channels {
			if strings.Contains(rack.Name, "coreos") {
				fmt.Printf("%d: ID-<<%s>> Name-<<%s>>\n", i, rack.ID, rack.Name)
				id = rack.ID
				name = rack.Name
				break
			}
		}
		return true, err
	})
	testerr(err)
	return id, name
}

func testerr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
