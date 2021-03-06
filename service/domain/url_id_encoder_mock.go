package domain

// Code generated by http://github.com/gojuno/minimock (dev). DO NOT EDIT.

import (
	"sync"
	mm_atomic "sync/atomic"
	mm_time "time"

	"github.com/gojuno/minimock/v3"
)

// UrlIDEncoderMock implements urlIDEncoder
type UrlIDEncoderMock struct {
	t minimock.Tester

	funcDecode          func(str string) (u1 uint64, err error)
	inspectFuncDecode   func(str string)
	afterDecodeCounter  uint64
	beforeDecodeCounter uint64
	DecodeMock          mUrlIDEncoderMockDecode

	funcEncode          func(id uint64) (s1 string, err error)
	inspectFuncEncode   func(id uint64)
	afterEncodeCounter  uint64
	beforeEncodeCounter uint64
	EncodeMock          mUrlIDEncoderMockEncode
}

// NewUrlIDEncoderMock returns a mock for urlIDEncoder
func NewUrlIDEncoderMock(t minimock.Tester) *UrlIDEncoderMock {
	m := &UrlIDEncoderMock{t: t}
	if controller, ok := t.(minimock.MockController); ok {
		controller.RegisterMocker(m)
	}

	m.DecodeMock = mUrlIDEncoderMockDecode{mock: m}
	m.DecodeMock.callArgs = []*UrlIDEncoderMockDecodeParams{}

	m.EncodeMock = mUrlIDEncoderMockEncode{mock: m}
	m.EncodeMock.callArgs = []*UrlIDEncoderMockEncodeParams{}

	return m
}

type mUrlIDEncoderMockDecode struct {
	mock               *UrlIDEncoderMock
	defaultExpectation *UrlIDEncoderMockDecodeExpectation
	expectations       []*UrlIDEncoderMockDecodeExpectation

	callArgs []*UrlIDEncoderMockDecodeParams
	mutex    sync.RWMutex
}

// UrlIDEncoderMockDecodeExpectation specifies expectation struct of the urlIDEncoder.Decode
type UrlIDEncoderMockDecodeExpectation struct {
	mock    *UrlIDEncoderMock
	params  *UrlIDEncoderMockDecodeParams
	results *UrlIDEncoderMockDecodeResults
	Counter uint64
}

// UrlIDEncoderMockDecodeParams contains parameters of the urlIDEncoder.Decode
type UrlIDEncoderMockDecodeParams struct {
	str string
}

// UrlIDEncoderMockDecodeResults contains results of the urlIDEncoder.Decode
type UrlIDEncoderMockDecodeResults struct {
	u1  uint64
	err error
}

// Expect sets up expected params for urlIDEncoder.Decode
func (mmDecode *mUrlIDEncoderMockDecode) Expect(str string) *mUrlIDEncoderMockDecode {
	if mmDecode.mock.funcDecode != nil {
		mmDecode.mock.t.Fatalf("UrlIDEncoderMock.Decode mock is already set by Set")
	}

	if mmDecode.defaultExpectation == nil {
		mmDecode.defaultExpectation = &UrlIDEncoderMockDecodeExpectation{}
	}

	mmDecode.defaultExpectation.params = &UrlIDEncoderMockDecodeParams{str}
	for _, e := range mmDecode.expectations {
		if minimock.Equal(e.params, mmDecode.defaultExpectation.params) {
			mmDecode.mock.t.Fatalf("Expectation set by When has same params: %#v", *mmDecode.defaultExpectation.params)
		}
	}

	return mmDecode
}

// Inspect accepts an inspector function that has same arguments as the urlIDEncoder.Decode
func (mmDecode *mUrlIDEncoderMockDecode) Inspect(f func(str string)) *mUrlIDEncoderMockDecode {
	if mmDecode.mock.inspectFuncDecode != nil {
		mmDecode.mock.t.Fatalf("Inspect function is already set for UrlIDEncoderMock.Decode")
	}

	mmDecode.mock.inspectFuncDecode = f

	return mmDecode
}

