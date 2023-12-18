package keeper_test

import (
	"testing"

	keepertest "github.com/amirh39/soarchain-core/testutil/keeper"
	"github.com/amirh39/soarchain-core/testutil/nullify"
	"github.com/amirh39/soarchain-core/x/poa/types"

	"github.com/stretchr/testify/require"
)

func Test_VerifyGeneratedNumber(t *testing.T) {
	keeper, ctx := keepertest.PoaKeeper(t)

	for _, tc := range []struct {
		desc    string
		request *types.QueryVerifyRandomNumberRequest
		err     error
	}{
		{
			desc: "Verify Valid Generated Number",
			request: &types.QueryVerifyRandomNumberRequest{
				Pubkey:  GeneratedNumber_Pubkey,
				Message: GeberatedNumber_Message,
				Vrv:     GeberatedNumber_Vrv,
				Proof:   GeberatedNumber_Proof,
			},
		},
	} {
		t.Run(tc.desc, func(t *testing.T) {
			response, err := keeper.VerifyGeneratedNumber(ctx, tc.request)
			if tc.err != nil {
				require.ErrorIs(t, err, tc.err)
			} else {
				require.NoError(t, err)
				require.Equal(t,
					nullify.Fill(true),
					nullify.Fill(response),
				)
				t.Log("response", response)
			}
		})
	}
}

func Test_NotVerifyGeneratedNumber_InvalidPubkey(t *testing.T) {
	keeper, ctx := keepertest.PoaKeeper(t)

	for _, tc := range []struct {
		desc    string
		request *types.QueryVerifyRandomNumberRequest
		err     error
	}{
		{
			desc: "Not Verified Generated Number",
			request: &types.QueryVerifyRandomNumberRequest{
				Pubkey:  "",
				Message: GeberatedNumber_Message,
				Vrv:     GeberatedNumber_Vrv,
				Proof:   GeberatedNumber_Proof,
			},
		},
	} {
		t.Run(tc.desc, func(t *testing.T) {
			response, err := keeper.VerifyGeneratedNumber(ctx, tc.request)
			if tc.err != nil {
				require.ErrorIs(t, err, tc.err)
			} else {
				require.NoError(t, err)
				require.Equal(t,
					nullify.Fill(false),
					nullify.Fill(response),
				)
				t.Log("response", response)
			}
		})
	}
}

func Test_NotVerifyGeneratedNumber_InvalidMessage(t *testing.T) {
	keeper, ctx := keepertest.PoaKeeper(t)

	for _, tc := range []struct {
		desc    string
		request *types.QueryVerifyRandomNumberRequest
		err     error
	}{
		{
			desc: "Not Verified Generated Number",
			request: &types.QueryVerifyRandomNumberRequest{
				Pubkey:  GeneratedNumber_Pubkey,
				Message: "",
				Vrv:     GeberatedNumber_Vrv,
				Proof:   GeberatedNumber_Proof,
			},
		},
	} {
		t.Run(tc.desc, func(t *testing.T) {
			response, err := keeper.VerifyGeneratedNumber(ctx, tc.request)
			if tc.err != nil {
				require.ErrorIs(t, err, tc.err)
			} else {
				require.NoError(t, err)
				require.Equal(t,
					nullify.Fill(false),
					nullify.Fill(response),
				)
				t.Log("response", response)
			}
		})
	}
}

func Test_NotVerifyGeneratedNumber_InvalidVRV(t *testing.T) {
	keeper, ctx := keepertest.PoaKeeper(t)

	for _, tc := range []struct {
		desc    string
		request *types.QueryVerifyRandomNumberRequest
		err     error
	}{
		{
			desc: "Not Verified Generated Number",
			request: &types.QueryVerifyRandomNumberRequest{
				Pubkey:  GeneratedNumber_Pubkey,
				Message: GeberatedNumber_Message,
				Vrv:     "",
				Proof:   GeberatedNumber_Proof,
			},
		},
	} {
		t.Run(tc.desc, func(t *testing.T) {
			response, err := keeper.VerifyGeneratedNumber(ctx, tc.request)
			if tc.err != nil {
				require.ErrorIs(t, err, tc.err)
			} else {
				require.NoError(t, err)
				require.Equal(t,
					nullify.Fill(false),
					nullify.Fill(response),
				)
				t.Log("response", response)
			}
		})
	}
}

func Test_NotVerifyGeneratedNumber_InvalidProof(t *testing.T) {
	keeper, ctx := keepertest.PoaKeeper(t)

	for _, tc := range []struct {
		desc    string
		request *types.QueryVerifyRandomNumberRequest
		err     error
	}{
		{
			desc: "Not Verified Generated Number",
			request: &types.QueryVerifyRandomNumberRequest{
				Pubkey:  GeneratedNumber_Pubkey,
				Message: GeberatedNumber_Message,
				Vrv:     GeberatedNumber_Vrv,
				Proof:   "",
			},
		},
	} {
		t.Run(tc.desc, func(t *testing.T) {
			response, err := keeper.VerifyGeneratedNumber(ctx, tc.request)
			if tc.err != nil {
				require.ErrorIs(t, err, tc.err)
			} else {
				require.NoError(t, err)
				require.Equal(t,
					nullify.Fill(false),
					nullify.Fill(response),
				)
				t.Log("response", response)
			}
		})
	}
}
