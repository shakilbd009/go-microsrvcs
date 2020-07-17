package github

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateRepoRequestAsJson(t *testing.T) {
	request := CreateRepoRequest{
		Name:        "golang intro",
		Description: "an intro to repo",
		Homepage:    "https://github.com",
		Private:     false,
		HasIssues:   false,
		HasProjects: true,
		HasWiki:     false,
	}
	bytes, err := json.Marshal(request)
	assert.Nil(t, err)
	assert.NotNil(t, bytes)
	assert.EqualValues(t, `{"name":"golang intro","description":"an intro to repo","homepage":"https://github.com","private":false,"has_issues":false,"has_projects":true,"has_wiki":false}`, string(bytes))
	//fmt.Println(string(bytes))
	var target CreateRepoRequest
	err = json.Unmarshal(bytes, &target)
	assert.Nil(t, err)
	assert.NotNil(t, target)
	assert.EqualValues(t, request.Name, target.Name)
	assert.EqualValues(t, request.HasIssues, target.HasIssues)
}
