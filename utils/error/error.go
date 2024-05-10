package error

import (
	"errors"
	msg "github.com/satryanararya/go-chefbot/constants/message"
)

var (
	// password
	ErrFailedHashingPassword = errors.New(msg.MsgFailedHashingPassword)
	ErrPasswordMismatch      = errors.New(msg.MsgPasswordMismatch)

	// token
	ErrFailedGeneratingToken = errors.New(msg.MsgFailedGeneratingToken)

	// external services
	ErrExternalService = errors.New(msg.MsgExternalServiceError)
	ErrItemNotFound    = errors.New(msg.MsgItemNotFound)
)