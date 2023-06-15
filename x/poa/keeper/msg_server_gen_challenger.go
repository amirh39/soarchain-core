package keeper

import (
	"context"
	"log"
	params "soarchain/app/params"
	"soarchain/x/poa/constants"
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

	if challengerType != constants.V2NChallengerType && challengerType != constants.V2XChallenger {
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
	errCert := k.validateCertificate(ctx, deviceCert)
	if errCert != nil {
		return nil, errCert
	}

	var newChallenger types.Challenger

	// Check challenger stake amount
	requiredStake := sdk.Coins{sdk.NewInt64Coin(params.BondDenom, constants.ChallengerAmount)}
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
		log.Println(transferErr)
		return nil, sdkerrors.Wrap(sdkerrors.ErrPanic, "Stake(challenger) funds couldn't be transferred to POA module!")
	}

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
