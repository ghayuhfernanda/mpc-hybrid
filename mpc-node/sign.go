package main

import (
    "encoding/json"
    "net/http"

    "github.com/bnb-chain/tss-lib/ecdsa/signing"
    "github.com/bnb-chain/tss-lib/tss"
)

type SignRequest struct {
    Message []byte `json:"message"`
}

func SignHandler(w http.ResponseWriter, r *http.Request) {
    var req SignRequest
    json.NewDecoder(r.Body).Decode(&req)

    // Load saved key share
    var saveData signing.LocalPartySaveData
    LoadPartyData("party_save.json", &saveData)

    // Peer IDs must match those used in key generation
    peerIDs := tss.GenerateTestPartyIDs(3)
    selfID := peerIDs[0] // unique per node

    params := makeParams(peerIDs, selfID, 2)

    outCh := make(chan tss.Message, 100)
    endCh := make(chan *signing.SignatureData, 1)

    party := signing.NewLocalParty(req.Message, params, saveData, outCh, endCh)

    go party.Start()

    sig := <-endCh

    // Return this node’s partial signature
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(sig)
}
