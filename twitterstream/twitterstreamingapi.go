package main

import (
	"anaconda"
	"fmt"
	"net/url"
	"os"
)

//TwitterTrack -
func TwitterTrack(maxtweet string, track string, consumerKey string, consumerSecret string, accessToken string, accessTokenSecret string) []TwitterList {
	//Mais sobre: https://github.com/ChimeraCoder/anaconda

	anaconda.SetConsumerKey(consumerKey)                             //Consumer Key
	anaconda.SetConsumerSecret(consumerSecret)                       //Consumer Secret
	client := anaconda.NewTwitterApi(accessToken, accessTokenSecret) //Access Token, Access Token Secret

	// setando os parametros utilizando url.Values
	v := url.Values{}
	v.Set("count", maxtweet)                    // ou v.Set("locations", "<Locations>")
	result, err := client.GetSearch(track, nil) //buscar por tweets que contenham o termo definido na track
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	// Ao menos que exista algo estranho, devemos ter ao menos 2 tweets
	if len(result.Statuses) < 2 {
		fmt.Printf("Esperado 2 ou mais tweets, foram encontrados %d", len(result.Statuses))
		os.Exit(1)
	}

	twitterList := make([]TwitterList, len(result.Statuses))

	// verificar a existência de tweet vazio
	for i, tweet := range result.Statuses {
		twitterList[i].Tweet = tweet.Text
	}

	return twitterList
}

//TwitterList -
type TwitterList struct {
	Tweet string
}
