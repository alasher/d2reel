/*\
|*|  fetcher.go
|*|  Input: string of match ids corresponding to Dota 2 matches.
|*|  Output: A download of each file (if possible) into /replays.
|*|  After the files are downloaded, the parser is called for each of the successful downloads
\*/

package fetcher

import(
	"fmt"
	"log"
	"io/ioutil"
	"encoding/json"
	"os"
	
	"github.com/Philipp15b/go-steam"
	"github.com/Philipp15b/go-steam/internal/steamlang"
)

const KEYCHAIN_PATH string = "/Users/austin/Developer/Go/src/github.com/alasher/d2reel/keychain.json"

type Keychain struct {
	WebAPIKey string
	SteamUsername string
	SteamPassword string
}

func Fetch(ids []string) {
	fmt.Printf("Fetching ids for the following matches: %v\n", ids)
	
	// First, we have to get the secret config from keychain.json
	keyFile, readError := ioutil.ReadFile(KEYCHAIN_PATH)
	if readError != nil {
		fmt.Println("Couldn't find your keychain.json file, sorry dude.")
		os.Exit(1)
	}
	
	var keychain Keychain
	jsonErr := json.Unmarshal(keyFile, &keychain)
	if jsonErr != nil {
		fmt.Println("We found your keychain.json file, but had trouble processing it. Check your syntax, brother.")
		os.Exit(1)
	}
	
	steamLoop(keychain)
}

func steamLoop(keychain Keychain) {
	loginInfo := new(steam.LogOnDetails)
	loginInfo.Username = keychain.SteamUsername
	loginInfo.Password = keychain.SteamPassword
	
	client := steam.NewClient()
	client.Connect()
	
	for event := range client.Events() {
		doExit := false
		switch e := event.(type) {
		case *steam.ConnectedEvent:
			client.Auth.LogOn(loginInfo)
		case *steam.MachineAuthUpdateEvent:
			ioutil.WriteFile("sentry", e.Hash, 0666)
		case *steam.LoggedOnEvent:
			client.Social.SetPersonaState(steamlang.EPersonaState_Online)
			doExit = true
			onLoggedIn(client)
		case steam.FatalErrorEvent:
			log.Print(e)
		case error:
			log.Print(e)
		}
		
		if doExit {
			break
		}
	}
}

func onLoggedIn(client *steam.Client) {
	fmt.Println("Successful Steam login!")
	fmt.Printf("I have %d friends!\n", client.Social.Friends.Count())
}