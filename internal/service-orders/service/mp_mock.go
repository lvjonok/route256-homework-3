package service

// Code generated by http://github.com/gojuno/minimock (dev). DO NOT EDIT.

//go:generate minimock -i gitlab.ozon.dev/lvjonok/homework-3/internal/service-orders/service.MarketplaceClient -o ./mp_mock.go -n MarketplaceClientMock

import (
	"context"
	"sync"
	mm_atomic "sync/atomic"
	mm_time "time"

	"github.com/gojuno/minimock/v3"
	mpAPI "gitlab.ozon.dev/lvjonok/homework-3/pkg/srv_marketplace/api"
)

// MarketplaceClientMock implements MarketplaceClient
type MarketplaceClientMock struct {
	t minimock.Tester

	funcGetCart          func(ctx context.Context, gp1 *mpAPI.GetCartRequest) (gp2 *mpAPI.GetCartResponse, err error)
	inspectFuncGetCart   func(ctx context.Context, gp1 *mpAPI.GetCartRequest)
	afterGetCartCounter  uint64
	beforeGetCartCounter uint64
	GetCartMock          mMarketplaceClientMockGetCart
}

// NewMarketplaceClientMock returns a mock for MarketplaceClient
func NewMarketplaceClientMock(t minimock.Tester) *MarketplaceClientMock {
	m := &MarketplaceClientMock{t: t}
	if controller, ok := t.(minimock.MockController); ok {
		controller.RegisterMocker(m)
	}

	m.GetCartMock = mMarketplaceClientMockGetCart{mock: m}
	m.GetCartMock.callArgs = []*MarketplaceClientMockGetCartParams{}

	return m
}

type mMarketplaceClientMockGetCart struct {
	mock               *MarketplaceClientMock
	defaultExpectation *MarketplaceClientMockGetCartExpectation
	expectations       []*MarketplaceClientMockGetCartExpectation

	callArgs []*MarketplaceClientMockGetCartParams
	mutex    sync.RWMutex
}

// MarketplaceClientMockGetCartExpectation specifies expectation struct of the MarketplaceClient.GetCart
type MarketplaceClientMockGetCartExpectation struct {
	mock    *MarketplaceClientMock
	params  *MarketplaceClientMockGetCartParams
	results *MarketplaceClientMockGetCartResults
	Counter uint64
}

// MarketplaceClientMockGetCartParams contains parameters of the MarketplaceClient.GetCart
type MarketplaceClientMockGetCartParams struct {
	ctx context.Context
	gp1 *mpAPI.GetCartRequest
}

// MarketplaceClientMockGetCartResults contains results of the MarketplaceClient.GetCart
type MarketplaceClientMockGetCartResults struct {
	gp2 *mpAPI.GetCartResponse
	err error
}

// Expect sets up expected params for MarketplaceClient.GetCart
func (mmGetCart *mMarketplaceClientMockGetCart) Expect(ctx context.Context, gp1 *mpAPI.GetCartRequest) *mMarketplaceClientMockGetCart {
	if mmGetCart.mock.funcGetCart != nil {
		mmGetCart.mock.t.Fatalf("MarketplaceClientMock.GetCart mock is already set by Set")
	}

	if mmGetCart.defaultExpectation == nil {
		mmGetCart.defaultExpectation = &MarketplaceClientMockGetCartExpectation{}
	}

	mmGetCart.defaultExpectation.params = &MarketplaceClientMockGetCartParams{ctx, gp1}
	for _, e := range mmGetCart.expectations {
		if minimock.Equal(e.params, mmGetCart.defaultExpectation.params) {
			mmGetCart.mock.t.Fatalf("Expectation set by When has same params: %#v", *mmGetCart.defaultExpectation.params)
		}
	}

	return mmGetCart
}

// Inspect accepts an inspector function that has same arguments as the MarketplaceClient.GetCart
func (mmGetCart *mMarketplaceClientMockGetCart) Inspect(f func(ctx context.Context, gp1 *mpAPI.GetCartRequest)) *mMarketplaceClientMockGetCart {
	if mmGetCart.mock.inspectFuncGetCart != nil {
		mmGetCart.mock.t.Fatalf("Inspect function is already set for MarketplaceClientMock.GetCart")
	}

	mmGetCart.mock.inspectFuncGetCart = f

	return mmGetCart
}

