package main

// 子功能的结构体
type SubFeature struct {
	ID             string `json:"id"`
	Description    string `json:"description"`
	ExpectedOutput string `json:"expectedOutput"`
	GroupID        int32  `json:"groupID"`
}

type SubFeatureGroup struct {
	GroupID          int    `json:"groupID"`
	GroupName        string `json:"groupName"`
	GroupDescription string `json:"groupDescription"`
}
