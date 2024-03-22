package main

import "strconv"

type Player struct {
	Role  string `json:"role"`
	Score int    `json:"score"`
	Name  string `json:"name"`
	Wins  int    `json:"wins"`
}

func (p Player) String() string {
	return p.Role + ": " + p.Name + "(" + strconv.Itoa(p.Wins) + ")\n" + strconv.Itoa(p.Score) + "\n"
}

type Match struct {
	Player1 Player `json:"player1"`
	Player2 Player `json:"player2"`
}

func (m Match) String() string {
	return "----------------------------\n" +
		m.Player1.String() +
		m.Player2.String() +
		"----------------------------\n"
}
func (m Match) Csv() string {
	return "\"" + m.Player1.Name + "\",\"" + m.Player1.Role + "\",\"" + strconv.Itoa(m.Player1.Wins) + "\"," + strconv.Itoa(m.Player1.Score) +
		",\"" + m.Player2.Name + "\",\"" + m.Player2.Role + "\",\"" + strconv.Itoa(m.Player2.Wins) + "\"," + strconv.Itoa(m.Player2.Score)

}

type GameRecord struct {
	MatchArr []*Match
}

func (g *GameRecord) AddMatch(m *Match) {
	g.MatchArr = append(g.MatchArr, m)
}

func (g *GameRecord) Csv() string {
	csvStr := "Player1,Role,Wins,Score,Player2,Role,Wins,Score\n"
	for _, m := range g.MatchArr {
		csvStr += m.Csv() + "\n"
	}
	return csvStr
}
