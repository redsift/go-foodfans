package foodfans_test

import (
	"math/rand"
	"testing"

	"github.com/redsift/go-foodfans"
)

func TestExampleFoodFans_New(t *testing.T) {
	want := []string {
		"cut_bland_cabbage",
		"separate_flavorful_cumquat",
		"separate_boiled_anchovies",
		"salt_chunked_coconut",
		"discard_clove_coated_cod",
		"divide_acrid_aero_chocolate_bar",
		"puree_chunked_cauliflower",
		"mound_gingery_buffalo_wings",
		"tilt_candied_basil",
		"soak_caramelized_cumquat",
	}

	foodfans.New()
	ff := foodfans.NewFoodFans(rand.NewSource(0))
	for i:=0; i<10; i++ {
		got := ff.New()
		if got != want[i] {
			t.Errorf("want %s, got %s", want[i], got)
		}
	}
}
