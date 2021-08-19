package counters

type StepCounter struct {
	Value int64
	Step  int64
}

func NewStepCounter(value, step int64) StepCounter {
	return StepCounter{
		Value: value,
		Step:  step,
	}
}

func (sc *StepCounter) Increment() {
	sc.Value += sc.Step
}
func (sc *StepCounter) Decrement() {
	sc.Value -= sc.Step
}
func (sc *StepCounter) GetValue() int64 {
	return sc.Value
}
