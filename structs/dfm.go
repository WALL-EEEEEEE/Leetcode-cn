package structs

type State int
type StateTran func(this *DFM, value interface{}) State

type DFM struct {
	state      State
	states_tbl map[State]StateTran
	extra      map[string]interface{}
}

func NewDFM(state State, states_tbl map[State]StateTran, extra map[string]interface{}) *DFM {
	dfm := &DFM{
		state:      state,
		states_tbl: states_tbl,
		extra:      extra,
	}
	return dfm
}

func (this *DFM) GetState() State {
	return this.state
}

func (this *DFM) setState(state State) {
	this.state = state
}

func (this *DFM) Step(value rune) {
	stateTran := this.states_tbl[this.state]
	this.setState(stateTran(this, value))
}

func (this *DFM) GetData(k string) interface{} {
	return this.extra[k]
}

func (this *DFM) SetData(k string, v interface{}) {
	this.extra[k] = v
}
