package pie

import (
	"fmt"
)

var (
	pies  = NewPies()
	Debug = false
)

type Pies struct {
	setc chan interface{}

	stateIdx int
	states   []interface{}

	refIdx int
	refs   []interface{}

	effectIdx int
	effects   [][]interface{}

	mounted bool
}

func NewPies() *Pies {
	println("NewPies")
	return &Pies{
		setc: make(chan interface{}),
	}
}

func Mount(fn func()) {
	pies.Mount(fn)
}
func (p *Pies) Mount(fn func()) {
	println("Mount")

	fn()

	p.mounted = true
	println("Mount: initial finished --------------------------------")

	for v := range p.setc {
		println("Mount: incoming update %v", v)

		p.stateIdx = 0
		p.refIdx = 0
		p.effectIdx = 0

		fn()

		println("Mount: re-render finished --------------------------------")
	}
}

func (p *Pies) set(v interface{}) {
	println("set: %v", v)
	p.setc <- v
}

func IntVar(i int) (int, func(int)) {
	return pies.IntVar(i)
}
func (p *Pies) IntVar(i int) (int, func(int)) {
	idx := p.stateIdx

	println("IntVar@%v", idx)

	if !p.mounted {
		println("IntVar@%v: initial %v", idx, i)
		p.states = append(p.states, i)
	}

	i2 := p.states[idx].(int)
	p.stateIdx++

	return i2, func(i2 int) {
		i := p.states[idx].(int)

		if i == i2 {
			println("IntVar@%v: ignoring; already %v", idx, i2)
			return
		}

		println("IntVar@%v: %v -> %v", idx, i, i2)
		p.states[idx] = i2

		go p.set(i2)
	}
}

func IntPtr(i int) *int {
	return pies.IntPtr(i)
}
func (p *Pies) IntPtr(i int) *int {
	idx := p.refIdx

	println("IntPtr@%v", idx)

	if !p.mounted {
		println("IntPtr@%v: initial %v", idx, i)
		p.refs = append(p.refs, &i)
	}

	i2 := p.refs[idx].(*int)
	p.refIdx++

	return i2
}

func Do(fn func(), deps ...interface{}) {
	pies.Do(fn, deps...)
}
func (p *Pies) Do(fn func(), deps ...interface{}) {
	idx := p.effectIdx

	println("Do@%v", idx)

	if !p.mounted {
		println("Do@%v: binding %v deps", idx, len(deps))
		p.effects = append(p.effects, deps)
	}

	oldDeps := p.effects[idx]
	p.effectIdx++

	changed := false
	for i := range oldDeps {
		changed = changed || oldDeps[i] != deps[i]
	}

	once := p.mounted && deps == nil
	if p.mounted && (once || !changed) {
		println("Do@%v: ignoring; mounted{%v} once{%v} changed{%v}", idx, p.mounted, once, !changed)
		return
	}

	p.effects[idx] = deps

	go fn()
}

func println(format string, a ...interface{}) {
	if !Debug {
		return
	}
	fmt.Printf("[pie] "+format+"\n", a...)
}
