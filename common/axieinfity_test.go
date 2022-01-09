package common

import (
	"github.com/ethereum/go-ethereum/accounts/abi/bind/backends"
)

func SLPSetup(client *backends.SimulatedBackend) {
	/*auth, err := bind.NewTransactor(strings.NewReader(key), "my awesome super secret password")
	if err != nil {
		return
	}
	// Deploy a new awesome contracts for the binding demo
	address, tx, token, err := DeployToken(auth, conn), new(big.Int), "Contracts in Go!!!", 0, "Go!")
	if err != nil {
		return
	}
	fmt.Printf("Contract pending deploy: 0x%x\n", address)
	fmt.Printf("Transaction waiting to be mined: 0x%x\n\n", tx.Hash())

	// Don't even wait, check its presence in the local pending state
	time.Sleep(250 * time.Millisecond) // Allow it to be processed by the local node :P

	name, err := token.Name(&bind.CallOpts{Pending: true})
	if err != nil {
		log.Fatalf("Failed to retrieve pending name: %v", err)
	}
	fmt.Println("Pending name:", name)
	*/
}
