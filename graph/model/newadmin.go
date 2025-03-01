package model

import (
	"github.com/GigaDesk/eardrum-prefix/validate"
)

// validates NewAdmin input data
func (n NewAdmin) Validate() error {
	//validate name
	if err := validate.ValidateName(n.Name); err != nil {
		return err
	}

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
