#Food Fans

Generates memorable references using food.

	import "github.com/redsift/go-foodfans"
	
	# Seed the rng if required in your main()
	rand.Seed(time.Now().UTC().UnixNano())
	
	s := foodfans.New()

## go generate

Relies on `stringer` to generate lookup table. If you change the constants, regenerate the table.

	# Install stringer if required
	go get golang.org/x/tools/cmd/stringer
	
	# Regenerate the lookup
	cd lookup
	go generate
