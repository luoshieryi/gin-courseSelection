package dao

import (
	"project/model"
)

// CreateMember 创建用户
func CreateMember(member model.Member) (int64, error) {

	err := model.DB.Create(&member).Error
	if err != nil {
		return 0, err
	}
	return member.ID, nil
}

// GetMemberByUsername 通过用户名查询用户
func GetMemberByUsername(username string) model.Member {
	member := model.Member{}

	model.DB.Find(&member, "username = ?", username)

	return member
}

// GetMemberByID 通过ID查询用户
func GetMemberByID(id int64) model.Member {
	member := model.Member{}
	model.DB.Find(&member, "id = ?", id)
	return member
}

func GetMemberList(offset int, limit int) ([]model.Member, error) {
	members := make([]model.Member, 0)
	err := model.DB.Limit(limit).Offset(offset).Find(&members).Error

	return members, err
}

// UpdateMember 更新用户信息
func UpdateMember(member model.Member) error {
	err := model.DB.Model(&member).Update(&member).Error

	return err
}
