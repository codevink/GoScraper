package utils

import "sync"

func AddWaitGroup(wg *sync.WaitGroup) {
    wg.Add(1)
}

func DoneWaitGroup(wg *sync.WaitGroup) {
    wg.Done()
}
