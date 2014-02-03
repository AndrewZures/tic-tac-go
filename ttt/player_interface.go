package ttt

type Player interface {
    MakeMove([]int) int;
    Symbol() string;
}
