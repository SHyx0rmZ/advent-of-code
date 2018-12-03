package day06

type Operation interface{ op() }

type Toggle struct{ Area }
type TurnOn struct{ Area }
type TurnOff struct{ Area }

func (Toggle) op()  {}
func (TurnOn) op()  {}
func (TurnOff) op() {}
