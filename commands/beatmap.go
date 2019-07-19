package commands

import (
	"github.com/thehowl/go-osuapi"
	"strconv"
	"strings"

	"github.com/bwmarrin/discordgo"
)

// Beatmap command
func Beatmap(session *discordgo.Session, message *discordgo.MessageCreate, args []string) {
	link := args[0]

	if strings.HasPrefix(link, "https://osu.ppy.sh/beatmapsets/") || strings.HasPrefix(link, "https://osu.ppy.sh/b/") {
		var id int64
		var err error
		var beatmaps []osuapi.Beatmap

		isMapset := false

		ids := strings.Split(link, "/")[4:]

		// Specific diff linked
		if len(ids) > 1 {
			id, _ = strconv.ParseInt(strings.Split(ids[0], "#")[0], 10, 32)
		} else { // Just the mapset linked
			id, _ = strconv.ParseInt(ids[0], 10, 32)
			isMapset = true
		}

		if isMapset {
			beatmaps, err = osuClient.GetBeatmaps(osuapi.GetBeatmapsOpts{
				BeatmapSetID: int(id),
			})
		} else {
			beatmaps, err = osuClient.GetBeatmaps(osuapi.GetBeatmapsOpts{
				BeatmapID: int(id),
			})
		}

		if err != nil {
			session.ChannelMessageSend(message.ChannelID, "Something went wrong!\n`"+err.Error()+"`")
		}

		session.ChannelMessageSend(message.ChannelID, "if you see this message, it works\nto verify, here's the map's title: "+beatmaps[0].Title)
	} else {
		session.ChannelMessageSend(message.ChannelID, "That's not a valid map link!")
	}
}
