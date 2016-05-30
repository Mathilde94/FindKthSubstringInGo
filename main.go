package main

import "fmt"
import "sort"

func common_prefix (a string, b string) int {
     var length int
     la := len(a)
     lb := len(b)
     for i,j := 0, 0; i<la && j<lb; i, j = i+1, j+1 {
     	  if a[i] != b[j] {
	      break
	  }
	  length ++
     }
     return length
}

func findRangeForIndex (indexes []int, k int, lower int, upper int) int {
     // find index where indexes[i] >=k and indexes[i-1]<k
     if k == indexes[upper] {
     	return upper
     }
     if k > indexes[upper] {
       return -1
     }
     if k <= indexes[lower] {
     	return lower
     }
     if (upper == lower) {
     	if k <= upper {
	   return upper
	   } else {
	     return upper + 1
	     }
     }
     middle := (upper + lower) / 2
     if k == indexes[middle] {
     	return middle
     }
     if k <= indexes[middle] {
     	return findRangeForIndex(indexes, k, lower, middle)
     } else {
       return findRangeForIndex(indexes, k, middle+1, upper)
     }
}

func main () {

     var n int
     fmt.Scan(&n)
     strings := make([]string, n)
     for i:=0; i<n; i++ {
     	  fmt.Scan(&strings[i])
     }

     // Let's build the suffix array of ALL the suffixes
     suffixes := make([]string, 1)
     for _, string := range strings {
     	  for i:=0; i< len(string); i++ {
	      suffixes = append(suffixes, string[i: len(string)])
	       }
     }
     sort.Strings(suffixes)     

     lengths := make([]int, len(suffixes))
     lcps := make([]int, len(suffixes))
     indexes := make([]int, len(suffixes))

     for i:=1; i<len(lengths); i++ {
     	  lengths[i] = len(suffixes[i])
	   if i>0 {
	      lcps[i] = common_prefix(suffixes[i-1], suffixes[i])
	      indexes[i] = indexes[i-1] + lengths[i] - lcps[i]
	   } else {
	     indexes[i] = lengths[i]
	   }
     }

     var Q int
     fmt.Scan(&Q)
     var query int

     for k:=0; k<Q; k++ {
     	  fmt.Scan(&query)
	  starting_string_index := findRangeForIndex(indexes, query, 0, len(indexes)-1)
	  if starting_string_index == -1 {
              fmt.Println("INVALID")
	  } else {		
	      final_index := lcps[starting_string_index] + query - indexes[starting_string_index-1]
	      fmt.Println(suffixes[starting_string_index][0:final_index])
          }
     }
}