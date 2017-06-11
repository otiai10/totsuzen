package totsuzen

import (
	"strings"

	runewidth "github.com/mattn/go-runewidth"
)

// Token ...
type Token struct {
	Raw        string
	Text       string
	innerWidth int
}

const (
	roofEdge  = "＿"
	roofTile  = "人"
	wallLeft  = "＞"
	wallRight = "＜"
	floorEdge = "￣"
	floorTile = "Y"
	floorBond = "^"
	// to make
	// ＿人人人人人人人人人人人＿
	// ＞ something like this ＜
	// ￣Y^Y^Y^Y^Y^Y^Y^Y^Y^Y^Y￣
)

// NewToken ...
func NewToken(raw string) *Token {
	line := strings.Split(raw, "\n")[0]
	return &Token{
		Raw:        raw,
		innerWidth: runewidth.StringWidth(" " + line + " "),
	}
}

// Totsuzenize ...
func (t *Token) Totsuzenize() *Token {
	t.Text = strings.Join([]string{t.Head(), t.Body(), t.Foot()}, "\n")
	return t
}

// Head creates head part of this token.
func (t *Token) Head() string {
	head := ""
	for runewidth.StringWidth(head) < t.innerWidth {
		head += "人"
	}
	return "＿" + head + "＿"
}

// Body creates body part of this token.
func (t *Token) Body() string {
	lines := strings.Split(t.Raw, "\n")
	for i, line := range lines {
		lines[i] = "＞ " + line + " ＜"
	}
	return strings.Join(lines, "\n")
}

// Foot creates foot part of this token.
func (t *Token) Foot() string {
	foot := ""
	y := true
	for runewidth.StringWidth(foot) < t.innerWidth {
		if y {
			foot += "Y"
		} else {
			foot += "^"
		}
		y = !y
	}
	return "￣" + foot + "￣"
}
