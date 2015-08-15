package gokatas

import (
	"fmt"
	"math/rand"
	"testing"
)

func TestGoroutine0(t *testing.T) {
	c := make(chan int)
	go func() {
		c <- 1
	}()
	if 1 != <-c {
		t.Fail()
	}
}

func TestGoroutine1(t *testing.T) {
	c := make(chan int)
	go func(c chan int) {
		c <- 1
	}(c)
	if 1 != <-c {
		t.Fail()
	}
}

func TestGoroutine2(t *testing.T) {
	expected := [5]int{1, 2, 3, 4, 5}
	c := make(chan int, 2)
	go func() {
		for _, i := range expected {
			c <- i
		}
	}()
	for _, i := range expected {
		if i != <-c {
			t.Fail()
		}
	}
}

func TestGoroutine3(t *testing.T) {
	expected := [5]int{1, 2, 3, 4, 5}
	c := make(chan int, 2)
	go func() {
		for _, i := range expected {
			c <- i
		}
	}()
	go func() {
		for _, i := range expected {
			if i != <-c {
				t.Fail()
			}
		}
	}()
}

func TestGoroutine4(t *testing.T) {

	expected := rand.Perm(100)

	c0 := make(chan (chan<- int), 2)
	c1 := make(chan (<-chan int), 2)

	go func(ci chan (chan<- int)) {
		c := <-ci
		for _, i := range expected {
			c <- i
		}
	}(c0)

	go func(ci chan (<-chan int)) {
		c := <-ci
		for _, i := range expected {
			if i != <-c {
				t.Fail()
			}
		}
	}(c1)

	ci := make(chan int, 2)
	c0 <- ci
	c1 <- ci
}

func TestGoroutine5(t *testing.T) {

	expected := [5]int{1, 2, 3, 4, 5}

	ci1 := make(chan int, 1)
	ci2 := make(chan int, 1)
	c1 := make(chan bool)
	c2 := make(chan bool)

	go func() {
		for _, i := range expected {
			fmt.Printf("Got %d from A\n", i)
			if i != <-ci1 {
				t.Fail()
			}
		}
		c1 <- true
	}()

	go func() {
		for _, i := range expected {
			fmt.Printf("Got %d from B\n", i)
			if i != <-ci2 {
				t.Fail()
			}
		}
		c2 <- true
	}()

	for _, i := range expected {
		ci1 <- i
		ci2 <- i
	}

	for count := 0; count < 2; {
		select {
		case <-c1:
			count++
		case <-c2:
			count++
		}
	}
}