// Return sets up results that will be returned by urlIDEncoder.Decode
func (mmDecode *mUrlIDEncoderMockDecode) Return(u1 uint64, err error) *UrlIDEncoderMock {
	if mmDecode.mock.funcDecode != nil {
		mmDecode.mock.t.Fatalf("UrlIDEncoderMock.Decode mock is already set by Set")
	}

	if mmDecode.defaultExpectation == nil {
		mmDecode.defaultExpectation = &UrlIDEncoderMockDecodeExpectation{mock: mmDecode.mock}
	}
	mmDecode.defaultExpectation.results = &UrlIDEncoderMockDecodeResults{u1, err}
	return mmDecode.mock
}

//Set uses given function f to mock the urlIDEncoder.Decode method
func (mmDecode *mUrlIDEncoderMockDecode) Set(f func(str string) (u1 uint64, err error)) *UrlIDEncoderMock {
	if mmDecode.defaultExpectation != nil {
		mmDecode.mock.t.Fatalf("Default expectation is already set for the urlIDEncoder.Decode method")
	}

	if len(mmDecode.expectations) > 0 {
		mmDecode.mock.t.Fatalf("Some expectations are already set for the urlIDEncoder.Decode method")
	}

	mmDecode.mock.funcDecode = f
	return mmDecode.mock
}

// When sets expectation for the urlIDEncoder.Decode which will trigger the result defined by the following
// Then helper
func (mmDecode *mUrlIDEncoderMockDecode) When(str string) *UrlIDEncoderMockDecodeExpectation {
	if mmDecode.mock.funcDecode != nil {
		mmDecode.mock.t.Fatalf("UrlIDEncoderMock.Decode mock is already set by Set")
	}

	expectation := &UrlIDEncoderMockDecodeExpectation{
		mock:   mmDecode.mock,
		params: &UrlIDEncoderMockDecodeParams{str},
	}
	mmDecode.expectations = append(mmDecode.expectations, expectation)
	return expectation
}

// Then sets up urlIDEncoder.Decode return parameters for the expectation previously defined by the When method
func (e *UrlIDEncoderMockDecodeExpectation) Then(u1 uint64, err error) *UrlIDEncoderMock {
	e.results = &UrlIDEncoderMockDecodeResults{u1, err}
	return e.mock
}

// Decode implements urlIDEncoder
func (mmDecode *UrlIDEncoderMock) Decode(str string) (u1 uint64, err error) {
	mm_atomic.AddUint64(&mmDecode.beforeDecodeCounter, 1)
	defer mm_atomic.AddUint64(&mmDecode.afterDecodeCounter, 1)

	if mmDecode.inspectFuncDecode != nil {
		mmDecode.inspectFuncDecode(str)
	}

	mm_params := &UrlIDEncoderMockDecodeParams{str}

	// Record call args
	mmDecode.DecodeMock.mutex.Lock()
	mmDecode.DecodeMock.callArgs = append(mmDecode.DecodeMock.callArgs, mm_params)
	mmDecode.DecodeMock.mutex.Unlock()

	for _, e := range mmDecode.DecodeMock.expectations {
		if minimock.Equal(e.params, mm_params) {
			mm_atomic.AddUint64(&e.Counter, 1)
			return e.results.u1, e.results.err
		}
	}

	if mmDecode.DecodeMock.defaultExpectation != nil {
		mm_atomic.AddUint64(&mmDecode.DecodeMock.defaultExpectation.Counter, 1)
		mm_want := mmDecode.DecodeMock.defaultExpectation.params
		mm_got := UrlIDEncoderMockDecodeParams{str}
		if mm_want != nil && !minimock.Equal(*mm_want, mm_got) {
			mmDecode.t.Errorf("UrlIDEncoderMock.Decode got unexpected parameters, want: %#v, got: %#v%s\n", *mm_want, mm_got, minimock.Diff(*mm_want, mm_got))
		}

		mm_results := mmDecode.DecodeMock.defaultExpectation.results
		if mm_results == nil {
			mmDecode.t.Fatal("No results are set for the UrlIDEncoderMock.Decode")
		}
		return (*mm_results).u1, (*mm_results).err
	}
	if mmDecode.funcDecode != nil {
		return mmDecode.funcDecode(str)
	}
	mmDecode.t.Fatalf("Unexpected call to UrlIDEncoderMock.Decode. %v", str)
	return
}

