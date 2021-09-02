package mysql

import "time"

type Community struct {
	ID            int       `json:"id" gorm:"primaryKey;column:id;type:int(11) auto_increment"`
	CommunityID   uint      `json:"community_id" gorm:"column:community_id;type:int(10) unsigned;not null;uniqueIndex:idx_community_id"`
	CommunityName string    `json:"community_name" gorm:"column:community_name;type:varchar(128);not null;uniqueIndex:idx_community_name"`
	Introduction  string    `json:"introduction" gorm:"column:introduction;type:varchar(256);not null"`
	CreatedAt     time.Time `json:"created_at" gorm:"column:created_at;type:timestamp;not null;default:CURRENT_TIMESTAMP"`
	UpdatedAt     time.Time `json:"updated_at" gorm:"column:updated_at;type:timestamp;not null;default:CURRENT_TIMESTAMP"`
	DeletedAt     time.Time `json:"deleted_at" gorm:"column:deleted_at;type:timestamp"`
}

func GetCommunityList() (data []*Community, err error) {
	err = db.Find(&data).Error
	if err != nil {
		return nil, err
	}
	return data, nil
}
func GetCommunityById(id int64) (data Community, err error) {
	//data = new(Community)
	err = db.Where("id = ?", id).First(&data).Error
	if err != nil {
		return
	}
	return data, nil
}
