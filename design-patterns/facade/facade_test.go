package facade_test

import (
	"os"

	"github.com/cndoit18/cs_study_plan/design-patterns/facade"
)

func ExampleNewShopSystem() {
	shopSystem := facade.NewShopSystem(os.Stdout)
	shopSystem.Buy()
	// Output:
	// Settlement completed
	// Payment of $20.00
}
