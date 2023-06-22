package main

import (
	"log"
)

const (
	POS_PREFIX = 1
	POS_SUFFIX = 2
	POS_ANY    = 3
	PosUnknown = 0
)

type QueryCondition struct {
	KeywordStr string
	Pos        int
	ZhLen      int // 0 means all
}

func createQuery(keystr string, position, zhlen int) *QueryCondition {
	if len(keystr) == 0 {
		log.Printf("INFO: createQuery Keyword NULL!\n")
		return nil
	}
	qc := new(QueryCondition)
	qc.KeywordStr = keystr
	qc.Pos = position
	qc.ZhLen = zhlen

	return qc
}
