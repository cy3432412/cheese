package logic

import "cheese/dao/mysql"

func GetCommunityList() (data []*mysql.Community, err error) {
	return mysql.GetCommunityList()
}
func GetSingleCommunity(id int64) (data mysql.Community, err error) {
	return mysql.GetCommunityById(id)
}
