package getInfoNodeRed

import (
	"context"
)

type NodeRedRepository interface {
	GetNodeRedInfo(ctx context.Context) ([]NodeRedResponse, error)
}

type NodeRedResponse struct {
	Id int `json:"id"`
	V1 int `json:"v1"`
	V2 int `json:"v2"`
	V3 int `json:"v3"`
}

type ShowResponse struct {
	TurbinaUno int `json:"turbinaUno"`
	Separador  int `json:"separador"`
	Enfriador  int `json:"enfriador"`
}

type ShowRequest struct {
	DateStart string `json:"dateStart"`
	DateEnd   string `json:"dateEnd"`
}

type ShowProcessedInfo interface {
	ShowNodeProcessedInfo(ctx context.Context, startDate string, endDate string) ([]ShowResponse, error)
}
