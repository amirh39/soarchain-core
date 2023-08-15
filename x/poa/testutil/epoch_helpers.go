package testutil

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/golang/mock/gomock"
)

func (epoch *MockEpochKeeper) ExpectAny(context context.Context) {

	epoch.EXPECT().GetEpochData(sdk.UnwrapSDKContext(context))
	epoch.EXPECT().SetEpochData(sdk.UnwrapSDKContext(context), gomock.Any())
}
