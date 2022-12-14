package types

import (
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var (
	ErrUserHasExisted       = sdkerrors.Register(ModuleName, 101, "user has existed")
	ErrRegister             = sdkerrors.Register(ModuleName, 102, "user register faild")
	ErrUserNotFound         = sdkerrors.Register(ModuleName, 103, "user not found")
	ErrAddressFormat        = sdkerrors.Register(ModuleName, 104, "address format error")
	ErrTransfer             = sdkerrors.Register(ModuleName, 105, "transfer error")
	ErrBurn                 = sdkerrors.Register(ModuleName, 107, "burn error")
	ErrUserUpdate           = sdkerrors.Register(ModuleName, 108, "user info update error")
	ErrMortgageAmount       = sdkerrors.Register(ModuleName, 109, "mortgage amount error")
	ErrAddressBookSet       = sdkerrors.Register(ModuleName, 110, "address set error")
	ErrGeneratingMobile     = sdkerrors.Register(ModuleName, 111, "error generating mobile")
	ErrGetMobile            = sdkerrors.Register(ModuleName, 112, "error get mobile")
	ErrGateway              = sdkerrors.Register(ModuleName, 113, "error gateway address")
	ErrMobileSetError       = sdkerrors.Register(ModuleName, 114, "error mobile set")
	ErrDevideError          = sdkerrors.Register(ModuleName, 115, "error devide")
	ErrParamsSetError       = sdkerrors.Register(ModuleName, 116, "error module params")
	ErrGetLastReveiveHeight = sdkerrors.Register(ModuleName, 117, "error get last receive height")
	ErrSetLastReveiveHeight = sdkerrors.Register(ModuleName, 118, "error set last receive height")
	ErrGetBonus             = sdkerrors.Register(ModuleName, 119, "error get bonus")
	ErrSetMortgageLog       = sdkerrors.Register(ModuleName, 120, "error set mortgage log")
	ErrGetMortgageLog       = sdkerrors.Register(ModuleName, 121, "error get mortgage log")
	ErrUserNotHaveMobile    = sdkerrors.Register(ModuleName, 122, "user not have mobile")
	ErrUserMobileCount      = sdkerrors.Register(ModuleName, 123, "user mobile count error")
	ErrValidatorNotFound    = sdkerrors.Register(ModuleName, 124, "validator not found")
	ErrPledgeFeeTransfer    = sdkerrors.Register(ModuleName, 125, "pledge fee transfer error")
	ErrPledgeFeeBurn        = sdkerrors.Register(ModuleName, 126, "pledge fee burn error")
	ErrPledgeFeeGateway     = sdkerrors.Register(ModuleName, 127, "pledge fee gateway error")
	ErrPledgeFeeDpos        = sdkerrors.Register(ModuleName, 128, "pledge fee dpos error")
	ErrPledgeDelegate       = sdkerrors.Register(ModuleName, 129, "pledge delegate error")
	ErrGetPledgeTotal       = sdkerrors.Register(ModuleName, 130, "get pledge total error")
	ErrPledgeFeeSet         = sdkerrors.Register(ModuleName, 131, "pledge fee set error")
	ErrChatInfoSet          = sdkerrors.Register(ModuleName, 132, "chat info set error")
)
