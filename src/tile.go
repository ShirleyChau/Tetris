package main

import "math/rand"

type BaseTile interface {
	init()
	status(rudder int)
}
type Tile struct {
	graph  [4][4]bool
	rudder int
}
type L struct { //L，
	Tile
}

func (tile *L) status(rudder int) {
	switch rudder {
	case 0:
		tile.graph = [4][4]bool{
			{false, true, false, false},
			{false, true, false, false},
			{false, true, true, false},
			{false, false, false, false},
		}
	case 1:
		tile.graph = [4][4]bool{
			{false, false, true, false},
			{true, true, true, false},
			{false, false, false, false},
			{false, false, false, false},
		}
	case 2:
		tile.graph = [4][4]bool{
			{false, true, true, false},
			{false, false, true, false},
			{false, false, true, false},
			{false, false, false, false},
		}
	case 3:
		tile.graph = [4][4]bool{
			{true, true, true, false},
			{false, false, true, false},
			{false, false, false, false},
			{false, false, false, false},
		}
	default:
		panic("LL error status!")
	}
}
func (tile *L) init() {
	tile.rudder = rand.Intn(4)
	tile.status(tile.rudder)
}

type T struct { //田
	Tile
}

func (tile *T) init() {
	tile.rudder = rand.Intn(4)
}
func (tile *T) status(rudder int) {
	tile.graph = [4][4]bool{
		{false, true, true, false},
		{false, true, true, false},
		{false, false, false, false},
		{false, false, false, false},
	}
}

type I struct { // |
	Tile
}

func (tile *I) status(rudder int) {
	if rudder%2 == 0 {
		tile.graph = [4][4]bool{
			{false, true, false, false},
			{false, true, false, false},
			{false, true, false, false},
			{false, true, false, false},
		}
	} else {
		tile.graph = [4][4]bool{
			{true, true, true, true},
			{false, false, false, false},
			{false, false, false, false},
			{false, false, false, false},
		}
	}
}
func (tile *I) init() {
	tile.rudder = rand.Intn(4)
	tile.status(tile.rudder)
}

type 土 struct {
	Tile
}
func (tile *土)status(rudder int){
	switch rudder {
	case 0:
		tile.graph = [4][4]bool{
			{false, true, false, false},
			{true, true, true, false},
			{false, false, false, false},
			{false, false, false, false},
		}
	case 1:
		tile.graph = [4][4]bool{
			{false, true, false, false},
			{false, true, true, false},
			{false, true, false, false},
			{false, false, false, false},
		}
	case 2:
		tile.graph = [4][4]bool{
			{false, false, false, false},
			{true, true, true, false},
			{false, true, false, false},
			{false, false, false, false},
		}
	case 3:
		tile.graph = [4][4]bool{
			{false, true, false, false},
			{true, true, false, false},
			{false, true, false, false},
			{false, false, false, false},
		}
	default:
		panic("LL error status!")
	}
}
func (tile *土)init(){
	tile.rudder = rand.Intn(4)
	tile.status(tile.rudder)
}
