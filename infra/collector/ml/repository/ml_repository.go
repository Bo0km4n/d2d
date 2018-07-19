package repository

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/Bo0km4n/d2d/infra/collector/ml"
	"github.com/Bo0km4n/d2d/infra/collector/model"
)

type MlRepository interface {
	Predict(*model.AggregatedLog) (interface{}, error)
}

func NewMlRepository() MlRepository {
	return &mlRepository{}
}

type mlRepository struct {
}

func (m *mlRepository) Predict(item *model.AggregatedLog) (interface{}, error) {
	b, err := json.Marshal(item)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest(
		"POST",
		"http://localhost:5000/api/v1/predict",
		bytes.NewBuffer(b),
	)
	if err != nil {
		return nil, err
	}

	// Content-Type 設定
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 201 {
		return map[string]string{
			"Predict": "none",
		}, nil
	}

	result, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	preResult := &ml.PredictResult{}
	if err := json.Unmarshal(result, preResult); err != nil {
		return nil, err
	}
	return preResult, nil
}
