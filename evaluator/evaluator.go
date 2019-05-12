
package evaluator

import (
	"interpreter/ast"
	"interpreter/object"
)

var (
	TRUE = &object.Boolean {Value: true}
	FALSE = &object.Boolean {Value: false}
	NULL = &object.Null{}
)

func Eval(node ast.Node) object.Object {
	switch node := node.(type) {

	case *ast.Program:
		return evalStatements(node.Statements)

	case *ast.ExpressionStatement:
		return Eval(node.Expression)

	case *ast.IntegerLiteral:
		return &object.Integer{ Value: node.Value }

	case *ast.Boolean:
		// Only 1 instance for true and false.
		// We are just returning a reference to them.
		return nativeBoolToBooleanObject(node.Value)
	}

	return nil
}

func nativeBoolToBooleanObject(input bool) *object.Boolean {
	if input { return TRUE }
	return FALSE
}

func evalStatements(stmts []ast.Statement) object.Object {
	var result object.Object
	for _, statement := range stmts {
		result = Eval(statement)
	}

	return result
}