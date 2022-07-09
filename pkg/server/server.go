package server

import (
	"context"
	"encoding/json"
	"github.com/shani1998/shani1998.github.io/pkg/pb"
	"io/ioutil"
)

const (
	templateDir = "github.com/shani1998/shani1998.github.io/templates"
)

func ListSkills(_ context.Context, _ *pb.SkillRequest) (*pb.SkillResponse, error) {
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
