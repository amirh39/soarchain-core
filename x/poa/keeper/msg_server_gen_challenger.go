package keeper

import (
	"context"
	"log"
	params "soarchain/app/params"
	"soarchain/x/poa/constants"
	"soarchain/x/poa/errors"
	"soarchain/x/poa/types"
	"strings"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func (k msgServer) GenChallenger(goctx context.Context, msg *types.MsgGenChallenger) (*types.MsgGenChallengerResponse, error) {
	ctx := sdk.UnwrapSDKContext(goctx)
	logger := k.Logger(ctx)

	log.Println("############## Generating a challenger Transaction Started ##############")

	challengerAddress, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, errors.ErrInvalidAddress)
	}

	challengerType := strings.ToLower(msg.ChallengerType)

	if challengerType != constants.V2NChallengerType && challengerType != constants.V2XChallenger {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "Invalid challenger type. Must be 'v2n' or 'v2x'.")
	}

	if logger != nil {
		logger.Info("Cahllenger type is valid.", "transaction", "GenChallenger")
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

	if logger != nil {
		logger.Info("Verifying Cahllenger successfully done.", "transaction", "GenChallenger")
	}

	//check if the address is uniqe
	isUniqueAddress := IsUniqueAddress(k, ctx, msg.Creator)
	if isUniqueAddress {
		return nil, sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest, "[GenChallenger][GetMotusWallet][GetChallengerUsingPubKey][GetRunnerUsingPubKey][GetClient] failed. Client with the address [ %T ] is already registered.", msg.Creator)
	}

	//check if the pubKey is uniqe, also check if msg.creator address have a motus wallet
	isUniquePubkey := IsUniquePubKey(k, ctx, msg.Creator, pubKeyHex)
	if isUniquePubkey {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "[GenChallenger][GetMotusWallet][GetChallengerUsingPubKey][GetRunnerUsingPubKey][GetClient] failed. Client PubKey is not uniqe OR Client is already registered.")
	}

	if logger != nil {
		logger.Info("Checking for Cahllenger address and pubKey successfully done.", "transaction", "GenChallenger")
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

		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidCoins, "Stake(challenger) funds couldn't be transferred to POA module!")
	}

	if logger != nil {
		logger.Info("Transfering coins successfully done.", "transaction", "GenChallenger")

	}

	newChallenger = types.Challenger{
		PubKey:       pubKeyHex,
		Address:      challengerAddress.String(),
		Score:        sdk.NewInt(50).String(), // Base Score
		StakedAmount: ChallengerStake.String(),
		NetEarnings:  sdk.NewCoin(params.BondDenom, sdk.ZeroInt()).String(),
		Type:         msg.ChallengerType,
		IpAddress:    msg.ChallengerIp,
	}

	k.SetChallenger(ctx, newChallenger)

	if logger != nil {
		logger.Info("Updating challenger successfully done.", "transaction", "GenChallenger")
	}

	log.Println("############## End of Gen Challenger Transaction ##############")

	return &types.MsgGenChallengerResponse{}, nil
}
