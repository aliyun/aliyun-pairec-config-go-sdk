package recallengine

import (
	"encoding/json"
	"reflect"
	"strings"
	"testing"
)

func TestRecallQueriesJSON(t *testing.T) {
	in := RecallConf{Count: 10, VersionId: "items", Queries: [][]float32{{1, 2}, {3, 4}}}
	data, err := json.Marshal(in)
	if err != nil {
		t.Fatal(err)
	}
	var out RecallConf
	if err = json.Unmarshal(data, &out); err != nil {
		t.Fatal(err)
	}
	if !reflect.DeepEqual(in, out) {
		t.Fatalf("lost matrix: %s", data)
	}
	data, _ = json.Marshal(RecallConf{Trigger: "1,2", Count: 10})
	if strings.Contains(string(data), "queries") {
		t.Fatalf("legacy payload changed: %s", data)
	}
}
