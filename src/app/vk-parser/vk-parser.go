package parser

import (
	"VkParser/src/app/models"

	"github.com/valyala/fasthttp"
)

// Struct for Vk Parser Service.
type VkParser struct {
	token   string // Private token for vk api
	version string // Vk api version
	url     string // vk api url
}

// Creates New VkParser structs
func NewVkParser(conf *models.Config) *VkParser {
	return &VkParser{
		token:   conf.Token,
		version: conf.Version,
		url:     conf.URL,
	}
}

// Get func for all vk-api methods.
// Returns responce from Vk.com if success else error.
func (vk *VkParser) GET(method string, params string) ([]byte, error) {
	req := vk.url + method + "?access_token=" + vk.token + "&v=" + vk.version + params
	//fmt.Println(req)
	status, body, err := fasthttp.Get(nil, req)
	if status != 200 || err != nil {
		return nil, err
	}

	return body, nil
}
