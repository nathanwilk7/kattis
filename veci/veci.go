package main

import (
       "errors"
       "fmt"
       "log"
       "strconv"
)

func main() {
     var input string
     fmt.Scanf("%s", &input)
     result, err := nextHighest(input)
     if err != nil {
     	fmt.Println(0)
     } else {
       fmt.Println(result)
     }
}

func nextHighest (n string) (int, error) {
     nArr, err := toIntSlice(n)
     if err != nil {
     	log.Fatal(err)
     }
     soFar := make([]int, 10)
     maxSoFar := -1
     for i := len(nArr) - 1; i >= 0; i-- {
     	 cur := nArr[i]
	 soFar[cur]++
	 if cur > maxSoFar {
	    maxSoFar = cur
	 } else if cur < maxSoFar {
	   replacement, err := minGreater(soFar, cur)
	   nArr[i] = replacement
	   if err != nil {
	      log.Fatal(err)
	   }
	   soFar[nArr[i]]--
	   sortedEnd := sortedSlice(soFar)
	   finalArr := combine(nArr[0:i+1], sortedEnd)
	   return arrToInt(finalArr), nil
	 }
     }
     return 0, errors.New("No greater int possible")
}

func arrToInt (intArr []int) int {
     res := 0
     factor := 1
     for i := len(intArr) - 1; i >= 0; i-- {
     	 res += intArr[i] * int(factor)
	 factor *= 10
     }
     return res
}

func combine (numArr []int, end []int) []int {
     res := []int{}
     for _, v := range numArr {
     	 res = append(res, v)
     }
     for _, v := range end {
     	 res = append(res, v)
     }
     return res
}

func sortedSlice (soFar []int) []int {
     s := []int{}
     for i, v := range soFar {
     	 for v != 0 {
	     s = append(s, i)
	     v--
	 }
     }
     return s
}

func minGreater (soFar []int, n int) (int, error) {
     for i, v := range soFar {
     	 if v != 0 && i > n {
	    return i, nil
	 }
     }
     return -1, errors.New(fmt.Sprintf("soFar %v does not have a val greater than %d", soFar, n))
}

func toIntSlice (n string) ([]int, error) {
     arr := []int{}
     for i := 0; i < len(n); i++ {
     	 j, err := strconv.Atoi(string(n[i]))
	 if err != nil {
	    return nil, err
	 }
     	 arr = append(arr, j)
     }     
     return arr, nil
}