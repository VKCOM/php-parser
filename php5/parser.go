package php5

import (
	"github.com/z7zmey/php-parser/node/expr"
	"io"

	"github.com/z7zmey/php-parser/comment"
	"github.com/z7zmey/php-parser/node"
	"github.com/z7zmey/php-parser/node/stmt"
	"github.com/z7zmey/php-parser/position"
	"github.com/z7zmey/php-parser/token"
)

var rootnode node.Node
var comments comment.Comments
var positions position.Positions
var positionBuilder position.Builder

func Parse(src io.Reader, fName string) (node.Node, comment.Comments, position.Positions) {
    yyDebug        = 0
    yyErrorVerbose = true
    rootnode = stmt.NewStmtList([]node.Node{}) //reset
    comments = comment.Comments{}
    positions = position.Positions{}
    positionBuilder = position.Builder{&positions}
    yyParse(newLexer(src, fName))
    return rootnode, comments, positions
}

func ListGetFirstNodeComments(list []node.Node) []comment.Comment {
	if len(list) == 0 {
		return nil
	}

	node := list[0]

	return comments[node]
}

type foreachVariable struct {
    node  node.Node
    byRef bool
}

type nodesWithEndToken struct {
    nodes []node.Node
    endToken token.Token
}

type boolWithToken struct {
    value bool
    token *token.Token
}

type simpleIndirectReference struct {
    all []*expr.Variable
    last *expr.Variable
}

const arrayFetchType = 0
const propertyFetchType = 1

type objectProperty struct {
    node node.Node
    fetchType int
}