package postgresql

import (
	"fmt"
)

type PostgresqlQuery struct {
	queries []interface{}
	args    []interface{}

	searchConditions string
}

func NewPostgresqlQuery(searchConditions string) *PostgresqlQuery {
	return &PostgresqlQuery{
		searchConditions: searchConditions,
	}
}

func (pq *PostgresqlQuery) Add(query string, arg interface{}) {
	pq.queries = append(pq.queries, query)
	pq.args = append(pq.args, arg)
}

func (pq *PostgresqlQuery) GetQueries() string {
	var result string
	for idx, query := range pq.queries {
		if len(pq.queries) > idx+1 {
			result += fmt.Sprintf(" %s %s", query, pq.searchConditions)
		} else {
			result += fmt.Sprintf(" %s", query)
		}
	}

	return result
}

func (pq *PostgresqlQuery) GetArgs() interface{} {
	return pq.args
}
