//doc https://www.youtube.com/watch?v=3DYIgTC4T1o&list=PL0q7mDmXPZm7s7weikYLpNZBKk5dCoWm6

package main
import (
	"fmt"
)


	//[obj] holds
		type MaxHeap struct{
			array []int
		}
	//[sub]
	//	insert
		func (h *MaxHeap) Insert(key int){
			h.array = append(h.array,key) //insert the value to the array?
			h.maxHeapifyUp(len(h.array)-1) //this supplies this func with the last index?

		}
	//	maxHeapifyUp
		func (h *MaxHeap) maxHeapifyUp(index int){
			for h.array[parent(index)] < h.array[index]{
				h.swap(parent(index),index)
				index = parent(index)
			}
		}
	//	get the parent index
		func parent(i int) int{
			return (i-1)/2 //search the position of the parent node
			//      		1[0]
			//		10[1]			20[2] <- position is (parent index) * 2 +2
			//	5[3] <- position is (parent index) * 2 +1
		}
	//	left
		func left(i int) int{
			return 2*i +1
		}
	//	right 
		func right(i int) int{
			return 2*i +1
		}
	//	swap
		func (h *MaxHeap) swap(i1,i2 int){
			h.array[i1], h.array[i2] = h.array[i2], h.array[i1]
		}


	//Extract
	func main() {
		m := &MaxHeap{}
		fmt.Println(m)

	}









