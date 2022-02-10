package resp

import "project/types"

func CreateMemberRes(code types.ErrNo, userId string) types.CreateMemberResponse {
	return types.CreateMemberResponse{Code: code, Data: struct{ UserID string }{UserID: userId}}
}

func GetMemberRes(code types.ErrNo, data types.TMember) types.GetMemberResponse {
	return types.GetMemberResponse{Code: code, Data: data}
}

func GetMemberListRes(code types.ErrNo, data []types.TMember) types.GetMemberListResponse {
	return types.GetMemberListResponse{Code: code, Data: struct{ MemberList []types.TMember }{MemberList: data}}
}

func UpdateMemberRes(code types.ErrNo) types.UpdateMemberResponse {
	return types.UpdateMemberResponse{Code: code}
}

func DeleteMemberRes(code types.ErrNo) types.DeleteMemberResponse {
	return types.DeleteMemberResponse{Code: code}
}

func LoginRes(code types.ErrNo, userId string) types.LoginResponse {
	return types.LoginResponse{Code: code, Data: struct{ UserID string }{UserID: userId}}
}

func LogoutRes(code types.ErrNo) types.LogoutResponse {
	return types.LogoutResponse{Code: code}
}

func WhoamiRes(code types.ErrNo, data types.TMember) types.WhoAmIResponse {
	return types.WhoAmIResponse{Code: code, Data: data}
}