// DecodeAfterCounter returns a count of finished UrlIDEncoderMock.Decode invocations
func (mmDecode *UrlIDEncoderMock) DecodeAfterCounter() uint64 {
	return mm_atomic.LoadUint64(&mmDecode.afterDecodeCounter)
}

// DecodeBeforeCounter returns a count of UrlIDEncoderMock.Decode invocations
func (mmDecode *UrlIDEncoderMock) DecodeBeforeCounter() uint64 {
	return mm_atomic.LoadUint64(&mmDecode.beforeDecodeCounter)
}

// Calls returns a list of arguments used in each call to UrlIDEncoderMock.Decode.
// The list is in the same order as the calls were made (i.e. recent calls have a higher index)
func (mmDecode *mUrlIDEncoderMockDecode) Calls() []*UrlIDEncoderMockDecodeParams {
	mmDecode.mutex.RLock()

	argCopy := make([]*UrlIDEncoderMockDecodeParams, len(mmDecode.callArgs))
	copy(argCopy, mmDecode.callArgs)

	mmDecode.mutex.RUnlock()

	return argCopy
}

// MinimockDecodeDone returns true if the count of the Decode invocations corresponds
// the number of defined expectations
func (m *UrlIDEncoderMock) MinimockDecodeDone() bool {
	for _, e := range m.DecodeMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			return false
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.DecodeMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterDecodeCounter) < 1 {
		return false
	}
	// if func was set then invocations count should be greater than zero
	if m.funcDecode != nil && mm_atomic.LoadUint64(&m.afterDecodeCounter) < 1 {
		return false
	}
	return true
}

// MinimockDecodeInspect logs each unmet expectation
func (m *UrlIDEncoderMock) MinimockDecodeInspect() {
	for _, e := range m.DecodeMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			m.t.Errorf("Expected call to UrlIDEncoderMock.Decode with params: %#v", *e.params)
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.DecodeMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterDecodeCounter) < 1 {
		if m.DecodeMock.defaultExpectation.params == nil {
			m.t.Error("Expected call to UrlIDEncoderMock.Decode")
		} else {
			m.t.Errorf("Expected call to UrlIDEncoderMock.Decode with params: %#v", *m.DecodeMock.defaultExpectation.params)
		}
	}
	// if func was set then invocations count should be greater than zero
	if m.funcDecode != nil && mm_atomic.LoadUint64(&m.afterDecodeCounter) < 1 {
		m.t.Error("Expected call to UrlIDEncoderMock.Decode")
	}
}

type mUrlIDEncoderMockEncode struct {
	mock               *UrlIDEncoderMock
	defaultExpectation *UrlIDEncoderMockEncodeExpectation
	expectations       []*UrlIDEncoderMockEncodeExpectation

	callArgs []*UrlIDEncoderMockEncodeParams
	mutex    sync.RWMutex
}

// UrlIDEncoderMockEncodeExpectation specifies expectation struct of the urlIDEncoder.Encode
type UrlIDEncoderMockEncodeExpectation struct {
	mock    *UrlIDEncoderMock
	params  *UrlIDEncoderMockEncodeParams
	results *UrlIDEncoderMockEncodeResults
	Counter uint64
}

// UrlIDEncoderMockEncodeParams contains parameters of the urlIDEncoder.Encode
type UrlIDEncoderMockEncodeParams struct {
	id uint64
}

// UrlIDEncoderMockEncodeResults contains results of the urlIDEncoder.Encode
type UrlIDEncoderMockEncodeResults struct {
	s1  string
	err error
}

// Expect sets up expected params for urlIDEncoder.Encode
func (mmEncode *mUrlIDEncoderMockEncode) Expect(id uint64) *mUrlIDEncoderMockEncode {
	if mmEncode.mock.funcEncode != nil {
		mmEncode.mock.t.Fatalf("UrlIDEncoderMock.Encode mock is already set by Set")
	}

	if mmEncode.defaultExpectation == nil {
		mmEncode.defaultExpectation = &UrlIDEncoderMockEncodeExpectation{}
	}

	mmEncode.defaultExpectation.params = &UrlIDEncoderMockEncodeParams{id}
	for _, e := range mmEncode.expectations {
		if minimock.Equal(e.params, mmEncode.defaultExpectation.params) {
			mmEncode.mock.t.Fatalf("Expectation set by When has same params: %#v", *mmEncode.defaultExpectation.params)
		}
	}

	return mmEncode
}

