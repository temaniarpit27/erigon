package cltrace

//go:generate go run github.com/eiiches/go-gen-proxy/cmd/go-gen-proxy --interface github.com/ledgerwatch/erigon/cl/abstract.BeaconState --package github.com/ledgerwatch/erigon/cl/cltrace --name BeaconStateProxy --output ./beaconstate.go

////go run github.com/roy2220/proxyz/cmd/proxyz -w proxy.go ../abstract BeaconState cltrace BeaconStateProxy