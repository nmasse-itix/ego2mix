# Ego2mix

This Go library calls the Eco2mix API from RTE and returns data about the French electricity grid.

## Usage

```go
package main

import (
	"fmt"

	ego2mix "github.com/nmasse-itix/ego2mix"
)

func main() {
	client := ego2mix.NewEco2mixClient("", nil)
	records, err := client.FetchNationalRealTimeData(1) // we want only the last result
	if err != nil {
		panic(err)
	}
	fmt.Printf("Intensité carbone à %s en France: %d gCO2eq / kWh\n", records[0].DateHeure, records[0].TauxCo2)
}
```
