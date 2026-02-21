package main

import (
    "github.com/bnb-chain/tss-lib/tss"
)

func makeParams(peerIDs []tss.PartyID, selfID *tss.PartyID, threshold int) *tss.Parameters {
    return tss.NewParameters(
        tss.S256(),              // secp256k1
        tss.NewPeerContext(peerIDs...),
        selfID,
        len(peerIDs),
        threshold,
    )
}