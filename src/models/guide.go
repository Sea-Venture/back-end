package models

import "gorm.io/gorm"

type Guide struct {
	gorm.Model
	GuideID          string `json:"guide_id"`
	F_Name 		 string `json:"f_name"`
	L_Name 		 string `json:"l_name"`
	Licence_Front string `json:"licence_front"`
	Licence_Back  string `json:"licence_back"`
	Area		string `json:"area"`
	NIC_Photo string `json:"nic_photo"`
	Phone_Number string `json:"phone_number"`
	Beach_ID 		uint `json:"beach_id"`
}