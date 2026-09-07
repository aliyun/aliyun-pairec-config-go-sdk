package experiments

import (
	"fmt"
	"strings"
	"testing"

	"github.com/aliyun/aliyun-pairec-config-go-sdk/v2/common"
	"github.com/aliyun/aliyun-pairec-config-go-sdk/v2/model"
)

const forceTestSceneName = "feed_scene"

// baseline exp_id of the fixture: uid=1000, no force plan
const forceBaselineExpId = "ER3_L5#EG7#E88_L6#EG9#E92_GL1#EG2#E15"

// forceFixture holds the object references of the test scene data:
// business scene feed_scene: room 3 (base, layer 5 + layer 6) and room 4 (normal, layer 8),
// global scene: room 1 (layer 1)
type forceFixture struct {
	scenes map[string]*model.Scene

	room3      *model.ExperimentRoom
	room4      *model.ExperimentRoom
	globalRoom *model.ExperimentRoom

	layer5      *model.Layer
	layer6      *model.Layer
	layer8      *model.Layer
	globalLayer *model.Layer

	group7       *model.ExperimentGroup
	group9       *model.ExperimentGroup
	group10      *model.ExperimentGroup
	group8       *model.ExperimentGroup
	globalGroup2 *model.ExperimentGroup

	exp87, exp88, exp89, exp90 *model.Experiment
	exp91, exp92               *model.Experiment
	exp15, exp16               *model.Experiment
}

