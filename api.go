package jmespath

import "strconv"

// JMESPath is the representation of a compiled JMES path query. A JMESPath is
// safe for concurrent use by multiple goroutines.
type JMESPath struct {
	ast  ASTNode
	intr *treeInterpreter
}

// Compile parses a JMESPath expression and returns, if successful, a JMESPath
// object that can be used to match against data.
func Compile(expression string) (*JMESPath, error) {
	intr := newInterpreter()

	parser := NewParser()
	ast, err := parser.Parse(expression, intr)
	if err != nil {
		return nil, err
	}
	jmespath := &JMESPath{ast: ast, intr: intr}
	return jmespath, nil
}

// CompileWithFunctions parses a JMESPath expression and returns, if successful, a JMESPath
// object that can be used to match against data.
func CompileWithFunctions(expression string, customFuncs []FunctionEntry) (*JMESPath, error) {
	intr := newInterpreterWithCustomFunctions(customFuncs)

	parser := NewParser()
	ast, err := parser.Parse(expression, intr)
	if err != nil {
		return nil, err
	}
	jmespath := &JMESPath{ast: ast, intr: intr}

	return jmespath, nil
}

// MustCompile is like Compile but panics if the expression cannot be parsed.
// It simplifies safe initialization of global variables holding compiled
// JMESPaths.
func MustCompile(expression string) *JMESPath {
	jmespath, err := Compile(expression)
	if err != nil {
		panic(`jmespath: Compile(` + strconv.Quote(expression) + `): ` + err.Error())
	}
	return jmespath
}

// Search evaluates a JMESPath expression against input data and returns the result.
func (jp *JMESPath) Search(data interface{}) (interface{}, error) {
	return jp.intr.Execute(jp.ast, data)
}

// Search evaluates a JMESPath expression against input data and returns the result.
func Search(expression string, data interface{}) (interface{}, error) {
	intr := newInterpreter()
	parser := NewParser()
	ast, err := parser.Parse(expression, intr)
	if err != nil {
		return nil, err
	}
	return intr.Execute(ast, data)
}
