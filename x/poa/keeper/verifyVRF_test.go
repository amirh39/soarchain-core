package keeper_test

import (
	"testing"

	keepertest "soarchain/testutil/keeper"
	"soarchain/testutil/nullify"
	"soarchain/x/poa/types"

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
				Pubkey:  "66eea999dcfb6fa4df8a5d2b22ea5e637d65ff9525e5f58f5e27bdac457c0450",
				Message: "soar19r5gmm7nqxy2v0pzm3c8ldkzax7ugqy5jwrv2y,1",
				Vrv:     "4afbf8af915f46626cadcff67ea7eee354fb6b8a3645de238126355fe524cd8c",
				Proof:   "e8cd528e10b85b629bd836b2f71a964cd4c2734f8136093d41e677b3c98fbb0e72f2f53371f6f4b068c3d05370d383f4b6e2ca59b5b71a745c7207c3dc754a0d58bd4cbbc630906c70c214cfdcbedfbd649627da37d8e53ce8cc14168b3e792b",
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
				Message: "soar19r5gmm7nqxy2v0pzm3c8ldkzax7ugqy5jwrv2y,1",
				Vrv:     "4afbf8af915f46626cadcff67ea7eee354fb6b8a3645de238126355fe524cd8c",
				Proof:   "e8cd528e10b85b629bd836b2f71a964cd4c2734f8136093d41e677b3c98fbb0e72f2f53371f6f4b068c3d05370d383f4b6e2ca59b5b71a745c7207c3dc754a0d58bd4cbbc630906c70c214cfdcbedfbd649627da37d8e53ce8cc14168b3e792b",
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
				Pubkey:  "66eea999dcfb6fa4df8a5d2b22ea5e637d65ff9525e5f58f5e27bdac457c0450",
				Message: "",
				Vrv:     "4afbf8af915f46626cadcff67ea7eee354fb6b8a3645de238126355fe524cd8c",
				Proof:   "e8cd528e10b85b629bd836b2f71a964cd4c2734f8136093d41e677b3c98fbb0e72f2f53371f6f4b068c3d05370d383f4b6e2ca59b5b71a745c7207c3dc754a0d58bd4cbbc630906c70c214cfdcbedfbd649627da37d8e53ce8cc14168b3e792b",
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
				Pubkey:  "66eea999dcfb6fa4df8a5d2b22ea5e637d65ff9525e5f58f5e27bdac457c0450",
				Message: "soar19r5gmm7nqxy2v0pzm3c8ldkzax7ugqy5jwrv2y,1",
				Vrv:     "",
				Proof:   "e8cd528e10b85b629bd836b2f71a964cd4c2734f8136093d41e677b3c98fbb0e72f2f53371f6f4b068c3d05370d383f4b6e2ca59b5b71a745c7207c3dc754a0d58bd4cbbc630906c70c214cfdcbedfbd649627da37d8e53ce8cc14168b3e792b",
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
				Pubkey:  "66eea999dcfb6fa4df8a5d2b22ea5e637d65ff9525e5f58f5e27bdac457c0450",
				Message: "soar19r5gmm7nqxy2v0pzm3c8ldkzax7ugqy5jwrv2y,1",
				Vrv:     "4afbf8af915f46626cadcff67ea7eee354fb6b8a3645de238126355fe524cd8c",
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