// Inspect accepts an inspector function that has same arguments as the urlIDEncoder.Encode
func (mmEncode *mUrlIDEncoderMockEncode) Inspect(f func(id uint64)) *mUrlIDEncoderMockEncode {
	if mmEncode.mock.inspectFuncEncode != nil {
		mmEncode.mock.t.Fatalf("Inspect function is already set for UrlIDEncoderMock.Encode")
	}

	mmEncode.mock.inspectFuncEncode = f

	return mmEncode
}

// Return sets up results that will be returned by urlIDEncoder.Encode
func (mmEncode *mUrlIDEncoderMockEncode) Return(s1 string, err error) *UrlIDEncoderMock {
	if mmEncode.mock.funcEncode != nil {
		mmEncode.mock.t.Fatalf("UrlIDEncoderMock.Encode mock is already set by Set")
	}

	if mmEncode.defaultExpectation == nil {
		mmEncode.defaultExpectation = &UrlIDEncoderMockEncodeExpectation{mock: mmEncode.mock}
	}
	mmEncode.defaultExpectation.results = &UrlIDEncoderMockEncodeResults{s1, err}
	return mmEncode.mock
}

//Set uses given function f to mock the urlIDEncoder.Encode method
func (mmEncode *mUrlIDEncoderMockEncode) Set(f func(id uint64) (s1 string, err error)) *UrlIDEncoderMock {
	if mmEncode.defaultExpectation != nil {
		mmEncode.mock.t.Fatalf("Default expectation is already set for the urlIDEncoder.Encode method")
	}

	if len(mmEncode.expectations) > 0 {
		mmEncode.mock.t.Fatalf("Some expectations are already set for the urlIDEncoder.Encode method")
	}

	mmEncode.mock.funcEncode = f
	return mmEncode.mock
}

// When sets expectation for the urlIDEncoder.Encode which will trigger the result defined by the following
// Then helper
func (mmEncode *mUrlIDEncoderMockEncode) When(id uint64) *UrlIDEncoderMockEncodeExpectation {
	if mmEncode.mock.funcEncode != nil {
		mmEncode.mock.t.Fatalf("UrlIDEncoderMock.Encode mock is already set by Set")
	}

	expectation := &UrlIDEncoderMockEncodeExpectation{
		mock:   mmEncode.mock,
		params: &UrlIDEncoderMockEncodeParams{id},
	}
	mmEncode.expectations = append(mmEncode.expectations, expectation)
	return expectation
}

// Then sets up urlIDEncoder.Encode return parameters for the expectation previously defined by the When method
func (e *UrlIDEncoderMockEncodeExpectation) Then(s1 string, err error) *UrlIDEncoderMock {
	e.results = &UrlIDEncoderMockEncodeResults{s1, err}
	return e.mock
}

// Encode implements urlIDEncoder
func (mmEncode *UrlIDEncoderMock) Encode(id uint64) (s1 string, err error) {
	mm_atomic.AddUint64(&mmEncode.beforeEncodeCounter, 1)
	defer mm_atomic.AddUint64(&mmEncode.afterEncodeCounter, 1)

	if mmEncode.inspectFuncEncode != nil {
		mmEncode.inspectFuncEncode(id)
	}

	mm_params := &UrlIDEncoderMockEncodeParams{id}

	// Record call args
	mmEncode.EncodeMock.mutex.Lock()
	mmEncode.EncodeMock.callArgs = append(mmEncode.EncodeMock.callArgs, mm_params)
	mmEncode.EncodeMock.mutex.Unlock()

	for _, e := range mmEncode.EncodeMock.expectations {
		if minimock.Equal(e.params, mm_params) {
			mm_atomic.AddUint64(&e.Counter, 1)
			return e.results.s1, e.results.err
		}
	}

	if mmEncode.EncodeMock.defaultExpectation != nil {
		mm_atomic.AddUint64(&mmEncode.EncodeMock.defaultExpectation.Counter, 1)
		mm_want := mmEncode.EncodeMock.defaultExpectation.params
		mm_got := UrlIDEncoderMockEncodeParams{id}
		if mm_want != nil && !minimock.Equal(*mm_want, mm_got) {
			mmEncode.t.Errorf("UrlIDEncoderMock.Encode got unexpected parameters, want: %#v, got: %#v%s\n", *mm_want, mm_got, minimock.Diff(*mm_want, mm_got))
		}

		mm_results := mmEncode.EncodeMock.defaultExpectation.results
		if mm_results == nil {
			mmEncode.t.Fatal("No results are set for the UrlIDEncoderMock.Encode")
		}
		return (*mm_results).s1, (*mm_results).err
	}
	if mmEncode.funcEncode != nil {
		return mmEncode.funcEncode(id)
	}
	mmEncode.t.Fatalf("Unexpected call to UrlIDEncoderMock.Encode. %v", id)
	return
}

