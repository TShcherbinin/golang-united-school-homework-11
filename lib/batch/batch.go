package batch

import (
	"math"
	"time"
)

type user struct {
	ID int64
}

func getOne(id int64) user {
	time.Sleep(time.Millisecond * 100)
	return user{ID: id}
}

func worker(idStart int64, idEnd int64, ch chan user) {
	for i := idStart; i < idEnd; i++ {
		ch <- getOne(i)
	}
}

func getBatch(n int64, pool int64) (res []user) {
	ch := make(chan user, pool)
	var remainUsers int64 = n
	var userPerPool int64 = int64(math.Ceil(float64(n) / float64(pool)))

	var i int64
	for i = 1; i <= pool; i++ {
		idStart := (i - 1) * userPerPool
		idEnd := idStart + userPerPool
		if remainUsers < userPerPool {
			idEnd = idStart + remainUsers
		}
		remainUsers -= userPerPool

		go worker(idStart, idEnd, ch)
	}

	res1 := make([]user, n, n)
	for i = 0; i < n; i++ {
		res1[i] = <-ch
	}

	return res1
}
