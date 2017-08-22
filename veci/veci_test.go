package main

import (
       "testing"
)

func Test_12 (t *testing.T) {
     str := "12"
     answer := 21
     res, err := nextHighest(str)
     if err != nil {
     	t.Fail()
     }
     if res != answer {
     	t.Fail()
     }
}

func Test_231 (t *testing.T) {
     str := "231"
     answer := 312
     res, err := nextHighest(str)
     if err != nil {
     	t.Fail()
     }
     if res != answer {
     	t.Fail()
     }
}

func Test_1432 (t *testing.T) {
     str := "1432"
     answer := 2134
     res, err := nextHighest(str)
     if err != nil {
     	t.Fail()
     }
     if res != answer {
     	t.Fail()
     }
}