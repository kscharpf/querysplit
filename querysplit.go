package querysplit

import (
        "fmt"
        "regexp"
)

type QuerySplitter struct {
  Base string
  splitterRegex *regexp.Regexp
  splitKey string
  otherRegex []*regexp.Regexp
  otherKeys []string
}

func NewQuerySplitter(base string, splitKey string, otherKeys []string) QuerySplitter{
  q := QuerySplitter{}

  q.Base = base
  q.splitterRegex = regexp.MustCompile(fmt.Sprintf(`%s=(\w+)`, splitKey))
  q.splitKey = splitKey
  q.otherRegex = make([] *regexp.Regexp, len(otherKeys))
  q.otherKeys = make([] string, len(otherKeys))
  for i := range otherKeys {
    q.otherKeys[i] = otherKeys[i]
    q.otherRegex[i] = regexp.MustCompile(fmt.Sprintf(`%s=(\w+)`, otherKeys[i]))
  }

  return q
}

func (q *QuerySplitter) Split(s string) ([]string, []string) {
  splitMatches :=  q.splitterRegex.FindAllStringSubmatch(s, -1)
  splits := make([]string, 1)
  for i := range splitMatches {
    splits = append(splits, splitMatches[i][1])
  }
  splits = splits[1:]

  responses := make([]string, len(q.otherKeys))
  for i := range q.otherKeys {
    t := q.otherRegex[i].FindStringSubmatch(s)
    responses[i] = t[1]
  }

  queries := make([] string, len(splits))
  for i := 0; i<len(queries); i++ {
    queries[i] = fmt.Sprintf("%s%s=%s", q.Base, q.splitKey, splits[i])

    for j := 0; j<len(responses); j++ {
      queries[i] = fmt.Sprintf("%s&%s=%s", queries[i], q.otherKeys[j], responses[j])
    }
  }
  return queries, splits
}




