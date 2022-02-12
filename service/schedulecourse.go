package service

import "project/types"

func dfs(u string, vis map[string]bool, res map[string]string, match map[string]string, group map[string][]string) bool {
	for v := range group[u] {
		if vis[string(v)] == false {
			vis[string(v)] = true
			if match[string(v)] == "" || dfs(match[string(v)], vis, res, match, group) {
				match[string(v)] = u
				res[u] = string(v)
				return true
			}
		}
	}
	return false
}

func GetScheduleCourse(request types.ScheduleCourseRequest) (types.ErrNo, map[string]string) {
	var res map[string]string //返回的答案
	res = make(map[string]string)
	var match map[string]string
	match = make(map[string]string)
	for u := range request.TeacherCourseRelationShip {
		var vis map[string]bool //属于谁
		vis = make(map[string]bool)
		dfs(u, vis, res, match, request.TeacherCourseRelationShip)
	}
	return types.OK, res
}
