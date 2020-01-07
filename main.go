package main

import "fmt"

func findShortestSubArray(nums []int) int {
  numMap := make([]int, 49999)
  inDegree := 0

  for _, num := range nums {
    numMap[num]++

    if numMap[num] > inDegree {
      inDegree = numMap[num]
    }
  }

  mSet := make(map[int]bool, len(numMap))

  for num, degree := range numMap {
    if degree == inDegree {
      mSet[num] = true
    }
  }

  fmt.Printf("mSet := %v\n", mSet)
  start := 0
  end := len(nums) - 1

  for start < end {

    _, partOfDegrees := mSet[nums[end]]

    if !partOfDegrees {
      end--
      continue
    }

    if partOfDegrees && len(mSet) > 1 {
      delete(mSet, nums[end])
      end--
      continue
    }

    _, partOfDegrees = mSet[nums[start]]

    if !partOfDegrees {
      start++
      continue
    }

    if partOfDegrees && len(mSet) > 1 {
      delete(mSet, nums[start])
      start++
      continue
    }

    break

  }

  fmt.Printf("start %d end %d\n", start, end)
  fmt.Printf("inDegree %d\n", inDegree)
  return (end-start) + 1
}

func main() {
  fmt.Printf("shortest sub array %d\n", findShortestSubArray([]int{2,1,1,2,1,3,3,3,1,3,1,3,2}))
}