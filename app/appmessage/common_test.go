package appmessage

import "github.com/c4ei/yunseokyeol/domain/consensus/model/externalapi"

// mainnetGenesisHash is the hash of the first block in the block DAG for the
// main network (genesis block).
var mainnetGenesisHash = externalapi.NewDomainHashFromByteArray(&[externalapi.DomainHashSize]byte{
	0xdc, 0x5f, 0x5b, 0x5b, 0x1d, 0xc2, 0xa7, 0x25,
	0x49, 0xd5, 0x1d, 0x4d, 0xee, 0xd7, 0xa4, 0x8b,
	0xaf, 0xd3, 0x14, 0x4b, 0x56, 0x78, 0x98, 0xb1,
	0x8c, 0xfd, 0x9f, 0x69, 0xdd, 0xcf, 0xbb, 0x63,
})

// simnetGenesisHash is the hash of the first block in the block DAG for the
// simulation test network.
var simnetGenesisHash = externalapi.NewDomainHashFromByteArray(&[externalapi.DomainHashSize]byte{
	0x9d, 0x89, 0xb0, 0x6e, 0xb3, 0x47, 0xb5, 0x6e,
	0xcd, 0x6c, 0x63, 0x99, 0x45, 0x91, 0xd5, 0xce,
	0x9b, 0x43, 0x05, 0xc1, 0xa5, 0x5e, 0x2a, 0xda,
	0x90, 0x4c, 0xf0, 0x6c, 0x4d, 0x5f, 0xd3, 0x62,
})

// mainnetGenesisMerkleRoot is the hash of the first transaction in the genesis
// block for the main network.
var mainnetGenesisMerkleRoot = externalapi.NewDomainHashFromByteArray(&[externalapi.DomainHashSize]byte{
	0x4a, 0x5e, 0x1e, 0x4b, 0xaa, 0xb8, 0x9f, 0x3a,
	0x32, 0x51, 0x8a, 0x88, 0xc3, 0x1b, 0xc8, 0x7f,
	0x61, 0x8f, 0x76, 0x67, 0x3e, 0x2c, 0xc7, 0x7a,
	0xb2, 0x12, 0x7b, 0x7a, 0xfd, 0xed, 0xa3, 0x3b,
})

var exampleAcceptedIDMerkleRoot = externalapi.NewDomainHashFromByteArray(&[externalapi.DomainHashSize]byte{
	0x09, 0x3B, 0xC7, 0xE3, 0x67, 0x11, 0x7B, 0x3C,
	0x30, 0xC1, 0xF8, 0xFD, 0xD0, 0xD9, 0x72, 0x87,
	0x7F, 0x16, 0xC5, 0x96, 0x2E, 0x8B, 0xD9, 0x63,
	0x65, 0x9C, 0x79, 0x3C, 0xE3, 0x70, 0xD9, 0x5F,
})

var exampleUTXOCommitment = externalapi.NewDomainHashFromByteArray(&[externalapi.DomainHashSize]byte{
	0x10, 0x3B, 0xC7, 0xE3, 0x67, 0x11, 0x7B, 0x3C,
	0x30, 0xC1, 0xF8, 0xFD, 0xD0, 0xD9, 0x72, 0x87,
	0x7F, 0x16, 0xC5, 0x96, 0x2E, 0x8B, 0xD9, 0x63,
	0x65, 0x9C, 0x79, 0x3C, 0xE3, 0x70, 0xD9, 0x5F,
})
