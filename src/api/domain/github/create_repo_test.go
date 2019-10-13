package github

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateRepoRequestAsJson(t *testing.T) {
	request := CreateRepoRequest{
		Name:        "Hello-word",
		Description: "This is our firts repo",
		Homepage:    "https://github.com",
		Private:     true,
		HasIssue:    true,
		HasProjects: true,
		HasWiki:     true,
	}

	//Marshal takes an input intarface and attemps to create a valid json string
	bytes, err := json.Marshal(request)
	assert.Nil(t, err)
	assert.NotNil(t, bytes)
	assert.EqualValues(t, `{"name":"Hello-word","description":"This is our firts repo","homepage":"https://github.com","private":true,"has_issue":true,"has_projects":true,"has_wiki":true}`, string(bytes))

	var target CreateRepoRequest

	//Unmarshal takes a input byte array and a *pointer* that we're trying to fill using this json
	err = json.Unmarshal(bytes, &target)
	assert.Nil(t, err)

}
