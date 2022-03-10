package source

import "math/rand"

type hpcDev struct {
}

func (dev *hpcDev) AddUserWithGroup(userName string, groupName string) (map[string]interface{}, error) {
	return map[string]interface{}{
		"success": "true",
		"data": map[string]interface{}{
			"gname": groupName,
			"gid":   rand.Intn(1000),
			"uname": userName,
			"uid":   rand.Intn(1000) + 1000,
		},
	}, nil
}

func (dev *hpcDev) AddUserToGroup(userName string, groupName string, gid int) (map[string]interface{}, error) {
	return map[string]interface{}{
		"success": "true",
		"data": map[string]interface{}{
			"uname": userName,
			"uid":   rand.Intn(1000) + 1000,
		},
	}, nil
}

func newDev(options *Options) HpcSource {
	return &hpcDev{}
}
