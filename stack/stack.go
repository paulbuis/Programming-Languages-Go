package stack

import (
    "errors"
)

type Stack interface {
    push(int)
    pop() (int, error)
    peek() (int, error)
}

type SliceStack struct {
    s []int
}

func New() SliceStack {
    ss := SliceStack{s: make([]int,0, 1)}
    return ss
}

func (ss *SliceStack) push(item int) {
    ss.s = append(ss.s, item)
}

func (ss *SliceStack) pop() (int, error) {
    if len(ss.s) == 0 {
        return 0, errors.New("empty")
    }
    top := ss.s[len(ss.s)-1]
    ss.s = ss.s[:len(ss.s)-1]
    return top, nil
}

func (ss SliceStack) peek() (int, error) {
    if len(ss.s) == 0 {
        return 0, errors.New("empty")
    }
    return ss.s[len(ss.s)-1], nil
}
