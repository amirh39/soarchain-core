package testutil

import (
	"context"

	"github.com/amirh39/soarchain-core/x/dpr/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/golang/mock/gomock"
)

func (bank *MockBankKeeper) ExpectAny(context context.Context) {
	bank.EXPECT().SendCoinsFromAccountToModule(sdk.UnwrapSDKContext(context), gomock.Any(), gomock.Any(), gomock.Any()).AnyTimes()
	bank.EXPECT().SendCoinsFromModuleToAccount(sdk.UnwrapSDKContext(context), gomock.Any(), gomock.Any(), gomock.Any()).AnyTimes()
	// bank.EXPECT().GetBalance(sdk.UnwrapSDKContext(context), gomock.Any(), gomock.Any())
}

func coinsOf(amount uint64) sdk.Coins {
	return sdk.Coins{
		sdk.Coin{
			Denom:  sdk.DefaultBondDenom,
			Amount: sdk.NewInt(int64(amount)),
		},
	}
}

func (bank *MockBankKeeper) ExpectPay(context context.Context, who string, amount uint64) *gomock.Call {
	whoAddr, err := sdk.AccAddressFromBech32(who)
	if err != nil {
		sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, "[AccAddressFromBech32] failed. Empty address string is not allowed.")
		return nil
	}
	return bank.EXPECT().SendCoinsFromAccountToModule(sdk.UnwrapSDKContext(context), whoAddr, types.ModuleName, coinsOf(amount))
}

func (bank *MockBankKeeper) ExpectRefund(context context.Context, who string, amount uint64) *gomock.Call {
	whoAddr, err := sdk.AccAddressFromBech32(who)
	if err != nil {
		sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "[AccAddressFromBech32] failed. Empty address string is not allowed.")
		return nil
	}
	return bank.EXPECT().SendCoinsFromModuleToAccount(sdk.UnwrapSDKContext(context), types.ModuleName, whoAddr, coinsOf(amount))
}
