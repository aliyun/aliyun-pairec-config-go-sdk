package model

import (
	"strconv"
	"sync"
)

// LayerParams offers Get* function to get value by the key
// If not found the key, defaultValue will return
type LayerParams interface {
	AddParam(key string, value interface{})

	AddParams(params map[string]interface{})

	Get(key string, defaultValue interface{}) interface{}

	GetString(key, defaultValue string) string

	GetInt(key string, defaultValue int) int

	GetFloat(key string, defaultValue float64) float64

	GetInt64(key string, defaultValue int64) int64

	ListParams() map[string]interface{}
}

type layerParams struct {
	mu         sync.RWMutex
	Parameters map[string]interface{}
}

func NewLayerParams() *layerParams {
	return &layerParams{
		Parameters: make(map[string]interface{}, 0),
	}
}

func (r *layerParams) AddParam(key string, value interface{}) {
	r.mu.Lock()
	r.Parameters[key] = value
	r.mu.Unlock()
}

func (r *layerParams) AddParams(params map[string]interface{}) {
	r.mu.Lock()
	for k, v := range params {
		r.Parameters[k] = v
	}
	r.mu.Unlock()
}

func (r *layerParams) Get(key string, defaultValue interface{}) interface{} {
	r.mu.RLock()
	defer r.mu.RUnlock()
	if val, ok := r.Parameters[key]; ok {
		return val
	}
	return defaultValue
}

func (r *layerParams) GetString(key, defaultValue string) string {
	r.mu.RLock()
	defer r.mu.RUnlock()
	val, ok := r.Parameters[key]
	if !ok {
		return defaultValue
	}

	switch value := val.(type) {
	case string:
		return value
	case int:
		return strconv.Itoa(value)
	case float64:
		return strconv.Itoa(int(value))
	case int32:
		return strconv.Itoa(int(value))
	case int64:
		return strconv.Itoa(int(value))
	}
	return defaultValue
}

func (r *layerParams) GetInt(key string, defaultValue int) int {
	r.mu.RLock()
	defer r.mu.RUnlock()
	val, ok := r.Parameters[key]
	if !ok {
		return defaultValue
	}
	switch value := val.(type) {
	case int:
		return value
	case float64:
		return int(value)
	case uint:
		return int(value)
	case int32:
		return int(value)
	case int64:
		return int(value)
	case string:
		if val, err := strconv.Atoi(value); err == nil {
			return val
		} else {
			return defaultValue
		}
	default:
		return defaultValue
	}
}
func (r *layerParams) GetFloat(key string, defaultValue float64) float64 {
	r.mu.RLock()
	defer r.mu.RUnlock()
	val, ok := r.Parameters[key]
	if !ok {
		return defaultValue
	}

	switch value := val.(type) {
	case float64:
		return value
	case int:
		return float64(value)
	case string:
		if f, err := strconv.ParseFloat(value, 64); err == nil {
			return f
		} else {
			return defaultValue
		}
	default:
		return defaultValue
	}
}
func (r *layerParams) GetInt64(key string, defaultValue int64) int64 {
	r.mu.RLock()
	defer r.mu.RUnlock()
	val, ok := r.Parameters[key]
	if !ok {
		return defaultValue
	}

	switch value := val.(type) {
	case int:
		return int64(value)
	case float64:
		return int64(value)
	case uint:
		return int64(value)
	case int32:
		return int64(value)
	case int64:
		return value
	case string:
		if val, err := strconv.ParseInt(value, 10, 64); err == nil {
			return val
		} else {
			return defaultValue
		}
	default:
		return defaultValue
	}
}

func (r *layerParams) ListParams() map[string]interface{} {
	r.mu.RLock()
	defer r.mu.RUnlock()
	p := make(map[string]interface{}, len(r.Parameters))

	for k, v := range r.Parameters {
		p[k] = v
	}
	return p
}

func MergeLayerParams(layersParamsMap map[string]LayerParams) LayerParams {
	mergedParams := NewLayerParams()
	for _, unmergedParams := range layersParamsMap {
		switch v := unmergedParams.(type) {
		case *layerParams:
			mergedParams.AddParams(v.ListParams())
		}
	}
	return LayerParams(mergedParams)
}
