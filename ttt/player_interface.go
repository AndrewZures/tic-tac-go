package ttt

type Player interface {
    MakeMove([]string) int;
    Symbol() string;
}
