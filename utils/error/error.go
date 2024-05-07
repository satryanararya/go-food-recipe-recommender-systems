package error

import (
	"errors"
	msg "github.com/satryanararya/go-chefbot/constants/message"
)

var (
	// Password
	ErrFailedHashingPassword = errors.New(msg.MsgFailedHashingPassword)
	ErrPasswordMismatch      = errors.New(msg.MsgPasswordMismatch)

	// Token
	ErrFailedGeneratingToken = errors.New(msg.MsgFailedGeneratingToken)
)