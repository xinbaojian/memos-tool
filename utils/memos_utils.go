package utils

import (
	"encoding/json"
	"fmt"
	"memos-tool/config"
	"memos-tool/enums"
)

type Memos struct {
	Content    string `json:"content"`
	Visibility string `json:"visibility"`
}

type MemosBody struct {
	ID           int    `json:"id"`
	RowStatus    string `json:"rowStatus"`
	CreatorID    int    `json:"creatorId"`
	CreatedTs    int    `json:"createdTs"`
	UpdatedTs    int    `json:"updatedTs"`
	DisplayTs    int    `json:"displayTs"`
	Content      string `json:"content"`
	Visibility   string `json:"visibility"`
	Pinned       bool   `json:"pinned"`
	CreatorName  string `json:"creatorName"`
	ResourceList []any  `json:"resourceList"`
	RelationList []any  `json:"relationList"`
}

// SendMemos 发送Memos消息
func SendMemos(memos *Memos) {
	var baseUrl = fmt.Sprintf("%s/api/v1/memo", config.Get("baseurl"))
	jsonStr, _ := json.Marshal(memos)
	openId := config.GetOpenId()
	url := fmt.Sprintf("%s?openId=%s", baseUrl, openId)
	body, _ := Post(url, jsonStr)
	var memosBody MemosBody
	err := json.Unmarshal(body, &memosBody)
	if err != nil {
		fmt.Println("add fail:", err)
	}
	fmt.Println("add success, ID: ", memosBody.ID)
}

// GetMemosList 获取Memos列表
func GetMemosList(status enums.RowStatus, limit int) []MemosBody {
	var apiUrl = fmt.Sprintf("%s/api/v1/memo?openId=%s&rowStatus=%s&limit=%d", config.Get("baseurl"), config.GetOpenId(), status.String(), limit)
	body, _ := Get(apiUrl)
	var memosList []MemosBody
	err := json.Unmarshal(body, &memosList)
	if err != nil {
		panic(fmt.Errorf("json.Unmarshal failed: %v", err))
	}
	return memosList
}