func buildForceFixture() *forceFixture {
	f := &forceFixture{}

	// global scene: room 1 / layer 1 / group 2 / experiments 15, 16
	// exp15 matches naturally (flow 100), exp16 never matches naturally (flow 0)
	f.exp15 = &model.Experiment{ExperimentId: 15, ExpGroupId: 2, LayerId: 1, ExpRoomId: 1, SceneId: 100, ExperimentName: "global_exp15",
		Type: common.Experiment_Type_Test, ExperimentFlow: 100, ExperimentConfig: `{"global_country":"gc0"}`}
	f.exp16 = &model.Experiment{ExperimentId: 16, ExpGroupId: 2, LayerId: 1, ExpRoomId: 1, SceneId: 100, ExperimentName: "global_exp16",
		Type: common.Experiment_Type_Test, ExperimentFlow: 0, ExperimentConfig: `{"global_country":"gc1"}`}
	f.globalGroup2 = &model.ExperimentGroup{ExpGroupId: 2, LayerId: 1, ExpRoomId: 1, SceneId: 100, ExpGroupName: "global_group",
		ExpGroupConfig: `{"global_group_param":"g2"}`}
	f.globalGroup2.AddExperiment(f.exp15)
	f.globalGroup2.AddExperiment(f.exp16)
	f.globalLayer = &model.Layer{LayerId: 1, ExpRoomId: 1, SceneId: 100, LayerName: "global_layer"}
	f.globalLayer.AddExperimentGroup(f.globalGroup2)
	f.globalRoom = &model.ExperimentRoom{ExpRoomId: 1, SceneId: 100, ExpRoomName: "global_room", Type: common.ExpRoom_Type_Base}
	f.globalRoom.AddLayer(f.globalLayer)
	globalScene := &model.Scene{SceneId: 100, SceneName: model.GlobalSceneName}
	globalScene.AddExperimentRoom(f.globalRoom)

	// room 3 (base) / layer 5 / group 7: default 87 + test 88 (flow 100) + test 89 (flow 0)
	f.exp87 = &model.Experiment{ExperimentId: 87, ExpGroupId: 7, LayerId: 5, ExpRoomId: 3, SceneId: 200, ExperimentName: "default_exp",
		Type: common.Experiment_Type_Default, ExperimentFlow: 100, ExperimentConfig: `{"default_param":"87"}`}
	f.exp88 = &model.Experiment{ExperimentId: 88, ExpGroupId: 7, LayerId: 5, ExpRoomId: 3, SceneId: 200, ExperimentName: "flow100_exp",
		Type: common.Experiment_Type_Test, ExperimentFlow: 100, ExperimentConfig: `{"forced_param":"88"}`}
	f.exp89 = &model.Experiment{ExperimentId: 89, ExpGroupId: 7, LayerId: 5, ExpRoomId: 3, SceneId: 200, ExperimentName: "flow0_exp",
		Type: common.Experiment_Type_Test, ExperimentFlow: 0, ExperimentConfig: `{"forced_param":"89"}`}
	f.group7 = &model.ExperimentGroup{ExpGroupId: 7, LayerId: 5, ExpRoomId: 3, SceneId: 200, ExpGroupName: "group7",
		ExpGroupConfig: `{"base_layer_param":"from_group7"}`}
	f.group7.AddExperiment(f.exp87)
	f.group7.AddExperiment(f.exp88)
	f.group7.AddExperiment(f.exp89)

	// room 3 / layer 5 / group 10: filter group matched only when the global param is forced to gc1
	f.exp90 = &model.Experiment{ExperimentId: 90, ExpGroupId: 10, LayerId: 5, ExpRoomId: 3, SceneId: 200, ExperimentName: "filter_exp",
		Type: common.Experiment_Type_Test, ExperimentFlow: 100, ExperimentConfig: `{"group10_exp_param":"90"}`}
	f.group10 = &model.ExperimentGroup{ExpGroupId: 10, LayerId: 5, ExpRoomId: 3, SceneId: 200, ExpGroupName: "group10",
		CrowdTargetType: common.CrowdTargetType_Filter, Filter: "global_country == 'gc1'",
		ExpGroupConfig: `{"group10_param":"g10"}`}
	f.group10.AddExperiment(f.exp90)
	f.layer5 = &model.Layer{LayerId: 5, ExpRoomId: 3, SceneId: 200, LayerName: "base_layer"}
	f.layer5.AddExperimentGroup(f.group7)
	f.layer5.AddExperimentGroup(f.group10)

	// room 3 / layer 6 / group 9
	f.exp92 = &model.Experiment{ExperimentId: 92, ExpGroupId: 9, LayerId: 6, ExpRoomId: 3, SceneId: 200, ExperimentName: "second_exp",
		Type: common.Experiment_Type_Test, ExperimentFlow: 100, ExperimentConfig: `{"second_param":"92"}`}
	f.group9 = &model.ExperimentGroup{ExpGroupId: 9, LayerId: 6, ExpRoomId: 3, SceneId: 200, ExpGroupName: "group9"}
	f.group9.AddExperiment(f.exp92)
	f.layer6 = &model.Layer{LayerId: 6, ExpRoomId: 3, SceneId: 200, LayerName: "second_layer"}
	f.layer6.AddExperimentGroup(f.group9)
	f.room3 = &model.ExperimentRoom{ExpRoomId: 3, SceneId: 200, ExpRoomName: "base_room", Type: common.ExpRoom_Type_Base}
	f.room3.AddLayer(f.layer5)
	f.room3.AddLayer(f.layer6)

	// room 4 (normal, no diversion bucket so never matched without force) / layer 8 / group 8
	f.exp91 = &model.Experiment{ExperimentId: 91, ExpGroupId: 8, LayerId: 8, ExpRoomId: 4, SceneId: 200, ExperimentName: "other_room_exp",
		Type: common.Experiment_Type_Test, ExperimentFlow: 100, ExperimentConfig: `{"test_room_param":"91"}`}
	f.group8 = &model.ExperimentGroup{ExpGroupId: 8, LayerId: 8, ExpRoomId: 4, SceneId: 200, ExpGroupName: "group8"}
	f.group8.AddExperiment(f.exp91)
	f.layer8 = &model.Layer{LayerId: 8, ExpRoomId: 4, SceneId: 200, LayerName: "test_layer"}
	f.layer8.AddExperimentGroup(f.group8)
	f.room4 = &model.ExperimentRoom{ExpRoomId: 4, SceneId: 200, ExpRoomName: "test_room", Type: common.ExpRoom_Type_Normal}
	f.room4.AddLayer(f.layer8)

	scene := &model.Scene{SceneId: 200, SceneName: forceTestSceneName}
	scene.AddExperimentRoom(f.room3)
	scene.AddExperimentRoom(f.room4)

	f.scenes = map[string]*model.Scene{
		forceTestSceneName:    scene,
		model.GlobalSceneName: globalScene,
	}

	for _, s := range f.scenes {
		for _, room := range s.ExperimentRooms {
			if err := room.Init(); err != nil {
				panic(err)
			}
			for _, layer := range room.Layers {
				for _, group := range layer.ExperimentGroups {
					if err := group.Init(); err != nil {
						panic(err)
					}
					for _, experiment := range group.Experiments {
						if err := experiment.Init(); err != nil {
							panic(err)
						}
					}
				}
			}
		}
		s.BuildExperimentIndex()
	}
	return f
}

