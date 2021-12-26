package account

// Define the Account type here.

type Account struct {
	value int64
	lock chan bool
}

var accountOpen = make(map[*Account]bool)

func Open(amount int64) *Account {

	if amount < 0 {
		return nil
	}

	a := &Account{
		value: amount,
		lock: make(chan bool, 1),
	}

	accountOpen[a] = true

	return a
}

func (a *Account) Balance() (int64, bool) {

	if open := accountOpen[a]; !open {
		return a.value, false
	}

	return a.value, true
}

func (a *Account) Deposit(amount int64) (int64, bool) {


	if open := accountOpen[a]; !open {
		return a.value, false
	}

	ok, bal := true, int64(0)

	a.lock <- true

	if a.value + amount < 0 {
		ok = false
	} else {
		a.value += amount
		bal = a.value
	}

	<-a.lock

	return bal, ok
}

func (a *Account) Close() (int64, bool) {

	// once account is closed, it won't be openend again => no need to lock map
	// but, it throws << fatal error: concurrent map read and map write >>
	// if open := accountOpen[a]; !open {
	// 	return a.value, false
	// }

	bal, ok := int64(0), true

	a.lock <- true
	bal = a.value

	if open := accountOpen[a]; open {
		ok = true
		accountOpen[a] = false
		a.value = 0
	} else {
		ok = false
	}
	<- a.lock

	return bal, ok
}
