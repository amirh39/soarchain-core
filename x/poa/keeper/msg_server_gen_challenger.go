package keeper

import (
	"context"
	params "soarchain/app/params"
	"soarchain/x/poa/types"
	"strings"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func (k msgServer) GenChallenger(goctx context.Context, msg *types.MsgGenChallenger) (*types.MsgGenChallengerResponse, error) {
	ctx := sdk.UnwrapSDKContext(goctx)

	challengerAddress, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "msg.Creator couldn't be parsed.")
	}

	challengerType := strings.ToLower(msg.Challengertype)

	if challengerType != "v2n" && challengerType != "v2x" {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "Invalid challenger type. Must be 'v2n' or 'v2x'.")
	}

	if msg.ChallengerStake == "" {
		return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, "Challenger Stake must be declared in the tx!")
	}

	if msg.Certificate == "" {
		return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, "Certificate must be declared in the tx!")
	}

	if msg.Signature == "" {
		return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, "Signature must be declared in the tx!")
	}

	deviceCert, err := k.CreateX509CertFromString(msg.Certificate)
	if err != nil {
		return nil, sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest, "[GenChallenger][CreateX509CertFromString] failed. Invalid device certificate. Error: [ %T ]", err)
	}

	pubKeyHex, err := k.VerifyX509CertByASN1AndExtractPubkey(msg.Creator, msg.Signature, deviceCert)
	if pubKeyHex == "" || err != nil {
		return nil, sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest, "[GenChallenger][VerifyX509CertByASN1AndExtractPubkey] failed. Invalid certificate validation. Error: [ %T ]", err)
	}

	// Check validity of certificate
	totalKeys := k.GetAllFactoryKeys(ctx)
	var validated bool = false
	var verificationError error = nil

	for i := uint64(0); i < uint64(len(totalKeys)); i++ {
		factoryKey, isFound := k.GetFactoryKeys(ctx, i)
		if isFound {
			factoryCert, err := k.CreateX509CertFromString(factoryKey.FactoryCert)
			if err != nil {
				return nil, sdkerrors.Wrap(sdkerrors.ErrPanic, "[GenChallenger][CreateX509CertFromString] failed. Factory certificate couldn't be created from the storage."+err.Error())
			}

			validated, err = k.ValidateX509Cert(deviceCert, factoryCert)
			if err != nil {
				verificationError = sdkerrors.Wrap(sdkerrors.ErrPanic, "[GenChallenger][ValidateX509Cert] failed. Couldn't validate factory certificate."+err.Error())
				continue // Try next certificate
			}

			if validated {
				verificationError = nil
				break
			}
		}
	}

	// No valid certificate found
	if verificationError != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrPanic, "[GenChallenger][ValidateX509Cert] failed. Device certificate couldn't be verified.")
	}

	var newChallenger types.Challenger

	// Check challenger stake amount
	requiredStake := sdk.Coins{sdk.NewInt64Coin(params.BondDenom, 1000000000)}
	ChallengerStake, err := sdk.ParseCoinsNormalized(msg.ChallengerStake)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidCoins, "Coins couldn't be parsed!")
	}
	if ChallengerStake.IsAllLT(requiredStake) || !ChallengerStake.DenomsSubsetOf(requiredStake) {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInsufficientFunds, "Sent amount of challenger: "+ChallengerStake.String()+" is below the required stake amount "+requiredStake.String())
	}

	// Transfer stakedAmount to poa modules account:
	transferErr := k.bankKeeper.SendCoinsFromAccountToModule(ctx, challengerAddress, types.ModuleName, requiredStake)
	if transferErr != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrPanic, "Stake(challenger) funds couldn't be transferred to POA module!")
	}

	// rewardMultiplier
	//var initialScore float64 = 50
	//rewardMultiplier := utility.CalculateRewardMultiplier(initialScore)

	newChallenger = types.Challenger{
		PubKey:       pubKeyHex,
		Address:      challengerAddress.String(),
		Score:        sdk.NewInt(50).String(), // Base Score
		StakedAmount: ChallengerStake.String(),
		NetEarnings:  sdk.NewCoin(params.BondDenom, sdk.ZeroInt()).String(),
		Type:         msg.Challengertype,
		IpAddr:       msg.ChallengerIp,
	}

	k.SetChallenger(ctx, newChallenger)

	return &types.MsgGenChallengerResponse{}, nil
}
