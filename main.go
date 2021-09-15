package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"text/tabwriter"

	loggly "loggly"
	"net/http"
	"time"
)

func main() {

	type Competition struct {
		ID   int `json:"id"`
		Area struct {
			ID   int    `json:"id"`
			Name string `json:"name"`
		} `json:"area"`
		Name        string    `json:"name"`
		Code        string    `json:"code"`
		Plan        string    `json:"plan"`
		LastUpdated time.Time `json:"lastUpdated"`
	}

	type Season struct {
		ID              int         `json:"id"`
		StartDate       string      `json:"startDate"`
		EndDate         string      `json:"endDate"`
		CurrentMatchday int         `json:"currentMatchday"`
		Winner          interface{} `json:"winner"`
	}

	type Standings struct {
		Stage string      `json:"stage"`
		Type  string      `json:"type"`
		Group interface{} `json:"group"`
		Table []struct {
			Position int `json:"position"`
			Team     struct {
				ID       int    `json:"id"`
				Name     string `json:"name"`
				CrestURL string `json:"crestUrl"`
			} `json:"team"`
			PlayedGames    int         `json:"playedGames"`
			Form           interface{} `json:"form"`
			Won            int         `json:"won"`
			Draw           int         `json:"draw"`
			Lost           int         `json:"lost"`
			Points         int         `json:"points"`
			GoalsFor       int         `json:"goalsFor"`
			GoalsAgainst   int         `json:"goalsAgainst"`
			GoalDifference int         `json:"goalDifference"`
		} `json:"table"`
	}

	type Data struct {
		Competition Competition `json:"competition"`
		Season      Season      `json:"season"`
		Standings   []Standings `json:"standings"`
		Message     string      `json:"message"`
		ErrorCode   int         `json:"errorCode"`
		Error       int         `json:"error"`
	}

	url := "http://api.football-data.org/v2/competitions/2021/standings"

	var api_Token = "0d5fb989868f4421bce51517a5bbb62d"

	// Create a new request using http
	req, err := http.NewRequest("GET", url, nil)

	// add authorization header to the req
	req.Header.Add("X-Auth-Token", api_Token)

	// Send req using http Client
	http_client := &http.Client{}
	resp, err := http_client.Do(req)

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	var data Data
	err_ := json.Unmarshal([]byte(body), &data)

	if err_ != nil {
		panic(err_)
	}

	var tag string
	tag = "My-Go-Demo"

	// Instantiate the client
	client := loggly.New(tag)

	//_, err = os.Stdout.Write(body)

	if data.Message != "" {
		fmt.Println("Message :", data.Message, "Error Code:", data.ErrorCode)
		client.EchoSend("error", "The football api returned an error message of "+data.Message)
	} else if data.Error != 0 {
		fmt.Println("Error :", data.Error)
		client.EchoSend("error", "The football api returned an error status : "+fmt.Sprint(data.Error))
	} else {
		fmt.Println("======================= " + data.Competition.Name + " =======================")
		fmt.Println("From " + data.Season.StartDate + " to " + data.Season.EndDate)
		fmt.Println("=======================TABLE=======================")
		writer := tabwriter.NewWriter(os.Stdout, 0, 8, 1, '\t', tabwriter.AlignRight)

		fmt.Fprintln(writer, "# Team\t\t\tPL\tW\tD\tL\tF\tA\tGD\tPts")

		table := data.Standings[0].Table

		for i := 0; i < len(table); i++ {
			team := table[i]
			fmt.Fprintln(writer, fmt.Sprint(team.Position)+" "+
				team.Team.Name+"\t\t\t"+
				fmt.Sprint(team.PlayedGames)+"\t"+
				fmt.Sprint(team.Won)+"\t"+
				fmt.Sprint(team.Draw)+"\t"+
				fmt.Sprint(team.Lost)+"\t"+
				fmt.Sprint(team.GoalsFor)+"\t"+
				fmt.Sprint(team.GoalsAgainst)+"\t"+
				fmt.Sprint(team.GoalDifference)+"\t"+
				fmt.Sprint(team.Points))
		}

		writer.Flush()
		client.EchoSend("info", "The api returned Competition "+data.Competition.Name+" is held from "+data.Season.StartDate+" to "+data.Season.EndDate)
	}

	if err != nil {
		log.Fatal(err)
	}
	////////////////////////////////////////////////////////////////

	// 	// Valid EchoSend (message echoed to console and no error returned)
	// 	err := client.EchoSend("info", "Good morning!")
	// 	fmt.Println("err:", err)

	// 	// Valid Send (no error returned)
	// 	err = client.Send("error", "Good morning! No echo.")
	// 	fmt.Println("err:", err)

	// 	// Invalid EchoSend -- message level error
	// 	err = client.EchoSend("blah", "blah")
	// 	fmt.Println("err:", err)
}
