package randomfts

/**
 * A class containing the result produced by follow-the-satoshi.
 */
type FtsResult struct {
	merkleProof []ProofEntry
	stakeholder StakeHolder
}

func NewFtsResult(merkleProof []ProofEntry, stakeholder StakeHolder ) FtsResult {
	return  FtsResult{merkleProof: merkleProof, stakeholder: stakeholder}
}

func (r FtsResult)getStakeHolder() StakeHolder {
	return r.stakeholder
}

func (r FtsResult)getMerkleProof() []ProofEntry {
	return r.merkleProof
}
