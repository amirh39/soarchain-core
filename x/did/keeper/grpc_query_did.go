package keeper

import (
	"context"
	"log"

	"soarchain/x/did/errors"
	"soarchain/x/did/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) DidAll(c context.Context, req *types.QueryAllDidRequest) (*types.QueryAllDidResponse, error) {

	log.Println("############## Fetching all did is Started ##############")

	if req == nil {
		return nil, status.Error(codes.InvalidArgument, errors.InvalidRequest)
	}

	ctx := sdk.UnwrapSDKContext(c)

	dids := k.GetAllDid(ctx)

	log.Println("############## End of fetching all dids ##############")

	return &types.QueryAllDidResponse{Did: dids}, nil
}

func (k Keeper) Did(c context.Context, req *types.QueryGetDidRequest) (*types.QueryGetDidResponse, error) {

	log.Println("############## Fetching a did is Started ##############")

	ctx := sdk.UnwrapSDKContext(c)

	if req == nil {
		return nil, status.Error(codes.InvalidArgument, errors.InvalidRequest)
	}

	did := string(req.Id)
	didDocument, found := k.GetDidDocumentWithSequence(ctx, did)

	if !found {
		return nil, sdkerrors.Wrap(sdkerrors.ErrNotFound, errors.ErrDidNotFound)
	}
	if didDocument.Empty() {
		return nil, sdkerrors.Wrap(sdkerrors.ErrNotFound, errors.ErrDidNotFound)
	}
	if didDocument.Deactivated() {
		return nil, sdkerrors.Wrap(sdkerrors.ErrNotFound, errors.ErrDidNotActive)
	}

	log.Println("############## End of fetching a did ##############")

	return &types.QueryGetDidResponse{DidDocument: &didDocument}, nil
}

// func (k Keeper) DidAll1(c context.Context, req *types.QueryAllDidRequest) (*types.QueryAllDidResponse, error) {

// 	ctx := sdk.UnwrapSDKContext(c)

// 	dids := k.GetAllDid(ctx)

// 	// case 1:
// 	if len(pins) == 1 {
// 		if pins[0] == 1 && val.Document.PidSupportedOneToTwnety {
// 			result = append(result, val)
// 		}
// 		if pins[0] == 2 && val.Document.PidSupportedTwentyOneToForthy {
// 			result = append(result, val)
// 		}
// 		if pins[0] == 3 && val.Document.PidSupportedForthyOneToSixty {
// 			result = append(result, val)
// 		}
// 	}

// 	// case 2:
// 	// fmt.Print("hereeeeeeeeeeeeeeeeeeeeee")
// 	if len(pins) == 2 {
// 		fmt.Print("xxxxxxxxxxxxxxxxxxxx", pins, val.Document.PidSupportedOneToTwnety, val.Document.PidSupportedTwentyOneToForthy)
// 		if pins[0] == 1 && pins[1] == 2 && val.Document.PidSupportedOneToTwnety && val.Document.PidSupportedTwentyOneToForthy {
// 			result = append(result, val)
// 		}
// 		if pins[0] == 1 && pins[1] == 3 && val.Document.PidSupportedOneToTwnety && val.Document.PidSupportedForthyOneToSixty {
// 			result = append(result, val)
// 		}
// 		if pins[0] == 2 && pins[1] == 3 && val.Document.PidSupportedTwentyOneToForthy && val.Document.PidSupportedForthyOneToSixty {
// 			result = append(result, val)
// 		}
// 	}

// 	// case 3:
// 	if len(pins) == 3 {
// 		if pins[0] == 1 && pins[1] == 2 && pins[2] == 3 && val.Document.PidSupportedOneToTwnety && val.Document.PidSupportedTwentyOneToForthy && val.Document.PidSupportedForthyOneToSixty {
// 			result = append(result, val)
// 		}
// 	}

// 	log.Println("############## End of fetching all dids ##############")

// 	return &types.QueryAllDidResponse{Did: dids}, nil
// }
