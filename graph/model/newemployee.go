package model

import (
	"github.com/AlekSi/pointer"
	"github.com/GigaDesk/eardrum-prefix/validate"
)

// validates New employee input data
func (n NewEmployee) Validate() error {
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

	//validate badge
	if n.Profilepicture != nil {
		if err := validate.ValidateBadge(pointer.GetString(n.Profilepicture)); err != nil {
			return err
		}
	}
	return nil
}
