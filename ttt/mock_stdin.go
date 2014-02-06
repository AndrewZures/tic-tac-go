package ttt

import "io"

type MockStdIO struct {
  reader io.Reader;
}