func newForceTestClient(sceneMap map[string]*model.Scene) *ExperimentClient {
	return &ExperimentClient{SceneMap: sceneMap}
}

func newForceTestContext(plan *model.ForcePlan) *model.ExperimentContext {
	return &model.ExperimentContext{
		RequestId:           "requestId",
		Uid:                 "1000",
		FilterParams:        map[string]interface{}{"country": "cn"},
		ForceExperimentPlan: plan,
	}
}

// case 1: no force plan, the baseline behavior
func TestForceExperiment_NoPlan(t *testing.T) {
	client := newForceTestClient(buildForceFixture().scenes)
	result := client.MatchExperiment(forceTestSceneName, newForceTestContext(nil))

	if got := result.GetExpId(); got != forceBaselineExpId {
		t.Fatalf("baseline exp_id = %q, want %q", got, forceBaselineExpId)
	}
	params := result.GetExperimentParams()
	if got := params.GetString("forced_param", ""); got != "88" {
		t.Errorf("forced_param = %q, want 88", got)
	}
	if got := params.GetString("second_param", ""); got != "92" {
		t.Errorf("second_param = %q, want 92", got)
	}
	if got := params.GetString("global_country", ""); got != "gc0" {
		t.Errorf("global_country = %q, want gc0", got)
	}
	if got := params.GetString("base_layer_param", ""); got != "from_group7" {
		t.Errorf("base_layer_param = %q, want from_group7", got)
	}
	if got := params.GetString("default_param", ""); got != "" {
		t.Errorf("default_param = %q, want empty", got)
	}
}

// case 2: force one experiment of the current room
func TestForceExperiment_SingleExperiment(t *testing.T) {
	client := newForceTestClient(buildForceFixture().scenes)

	plan, err := client.ResolveExperimentIds(forceTestSceneName, []int64{89})
	if err != nil {
		t.Fatal(err)
	}
	result := client.MatchExperiment(forceTestSceneName, newForceTestContext(plan))

	want := "ER3_L5#EG7#E89_L6#EG9#E92_GL1#EG2#E15"
	if got := result.GetExpId(); got != want {
		t.Fatalf("exp_id = %q, want %q", got, want)
	}
	params := result.GetExperimentParams()
	// the forced experiment takes effect
	if got := params.GetString("forced_param", ""); got != "89" {
		t.Errorf("forced_param = %q, want 89", got)
	}
	// the unforced layer keeps the baseline behavior
	if got := params.GetString("second_param", ""); got != "92" {
		t.Errorf("second_param = %q, want 92", got)
	}
	if err := plan.VerifyApplied(result); err != nil {
		t.Errorf("VerifyApplied = %v, want nil", err)
	}
}

// case 3: force an experiment of another room
func TestForceExperiment_OtherRoom(t *testing.T) {
	client := newForceTestClient(buildForceFixture().scenes)

	plan, err := client.ResolveExperimentIds(forceTestSceneName, []int64{91})
	if err != nil {
		t.Fatal(err)
	}
	result := client.MatchExperiment(forceTestSceneName, newForceTestContext(plan))

	want := "ER4_L8#EG8#E91_GL1#EG2#E15"
	if got := result.GetExpId(); got != want {
		t.Fatalf("exp_id = %q, want %q", got, want)
	}
	params := result.GetExperimentParams()
	if got := params.GetString("test_room_param", ""); got != "91" {
		t.Errorf("test_room_param = %q, want 91", got)
	}
	// params of the layers of the old room must not be merged
	for _, key := range []string{"base_layer_param", "forced_param", "second_param", "default_param"} {
		if got := params.GetString(key, ""); got != "" {
			t.Errorf("%s = %q, want empty", key, got)
		}
	}
	if err := plan.VerifyApplied(result); err != nil {
		t.Errorf("VerifyApplied = %v, want nil", err)
	}
}

