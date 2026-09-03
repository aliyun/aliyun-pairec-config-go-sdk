package model

import (
	"errors"
	"fmt"
	"sort"
	"strconv"
)

// ForceScope is the force plan of one domain (business scene or global scene)
type ForceScope struct {
	// Room is the experiment room the forced experiments belong to
	Room *ExperimentRoom

	// ByLayer holds the forced experiment path of each layer, key is LayerId
	// to avoid the ambiguity of duplicated layer names across rooms
	ByLayer map[int64]*ExperimentPath
}

// ForcePlan is the complete force plan of one request,
// it holds at most two scopes: the business scene and the global scene
type ForcePlan struct {
	scopes map[string]*ForceScope // key: sceneName
}

// NewForcePlan returns an empty ForcePlan
func NewForcePlan() *ForcePlan {
	return &ForcePlan{scopes: make(map[string]*ForceScope)}
}

// AddScope adds the scope of the scene to the plan
func (p *ForcePlan) AddScope(sceneName string, scope *ForceScope) {
	p.scopes[sceneName] = scope
}

// ScopeFor returns the scope of the scene, nil is returned if not exists
func (p *ForcePlan) ScopeFor(sceneName string) *ForceScope {
	if p == nil {
		return nil
	}
	return p.scopes[sceneName]
}

// VerifyApplied asserts every forced item of the plan is reflected in the match result,
// it compares the experiment group and experiment pointers layer by layer in each domain,
// returns a descriptive error if any forced item is not applied, nil otherwise.
// nil plan returns nil.
func (p *ForcePlan) VerifyApplied(result *ExperimentResult) error {
	if p == nil {
		return nil
	}
	if result == nil {
		return errors.New("force experiment not applied: nil experiment result")
	}
	if scope, found := p.scopes[result.SceneName]; found {
		if err := scope.VerifyApplied(result); err != nil {
			return err
		}
	}
	if result.SceneName != GlobalSceneName {
		if scope, found := p.scopes[GlobalSceneName]; found {
			// a nil global result means the forced global experiment was not applied,
			// silent pass would violate the contract of this method
			if result.GlobalSceneExperimentResult == nil {
				return fmt.Errorf("force experiment not applied in scene %s: nil global scene experiment result", result.SceneName)
			}
			if err := scope.VerifyApplied(result.GlobalSceneExperimentResult); err != nil {
				return err
			}
		}
	}
	return nil
}

// VerifyApplied asserts the forced items of the scope are reflected in the match result
func (s *ForceScope) VerifyApplied(result *ExperimentResult) error {
	if s == nil || result == nil {
		return nil
	}

	layerIds := make([]int64, 0, len(s.ByLayer))
	for layerId := range s.ByLayer {
		layerIds = append(layerIds, layerId)
	}
	sort.Slice(layerIds, func(i, j int) bool { return layerIds[i] < layerIds[j] })

	for _, layerId := range layerIds {
		path := s.ByLayer[layerId]
		layerName := path.Layer.LayerName
		gotExperiment := result.layer2Experiment[layerName]
		if result.layer2ExperimentGroup[layerName] != path.Group || gotExperiment != path.Experiment {
			got := "none"
			if gotExperiment != nil {
				got = strconv.FormatInt(gotExperiment.ExperimentId, 10)
			}
			return fmt.Errorf("force experiment not applied in scene %s: experiment %d (layer %s) expected, got %s",
				result.SceneName, path.Experiment.ExperimentId, layerName, got)
		}
	}
	return nil
}
