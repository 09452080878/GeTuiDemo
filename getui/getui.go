package main

import (
	"GetuiDemo/getui/push"
	"GetuiDemo/getui/query"
	"GetuiDemo/getui/style"
	"GetuiDemo/getui/token"
	"GetuiDemo/getui/tool"
	"log"
	"time"
)

var (
	appId        string = "4hva40KrHd9aGZaIVjG9S2"
	appKey       string = "pobcc0EKOP8Eq6MRu7uV8A"
	appSecret    string = "A3tbszbNpJAHD3iObCQl4A"
	masterSecret string = "AOPHaK7F3yAqiNKe62hHd3"
	cid          string = "d40cd1f6052824d37d2811425e5504a2"
)

func main() {

	tokenStr, err := getGeTuiToken()
	if err != nil {
		log.Printf("get getui sign token err : %+v\n", err.Error())
		return
	}

	saveListBodyParmar := GetSaveListBodyParmar(appKey)
	saveRes, err := SaveListBody(appId, tokenStr, saveListBodyParmar)
	if err != nil {
		log.Printf("save list body  err : %+v\n", err.Error())
		return
	}

	parmar := GetPushListParmar(saveRes.TaskId, []string{cid})
	_, err = push.PushList(appId, tokenStr, parmar)
	if err != nil {
		log.Printf("save list body  err : %+v\n", err.Error())
		return
	}

	pushSingleResult, err := pushSingle(tokenStr)
	if err != nil {
		log.Printf("get push single err : %+v\n", err.Error())
		return
	}

	_, err = getPushResult(tokenStr, pushSingleResult.TaskId)
	if err != nil {
		log.Printf("query push result err : %+v\n", err)
		return
	}

}

func GetPushListParmar(taskId string, cids []string) *push.PushListParmar {

	pushListParmar := &push.PushListParmar{
		TaskId:     taskId,
		Cid:        cids,
		NeedDetail: true,
	}

	return pushListParmar
}

func SaveListBody(appId string, auth_token string, parmar *push.SaveListBodyParmar) (*push.SaveListBodyResult, error) {

	saveListBodyResult, err := push.SaveListBody(appId, auth_token, parmar)
	if err != nil {
		log.Printf("get push single err : %+v\n", err)
		return saveListBodyResult, err
	}
	log.Printf("saveListBodyResult: %+v\n", saveListBodyResult)
	return saveListBodyResult, err
}

func GetSaveListBodyParmar(appKey string) *push.SaveListBodyParmar {

	message := tool.GetMessage()
	message.SetAppKey(appKey)
	message.SetMsgType("notification")

	notification := tool.GetNotification()
	notification.SetTransmissionContent("透传内容")

	unWindStyle := style.GetUnwindStyle("检测到可疑人员", "警告通知")
	unWindStyle.SetBigStyle("1")
	unWindStyle.SetBigImageUrl("http://s0.hao123img.com/res/r/image/2016-04-14/2a3b604cdc47bdc4e2ffa252d31179d1.jpg")

	notification.SetNotifyStyle(unWindStyle)

	saveListBodyParmar := &push.SaveListBodyParmar{
		Message:      message,
		Notification: notification,
		TaskName:     time.Now().Format("20160102150405"),
	}
	log.Printf("saveListBodyParmar: %+v\n", saveListBodyParmar)
	return saveListBodyParmar
}

func getPushResult(auth_token string, taskId string) (*query.PushRESResult, error) {
	pushRESParmar := &query.PushRESParmar{
		TaskIdList: []string{taskId},
	}

	PushRESResult, err := query.PushResult(appId, auth_token, pushRESParmar)
	if err != nil {
		log.Printf("query push result err : %+v\n", err.Error())
		return PushRESResult, err
	}
	return PushRESResult, nil
}

//单推
func pushSingle(auth_token string) (*push.PushSingleResult, error) {

	message := tool.GetMessage()
	message.SetAppKey(appKey)
	message.SetMsgType("notification")

	notification := tool.GetNotification()
	notification.SetTransmissionContent("透传内容")

	unWindStyle := style.GetUnwindStyle("检测到可疑人员", "警告通知")
	unWindStyle.SetBigStyle("1")
	unWindStyle.SetBigImageUrl("http://s0.hao123img.com/res/r/image/2016-04-14/2a3b604cdc47bdc4e2ffa252d31179d1.jpg")

	notification.SetNotifyStyle(unWindStyle)

	pushSingleParmar := &push.PushSingleParmar{
		Message:      message,
		Notification: notification,
		Cid:          cid,
		RequestId:    time.Now().Format("20160102150405"),
	}

	log.Printf("pushSingleParmar: %+v\n", pushSingleParmar)

	pushSingleResult, err := push.PushSingle(appId, auth_token, pushSingleParmar)
	if err != nil {
		log.Printf("get push single err : %+v\n", err.Error())
		return pushSingleResult, err
	}

	log.Printf("push single:\n result:%+v\n status:%+v\n taskId:%+v\n", pushSingleResult.Result, pushSingleResult.Status, pushSingleResult.TaskId)

	return pushSingleResult, nil
}

func getGeTuiToken() (string, error) {
	tokenResult, err := token.GetAuthSign(appId, appKey, masterSecret)
	if err != nil {
		log.Printf("get getui sign token err : %+v\n", err.Error())
		return "", err
	}

	log.Printf("tokenResult: %+v\n", tokenResult)

	return tokenResult.AuthToken, nil
}
