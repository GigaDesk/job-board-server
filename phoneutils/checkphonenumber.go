package phoneutils

import (
	"errors"
	"log"

	"github.com/GigaDesk/eardrum-server/graph/model"
	"gorm.io/gorm"
)

// checks if phone number exists in both the unverified school table and the school table.
// Takes database instance and phone number as arguments
func CheckSchoolPhoneNumber(db *gorm.DB, phone string) (*model.PhoneNumberExists, error) {
	var verifiedcount int64
	var unverifiedcount int64

	if err := db.Model(&model.School{}).Where("phone_number = ?", phone).Count(&verifiedcount).Error; err != nil {
		log.Println(err)
		return nil, errors.New("error checking verified phone number existence")
	} 
	if err := db.Model(&model.UnverifiedSchool{}).Where("phone_number = ?", phone).Count(&unverifiedcount).Error; err != nil {
		log.Println(err)
		return nil, errors.New("error checking unverified phone number existence")
	} 
	if verifiedcount > 0 && unverifiedcount > 0{
		log.Println("phonenumber exists in both verified and unverified states")
	}
	
	phoneExists:= &model.PhoneNumberExists{
		Verified: verifiedcount > 0,
		Unverified: unverifiedcount > 0,
	}
		
	return phoneExists, nil
}