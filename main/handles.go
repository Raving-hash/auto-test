package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

// 获取所有子功能
func GetSubFeatures(w http.ResponseWriter, r *http.Request) {
	rows, err := db.Query("SELECT SubFeatureID, GroupID, Description, ExpectedOutput FROM SubFeatures")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var subFeatures []SubFeature
	for rows.Next() {
		var s SubFeature
		if err := rows.Scan(&s.ID, &s.GroupID, &s.Description, &s.ExpectedOutput); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		subFeatures = append(subFeatures, s)
	}

	json.NewEncoder(w).Encode(subFeatures)
}

// 创建子功能的简化示例
func CreateSubFeature(w http.ResponseWriter, r *http.Request) {
	// 接收和解析请求体略...
	var newSubFeature SubFeature
	// 使用json.NewDecoder().Decode()从请求体中解析JSON到newSubFeature
	if err := json.NewDecoder(r.Body).Decode(&newSubFeature); err != nil {
		// 如果解析出错，返回错误信息和HTTP状态码400（Bad Request）
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	// 确保请求体关闭（Close the request body）
	defer r.Body.Close()

	result, err := db.Exec("INSERT INTO SubFeatures (GroupID, Description, ExpectedOutput) VALUES (?, ?, ?)",
		newSubFeature.GroupID, newSubFeature.Description, newSubFeature.ExpectedOutput)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	id, err := result.LastInsertId()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	fmt.Fprintf(w, "SubFeature with ID %d created", id)
}

// UpdateSubFeature 更新子功能的HTTP处理函数
func UpdateSubFeature(w http.ResponseWriter, r *http.Request) {
	// 从URL路径中解析出SubFeatureID
	vars := mux.Vars(r)
	subFeatureID := vars["id"]

	// 创建一个SubFeature实例用于存储请求体中的数据
	var updatedSubFeature SubFeature

	// 从请求体中解析JSON到updatedSubFeature
	if err := json.NewDecoder(r.Body).Decode(&updatedSubFeature); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	// 使用SQL更新语句，根据SubFeatureID更新子功能信息
	// 注意：这里简化地只更新了Description和ExpectedOutput字段，实际根据需求可能需要更新更多字段
	stmt, err := db.Prepare("UPDATE SubFeatures SET Description = ?, ExpectedOutput = ? WHERE SubFeatureID = ?")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer stmt.Close()

	// 执行更新操作
	_, err = stmt.Exec(updatedSubFeature.Description, updatedSubFeature.ExpectedOutput, subFeatureID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// 返回更新成功的信息
	fmt.Fprintf(w, "SubFeature with ID %s updated successfully", subFeatureID)
}

// 删除子功能（示例简化处理）
func DeleteSubFeature(w http.ResponseWriter, r *http.Request) {
	// 实现略
}

// 执行子功能测试
func TestSubFeature(w http.ResponseWriter, r *http.Request) {
	// 实现略
}
