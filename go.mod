module github.com/bhashimoto/go-pokedex-bootdev

go 1.22.5

require internal/pokeapi v1.0.0
replace internal/pokeapi => ./internal/pokeapi
require internal/pokecache v1.0.0
replace internal/pokecache => ./internal/pokecache