// Return sets up results that will be returned by MarketplaceClient.GetCart
func (mmGetCart *mMarketplaceClientMockGetCart) Return(gp2 *mpAPI.GetCartResponse, err error) *MarketplaceClientMock {
	if mmGetCart.mock.funcGetCart != nil {
		mmGetCart.mock.t.Fatalf("MarketplaceClientMock.GetCart mock is already set by Set")
	}

	if mmGetCart.defaultExpectation == nil {
		mmGetCart.defaultExpectation = &MarketplaceClientMockGetCartExpectation{mock: mmGetCart.mock}
	}
	mmGetCart.defaultExpectation.results = &MarketplaceClientMockGetCartResults{gp2, err}
	return mmGetCart.mock
}

//Set uses given function f to mock the MarketplaceClient.GetCart method
func (mmGetCart *mMarketplaceClientMockGetCart) Set(f func(ctx context.Context, gp1 *mpAPI.GetCartRequest) (gp2 *mpAPI.GetCartResponse, err error)) *MarketplaceClientMock {
	if mmGetCart.defaultExpectation != nil {
		mmGetCart.mock.t.Fatalf("Default expectation is already set for the MarketplaceClient.GetCart method")
	}

	if len(mmGetCart.expectations) > 0 {
		mmGetCart.mock.t.Fatalf("Some expectations are already set for the MarketplaceClient.GetCart method")
	}

	mmGetCart.mock.funcGetCart = f
	return mmGetCart.mock
}

// When sets expectation for the MarketplaceClient.GetCart which will trigger the result defined by the following
// Then helper
func (mmGetCart *mMarketplaceClientMockGetCart) When(ctx context.Context, gp1 *mpAPI.GetCartRequest) *MarketplaceClientMockGetCartExpectation {
	if mmGetCart.mock.funcGetCart != nil {
		mmGetCart.mock.t.Fatalf("MarketplaceClientMock.GetCart mock is already set by Set")
	}

	expectation := &MarketplaceClientMockGetCartExpectation{
		mock:   mmGetCart.mock,
		params: &MarketplaceClientMockGetCartParams{ctx, gp1},
	}
	mmGetCart.expectations = append(mmGetCart.expectations, expectation)
	return expectation
}

// Then sets up MarketplaceClient.GetCart return parameters for the expectation previously defined by the When method
func (e *MarketplaceClientMockGetCartExpectation) Then(gp2 *mpAPI.GetCartResponse, err error) *MarketplaceClientMock {
	e.results = &MarketplaceClientMockGetCartResults{gp2, err}
	return e.mock
}

// GetCart implements MarketplaceClient
func (mmGetCart *MarketplaceClientMock) GetCart(ctx context.Context, gp1 *mpAPI.GetCartRequest) (gp2 *mpAPI.GetCartResponse, err error) {
	mm_atomic.AddUint64(&mmGetCart.beforeGetCartCounter, 1)
	defer mm_atomic.AddUint64(&mmGetCart.afterGetCartCounter, 1)

	if mmGetCart.inspectFuncGetCart != nil {
		mmGetCart.inspectFuncGetCart(ctx, gp1)
	}

	mm_params := &MarketplaceClientMockGetCartParams{ctx, gp1}

	// Record call args
	mmGetCart.GetCartMock.mutex.Lock()
	mmGetCart.GetCartMock.callArgs = append(mmGetCart.GetCartMock.callArgs, mm_params)
	mmGetCart.GetCartMock.mutex.Unlock()

	for _, e := range mmGetCart.GetCartMock.expectations {
		if minimock.Equal(e.params, mm_params) {
			mm_atomic.AddUint64(&e.Counter, 1)
			return e.results.gp2, e.results.err
		}
	}

	if mmGetCart.GetCartMock.defaultExpectation != nil {
		mm_atomic.AddUint64(&mmGetCart.GetCartMock.defaultExpectation.Counter, 1)
		mm_want := mmGetCart.GetCartMock.defaultExpectation.params
		mm_got := MarketplaceClientMockGetCartParams{ctx, gp1}
		if mm_want != nil && !minimock.Equal(*mm_want, mm_got) {
			mmGetCart.t.Errorf("MarketplaceClientMock.GetCart got unexpected parameters, want: %#v, got: %#v%s\n", *mm_want, mm_got, minimock.Diff(*mm_want, mm_got))
		}

		mm_results := mmGetCart.GetCartMock.defaultExpectation.results
		if mm_results == nil {
			mmGetCart.t.Fatal("No results are set for the MarketplaceClientMock.GetCart")
		}
		return (*mm_results).gp2, (*mm_results).err
	}
	if mmGetCart.funcGetCart != nil {
		return mmGetCart.funcGetCart(ctx, gp1)
	}
	mmGetCart.t.Fatalf("Unexpected call to MarketplaceClientMock.GetCart. %v %v", ctx, gp1)
	return
}

