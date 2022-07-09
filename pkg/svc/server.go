package svc

import (
	"context"
	"encoding/json"
	"io/ioutil"

	"github.com/shani1998/shani1998.github.io/pkg/pb"
)

const (
	templateDir = "/Users/spathak/projects/shani1998.github.io/internal/store"
)

type PortfolioServer struct {
	pb.PortfolioServer
}

// ListSkills returns a list of skills with corresponding proficiency percentage
func (ps *PortfolioServer) ListSkills(_ context.Context, _ *pb.SkillRequest) (*pb.SkillResponse, error) {
	file, err := ioutil.ReadFile(templateDir + "/skills.json")
	if err != nil {
		return nil, err
	}
	resp := &pb.SkillResponse{}
	err = json.Unmarshal(file, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
