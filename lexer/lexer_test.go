/**--------------------------------------------------------
 * Author: jiang5630@outlook.com 2024-10-27
 * Description:
 * This file is part of the MyGoScript project.
 * --------------------------------------------------------
 * 作者：jiang5630@outlook.com  2024年10月27日
 * 描述：
 --------------------------------------------------------*/

package lexer

import (
	"MyGoScript/token"
	"testing"
)

func TestNextToken(t *testing.T) {
	input := "=+(){},;"
	tests := []struct {
		expectedType    token.TTokenType
		expectedLiteral string
	}{
		{token.ASSIGN, "="},
		{token.PLUS, "+"},
		{token.LPAREN, "("},
		{token.RPAREN, ")"},
		{token.LBRACE, "{"},
		{token.RBRACE, "}"},
		{token.SEMICOLON, ";"},
		{token.COMMA, ","},
		{token.RBRACE, "}"},
		{token.SEMICOLON, ";"},
		{token.EOF, ""},
	}
	l := New(input)

}
