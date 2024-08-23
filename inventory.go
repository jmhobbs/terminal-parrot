package main

var Animations map[string]Animation

func init() {
	Animations = make(map[string]Animation)
	parrotAnimation, err := LoadFromBytes(parrot)
	if err != nil {
		panic(err)
	}
	Animations["parrot"] = *parrotAnimation
}
