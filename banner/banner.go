package banner

import (
	"fmt"
)

// prints the version message
const version = "v0.0.3"

func PrintVersion() {
	fmt.Printf("Current org2asn version %s\n", version)
}

// Prints the Colorful banner
func PrintBanner() {
	banner := `
                      ___                    
  ____   _____ ____ _|__ \ ____ _ _____ ____ 
 / __ \ / ___// __  /__/ // __  // ___// __ \
/ /_/ // /   / /_/ // __// /_/ /(__  )/ / / /
\____//_/    \__, //____/\__,_//____//_/ /_/ 
            /____/
`
	fmt.Printf("%s\n%50s\n\n", banner, "Current org2asn version "+version)
}
