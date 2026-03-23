package recallengine

import (
	"fmt"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func getTestClient() *Client {
	endpoint := os.Getenv("RECALL_ENGINE_SERVICE_ENDPOINT")
	username := os.Getenv("RECALL_ENGINE_SERVICE_USERNAME")
	password := os.Getenv("RECALL_ENGINE_SERVICE_PASSWORD")
	token := os.Getenv("RECALL_ENGINE_SERVICE_TOKEN")
	return NewClient(endpoint, username, password, WithRequestHeader("Authorization", token))
}

func TestRecall(t *testing.T) {
	client := getTestClient()
	if client.Endpoint == "" {
		t.Skip("RECALL_ENGINE_SERVICE_ENDPOINT not set, skipping test")
	}

	instanceId := os.Getenv("RECALL_ENGINE_INSTANCE_ID")
	t.Run("u2i recall with V2 version", func(t *testing.T) {
		request := &RecallRequest{
			InstanceId: instanceId,
			Service:    "recall_test",
			Version:    "V2",
			Uid:        "123",
			Recalls: map[string]RecallConf{
				"u2i_recall": {
					Trigger: "123",
					Count:   100,
				},
			},
		}

		resp, err := client.Recall(request)
		if err != nil {
			t.Fatalf("Failed to recall: %v", err)
		}

		record := resp.Result
		assert.Equal(t, 100, record.Size())
		//t.Log(record)

		for i := 0; i < record.Size(); i++ {
			assert.Equal(t, record.GetColumn("item_id").(*Column[string]).SafeGet(i), fmt.Sprintf("item_%d", record.Len()-i-1))
			assert.Equal(t, record.GetColumn("score").(*Column[float32]).SafeGet(i), float32(float64(record.Len()-i-1)*0.1))
			assert.Equal(t, record.GetColumn("recall_name").(*Column[string]).SafeGet(i), "u2i_recall")
			assert.Equal(t, record.GetColumn("tag").(*Column[int32]).SafeGet(i), int32(record.Len()-i-1))
			assert.Equal(t, record.GetColumn("category").(*Column[string]).SafeGet(i), fmt.Sprintf("category_%d", record.Len()-i-1))
			assert.Equal(t, record.GetColumn("list").(*Column[[]string]).SafeGet(i),
				[]string{fmt.Sprintf("a_%d", record.Len()-i-1), fmt.Sprintf("b_%d", record.Len()-i-1), fmt.Sprintf("c_%d", record.Len()-i-1), fmt.Sprintf("bool_%v", (record.Len()-i-1)%2)})
		}
	})
	t.Run("u2i recall with V2 version and use retain_fields", func(t *testing.T) {
		request := &RecallRequest{
			InstanceId: instanceId,
			Service:    "recall_test",
			Version:    "V2",
			Uid:        "123",
			Recalls: map[string]RecallConf{
				"u2i_recall": {
					Trigger: "123",
					Count:   100,
				},
			},
			RetainFields: []string{"tag"},
		}

		resp, err := client.Recall(request)
		if err != nil {
			t.Fatalf("Failed to recall: %v", err)
		}

		record := resp.Result
		assert.Equal(t, 100, record.Size())
		//t.Log(record)

		for i := 0; i < record.Size(); i++ {
			assert.Equal(t, record.GetColumn("item_id").(*Column[string]).SafeGet(i), fmt.Sprintf("item_%d", record.Len()-i-1))
			assert.Equal(t, record.GetColumn("score").(*Column[float32]).SafeGet(i), float32(float64(record.Len()-i-1)*0.1))
			assert.Equal(t, record.GetColumn("recall_name").(*Column[string]).SafeGet(i), "u2i_recall")
			assert.Equal(t, record.GetColumn("tag").(*Column[int32]).SafeGet(i), int32(record.Len()-i-1))
		}
		assert.Equal(t, record.GetColumn("category"), nil)
		assert.Equal(t, record.GetColumn("list"), nil)
	})
}
