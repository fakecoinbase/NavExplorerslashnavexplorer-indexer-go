package consensus

import (
	"github.com/NavExplorer/navexplorer-indexer-go/pkg/explorer"
	"log"
)

type consensusParameters []*explorer.ConsensusParameter

var Parameters consensusParameters

type Parameter int

var (
	VOTING_CYCLE_LENGTH Parameter = 0

	CONSULTATION_MIN_SUPPORT        Parameter = 1
	CONSULTATION_ANSWER_MIN_SUPPORT Parameter = 2

	CONSULTATION_MIN_CYCLES         Parameter = 3
	CONSULTATION_MAX_VOTING_CYCLES  Parameter = 4
	CONSULTATION_MAX_SUPPORT_CYCLES Parameter = 5
	CONSULTATION_REFLECTION_LENGTH  Parameter = 6
	CONSULTATION_MIN_FEE            Parameter = 7

	CONSULTATION_ANSWER_MIN_FEE Parameter = 8

	PROPOSAL_MIN_QUORUM        Parameter = 9
	PROPOSAL_MIN_ACCEPT        Parameter = 10
	PROPOSAL_MIN_REJECT        Parameter = 11
	PROPOSAL_MIN_FEE           Parameter = 12
	PROPOSAL_MAX_VOTING_CYCLES Parameter = 13

	PAYMENT_REQUEST_MIN_QUORUM        Parameter = 14
	PAYMENT_REQUEST_MIN_ACCEPT        Parameter = 15
	PAYMENT_REQUEST_MIN_REJECT        Parameter = 16
	PAYMENT_REQUEST_MIN_FEE           Parameter = 17
	PAYMENT_REQUEST_MAX_VOTING_CYCLES Parameter = 18

	FUND_SPREAD_ACCUMULATION Parameter = 19
	FUND_PERCENT_PER_BLOCK   Parameter = 20

	GENERATION_PER_BLOCK                    Parameter = 21
	NAVNS_FEE                               Parameter = 22
	CONSENSUS_PARAMS_DAO_VOTE_LIGHT_MIN_FEE Parameter = 23
)

func (p *consensusParameters) Get(parameter Parameter) *explorer.ConsensusParameter {
	for _, c := range Parameters {
		if parameter == Parameter(c.Id) {
			return c
		}
	}

	log.Fatalf("Failed to get consensus parameter %s", string(parameter))

	return nil
}

func (p *consensusParameters) GetById(id int) *explorer.ConsensusParameter {
	for _, c := range Parameters {
		if id == c.Id {
			return c
		}
	}

	log.Fatalf("Failed to get consensus parameter %d", id)

	return nil
}
