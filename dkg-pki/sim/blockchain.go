package sim

import (
	"fmt"
	"os"
	"bytes"
	"github.com/ethereum/go-ethereum/common"
	"github.com/dfinity/dkg-pki/bls"
	"github.com/dfinity/dkg-pki/state"
	"github.com/ethereum/go-ethereum/crypto"
)

var caflag bool
var issuerNum int
// BlockchainSimulator -- Encodes the state of all processes, groups and the blockchain
type BlockchainSimulator struct {
	groupSize uint16
	threshold uint16
	seed      bls.Rand
	proc      []ProcessSimulator
	group     []GroupSimulator
	grpmap    map[common.Address]*GroupSimulator
	chain     []state.State
}

// DoubleCheck -- enable optional double-checks for verification
var DoubleCheck = true

// Vvec -- enable checks involving the verification vectors
var Vvec = true

// Timing -- enable output of timing information
var Timing = false

// InitProcs -- initialize the individual processes for the genesis block
func (sim *BlockchainSimulator) InitProcs(n uint) {
	sim.proc = make([]ProcessSimulator, n)
	rsec := sim.seed.Ders("InitProcs_sec")
	rseed := sim.seed.Ders("InitProcs_seed")
//	fmt.Printf("(rsec)%x\n", rsec.Bytes())
//	fmt.Printf("(rseed)%x\n",rseed.Bytes())

	for i := 0; i < int(n); i++ {
		sim.proc[i] = NewProcessSimulator(bls.SeckeyFromRand(rsec.Deri(i)), rseed.Deri(i))
	//	fmt.Println(sim.proc[i].String())
	}
}

// InitGroups -- initialize the groups for the genesis block
func (sim *BlockchainSimulator) InitGroups(n uint16) {
	sim.group = make([]GroupSimulator, n)
	sim.grpmap = make(map[common.Address]*GroupSimulator)
	r := sim.seed.Ders("InitGroups")
	// build a temporary state datastructure from processes
	/* s := state.NewState()
	for _, p := range sim.proc {
		s.AddNode(p.reginfo)
	} */
	// create n groups
	for i := 0; i < int(n); i++ {
		// choose members based on r
		/* groupinfo := s.NewRandomGroup(r.Deri(i), sim.groupSize)
		   groupinfo.Log() */
		// LATER: replace the following using groupinfo
		indices := r.Deri(i).RandomPerm(len(sim.proc), int(sim.groupSize))
		members := make([]*ProcessSimulator, sim.groupSize)
		for j, idx := range indices {
			members[j] = &(sim.proc[idx])
		}
		sim.group[i] = NewGroupSimulator(members, sim.threshold)
		sim.grpmap[sim.group[i].Address()] = &sim.group[i]
		fmt.Println(sim.group[i].String())
	}
}

// NewBlockchainSimulator -- create a new blockchain simulation
// set the seed and define parameters like group size, threshold, number of processes etc.
func NewBlockchainSimulator(seed bls.Rand, groupSize uint16, threshold uint16, nProcesses uint, nGroups uint16) BlockchainSimulator {
	sim := BlockchainSimulator{seed: seed, groupSize: groupSize, threshold: threshold}
//	sim.Log()

	// Start the processes first
//	fmt.Printf("--- Process setup: (N)%d\n", nProcesses)
//	fmt.Printf("(N)%d (k)%d\n", nProcesses,sim.threshold)
	sim.InitProcs(nProcesses)

	// Start the groups
//	fmt.Printf("--- Group setup: (m)%d\n", nGroups)
	sim.InitGroups(nGroups)

	// Build the genesis block
	genesis := state.NewState()
	for _, p := range sim.proc {
		genesis.AddNode(p.reginfo)
		// this includes verification of proof-of-possession
	}
	for _, g := range sim.group {
		genesis.AddGroup(g.reginfo)
	}
	// the sig field remains empty because the genesis block is not signed

	// print op counts
	if Timing {
		bls.PrintCtrs()
	}

	// Build the chain with 1 block
	sim.chain = append(sim.chain, genesis)

	return sim
}

