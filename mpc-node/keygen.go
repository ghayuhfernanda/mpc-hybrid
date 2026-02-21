package main

import (
	"encoding/json"
	"net/http"
	"github.com/bnb-chain/tss-lib/ecdsa/keygen"
    "github.com/bnb-chain/tss-lib/tss"
)

// KeygenHandler handles key generation
func KeygenHandler(w http.ResponseWriter, r *http.Request) {
	// Peer IDs must be shared/configured beforehand
    peerIDs := tss.GenerateTestPartyIDs(3)
    selfID := peerIDs[0] // vary per node (A,B,C)

    params := makeParams(peerIDs, selfID, 2)

    outCh := make(chan tss.Message, 100)
    endCh := make(chan keygen.LocalPartySaveData, 1)

    party := keygen.NewLocalParty(params, outCh, endCh)

    go party.Start()

    saveData := <-endCh

    // Save this node’s share for use in signing
    SavePartyData("party_save.json", saveData)

    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(map[string]string{
        "status": "keygen completed",
    })
}
