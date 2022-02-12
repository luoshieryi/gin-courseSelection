package service

import (
	"project/dao"
	"project/model"
	"project/types"
	"strconv"
)

// CreateMember 创建用户
func CreateMember(request types.CreateMemberRequest) (string, types.ErrNo) {
	member := dao.GetMemberByUsername(request.Username)
	if member.ID != 0 {
		return strconv.FormatInt(member.ID, 10), types.UserHasExisted
	}

	entity := model.Member{
		Nickname: request.Nickname,
		Username: request.Username,
		Password: request.Password,
		UserType: int(request.UserType),
		Deleted:  false,
	}
	id, err := dao.CreateMember(entity)
	if err != nil {
		return "", types.UnknownError
	}
	return strconv.FormatInt(id, 10), types.OK
}

func GetMember(request types.GetMemberRequest) (types.TMember, types.ErrNo) {
	_id, err := strconv.ParseInt(request.UserID, 10, 64)
	if err != nil {
		return types.TMember{}, types.UnknownError
	}

	member := dao.GetMemberByID(_id)
	if member.ID == 0 {
		return types.TMember{}, types.UserNotExisted
	}

	res := types.TMember{
		UserID:   strconv.FormatInt(member.ID, 10),
		Nickname: member.Nickname,
		Username: member.Username,
		UserType: types.UserType(member.UserType),
	}
	if member.Deleted {
		return res, types.UserHasDeleted
	}
	return res, types.OK
}

func GetMemberList(request types.GetMemberListRequest) ([]types.TMember, types.ErrNo) {
	members, err := dao.GetMemberList(request.Offset, request.Limit)
	if err != nil {
		return []types.TMember{}, types.UnknownError
	}

	tMembers := make([]types.TMember, 0)
	for _, value := range members {
		tMembers = append(tMembers, types.TMember{
			UserID:   strconv.FormatInt(value.ID, 10),
			Nickname: value.Nickname,
			Username: value.Username,
			UserType: types.UserType(value.UserType),
		})
	}
	return tMembers, types.OK
}

func UpdateMember(request types.UpdateMemberRequest) types.ErrNo {
	_id, err := strconv.ParseInt(request.UserID, 10, 64)
	if err != nil {
		return types.UnknownError
	}

	member := dao.GetMemberByID(_id)
	if member.ID == 0 {
		return types.UserNotExisted
	}

	member.Nickname = request.Nickname

	err = dao.UpdateMember(member)
	if err != nil {
		return types.UnknownError
	}

	return types.OK
}

func DeleteMember(request types.DeleteMemberRequest) types.ErrNo {
	_id, err := strconv.ParseInt(request.UserID, 10, 64)
	if err != nil {
		return types.UnknownError
	}

	member := dao.GetMemberByID(_id)
	if member.ID == 0 {
		return types.UserNotExisted
	}

	member.Deleted = true

	err = dao.UpdateMember(member)
	if err != nil {
		return types.UnknownError
	}

	return types.OK
}
