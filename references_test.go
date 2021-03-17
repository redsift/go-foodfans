package foodfans_test

import (
	"math/rand"
	"testing"

	"github.com/redsift/go-foodfans"
)

func TestExampleFoodFans_New(t *testing.T) {
	want := []string{
		"cut_heavenly_fruit_pie",
		"separate_warm_corn_flakes",
		"separate_boiled_funnel_cake",
		"salt_chunked_pollock",
		"discard_lemon_mayonnaise",
		"divide_acrid_fruit",
		"puree_layered_guacamole",
		"mound_tender_uglifruit",
		"tilt_infused_peanuts",
		"soak_caramelized_yam",
	}

	foodfans.New()
	rand.Seed(0)
	for i := 0; i < 10; i++ {
		got := foodfans.New()
		if got != want[i] {
			t.Errorf("want %s, got %s", want[i], got)
		}
	}
}
