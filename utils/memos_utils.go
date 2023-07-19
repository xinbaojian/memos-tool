package utils

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"memos-tool/config"
	"net/http"
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
	req, _ := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, _ := client.Do(req)
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(resp.Body)

	body, _ := io.ReadAll(resp.Body)
	var memosBody MemosBody
	err := json.Unmarshal(body, &memosBody)
	if err != nil {
		fmt.Println("add fail:", err)
	}
	fmt.Println("add success, ID: ", memosBody.ID)
}
