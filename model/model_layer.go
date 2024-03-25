package model

import (
	"crypto/md5"
	"fmt"
	"hash/fnv"
	"strconv"

	"github.com/aliyun/aliyun-pairec-config-go-sdk/v2/common"
)

type Layer struct {
	LayerId   int64  `json:"layer_id,omitempty"`
	ExpRoomId int64  `json:"exp_room_id"`
	SceneId   int64  `json:"scene_id"`
	LayerName string `json:"layer_name"`
	LayerInfo string `json:"layer_info"`

	ExperimentGroups []*ExperimentGroup `json:"experiment_groups"`
}

func (l *Layer) AddExperimentGroup(g *ExperimentGroup) {
	l.ExperimentGroups = append(l.ExperimentGroups, g)
}

func (l *Layer) FindMatchExperimentGroup(experimentContext *ExperimentContext) *ExperimentGroup {
	// first find debug users
	for _, group := range l.ExperimentGroups {
		if group.MatchDebugUsers(experimentContext) {
			return group
		}
	}

	// find filter or crowdid experiment group
	for _, group := range l.ExperimentGroups {
		if group.CrowdTargetType == "" {
			if group.Filter != "" || group.CrowdId > 0 {
				if group.Match(experimentContext) {
					return group
				}
			}

		} else if group.CrowdTargetType == common.CrowdTargetType_Filter || group.CrowdTargetType == common.CrowdTargetType_CrowdId {
			if group.Match(experimentContext) {
				return group
			}
		}
	}

	// find random experiment group
	for _, group := range l.ExperimentGroups {
		if group.CrowdTargetType == common.CrowdTargetType_Random {
			hashKey := fmt.Sprintf("%s_EXPROOM%d_LAYER%d", experimentContext.Uid, l.ExpRoomId, l.LayerId)
			hashValue := hashValue(hashKey)
			//e.logInfo("match experiment hash key:%s, value:%d", hashKey, hashValue)
			hashValueStr := strconv.FormatUint(hashValue, 10)
			if group.Match(&ExperimentContext{Uid: hashValueStr}) {
				return group
			}
		}
	}

	for _, group := range l.ExperimentGroups {
		if group.CrowdTargetType == "" {
			if group.Filter == "" && group.CrowdId == 0 {
				return group
			}

		} else if group.CrowdTargetType == common.CrowdTargetType_ALL {
			return group
		}
	}

	return nil
}
func hashValue(hashKey string) uint64 {
	md5 := md5.Sum([]byte(hashKey))
	hash := fnv.New64()
	hash.Write(md5[:])

	return hash.Sum64()
}
