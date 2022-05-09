package object_pool

import (
	"strconv"
	"testing"
)

func Test_Pool(t *testing.T) {
	poolLen := 3
	poolObjectArr := make([]iPoolObject, 0)

	for i := 0; i < poolLen; i++ {
		c := &connection{id: strconv.Itoa(i)}
		poolObjectArr = append(poolObjectArr, c)
	}

	pool, err := InitPool(poolObjectArr)
	if err != nil {
		t.Errorf("Init Pool Error: %s", err)
	}

	connArr := make([]iPoolObject, 0)
	for i := 0; i < poolLen; i++ {
		conn, err := pool.Loan()
		if err != nil {
			t.Errorf("Pool Loan Error: %s", err)
		}
		
		connArr = append(connArr, conn)
	}

	_, err = pool.Loan()
	if err == nil {
		t.Error("More loans than the limited number of pools.")
	}

	pool.Receive(connArr[poolLen-1])

	_, err = pool.Loan()
	if err != nil {
		t.Errorf("Pool Loan Error: %s", err)
	}
}