// Advance -- carry out the simulation for the given number of steps (blocks)
func (sim *BlockchainSimulator) Advance(n uint, verbose bool) {
	if n == 0 {
		return
	}
	// choose tip
	tip := sim.Tip()
	// select pre-determined random group from tip
	a := tip.SelectedGroupAddress()
	g := sim.grpmap[a]
	// get new group signature
	var buffer bytes.Buffer
	if !caflag{
	var dst []byte
	fmt.Sscanf("3101010000A000000333181230041140" +
		"37710FEB7CC3617767874E85509C268E8F931D68773E93A" +
		"89F39A4247DFE2D280FC5BC838353885B6DAD447C8F90116BD" +
		"9D314047591989F67F319544D42A48B", "%X", &dst)
	h := crypto.Keccak256Hash(dst)
	sig := g.Sign(h.Bytes())
	if DoubleCheck {
		if !bls.VerifySig(tip.GroupPubkey(a), h.Bytes(), sig) {
			fmt.Println("Error: group signature not valid.")
		}
	}
	fmt.Printf("\nmsgHash: %x\n", h.Bytes())
	buffer.Write(dst)
	fmt.Printf("sig: %x\n", sig.Rand().Bytes())
	buffer.Write(sig.Rand().Bytes())
	ca :=buffer.Bytes()
	fmt.Printf("ca: %x\n", ca)
	file,err := os.OpenFile("01010000.C18", os.O_CREATE|os.O_RDWR, os.ModePerm)
	if err != nil {
		defer file.Close()
		os.Exit(0)
	}
	_ , err1 := file.Write(ca)
	if err1 != nil {
		defer file.Close()
		os.Exit(0)
	}
	defer file.Close()
	fmt.Println("Successfully generate the root certificate!")
	caflag =true
	}
	var dst1 []byte
	issueName :=fmt.Sprintf("%.6X",0x348+issuerNum)
	dataMsg :="12621462FF1230"
	dataMsg +=issueName
	dataMsg +="040011403838B10B85D96C2EB1D07DE85EC87D44C6ADD77F94A1F90C310897BF29FC784D71FB87D6C12137C25F0D990B776095EAA2EE45E4FA7297E20349C5A50E905351"
	fmt.Sscanf(dataMsg, "%X", &dst1)
//	fmt.Printf("dst1:%x \n",dst1)
	buffer.Reset()
	h := crypto.Keccak256Hash(dst1)
	sig := g.Sign(h.Bytes())
	fmt.Printf("\nmsgHash: %x\n", h.Bytes())
	buffer.Write(dst1)
	fmt.Printf("sig: %x\n", sig.Rand().Bytes())
	buffer.Write(sig.Rand().Bytes())
	issuer :=buffer.Bytes()
	fmt.Printf("issuer: %x\n", issuer)
	issueName +=".I18"
	file1,err2 := os.OpenFile(issueName, os.O_CREATE|os.O_RDWR, os.ModePerm)
	if err2 != nil {
		defer file1.Close()
		os.Exit(0)
	}
	_ , err3 := file1.Write(issuer)
	if err3 != nil {
		defer file1.Close()
		os.Exit(0)
	}
	defer file1.Close()

	fmt.Printf("Successfully generate the issuer%d certificate!\n",issuerNum+1)
	issuerNum +=1

	// the new state is identical to the curren tip, except that we overwrite the signature
	newstate := tip

	// sign new state by group
	newstate.SetSignature(sig)

	// append new state
	sim.chain = append(sim.chain, newstate)

	// recurse
	sim.Advance(n-1, verbose)
	return
}

// Log -- print out a short form of the current state of the random beacon
func (sim *BlockchainSimulator) Log() {
	seed := sim.seed.Bytes()
	fmt.Printf("BlkCh: (n)%d (k)%d (seed)%x\n", sim.groupSize, sim.threshold, seed[:8])
	/*
		fmt.Println("  groups: ", len(sim.group))
		fmt.Println("  processes: ", len(sim.proc))
		fmt.Println("  chain height: ", len(sim.chain))
		sim.chain[len(sim.chain)-1].Log()
		for _, p := range sim.proc {
			p.Log()
		}
		for _, g := range sim.group {
			g.Log()
		}
	*/
}

// Length -- return the current block height
func (sim *BlockchainSimulator) Length() int {
	return len(sim.chain)
}

// Tip -- return the current state at the tip of the chain
func (sim *BlockchainSimulator) Tip() state.State {
	return sim.chain[len(sim.chain)-1]
}
