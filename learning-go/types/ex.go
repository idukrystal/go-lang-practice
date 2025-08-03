package main

import (
	"fmt"
	"io"
	"slices"
)

type Team struct {
	Name string
	Players []string
}

type League struct {
	Teams []Team
	Wins map[string]int
}

func (l League) getTeam(name string) (Team, bool) {
	for _, team := range l.Teams {
		if team.Name == name {
			return team, true
		}
	}
	return Team{}, false
}

type writerAdapter  func(a... any) (int, error)

func (wa writerAdapter) Write(a []byte) (int, error) {
	return wa(string(a))
}

func (l League) MatchResult(teamA string, aScore int, teamB string, bScore int) {
	// just to confirm teams are in legue
	_, ok := l.getTeam(teamA)
	if !ok {
		fmt.Println("Team-A not in legue")
		return
	}

	_, ok = l.getTeam(teamB)
	if !ok {
		fmt.Println("Team-B not in legue")
		return
	}

	if aScore > bScore {
		l.Wins[teamA]++
	}
	if aScore < bScore {
		l.Wins[teamB]++
	}
	if aScore == bScore {
		//handle a draw
	}
}

func (l League) Ranking() []string {
	teamNames := make([]string, 0, len(l.Teams))
	for _, team := range l.Teams {
		teamNames = append(teamNames, team.Name)
	}

	slices.SortFunc(teamNames,
		func(nameA, nameB string) int {
			scoreA := l.Wins[nameA]
			scoreB := l.Wins[nameB]
			switch {
			case scoreA > scoreB:
				return -1
			case scoreA < scoreB:
				return 1
			default:
				return 0
			}
		},
	)
	return teamNames
}

type Ranker interface {
	Ranking() []string
}

func RankPrinter(ranker Ranker, writer io.Writer) {
	ranking := ranker.Ranking()
	for _, rank := range ranking {
		io.WriteString(writer, rank + "\n")
	}
}

func main() {
	league := League{
		[]Team {
			Team{"Animals", []string{"dog", "cat", "ram"}},
			Team{"Humans", []string{"black", "white", "brown"}},
			Team{"Plants", []string{"vegetable", "fruits", "roots"}},
		},
		map[string]int{},
	}
	fmt.Println(league.Ranking())
	league.MatchResult("Animals", 50, "Plants", 70)
	fmt.Println(league.Wins)
	writerFunc := writerAdapter(fmt.Print)
	RankPrinter(league, writerFunc)
}
