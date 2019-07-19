package commands

import (
	"fmt"
	osuapi "github.com/thehowl/go-osuapi"
	"io/ioutil"
	"osuBot/util"
)

var (
	osuToken  string
	osuClient osuapi.Client
)

func init() {
	token, err := ioutil.ReadFile("osu_api_token")
	util.CheckForErrors(err)

	osuToken = string(token)
	osuClient := osuapi.NewClient(osuToken)
	fmt.Println(osuClient)
}
