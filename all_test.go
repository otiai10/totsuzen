package totsuzen

import (
	"testing"

	. "github.com/otiai10/mint"
)

func TestNewToken(t *testing.T) {
	Expect(t, NewToken("test")).TypeOf("*totsuzen.Token")
}

func TestToken_Totsuzenize(t *testing.T) {
	token := NewToken("test")
	Expect(t, token.Text).ToBe("")
	token.Totsuzenize()
	Expect(t, token.Text).ToBe(
		`＿人人人＿
＞ test ＜
￣Y^Y^Y^￣`)
}

func TestToken_withJapanese(t *testing.T) {
	text := NewToken("日本語もいけます").Totsuzenize().Text
	Expect(t, text).ToBe(
		`＿人人人人人人人人人＿
＞ 日本語もいけます ＜
￣Y^Y^Y^Y^Y^Y^Y^Y^Y^￣`)
}

func TestToken_withMultiLines(t *testing.T) {
	text := NewToken("test").Totsuzenize().Text
	Expect(t, NewToken(text).Totsuzenize().Text).ToBe(`＿人人人人人人＿
＞ ＿人人人＿ ＜
＞ ＞ test ＜ ＜
＞ ￣Y^Y^Y^￣ ＜
￣Y^Y^Y^Y^Y^Y^￣`)
}
