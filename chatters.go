package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type Chatters struct {
	Chatters struct {
		Broadcaster []string `json:"broadcaster"`
		Vips        []string `json:"vips"`
		Moderators  []string `json:"moderators"`
		Staff       []string `json:"staff"`
		Admins      []string `json:"admins"`
		Global_mods []string `json:"global_mods"`
		Viewers     []string `json:"viewers"`
	} `json:"chatters"`
}

func main() {
	fmt.Println("Enter channel name:")
	var channel string
	fmt.Scanln(&channel)

	url := "http://tmi.twitch.tv/group/user/" + channel + "/chatters"
	resp, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	var chatters Chatters
	err = json.Unmarshal(body, &chatters)
	if err != nil {
		log.Fatal(err)
	}

	totalChatters := len(chatters.Chatters.Broadcaster) + len(chatters.Chatters.Vips) + len(chatters.Chatters.Moderators) + len(chatters.Chatters.Staff) + len(chatters.Chatters.Admins) + len(chatters.Chatters.Global_mods) + len(chatters.Chatters.Viewers)
	fmt.Println("Total Chatters:", totalChatters)

	totalVips := len(chatters.Chatters.Vips)
	fmt.Println("Total Vips:", totalVips)

	totalMods := len(chatters.Chatters.Moderators)
	fmt.Println("Total Mods:", totalMods)

	totalStaff := len(chatters.Chatters.Staff)
	fmt.Println("Total Staff:", totalStaff)

	totalAdmins := len(chatters.Chatters.Admins)
	fmt.Println("Total Admins:", totalAdmins)

	totalGlobalMods := len(chatters.Chatters.Global_mods)
	fmt.Println("Total Global Mods:", totalGlobalMods)

	totalViewers := len(chatters.Chatters.Viewers)
	fmt.Println("Total Viewers:", totalViewers)

	// if output is over x lines, it will be saved to a file and not printed to the console
	if len(chatters.Chatters.Broadcaster)+len(chatters.Chatters.Vips)+len(chatters.Chatters.Moderators)+len(chatters.Chatters.Staff)+len(chatters.Chatters.Admins)+len(chatters.Chatters.Global_mods)+len(chatters.Chatters.Viewers) > 0 {
		file, err := json.MarshalIndent(chatters, "", " ")
		if err != nil {
			log.Fatal(err)
		}
		err = ioutil.WriteFile("chatters.json", file, 0644)
		if err != nil {
			log.Fatal(err)
		}

		// fmt.Println("Output is over 0 lines, it has been saved to chatters.json")
		return
	}
	fmt.Println("Broadcaster:")
	for _, broadcaster := range chatters.Chatters.Broadcaster {
		fmt.Println(broadcaster)
	}
	fmt.Println("Vips:")
	for _, vip := range chatters.Chatters.Vips {
		fmt.Println(vip)
	}
	fmt.Println("Moderators:")
	for _, mod := range chatters.Chatters.Moderators {
		fmt.Println(mod)
	}
	fmt.Println("Staff:")
	for _, staff := range chatters.Chatters.Staff {
		fmt.Println(staff)
	}
	fmt.Println("Admins:")
	for _, admin := range chatters.Chatters.Admins {
		fmt.Println(admin)
	}
	fmt.Println("Global Mods:")
	for _, global_mod := range chatters.Chatters.Global_mods {
		fmt.Println(global_mod)
	}
	fmt.Println("Viewers:")
	for _, viewer := range chatters.Chatters.Viewers {
		fmt.Println(viewer)
	}
}
