package types

import (
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	sdkcodec "github.com/cosmos/cosmos-sdk/codec"

	"soarchain/x/did/errors"
)

type RunnerDidDocumentOption func(opts *RunnerDid)

func NewRunnerDidDocument(id string, pubkey string, address string, opts ...RunnerDidDocumentOption) RunnerDid {
	doc := RunnerDid{
		Id:      id,
		PubKey:  pubkey,
		Address: address,
	}

	for _, opt := range opts {
		opt(&doc)
	}
	return doc
}

func NewRunnerDidDocumentWithSeq(doc *RunnerDid, seq uint64) RunnerDidWithSeq {
	return RunnerDidWithSeq{
		Document: doc,
		Sequence: seq,
	}
}

func (doc RunnerDid) Empty() bool {
	return EmptyDid(doc.Id)
}

func (d RunnerDidWithSeq) Empty() bool {
	return d.Document == nil || d.Document.Empty() && d.Sequence == InitialSequence
}

func (d RunnerDidWithSeq) Deactivated() bool {
	return d.Document.Empty() && d.Sequence != InitialSequence
}

func (doc RunnerDid) RunnerVerificationMethodByID(id string) (VerificationMethod, bool) {
	for _, verificationMethod := range doc.VerificationMethods {
		if verificationMethod.Id == id {
			return *verificationMethod, true
		}
	}
	return VerificationMethod{}, false
}

func (doc RunnerDid) VerificationMethodFrom(relationships []VerificationRelationship, id string) (VerificationMethod, bool) {
	for _, relationship := range relationships {
		if relationship.hasDedicatedMethod() {
			veriMethod := relationship.GetVerificationMethod()
			if veriMethod.Id == id {
				return *veriMethod, true
			}
		} else {
			veriMethodID := relationship.GetVerificationMethodId()
			if veriMethodID == id {
				return doc.RunnerVerificationMethodByID(veriMethodID)
			}
		}
	}

	return VerificationMethod{}, false
}

func WithRunnerVerificationMethods(verificationMethods []*VerificationMethod) RunnerDidDocumentOption {
	return func(opts *RunnerDid) {
		opts.VerificationMethods = verificationMethods
	}
}

func WithRunnerAuthentications(authentications []VerificationRelationship) RunnerDidDocumentOption {
	return func(opts *RunnerDid) {
		opts.Authentications = authentications
	}
}

func WithRunnerKeys(soarchainPublicKey *Keys) RunnerDidDocumentOption {
	return func(opts *RunnerDid) {
		opts.Keys = soarchainPublicKey
	}
}

func (d RunnerDidWithSeq) Deactivate(newSeq uint64) RunnerDidWithSeq {
	return NewRunnerDidDocumentWithSeq(&RunnerDid{}, newSeq)
}

func mustGetRunnerSignBytesWithSeq(signableData sdkcodec.ProtoMarshaler, seq uint64) ([]byte, error) {
	dAtA, err := signableData.Marshal()
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrNotFound, "[mustGetSignBytesWithSeq][Marshal] failed. SignableData is not valid.")
	}
	dataWithSeq := RunnerDataWithSeq{
		Data:     dAtA,
		Sequence: seq,
	}

	dAtA, err = dataWithSeq.Marshal()

	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrNotFound, errors.ErrInvalidDidDocument)
	}
	return dAtA, nil
}
