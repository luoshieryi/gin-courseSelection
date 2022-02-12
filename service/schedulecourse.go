package service

import (
	"project/types"
)

func dfs(u string, vis map[string]bool, res map[string]string, match map[string]string, group map[string][]string) bool {
	for _, v := range group[u] {
		if vis[v] == false {
			vis[v] = true
			if match[v] == "" || dfs(match[v], vis, res, match, group) {
				match[v] = u
				res[u] = v
				return true
			}
		}
	}
	return false
}

func GetScheduleCourse(request types.ScheduleCourseRequest) (types.ErrNo, map[string]string) {
	var res map[string]string //返回的答案
	res = make(map[string]string)
	var match map[string]string //课程对应的人
	match = make(map[string]string)
	for u := range request.TeacherCourseRelationShip {
		var vis map[string]bool
		vis = make(map[string]bool)
		dfs(u, vis, res, match, request.TeacherCourseRelationShip)
	}
	return types.OK, res
}
