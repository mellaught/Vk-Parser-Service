package parser

import (
	"log"
	"testing"
)

var (
	token   = ""
	version = ""
	url     = ""
)

func TestGetProfiles(t *testing.T) {
	vk := VkParser{
		token:   token,
		version: version,
		url:     url,
	}

	groups, _, err := vk.GetUserGroups("365133336")
	if err != nil {
		t.Fatal(err)
	}

	resp, err := vk.GetMembers(groups)
	if err != nil {
		t.Fatal(err)
	} else {
		log.Println(resp)
	}
}
