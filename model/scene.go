package model

// ExperimentPath locates an experiment by its position in the scene hierarchy
type ExperimentPath struct {
	Room       *ExperimentRoom
	Layer      *Layer
	Group      *ExperimentGroup
	Experiment *Experiment
}

type Scene struct {
	SceneId int64 `json:"scene_id,omitempty"`
	SceneName string `json:"scene_name"`
	SceneInfo string `json:"scene_info"`
	ExperimentRooms []*ExperimentRoom `json:"experiment_rooms"`

	// experimentIndex maps the experiment id to its position in the scene,
	// it is built by BuildExperimentIndex when the scene data is assembled
	experimentIndex map[int64]*ExperimentPath
}

func (s *Scene) AddExperimentRoom(room *ExperimentRoom) {
	s.ExperimentRooms = append(s.ExperimentRooms, room)
}

// BuildExperimentIndex builds the experiment id index of the scene,
// it should be invoked after the scene data is assembled.
// A SceneMap assembled by hand must invoke it too, otherwise
// FindExperimentPath finds nothing.
func (s *Scene) BuildExperimentIndex() {
	index := make(map[int64]*ExperimentPath)
	for _, room := range s.ExperimentRooms {
		for _, layer := range room.Layers {
			for _, group := range layer.ExperimentGroups {
				for _, experiment := range group.Experiments {
					index[experiment.ExperimentId] = &ExperimentPath{
						Room:       room,
						Layer:      layer,
						Group:      group,
						Experiment: experiment,
					}
				}
			}
		}
	}
	s.experimentIndex = index
}

// FindExperimentPath returns the experiment path of the experiment id,
// nil is returned if the experiment is not found in the scene
func (s *Scene) FindExperimentPath(experimentId int64) *ExperimentPath {
	return s.experimentIndex[experimentId]
}