// case 4: force an experiment whose group conditions do not match the uid
func TestForceExperiment_GroupConditionNotMatch(t *testing.T) {
	client := newForceTestClient(buildForceFixture().scenes)

	// group10 matches only when global_country == gc1, the natural value is gc0
	plan, err := client.ResolveExperimentIds(forceTestSceneName, []int64{90})
	if err != nil {
		t.Fatal(err)
	}
	result := client.MatchExperiment(forceTestSceneName, newForceTestContext(plan))

	want := "ER3_L5#EG10#E90_L6#EG9#E92_GL1#EG2#E15"
	if got := result.GetExpId(); got != want {
		t.Fatalf("exp_id = %q, want %q", got, want)
	}
	params := result.GetExperimentParams()
	if got := params.GetString("group10_param", ""); got != "g10" {
		t.Errorf("group10_param = %q, want g10", got)
	}
	if err := plan.VerifyApplied(result); err != nil {
		t.Errorf("VerifyApplied = %v, want nil", err)
	}
}

// case 5: force a global scene experiment, the business group filter follows the forced global param
func TestForceExperiment_GlobalScene(t *testing.T) {
	client := newForceTestClient(buildForceFixture().scenes)

	plan, err := client.ResolveExperimentIds(forceTestSceneName, []int64{16})
	if err != nil {
		t.Fatal(err)
	}
	if plan.ScopeFor(forceTestSceneName) != nil {
		t.Fatal("plan should not contain the business scope")
	}
	result := client.MatchExperiment(forceTestSceneName, newForceTestContext(plan))

	// global_country becomes gc1, so the business layer 5 matches group10 naturally
	want := "ER3_L5#EG10#E90_L6#EG9#E92_GL1#EG2#E16"
	if got := result.GetExpId(); got != want {
		t.Fatalf("exp_id = %q, want %q", got, want)
	}
	params := result.GetExperimentParams()
	if got := params.GetString("global_country", ""); got != "gc1" {
		t.Errorf("global_country = %q, want gc1", got)
	}
	if got := params.GetString("group10_param", ""); got != "g10" {
		t.Errorf("group10_param = %q, want g10", got)
	}
	if err := plan.VerifyApplied(result); err != nil {
		t.Errorf("VerifyApplied = %v, want nil", err)
	}
}

// case 6: force a business scene experiment and a global scene experiment together
func TestForceExperiment_BusinessAndGlobal(t *testing.T) {
	client := newForceTestClient(buildForceFixture().scenes)

	plan, err := client.ResolveExperimentIds(forceTestSceneName, []int64{88, 16})
	if err != nil {
		t.Fatal(err)
	}
	result := client.MatchExperiment(forceTestSceneName, newForceTestContext(plan))

	want := "ER3_L5#EG7#E88_L6#EG9#E92_GL1#EG2#E16"
	if got := result.GetExpId(); got != want {
		t.Fatalf("exp_id = %q, want %q", got, want)
	}
	if err := plan.VerifyApplied(result); err != nil {
		t.Errorf("VerifyApplied = %v, want nil", err)
	}
}

