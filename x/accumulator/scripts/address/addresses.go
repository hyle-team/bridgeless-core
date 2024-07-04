package address

import (
	"flag"
	"fmt"
	"os"
	"strings"

	sdk "github.com/cosmos/cosmos-sdk/types"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
)

func main() {
	moduleNames := flag.String("modules", "", "Comma-separated list of module names")
	outputFile := flag.String("output", "module_addresses.txt", "Output file to save module addresses")
	flag.Parse()

	if *moduleNames == "" {
		fmt.Println("Please provide a comma-separated list of module names using the -modules flag.")
		return
	}

	moduleNamesSlice := splitModuleNames(*moduleNames)

	file, err := os.Create(*outputFile)
	if err != nil {
		fmt.Printf("Error creating file: %v\n", err)
		return
	}
	defer file.Close()

	for _, moduleName := range moduleNamesSlice {
		address, err := generateModuleAddressWithPrefix(moduleName, "evmos")
		if err != nil {
			fmt.Printf("Error generating address for module %s: %v\n", moduleName, err)
			continue
		}
		line := fmt.Sprintf("%s: %s\n", moduleName, address)
		_, err = file.WriteString(line)
		if err != nil {
			fmt.Printf("Error writing to file: %v\n", err)
			return
		}
	}

	fmt.Printf("Module addresses have been saved to %s\n", *outputFile)
}

// splitModuleNames splits the comma-separated module names into a slice
func splitModuleNames(moduleNames string) []string {
	return strings.Split(moduleNames, ",")
}

// generateModuleAddressWithPrefix generates an address for a given module name with a specified prefix
func generateModuleAddressWithPrefix(moduleName, prefix string) (string, error) {
	key := moduleName + "address"

	address := authtypes.NewModuleAddress(key)
	bech32Addr, err := sdk.Bech32ifyAddressBytes(prefix, address)
	if err != nil {
		return "", err
	}
	return bech32Addr, nil
}