// GetCartAfterCounter returns a count of finished MarketplaceClientMock.GetCart invocations
func (mmGetCart *MarketplaceClientMock) GetCartAfterCounter() uint64 {
	return mm_atomic.LoadUint64(&mmGetCart.afterGetCartCounter)
}

// GetCartBeforeCounter returns a count of MarketplaceClientMock.GetCart invocations
func (mmGetCart *MarketplaceClientMock) GetCartBeforeCounter() uint64 {
	return mm_atomic.LoadUint64(&mmGetCart.beforeGetCartCounter)
}

// Calls returns a list of arguments used in each call to MarketplaceClientMock.GetCart.
// The list is in the same order as the calls were made (i.e. recent calls have a higher index)
func (mmGetCart *mMarketplaceClientMockGetCart) Calls() []*MarketplaceClientMockGetCartParams {
	mmGetCart.mutex.RLock()

	argCopy := make([]*MarketplaceClientMockGetCartParams, len(mmGetCart.callArgs))
	copy(argCopy, mmGetCart.callArgs)

	mmGetCart.mutex.RUnlock()

	return argCopy
}

// MinimockGetCartDone returns true if the count of the GetCart invocations corresponds
// the number of defined expectations
func (m *MarketplaceClientMock) MinimockGetCartDone() bool {
	for _, e := range m.GetCartMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			return false
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.GetCartMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterGetCartCounter) < 1 {
		return false
	}
	// if func was set then invocations count should be greater than zero
	if m.funcGetCart != nil && mm_atomic.LoadUint64(&m.afterGetCartCounter) < 1 {
		return false
	}
	return true
}

// MinimockGetCartInspect logs each unmet expectation
func (m *MarketplaceClientMock) MinimockGetCartInspect() {
	for _, e := range m.GetCartMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			m.t.Errorf("Expected call to MarketplaceClientMock.GetCart with params: %#v", *e.params)
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.GetCartMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterGetCartCounter) < 1 {
		if m.GetCartMock.defaultExpectation.params == nil {
			m.t.Error("Expected call to MarketplaceClientMock.GetCart")
		} else {
			m.t.Errorf("Expected call to MarketplaceClientMock.GetCart with params: %#v", *m.GetCartMock.defaultExpectation.params)
		}
	}
	// if func was set then invocations count should be greater than zero
	if m.funcGetCart != nil && mm_atomic.LoadUint64(&m.afterGetCartCounter) < 1 {
		m.t.Error("Expected call to MarketplaceClientMock.GetCart")
	}
}

// MinimockFinish checks that all mocked methods have been called the expected number of times
func (m *MarketplaceClientMock) MinimockFinish() {
	if !m.minimockDone() {
		m.MinimockGetCartInspect()
		m.t.FailNow()
	}
}

// MinimockWait waits for all mocked methods to be called the expected number of times
func (m *MarketplaceClientMock) MinimockWait(timeout mm_time.Duration) {
	timeoutCh := mm_time.After(timeout)
	for {
		if m.minimockDone() {
			return
		}
		select {
		case <-timeoutCh:
			m.MinimockFinish()
			return
		case <-mm_time.After(10 * mm_time.Millisecond):
		}
	}
}

func (m *MarketplaceClientMock) minimockDone() bool {
	done := true
	return done &&
		m.MinimockGetCartDone()
}
