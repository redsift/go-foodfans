package foodfans_test

import (
	"fmt"
	"github.com/redsift/go-foodfans"
	"math/rand"
)

func ExampleFoodFans_New() {
	foodfans.New()
	ff := foodfans.NewFoodFans(rand.NewSource(0))
	for i:=0; i<10; i++ {
		fmt.Println(ff.New())
	}
	// Output:
	// cut_bland_cabbage
	// separate_flavorful_cumquat
	// separate_boiled_anchovies
	// salt_chunked_coconut
	// discard_clove_coated_cod
	// divide_acrid_aero_chocolate_bar
	// puree_chunked_cauliflower
	// mound_gingery_buffalo_wings
	// tilt_candied_basil
	// soak_caramelized_cumquat
}
