package main

func main() {
	eng := Engine{}
	eng.Init(InitState, 0x1000)
	for i := 0; i < Threads; i++ {
		go eng.Thread()
	}
	eng.Wait()
	for InitState.Final() == false {
		InitState = eng.Solve(InitState)
	}
}
