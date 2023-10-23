package cli

import (
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

func TestTodoRequester_GetTodoList(t *testing.T) {
	var data = `{"contents":[{"id":506,"title":"hello","order_num":44,"created_at":"2023-10-23T09:00:16.62752Z","is_deleted":false},{"id":505,"title":"hello","order_num":44,"created_at":"2023-10-23T09:00:16.066971Z","is_deleted":false},{"id":504,"title":"hello","order_num":44,"created_at":"2023-10-23T08:56:49.629324Z","is_deleted":false},{"id":503,"title":"hello","order_num":44,"created_at":"2023-10-23T08:56:32.729941Z","is_deleted":false},{"id":493,"title":"angry","order_num":1,"created_at":"2023-10-20T04:23:30.67971Z","is_deleted":false}],"total_content":5,"current_page":0,"last_page":0}`
	result := &listResDto{}
	err := json.NewDecoder(strings.NewReader(data)).Decode(&result)
	assert.NoError(t, err)
	var d2 = `{"title":"hello","content2":"con?"}`
	testStruct := testSt{}
	err = json.Unmarshal([]byte(d2), &testStruct)
	assert.NoError(t, err)
}

type testSt struct {
	Title   string `json:"title"`
	Contetn string `json:"content2"`
}
