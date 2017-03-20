// This file was generated by counterfeiter
package routefakes

import (
	"route-sync/route"
	"sync"
)

type FakeRouter struct {
	TCPStub        func(routes []*route.TCP) error
	tCPMutex       sync.RWMutex
	tCPArgsForCall []struct {
		routes []*route.TCP
	}
	tCPReturns struct {
		result1 error
	}
	HTTPStub        func(routes []*route.HTTP) error
	hTTPMutex       sync.RWMutex
	hTTPArgsForCall []struct {
		routes []*route.HTTP
	}
	hTTPReturns struct {
		result1 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeRouter) TCP(routes []*route.TCP) error {
	var routesCopy []*route.TCP
	if routes != nil {
		routesCopy = make([]*route.TCP, len(routes))
		copy(routesCopy, routes)
	}
	fake.tCPMutex.Lock()
	fake.tCPArgsForCall = append(fake.tCPArgsForCall, struct {
		routes []*route.TCP
	}{routesCopy})
	fake.recordInvocation("TCP", []interface{}{routesCopy})
	fake.tCPMutex.Unlock()
	if fake.TCPStub != nil {
		return fake.TCPStub(routes)
	} else {
		return fake.tCPReturns.result1
	}
}

func (fake *FakeRouter) TCPCallCount() int {
	fake.tCPMutex.RLock()
	defer fake.tCPMutex.RUnlock()
	return len(fake.tCPArgsForCall)
}

func (fake *FakeRouter) TCPArgsForCall(i int) []*route.TCP {
	fake.tCPMutex.RLock()
	defer fake.tCPMutex.RUnlock()
	return fake.tCPArgsForCall[i].routes
}

func (fake *FakeRouter) TCPReturns(result1 error) {
	fake.TCPStub = nil
	fake.tCPReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeRouter) HTTP(routes []*route.HTTP) error {
	var routesCopy []*route.HTTP
	if routes != nil {
		routesCopy = make([]*route.HTTP, len(routes))
		copy(routesCopy, routes)
	}
	fake.hTTPMutex.Lock()
	fake.hTTPArgsForCall = append(fake.hTTPArgsForCall, struct {
		routes []*route.HTTP
	}{routesCopy})
	fake.recordInvocation("HTTP", []interface{}{routesCopy})
	fake.hTTPMutex.Unlock()
	if fake.HTTPStub != nil {
		return fake.HTTPStub(routes)
	} else {
		return fake.hTTPReturns.result1
	}
}

func (fake *FakeRouter) HTTPCallCount() int {
	fake.hTTPMutex.RLock()
	defer fake.hTTPMutex.RUnlock()
	return len(fake.hTTPArgsForCall)
}

func (fake *FakeRouter) HTTPArgsForCall(i int) []*route.HTTP {
	fake.hTTPMutex.RLock()
	defer fake.hTTPMutex.RUnlock()
	return fake.hTTPArgsForCall[i].routes
}

func (fake *FakeRouter) HTTPReturns(result1 error) {
	fake.HTTPStub = nil
	fake.hTTPReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeRouter) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.tCPMutex.RLock()
	defer fake.tCPMutex.RUnlock()
	fake.hTTPMutex.RLock()
	defer fake.hTTPMutex.RUnlock()
	return fake.invocations
}

func (fake *FakeRouter) recordInvocation(key string, args []interface{}) {
	fake.invocationsMutex.Lock()
	defer fake.invocationsMutex.Unlock()
	if fake.invocations == nil {
		fake.invocations = map[string][][]interface{}{}
	}
	if fake.invocations[key] == nil {
		fake.invocations[key] = [][]interface{}{}
	}
	fake.invocations[key] = append(fake.invocations[key], args)
}

var _ route.Router = new(FakeRouter)
