package types

import (
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	sdkcodec "github.com/cosmos/cosmos-sdk/codec"

	"soarchain/x/did/errors"
)

type ChallengerDidDocumentOption func(opts *ChallengerDidDocument)

func NewChallengerDidDocument(id string, pubkey string, address string, opts ...ChallengerDidDocumentOption) ChallengerDidDocument {
	doc := ChallengerDidDocument{
		Id:      id,
		PubKey:  pubkey,
		Address: address,
	}

	for _, opt := range opts {
		opt(&doc)
	}
	return doc
}

func NewChallengerDidDocumentWithSeq(doc *ChallengerDidDocument, seq uint64) ChallengerDidDocumentWithSeq {
	return ChallengerDidDocumentWithSeq{
		Document: doc,
		Sequence: seq,
	}
}

func (doc ChallengerDidDocument) Empty() bool {
	return EmptyDid(doc.Id)
}

func (d ChallengerDidDocumentWithSeq) Empty() bool {
	return d.Document == nil || d.Document.Empty() && d.Sequence == InitialSequence
}

func (d ChallengerDidDocumentWithSeq) Deactivated() bool {
	return d.Document.Empty() && d.Sequence != InitialSequence
}

func (doc ChallengerDidDocument) ChallengerVerificationMethodByID(id string) (VerificationMethod, bool) {
	for _, verificationMethod := range doc.VerificationMethods {
		if verificationMethod.Id == id {
			return *verificationMethod, true
		}
	}
	return VerificationMethod{}, false
}

func (doc ChallengerDidDocument) VerificationMethodFrom(relationships []VerificationRelationship, id string) (VerificationMethod, bool) {
	for _, relationship := range relationships {
		if relationship.hasDedicatedMethod() {
			veriMethod := relationship.GetVerificationMethod()
			if veriMethod.Id == id {
				return *veriMethod, true
			}
		} else {
			veriMethodID := relationship.GetVerificationMethodId()
			if veriMethodID == id {
				return doc.ChallengerVerificationMethodByID(veriMethodID)
			}
		}
	}

	return VerificationMethod{}, false
}

func WithChallengerVerificationMethods(verificationMethods []*VerificationMethod) ChallengerDidDocumentOption {
	return func(opts *ChallengerDidDocument) {
		opts.VerificationMethods = verificationMethods
	}
}

func WithChallengerAuthentications(authentications []VerificationRelationship) ChallengerDidDocumentOption {
	return func(opts *ChallengerDidDocument) {
		opts.Authentications = authentications
	}
}

func WithChallengerKeys(soarchainPublicKey *Keys) ChallengerDidDocumentOption {
	return func(opts *ChallengerDidDocument) {
		opts.Keys = soarchainPublicKey
	}
}

func (d ChallengerDidDocumentWithSeq) Deactivate(newSeq uint64) ChallengerDidDocumentWithSeq {
	return NewChallengerDidDocumentWithSeq(&ChallengerDidDocument{}, newSeq)
}

func mustGetChallengerSignBytesWithSeq(signableData sdkcodec.ProtoMarshaler, seq uint64) ([]byte, error) {
	dAtA, err := signableData.Marshal()
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrNotFound, "[mustGetSignBytesWithSeq][Marshal] failed. SignableData is not valid.")
	}
	dataWithSeq := ChallengerDataWithSeq{
		Data:     dAtA,
		Sequence: seq,
	}

	dAtA, err = dataWithSeq.Marshal()

	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrNotFound, errors.ErrInvalidDidDocument)
	}
	return dAtA, nil
}
