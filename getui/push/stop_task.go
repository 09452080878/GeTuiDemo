package push

import (
	"encoding/json"
	util "github.com/09452080878/GetuiDemo/getui/util"
)

type StopTaskResult struct {
	Result string `json:"result"`
	TaskId string `json:"taskid"` //	任务标识号
}

func StopTask(appId string, auth_token string, taskid string) (*StopTaskResult, error) {

	url := util.TokenDomain + appId + "/stop_task/" + taskid

	result, err := util.Delete(url, auth_token, nil)
	if err != nil {
		return nil, err
	}

	stopTaskResult := new(StopTaskResult)
	if err := json.Unmarshal([]byte(result), &stopTaskResult); err != nil {
		return nil, err
	}

	return stopTaskResult, err
}
