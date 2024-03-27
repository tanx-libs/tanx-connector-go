package client

import "fmt"

var (
	ErrMarketNotProvided     = fmt.Errorf("provide a market")
	ErrInvalidAmount         = fmt.Errorf("amount should be greater than 0")
	ErrNotLoggedIn           = fmt.Errorf("not logged in, please login first as this is a protected endpoint")
	ErrCoinNotFound          = fmt.Errorf("coin not found")
	ErrInvalidNetwork        = fmt.Errorf("invalid network")
	ErrAlreadyLoggedIn       = fmt.Errorf("already logged in")
)



// json decoding error
type ErrJSONDecoding struct {
	Err error
}

func (e *ErrJSONDecoding) Error() string {
	return fmt.Sprintf("error: %s\n", e.Err)
}

// 4xx error
type ErrClient struct {
	Status int
	Err error
}

func (e *ErrClient) Error() string {
	return fmt.Sprintf("status: %d\nerror: %s\n", e.Status, e.Err)
}

// 5xx error
type ErrServer struct {
	Status int
	Err error
}

func (e *ErrServer) Error() string {
	return fmt.Sprintf("status: %d\nerror: %s\n", e.Status, e.Err)
}


// ErrInsufficientAllowance error
type ErrInsufficientAllowance struct {
	Allowance float64
	Amount float64
	Err error
}

func (e *ErrInsufficientAllowance) Error() string {
	return fmt.Sprintf("current allowance of %v which is less than deposit amount %v. call SetAllowance function to change the allowance amount", e.Allowance, e.Amount)
}
