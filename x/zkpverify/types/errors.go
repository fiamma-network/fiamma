package types

// DONTCOVER

import (
	sdkerrors "cosmossdk.io/errors"
)

// x/zkpverify module sentinel errors
var (
	ErrInvalidSigner      = sdkerrors.Register(ModuleName, 1100, "expected gov account as only signer for proposal message")
	ErrInvalidVerifyData  = sdkerrors.Register(ModuleName, 1101, "invalid proof verify data")
	ErrSubmitProof        = sdkerrors.Register(ModuleName, 1102, "error submitting proof to DA")
	ErrInvalidProofSystem = sdkerrors.Register(ModuleName, 1103, "invalid proof system")
	ErrDecodeProof        = sdkerrors.Register(ModuleName, 1104, "base64 decoding proof error")
	ErrDecodePublicInput  = sdkerrors.Register(ModuleName, 1105, "base64 decoding public input error")
	ErrDecodeVk           = sdkerrors.Register(ModuleName, 1106, "base64 decoding verification key error")
)