// case 7: force multiple experiments of different layers in the same room
func TestForceExperiment_MultipleLayers(t *testing.T) {
	client := newForceTestClient(buildForceFixture().scenes)

	plan, err := client.ResolveExperimentIds(forceTestSceneName, []int64{88, 92})
	if err != nil {
		t.Fatal(err)
	}
	result := client.MatchExperiment(forceTestSceneName, newForceTestContext(plan))

	if got := result.GetExpId(); got != forceBaselineExpId {
		t.Fatalf("exp_id = %q, want %q", got, forceBaselineExpId)
	}
	params := result.GetExperimentParams()
	if got := params.GetString("forced_param", ""); got != "88" {
		t.Errorf("forced_param = %q, want 88", got)
	}
	if got := params.GetString("second_param", ""); got != "92" {
		t.Errorf("second_param = %q, want 92", got)
	}
	if err := plan.VerifyApplied(result); err != nil {
		t.Errorf("VerifyApplied = %v, want nil", err)
	}
}

// case 8: force a default type experiment, no E<id> segment in exp_id
func TestForceExperiment_DefaultTypeExperiment(t *testing.T) {
	client := newForceTestClient(buildForceFixture().scenes)

	plan, err := client.ResolveExperimentIds(forceTestSceneName, []int64{87})
	if err != nil {
		t.Fatal(err)
	}
	result := client.MatchExperiment(forceTestSceneName, newForceTestContext(plan))

	want := "ER3_L5#EG7_L6#EG9#E92_GL1#EG2#E15"
	if got := result.GetExpId(); got != want {
		t.Fatalf("exp_id = %q, want %q", got, want)
	}
	params := result.GetExperimentParams()
	if got := params.GetString("default_param", ""); got != "87" {
		t.Errorf("default_param = %q, want 87", got)
	}
	if err := plan.VerifyApplied(result); err != nil {
		t.Errorf("VerifyApplied = %v, want nil", err)
	}
}

// case 9: the plan keeps working after SceneMap is reloaded
func TestForceExperiment_SceneMapReplaced(t *testing.T) {
	client := newForceTestClient(buildForceFixture().scenes)

	plan, err := client.ResolveExperimentIds(forceTestSceneName, []int64{89})
	if err != nil {
		t.Fatal(err)
	}
	want := "ER3_L5#EG7#E89_L6#EG9#E92_GL1#EG2#E15"

	// (a) new snapshot without experiment 89
	fresh := buildForceFixture()
	fresh.group7.Experiments = fresh.group7.Experiments[:2] // drop exp89
	fresh.scenes[forceTestSceneName].BuildExperimentIndex()
	client.SceneMap = fresh.scenes
	result := client.MatchExperiment(forceTestSceneName, newForceTestContext(plan))
	if got := result.GetExpId(); got != want {
		t.Fatalf("(a) exp_id = %q, want %q", got, want)
	}

	// (b) new snapshot without the business scene
	client.SceneMap = map[string]*model.Scene{model.GlobalSceneName: fresh.scenes[model.GlobalSceneName]}
	result = client.MatchExperiment(forceTestSceneName, newForceTestContext(plan))
	if got := result.GetExpId(); got != want {
		t.Fatalf("(b) exp_id = %q, want %q", got, want)
	}

	// (c) empty SceneMap
	client.SceneMap = map[string]*model.Scene{}
	result = client.MatchExperiment(forceTestSceneName, newForceTestContext(plan))
	wantC := "ER3_L5#EG7#E89_L6#EG9#E92"
	if got := result.GetExpId(); got != wantC {
		t.Fatalf("(c) exp_id = %q, want %q", got, wantC)
	}
}

