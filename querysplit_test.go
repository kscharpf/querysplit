package querysplit

import (
  "testing"
)

func TestQuery(t *testing.T) {
  otherkeys := make([] string, 2)
  otherkeys[0] = "startkey"
  otherkeys[1] = "endkey"
  q := NewQuerySplitter(`http://www.example.com/get?`, "splitkey", otherkeys)

  var queries []string
  queries, splits := q.Split(`http://www.example.com/get?splitkey=V1&splitkey=V2&splitkey=V3&splitkey=V4&startkey=444333&endkey=888777`)

  expQ0 := `http://www.example.com/get?splitkey=V1&startkey=444333&endkey=888777`
  expQ1 := `http://www.example.com/get?splitkey=V2&startkey=444333&endkey=888777`
  expQ2 := `http://www.example.com/get?splitkey=V3&startkey=444333&endkey=888777`
  expQ3 := `http://www.example.com/get?splitkey=V4&startkey=444333&endkey=888777`
  if queries[0] != expQ0 {
    t.Errorf("Received: %v Expected %v", queries[0], expQ0)
  } else if queries[1] != expQ1 {
    t.Errorf("Received: %v Expected %v", queries[1], expQ1)
  } else if queries[2] != expQ2 {
    t.Errorf("Received: %v Expected %v", queries[2], expQ2)
  } else if queries[3] != expQ3 {
    t.Errorf("Received: %v Expected %v", queries[3], expQ3)
  }

  if splits[0] != "V1" {
    t.Errorf("Split Received: %v Expected V1", splits[0])
  } else if splits[1] != "V2" {
    t.Errorf("Split Received: %v Expected V2", splits[1])
  } else if splits[2] != "V3" {
    t.Errorf("Split Received: %v Expected V3", splits[2])
  } else if splits[3] != "V4" {
    t.Errorf("Split Received: %v Expected V4", splits[3])
  }
}
