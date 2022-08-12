package game

import "github.com/rafaelkperes/wordle/internal/dict"

type State int

const (
	StatePlaying State = iota
	StateWin
	StateLoss
)

type Match int

const (
	MatchNone Match = iota
	MatchLetter
	MatchPosition
)

type Game struct {
	d dict.Dict

	board      []string
	matchBoard [][]Match
	curRow     int

	tries      int
	solution   string
	solLetters map[rune]struct{}
	state      State
}

func NewGame(d dict.Dict) *Game {
	return &Game{
		d: d,
	}
}

func (g *Game) Board() []string {
	return g.board
}

func (g *Game) Start(tries int) {
	g.tries = tries
	g.solution = g.d.RandomWord()
	g.solLetters = make(map[rune]struct{}, dict.WordLen)
	for _, letter := range g.solution {
		g.solLetters[letter] = struct{}{}
	}

	g.board = make([]string, tries)
	g.matchBoard = make([][]Match, tries)
	for i := 0; i < tries; i++ {
		g.matchBoard[i] = make([]Match, dict.WordLen)
	}

	g.curRow = 0
	g.state = StatePlaying
}

func (g *Game) Type(letter string) {
	if len(letter) > 1 {
		panic("unexpected typing length > 1")
	}

	curWord := g.board[g.curRow]
	if len(curWord) < dict.WordLen {
		g.board[g.curRow] = curWord + letter
	}
}

func (g *Game) Backspace() {
	curWord := g.board[g.curRow]
	if len(curWord) > 0 {
		g.board[g.curRow] = curWord[:len(curWord)-1]
	}
}

// Enter returns wheter the word was complete or not and it could move to next row.
// The state of the game should be checked afterwards, to check whether it has got the right word
// or reached the end.
func (g *Game) Enter() bool {
	curWord := g.board[g.curRow]
	if len(curWord) < dict.WordLen {
		return false
	}

	if curWord == g.solution {
		g.state = StateWin
		return true
	}
	if g.curRow == g.tries-1 {
		g.state = StateLoss
		return true
	}

	g.fillMatches()
	g.curRow++
	return true
}

func (g *Game) fillMatches() {
	curWord := g.board[g.curRow]
	for idx, letter := range curWord {
		m := MatchNone
		if letter == rune(g.solution[idx]) {
			m = MatchPosition
		} else if _, ok := g.solLetters[letter]; ok {
			m = MatchLetter
		}
		g.matchBoard[g.curRow][idx] = m
	}
}

func (g *Game) State() State {
	return StatePlaying
}
