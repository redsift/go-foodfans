#Food Fans

Generates memorable references using food.

	import "github.com/redsift/go-foodfans"
	s := foodfans.NewReference()

## go generate

Relies on `stringer` to generate lookup table. If you change the constants, regenerate the table.

	# Install stringer if required
	go get golang.org/x/tools/cmd/stringer
	
	# Regenerate the lookup
	cd lookup
	go generate
