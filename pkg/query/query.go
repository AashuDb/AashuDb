package query

// Defines the SQL Query itself i.e. Type of Query, Table Name, Conditions,
// Insertion Values, Fields, etc.
type Query struct {
	Type       Query_t
	Table      string
	Conditions []Condition_t

	// Not Required Yet
	// Updates		map[string]string

	Insertions [][]string
	Fields     []string

	// Not Implemented Yet
	// Aliases		map[string]string
}

// Defines the type of SQL Query i.e. Select, Insert, etc.
type Query_t uint8

const (
	INVALID_QUERY_TYPE Query_t = iota
	SELECT
	INSERT
)

var Query_String = []string{
	"Invalid Query",
	"Select",
	"Insert",
}

type Condition_t struct {
	Operand1   string
	IsOp1Field bool

	Operator Operator_t

	Operand2   string
	IsOp2Field bool
}

// Defines the type of Operator i.e Equals(=), Not Equals(!=), etc.
type Operator_t uint8

const (
	INVALID_OPERATOR_TYPE Operator_t = iota
	EQUALS
	NOT_EQUALS
	GREATER_THAN
	LESS_THAN
	GREATER_THAN_EQUALS
	LESS_THAN_EQUALS
)

var Operator_String = []string{
	"Invalid Operator",
	"Equals",
	"Not Equals",
	"Greater Than",
	"Less Than",
	"Greater Than Equals",
	"Less Than Equals",
}
