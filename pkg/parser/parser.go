package parser

import (
	"fmt"

	"github.com/AashuDb/AashuDb/pkg/query"
)

func Parse(sql string) (query.Query, error) {
	q, err := parse(sql)

	return q, err
}

func parse(sql string) (query.Query, error) {
	//implement parser
	fmt.Println("Query:", sql)

	return query.Query{}, nil
}