// EncodeAfterCounter returns a count of finished UrlIDEncoderMock.Encode invocations
func (mmEncode *UrlIDEncoderMock) EncodeAfterCounter() uint64 {
	return mm_atomic.LoadUint64(&mmEncode.afterEncodeCounter)
}

// EncodeBeforeCounter returns a count of UrlIDEncoderMock.Encode invocations
func (mmEncode *UrlIDEncoderMock) EncodeBeforeCounter() uint64 {
	return mm_atomic.LoadUint64(&mmEncode.beforeEncodeCounter)
}

// Calls returns a list of arguments used in each call to UrlIDEncoderMock.Encode.
// The list is in the same order as the calls were made (i.e. recent calls have a higher index)
func (mmEncode *mUrlIDEncoderMockEncode) Calls() []*UrlIDEncoderMockEncodeParams {
	mmEncode.mutex.RLock()

	argCopy := make([]*UrlIDEncoderMockEncodeParams, len(mmEncode.callArgs))
	copy(argCopy, mmEncode.callArgs)

	mmEncode.mutex.RUnlock()

	return argCopy
}

// MinimockEncodeDone returns true if the count of the Encode invocations corresponds
// the number of defined expectations
func (m *UrlIDEncoderMock) MinimockEncodeDone() bool {
	for _, e := range m.EncodeMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			return false
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.EncodeMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterEncodeCounter) < 1 {
		return false
	}
	// if func was set then invocations count should be greater than zero
	if m.funcEncode != nil && mm_atomic.LoadUint64(&m.afterEncodeCounter) < 1 {
		return false
	}
	return true
}

// MinimockEncodeInspect logs each unmet expectation
func (m *UrlIDEncoderMock) MinimockEncodeInspect() {
	for _, e := range m.EncodeMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			m.t.Errorf("Expected call to UrlIDEncoderMock.Encode with params: %#v", *e.params)
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.EncodeMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterEncodeCounter) < 1 {
		if m.EncodeMock.defaultExpectation.params == nil {
			m.t.Error("Expected call to UrlIDEncoderMock.Encode")
		} else {
			m.t.Errorf("Expected call to UrlIDEncoderMock.Encode with params: %#v", *m.EncodeMock.defaultExpectation.params)
		}
	}
	// if func was set then invocations count should be greater than zero
	if m.funcEncode != nil && mm_atomic.LoadUint64(&m.afterEncodeCounter) < 1 {
		m.t.Error("Expected call to UrlIDEncoderMock.Encode")
	}
}

// MinimockFinish checks that all mocked methods have been called the expected number of times
func (m *UrlIDEncoderMock) MinimockFinish() {
	if !m.minimockDone() {
		m.MinimockDecodeInspect()

		m.MinimockEncodeInspect()
		m.t.FailNow()
	}
}

// MinimockWait waits for all mocked methods to be called the expected number of times
func (m *UrlIDEncoderMock) MinimockWait(timeout mm_time.Duration) {
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

func (m *UrlIDEncoderMock) minimockDone() bool {
	done := true
	return done &&
		m.MinimockDecodeDone() &&
		m.MinimockEncodeDone()
}
