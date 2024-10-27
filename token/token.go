/**--------------------------------------------------------
 * Author: jiang5630@outlook.com 2024-10-27
 * Description:
 * This file is part of the MyGoScript project.
 * --------------------------------------------------------
 * 作者：jiang5630@outlook.com  2024年10月27日
 * 描述：定义词法单元的Token数据结构
 --------------------------------------------------------*/

package token

// TTokenType 词法单元类型
type TTokenType string

// TToken 词法单元
type TToken struct {
	Type    TTokenType
	Literal string
}
