package keeper

import (
	"context"
	"encoding/base64"
	"strconv"

	"fiamma/x/zkpverify/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) SendTask(goCtx context.Context, msg *types.MsgSendTask) (*types.MsgSendTaskResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	proofSystemId, err := types.ProofSystemIdFromString(msg.ProofSystem)
	if err != nil {
		k.Logger().Info("Error parsing proof system:", "error", err)
		return nil, types.ErrInvalidProofSystem

	}

	proofBuffer, err := base64.StdEncoding.DecodeString(msg.Proof)
	if err != nil {
		return nil, types.ErrDecodeProof
	}

	publicInputBuffer, err := base64.StdEncoding.DecodeString(msg.PublicInput)
	if err != nil {
		return nil, types.ErrDecodeProof
	}

	verifyKey, err := base64.StdEncoding.DecodeString(msg.Vk)
	if err != nil {
		return nil, types.ErrDecodeProof

	}

	verfiyData := types.VerifyData{
		ProofSystem: uint64(proofSystemId),
		Proof:       proofBuffer,
		PublicInput: publicInputBuffer,
		Vk:          verifyKey,
	}
	// submit proof data to DA
	verifyId, dataCommitments, err := k.SubmitVerifyDataToDA(ctx, verfiyData)
	if err != nil {
		k.Logger().Info("Error submitting proof to DA:", "error", err)
		return nil, types.ErrSubmitProof
	}

	verifyIdStr := base64.StdEncoding.EncodeToString(verifyId[:])
	dataCommitmentStr := base64.StdEncoding.EncodeToString(dataCommitments[0])

	// The chain first verifies the correctness of the proofs submitted by the user, and saves the results.
	// The observer may challenge the result at a later stage.
	result := k.verifyProof(&verfiyData)

	k.Logger().Info("Proof verification:", "result", result)

	// store verify data in the store
	verifyResult := types.VerifyResult{
		VerifyId:       verifyIdStr,
		DataCommitment: dataCommitmentStr,
		Result:         result,
	}

	k.SetVerifyResult(ctx, verifyId[:], verifyResult)

	event := sdk.NewEvent("verifyFinished",
		sdk.NewAttribute("verifyId", verifyIdStr),
		sdk.NewAttribute("dataCommitment", dataCommitmentStr),
		sdk.NewAttribute("verifyResult", strconv.FormatBool(result)),
		sdk.NewAttribute("proofSystem", msg.ProofSystem))
	ctx.EventManager().EmitEvent(event)

	return &types.MsgSendTaskResponse{}, nil
}
