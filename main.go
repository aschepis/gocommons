//
//  main.go
//
//  Created by Hicham Bouabdallah
//  Copyright (c) 2012 SimpleRocket LLC
//
//  Permission is hereby granted, free of charge, to any person
//  obtaining a copy of this software and associated documentation
//  files (the "Software"), to deal in the Software without
//  restriction, including without limitation the rights to use,
//  copy, modify, merge, publish, distribute, sublicense, and/or sell
//  copies of the Software, and to permit persons to whom the
//  Software is furnished to do so, subject to the following
//  conditions:
//
//  The above copyright notice and this permission notice shall be
//  included in all copies or substantial portions of the Software.
//
//  THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND,
//  EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES
//  OF MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND
//  NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT
//  HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY,
//  WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING
//  FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR
//  OTHER DEALINGS IN THE SOFTWARE.
//

package main

import "fmt"
import "github.com/hishboy/gocommons/lang"

func main() { 
	queue := lang.NewQueue()
	queue.Push("Hello")
	queue.Push(4)
	queue.Push(5)
	queue.Push(8)
	queue.Push(6)
	queue.Push("World")
	fmt.Println("Total items before poll: ", queue.Len())
	fmt.Println(queue.Poll(), queue.Poll(), queue.Poll(), queue.Poll(), queue.Poll(), queue.Poll(), queue.Poll())
	fmt.Println("Total items after poll: ", queue.Len())
	
	fmt.Println("\n")
	
	stack := lang.NewStack()
	stack.Push("World")
	stack.Push(4)
	stack.Push(5)
	stack.Push(8)
	stack.Push(6)
	stack.Push("Hello")
	fmt.Println("Total items before pop: ", stack.Len())
	fmt.Println(stack.Pop(), stack.Pop(), stack.Pop(), stack.Pop(), stack.Pop(), stack.Pop(), stack.Pop())
	fmt.Println("Total items after pop: ", stack.Len())
	
	
	array := lang.NewArrayList()
	array.Add("hello")
	array.Add("world")
	array.Add("Hicham")
	fmt.Println("total items in array:", array.Len())
	fmt.Println("array:", array)
	fmt.Println("array[1]:", array.Get(1))
	fmt.Println("array.indexOf(world):", array.IndexOf("world"))
	fmt.Println("array.Contains(world):", array.Contains("world"), array.Contains("random"))
	
	for i := 0; i < array.Len(); i++ {
		fmt.Println("item(", i, "):", array.Get(i))
	}
	
	fmt.Println("array.Remove(world):", array.Remove("world"), array)
	fmt.Println("array.Remove(Hicham):", array.Remove("Hicham"), array)
	fmt.Println("array.IsEmpty():", array.IsEmpty(), array.Remove("hello"), array.IsEmpty(), array.Len())
	
	for i := 0; i < 30; i++ {
		array.Add(i)
	}
	
	fmt.Println("array.indexOf(5):", array.IndexOf(5), array.Len())
	array.Clear()
	array.Add("one", "two", "three")
	
	array2 := lang.NewArrayList()
	array2.Add("four", "five", "six")
	
	array.AddFromArrayList(array2)
	
	fmt.Println("array:", array, array.Len())
	fmt.Println("array.First/Last:", array.First(), array.Last())
	fmt.Println("array.ToSlice:", array.ToSlice())
	
	
	set := lang.NewHashSet()
	set.Add("hello", "world", "hello")
	fmt.Println("total items in set:", set.Len(), set)
	setToSlice := set.ToSlice()
	fmt.Println("set.ToSlice:", setToSlice, len(setToSlice))
	
}