// case 10: ResolveExperimentIds errors
func TestResolveExperimentIds_Errors(t *testing.T) {
	client := newForceTestClient(buildForceFixture().scenes)

	cases := []struct {
		name    string
		scene   string
		ids     []int64
		wantErr string
	}{
		{"R1 invalid id", forceTestSceneName, []int64{88, -1},
			"experiment_ids contains invalid experiment id: -1"},
		{"R2 unavailable id", forceTestSceneName, []int64{999},
			"experiment_ids contains unavailable experiment id: 999"},
		{"R3 different rooms", forceTestSceneName, []int64{88, 91},
			"experiment_ids conflict in scene feed_scene: experiment 88 belongs to exp_room 3, but experiment 91 belongs to exp_room 4"},
		{"R4 same group", forceTestSceneName, []int64{88, 89},
			"experiment_ids conflict in scene feed_scene: experiment 88 and experiment 89 belong to the same experiment group 7 (layer base_layer)"},
		{"R5 same layer different groups", forceTestSceneName, []int64{88, 90},
			"experiment_ids conflict in scene feed_scene: experiment 88 (exp_group 7) and experiment 90 (exp_group 10) belong to the same layer base_layer"},
		{"R6 conflict in global scene", forceTestSceneName, []int64{15, 16},
			"experiment_ids conflict in global scene: experiment 15 and experiment 16 belong to the same experiment group 2 (layer global_layer)"},
		{"R7 scene not found", "no_such_scene", []int64{88},
			"experiment_ids not supported: abtest scene no_such_scene not found"},
	}
	for _, c := range cases {
		plan, err := client.ResolveExperimentIds(c.scene, c.ids)
		if err == nil || err.Error() != c.wantErr {
			t.Errorf("%s: err = %v, want %q", c.name, err, c.wantErr)
		}
		if plan != nil {
			t.Errorf("%s: plan = %v, want nil", c.name, plan)
		}
	}

	// duplicated ids are deduplicated
	plan, err := client.ResolveExperimentIds(forceTestSceneName, []int64{88, 88, 88})
	if err != nil {
		t.Fatal(err)
	}
	if scope := plan.ScopeFor(forceTestSceneName); scope == nil || len(scope.ByLayer) != 1 {
		t.Errorf("duplicated ids: scope = %v, want one layer", scope)
	}

	// empty ids
	plan, err = client.ResolveExperimentIds(forceTestSceneName, nil)
	if plan != nil || err != nil {
		t.Errorf("nil ids: plan = %v, err = %v, want nil, nil", plan, err)
	}
	plan, err = client.ResolveExperimentIds(forceTestSceneName, []int64{})
	if plan != nil || err != nil {
		t.Errorf("empty ids: plan = %v, err = %v, want nil, nil", plan, err)
	}

	// global ids fall to R2 when the global scene is not loaded
	f := buildForceFixture()
	delete(f.scenes, model.GlobalSceneName)
	client = newForceTestClient(f.scenes)
	if _, err := client.ResolveExperimentIds(forceTestSceneName, []int64{15}); err == nil ||
		err.Error() != "experiment_ids contains unavailable experiment id: 15" {
		t.Errorf("no global scene: err = %v, want unavailable", err)
	}
}

// case 11: a handcrafted plan with a nil Room is treated as no force
func TestForceExperiment_NilRoomPlan(t *testing.T) {
	f := buildForceFixture()
	client := newForceTestClient(f.scenes)

	plan := model.NewForcePlan()
	plan.AddScope(forceTestSceneName, &model.ForceScope{
		ByLayer: map[int64]*model.ExperimentPath{
			5: {Room: f.room3, Layer: f.layer5, Group: f.group7, Experiment: f.exp88},
		},
	})
	result := client.MatchExperiment(forceTestSceneName, newForceTestContext(plan))

	if got := result.GetExpId(); got != forceBaselineExpId {
		t.Fatalf("exp_id = %q, want baseline %q", got, forceBaselineExpId)
	}
}

// case 12: BuildExperimentIndex covers all experiments and points to fresh objects after rebuild
func TestSceneBuildExperimentIndex(t *testing.T) {
	f := buildForceFixture()
	scene := f.scenes[forceTestSceneName]

	for _, room := range scene.ExperimentRooms {
		for _, layer := range room.Layers {
			for _, group := range layer.ExperimentGroups {
				for _, experiment := range group.Experiments {
					path := scene.FindExperimentPath(experiment.ExperimentId)
					if path == nil || path.Room != room || path.Layer != layer ||
						path.Group != group || path.Experiment != experiment {
						t.Errorf("experiment %d: path mismatch", experiment.ExperimentId)
					}
				}
			}
		}
	}
	if scene.FindExperimentPath(999) != nil {
		t.Error("FindExperimentPath(999) should be nil")
	}

	// a rebuilt scene points to the new objects
	freshScene := buildForceFixture().scenes[forceTestSceneName]
	if freshScene.FindExperimentPath(88).Experiment == scene.FindExperimentPath(88).Experiment {
		t.Error("rebuild should point to new experiment objects")
	}
}

