package main

// 子功能的结构体
type SubFeature struct {
	ID             string `json:"id"`
	Description    string `json:"description"`
	ExpectedOutput string `json:"expectedOutput"`
	GroupID        int32  `json:"groupID"`
}

// 假设的数据库存储（请使用真实数据库替代）
var subFeatures []SubFeature = []SubFeature{}
