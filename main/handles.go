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

// DeleteSubFeature 删除子功能的HTTP处理函数
func DeleteSubFeature(w http.ResponseWriter, r *http.Request) {
	// 使用mux.Vars来从请求URL中获取子功能的ID
	vars := mux.Vars(r)
	subFeatureID := vars["id"]

	// 使用SQL删除语句，根据SubFeatureID来删除特定的子功能
	stmt, err := db.Prepare("DELETE FROM SubFeatures WHERE SubFeatureID = ?")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer stmt.Close()

	// 执行删除操作
	result, err := stmt.Exec(subFeatureID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// 检查是否有行被删除
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if rowsAffected == 0 {
		http.Error(w, fmt.Sprintf("No SubFeature found with ID %s", subFeatureID), http.StatusNotFound)
		return
	}

	// 返回删除成功的信息
	fmt.Fprintf(w, "SubFeature with ID %s deleted successfully", subFeatureID)
}

// 执行子功能测试
func TestSubFeature(w http.ResponseWriter, r *http.Request) {
	// 实现略 todo
}

func CreateSubFeatureGroup(w http.ResponseWriter, r *http.Request) {
	var newGroup SubFeatureGroup
	if err := json.NewDecoder(r.Body).Decode(&newGroup); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	stmt, err := db.Prepare("INSERT INTO SubFeatureGroups (GroupName, GroupDescription) VALUES (?, ?)")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer stmt.Close()

	result, err := stmt.Exec(newGroup.GroupName, newGroup.GroupDescription)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	groupID, err := result.LastInsertId()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	fmt.Fprintf(w, "SubFeatureGroup with ID %d created successfully", groupID)
}

func DeleteSubFeatureGroup(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	groupID := vars["groupId"]

	stmt, err := db.Prepare("DELETE FROM SubFeatureGroups WHERE GroupID = ?")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer stmt.Close()

	result, err := stmt.Exec(groupID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if rowsAffected, err := result.RowsAffected(); err == nil && rowsAffected == 0 {
		http.Error(w, fmt.Sprintf("No SubFeatureGroup found with ID %s", groupID), http.StatusNotFound)
		return
	}

	fmt.Fprintf(w, "SubFeatureGroup with ID %s deleted successfully", groupID)
}

func UpdateSubFeatureGroup(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	groupID := vars["groupId"]

	var updatedGroup SubFeatureGroup
	if err := json.NewDecoder(r.Body).Decode(&updatedGroup); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	stmt, err := db.Prepare("UPDATE SubFeatureGroups SET GroupName = ?, GroupDescription = ? WHERE GroupID = ?")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer stmt.Close()

	_, err = stmt.Exec(updatedGroup.GroupName, updatedGroup.GroupDescription, groupID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	fmt.Fprintf(w, "SubFeatureGroup with ID %s updated successfully", groupID)
}

func GetSubFeatureGroups(w http.ResponseWriter, r *http.Request) {
	rows, err := db.Query("SELECT GroupID, GroupName, GroupDescription FROM SubFeatureGroups")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var groups []SubFeatureGroup
	for rows.Next() {
		var group SubFeatureGroup
		if err := rows.Scan(&group.GroupID, &group.GroupName, &group.GroupDescription); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		groups = append(groups, group)
	}

	json.NewEncoder(w).Encode(groups)
}