// case 13: VerifyApplied detects mismatches
func TestForcePlanVerifyApplied(t *testing.T) {
	f := buildForceFixture()
	client := newForceTestClient(f.scenes)

	// nil receiver
	var nilPlan *model.ForcePlan
	if err := nilPlan.VerifyApplied(nil); err != nil {
		t.Errorf("nil plan should return nil, got %v", err)
	}

	// normal result: default type experiment, unforced layer and global scene force
	plan, err := client.ResolveExperimentIds(forceTestSceneName, []int64{87, 16})
	if err != nil {
		t.Fatal(err)
	}
	result := client.MatchExperiment(forceTestSceneName, newForceTestContext(plan))
	if err := plan.VerifyApplied(result); err != nil {
		t.Errorf("VerifyApplied = %v, want nil", err)
	}

	plan88, err := client.ResolveExperimentIds(forceTestSceneName, []int64{88})
	if err != nil {
		t.Fatal(err)
	}

	// the layer matched another experiment
	otherResult := model.NewExperimentResult(forceTestSceneName, nil)
	otherResult.AddMatchExperimentGroup("base_layer", f.group7)
	otherResult.AddMatchExperiment("base_layer", f.exp91)
	err = plan88.VerifyApplied(otherResult)
	wantErr := "force experiment not applied in scene feed_scene: experiment 88 (layer base_layer) expected, got 91"
	if err == nil || err.Error() != wantErr {
		t.Errorf("err = %v, want %q", err, wantErr)
	}

	// the layer matched no experiment
	noneResult := model.NewExperimentResult(forceTestSceneName, nil)
	noneResult.AddMatchExperimentGroup("base_layer", f.group7)
	err = plan88.VerifyApplied(noneResult)
	wantErr = "force experiment not applied in scene feed_scene: experiment 88 (layer base_layer) expected, got none"
	if err == nil || err.Error() != wantErr {
		t.Errorf("err = %v, want %q", err, wantErr)
	}

	// mismatch in the global scene domain
	plan16, err := client.ResolveExperimentIds(forceTestSceneName, []int64{16})
	if err != nil {
		t.Fatal(err)
	}
	globalResult := model.NewExperimentResult(model.GlobalSceneName, nil)
	globalResult.AddMatchExperimentGroup("global_layer", f.globalGroup2)
	globalResult.AddMatchExperiment("global_layer", f.exp15)
	bizResult := model.NewExperimentResult(forceTestSceneName, nil)
	bizResult.GlobalSceneExperimentResult = globalResult
	err = plan16.VerifyApplied(bizResult)
	wantErr = "force experiment not applied in scene pairec_abtest_global_scene: experiment 16 (layer global_layer) expected, got 15"
	if err == nil || err.Error() != wantErr {
		t.Errorf("err = %v, want %q", err, wantErr)
	}
}

// MatchExperiment self-checks the force plan and logs an error on mismatch (the post-check)
func TestForceExperiment_MatchExperimentSelfCheck(t *testing.T) {
	f := buildForceFixture()
	client := newForceTestClient(f.scenes)

	var logs []string
	client.Logger = LoggerFunc(func(format string, args ...interface{}) {
		logs = append(logs, fmt.Sprintf(format, args...))
	})

	// a handcrafted plan whose layer does not belong to the room of the scope
	plan := model.NewForcePlan()
	plan.AddScope(forceTestSceneName, &model.ForceScope{
		Room: f.room3,
		ByLayer: map[int64]*model.ExperimentPath{
			8: {Room: f.room4, Layer: f.layer8, Group: f.group8, Experiment: f.exp91},
		},
	})
	result := client.MatchExperiment(forceTestSceneName, newForceTestContext(plan))

	found := false
	for _, msg := range logs {
		if strings.Contains(msg, "force experiment not applied") {
			found = true
		}
	}
	if !found {
		t.Error("MatchExperiment should log the self-check error")
	}
	if result == nil {
		t.Fatal("result should not be nil")
	}
}
