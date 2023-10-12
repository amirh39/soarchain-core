package keeper_test

import (
	"fmt"
	keepertest "soarchain/testutil/keeper"
	"soarchain/x/dpr/types"
	"strconv"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func Test_GetDpr(t *testing.T) {
	keeper, ctx := keepertest.DprKeeper(t)
	wctx := sdk.WrapSDKContext(ctx)
	msgs := SetupSecondDpr(2)
	keeper.SetDpr(ctx, msgs[0])
	keeper.SetDpr(ctx, msgs[1])

	for _, tc := range []struct {
		desc     string
		request  *types.QueryGetDprRequest
		response *types.QueryGetDprResponse
		err      error
	}{
		{
			desc: "First",
			request: &types.QueryGetDprRequest{
				Id: msgs[0].Id,
			},
			response: &types.QueryGetDprResponse{Dpr: msgs[0]},
		},
		{
			desc: "Second",
			request: &types.QueryGetDprRequest{
				Id: msgs[1].Id,
			},
			response: &types.QueryGetDprResponse{Dpr: msgs[1]},
		},
		{
			desc: "KeyNotFound",
			request: &types.QueryGetDprRequest{
				Id: strconv.Itoa(100000),
			},
			err: status.Error(codes.NotFound, "not found"),
		},
		{
			desc: "InvalidRequest",
			err:  status.Error(codes.InvalidArgument, "invalid request"),
		},
	} {
		t.Run(tc.desc, func(t *testing.T) {
			response, err := keeper.Dpr(wctx, tc.request)
			if err != nil {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
				require.NotNil(t, response)
			}
		})
	}
}

func Test_GetAllDpr(t *testing.T) {
	keeper, ctx := keepertest.DprKeeper(t)
	msgs := SetupSecondDpr(2)
	keeper.SetDpr(ctx, msgs[0])
	keeper.SetDpr(ctx, msgs[1])

	allDprs := keeper.GetAllDpr(ctx)
	require.NotNil(t, allDprs)
}

func Test_DPRsByClientPubkey(t *testing.T) {
	keeper, ctx := keepertest.DprKeeper(t)
	wctx := sdk.WrapSDKContext(ctx)

	dpr1 := types.Dpr{
		Id:            "dpr1",
		ClientPubkeys: []string{"pubkey1", "pubkey2"},
	}
	dpr2 := types.Dpr{
		Id:            "dpr2",
		ClientPubkeys: []string{"pubkey3", "pubkey4"},
	}
	dpr3 := types.Dpr{
		Id:            "dpr3",
		ClientPubkeys: []string{"pubkey5", "pubkey1"},
	}
	keeper.SetDpr(ctx, dpr1)
	keeper.SetDpr(ctx, dpr2)
	keeper.SetDpr(ctx, dpr3)

	response, err := keeper.DPRsByClientPubkey(wctx, &types.QueryDPRsByClientPubkeyRequest{
		ClientPubkey: "pubkey1",
	})
	require.NoError(t, err)

	require.NotEmpty(t, response.Dpr)

	fmt.Println(response.Dpr)

	for _, dpr := range response.Dpr {
		require.Contains(t, dpr.ClientPubkeys, "pubkey1")
	}
}
