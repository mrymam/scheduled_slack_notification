package model

import "fmt"

type Metric struct {
	Name        string
	Description string
	Value       interface{}
}

func (m Metric) GetValueStr() (string, error) {

	switch t := m.Value.(type) {
	case string:
		return t, nil
	case int64, int32, int:
		return fmt.Sprintf("%d", t), nil
	case float64, float32:
		return fmt.Sprintf("%f", t), nil
	}

	return "", fmt.Errorf("cast error")
}
