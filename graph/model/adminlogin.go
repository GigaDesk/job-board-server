package model

import "github.com/GigaDesk/eardrum-prefix/validate"

// validates admin login input data
func (n AdminLogin) Validate() error {
	//validate phone number in kenyan conditions
	if err := validate.ValidateKenyanPhoneNumber(n.PhoneNumber); err != nil {
		return err
	}
	//validate password
    if err := validate.ValidatePassword(n.Password); err != nil {
		return err
	}
	return nil

}