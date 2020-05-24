package model

type (

	/**
	用户表 结构体
	 */
	User struct {
		Id uint64  `json:"id" gorm:"primary_key"`
		Name string `json:"name"`
		Address string `json:"address"`
	}

)
