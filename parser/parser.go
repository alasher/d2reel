package parser

import (
	"fmt"
	"time"
	
	"github.com/dotabuff/yasha"
)

const DEMO_PATH string = "/Users/austin/Developer/Go/src/github.com/alasher/d2reel/replays"

func Parse(ids []string) {
	fmt.Printf("Match ids are: %v\n", ids)
	
	for _, id := range ids {
		
		path := DEMO_PATH + "/" + id + ".dem"
		fmt.Printf("parsing file %v\n", path)
		parser := yasha.ParserFromFile(DEMO_PATH + "/1458895412.dem")

		var now time.Duration
		var gameTime, preGameStarttime float64
		var killsTotal int = 0


		parser.OnEntityPreserved = func(pe *yasha.PacketEntity) {
			if pe.Name == "DT_DOTAGamerulesProxy" {
				gameTime = pe.Values["DT_DOTAGamerules.m_fGameTime"].(float64)
				preGameStarttime = pe.Values["DT_DOTAGamerules.m_flPreGameStartTime"].(float64)
				now = time.Duration(gameTime-preGameStarttime) * time.Second
			}
		}

		parser.OnCombatLog = func(entry yasha.CombatLogEntry) {
			switch log := entry.(type) {
			case *yasha.CombatLogDeath:
				if log.TargetIsHero {
					killsTotal++
				}
				if bothAreHeroes(log) {
					fmt.Printf("%7s | %s just killed %s!\n", now, log.Source, log.Target)
				}
			}
		}
		
		parser.Parse()
		fmt.Printf("Looks like %d kills total!\n", killsTotal)
	}
	
}

func bothAreHeroes(entry *yasha.CombatLogDeath) bool {
	return entry.AttackerIsHero && entry.TargetIsHero
}