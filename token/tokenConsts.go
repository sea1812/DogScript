/**--------------------------------------------------------
 * Author: jiang5630@outlook.com 2024-10-27
 * Description:
 * This file is part of the MyGoScript project.
 * --------------------------------------------------------
 * 作者：jiang5630@outlook.com  2024年10月27日
 * 描述：TTokenType的常量定义
 --------------------------------------------------------*/

package token

const (
	ILLEGAL = "ILLEGAL" //表示未知的词法单元
	EOF     = "EOF"     //表示文件结尾

	// 标识符和字面量
	IDENT = "IDENT"
	INT   = "INT"

	// 运算符
	ASSIGN   = "="
	PLUS     = "+"
	MINUS    = "-"
	BANG     = "!"
	ASTERISK = "*"
	SLASH    = "/"

	// 比较符
	LT     = "<"
	GT     = ">"
	EQ     = "=="
	NOT_EQ = "!="

	// 分隔符
	COMMA     = ","
	SEMICOLON = ";"
	LPAREN    = "("
	RPAREN    = ")"
	LBRACE    = "{"
	RBRACE    = "}"
	LBRACKET  = "["
	RBRACKET  = "]"

	// 关键字
	FUNCTION = "FUNCTION"
	LET      = "LET"
	TRUE     = "TRUE"
	FALSE    = "FALSE"
	IF       = "IF"
	ELSE     = "ELSE"
	RETURN   = "RETURN"
)

var keywords = map[string]TTokenType{
	"fn":     FUNCTION,
	"let":    LET,
	"true":   TRUE,
	"false":  FALSE,
	"if":     IF,
	"else":   ELSE,
	"return": RETURN,
}

// LookupIdent 查找标识符
func LookupIdent(ident string) TTokenType {
	if mtok, ok := keywords[ident]; ok {
		return mtok
	}
	return IDENT
}
