package keeper

import (
	"context"
	"crypto/ecdsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/hex"
	"errors"
	"fmt"
	params "soarchain/app/params"
	"soarchain/x/poa/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func (k msgServer) GenChallenger(goctx context.Context, msg *types.MsgGenChallenger) (*types.MsgGenChallengerResponse, error) {
	ctx := sdk.UnwrapSDKContext(goctx)

	deviceCert, err := k.CreateX509CertFromString(msg.Certificate)
	if err != nil {
		return nil, sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest, "[GenChallenger][CreateX509CertFromString] failed. Invalid device certificate. Error: [ %T ]", err)
	}
	fmt.Print("hhhhhhhhhhhhhhhh---deviceCert")
	fmt.Print(deviceCert)

	pubKeyFromCertificate, err := x509.MarshalPKIXPublicKey(deviceCert.PublicKey)
	if err != nil {
		return nil, sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest, "[GenChallenger][MarshalPKIXPublicKey] failed. Couldn't convert a public key to PKIX.Error: [ %T ]", err)
	}

	fmt.Print("ccccccccccccccccccccc-----pubKeyFromCertificate")
	fmt.Print(pubKeyFromCertificate)

	pubKeyHex := hex.EncodeToString(pubKeyFromCertificate)

	fmt.Print("uuuuuuuuuuuuuuuuuuuu------pubKeyHex")
	fmt.Print(pubKeyHex)

	signature, err := hex.DecodeString(msg.Signature)
	if err != nil {
		return nil, sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest, "[GenChallenger][DecodeString] failed. Invalid signature encoding.Error: [ %T ]", err)
	}

	fmt.Print("aaaaaaaaaaaaaaaaa----------------signature")
	fmt.Print(signature)

	hashedAddr := sha256.Sum256([]byte(msg.Creator))

	fmt.Print("rrrrrrrrrrrrrrrrr--------hashedAddr")
	fmt.Print(signature)

	if deviceCert.PublicKeyAlgorithm == x509.ECDSA {

		if ecdsaPubKey, ok := deviceCert.PublicKey.(*ecdsa.PublicKey); ok {

			if ecdsa.VerifyASN1(ecdsaPubKey, hashedAddr[:], signature) {
				// signature is valid
			} else {
				return nil, sdkerrors.Wrapf(sdkerrors.ErrUnauthorized, "[GenChallenger][VerifyASN1] failed. Signature verification failed. Error: [ %T ]", err)
			}
		} else {
			return nil, sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest, "[GenChallenger] failed. Invalid public key type. Error: [ %T ]", err)
		}
	}

	msgSenderAddress, addrErr := sdk.AccAddressFromBech32(msg.Creator)
	if addrErr != nil {
		if errors.Is(addrErr, sdkerrors.ErrInvalidAddress) {
			return nil, sdkerrors.Wrap(addrErr, "msg.Creator couldn't be parsed.")
		}
		return nil, addrErr
	}

	if msg.ChallengerAddr == "" {
		return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, "Challenger Address must be declared in the tx!")
	}

	if msg.ChallengerPubKey == "" {
		return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, "Challenger Public Key must be declared in the tx!")
	}

	if msg.ChallengerStake == "" {
		return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, "Challenger Stake must be declared in the tx!")
	}

	var newChallenger types.Challenger

	challengers := k.GetAllRunner(ctx)
	for _, runner := range challengers {
		if msg.ChallengerPubKey == runner.PubKey {
			sdkerrors.Wrap(sdkerrors.ErrConflict, "Challenger is already registered in storage.")
			break
		}
	}

	runners := k.GetAllRunner(ctx)
	for _, runner := range runners {
		if msg.ChallengerPubKey == runner.PubKey {
			sdkerrors.Wrap(sdkerrors.ErrConflict, "Challenger is already registered as runner in storage.")
			break
		}
	}

	clients := k.GetAllClient(ctx)
	for _, client := range clients {
		if msg.ChallengerPubKey == client.Index {
			sdkerrors.Wrap(sdkerrors.ErrConflict, "Challenger is already registered as client in storage.")
			break
		}
	}

	ChallengerAddr, err := sdk.AccAddressFromBech32(msg.ChallengerAddr)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "Invalid challenger address!")
	}
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
	transferErr := k.bankKeeper.SendCoinsFromAccountToModule(ctx, msgSenderAddress, types.ModuleName, requiredStake)
	if transferErr != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrPanic, "Stake(challenger) funds couldn't be transferred to POA module!")
	}

	// rewardMultiplier
	//var initialScore float64 = 50
	//rewardMultiplier := utility.CalculateRewardMultiplier(initialScore)

	newChallenger = types.Challenger{
		PubKey:       pubKeyHex,
		Address:      ChallengerAddr.String(),
		Score:        sdk.NewInt(50).String(), // Base Score
		StakedAmount: ChallengerStake.String(),
		NetEarnings:  sdk.NewCoin(params.BondDenom, sdk.ZeroInt()).String(),
		Type:         msg.Challengertype,
		IpAddr:       msg.ChallengerIp,
	}

	k.SetChallenger(ctx, newChallenger)

	return &types.MsgGenChallengerResponse{}, nil
}
