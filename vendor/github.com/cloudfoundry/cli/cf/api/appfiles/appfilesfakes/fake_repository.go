// This file was generated by counterfeiter
package appfilesfakes

import (
	"sync"

	"github.com/cloudfoundry/cli/cf/api/appfiles"
)

type FakeRepository struct {
	ListFilesStub        func(appGUID string, instance int, path string) (files string, apiErr error)
	listFilesMutex       sync.RWMutex
	listFilesArgsForCall []struct {
		appGUID  string
		instance int
		path     string
	}
	listFilesReturns struct {
		result1 string
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeRepository) ListFiles(appGUID string, instance int, path string) (files string, apiErr error) {
	fake.listFilesMutex.Lock()
	fake.listFilesArgsForCall = append(fake.listFilesArgsForCall, struct {
		appGUID  string
		instance int
		path     string
	}{appGUID, instance, path})
	fake.recordInvocation("ListFiles", []interface{}{appGUID, instance, path})
	fake.listFilesMutex.Unlock()
	if fake.ListFilesStub != nil {
		return fake.ListFilesStub(appGUID, instance, path)
	} else {
		return fake.listFilesReturns.result1, fake.listFilesReturns.result2
	}
}

func (fake *FakeRepository) ListFilesCallCount() int {
	fake.listFilesMutex.RLock()
	defer fake.listFilesMutex.RUnlock()
	return len(fake.listFilesArgsForCall)
}

func (fake *FakeRepository) ListFilesArgsForCall(i int) (string, int, string) {
	fake.listFilesMutex.RLock()
	defer fake.listFilesMutex.RUnlock()
	return fake.listFilesArgsForCall[i].appGUID, fake.listFilesArgsForCall[i].instance, fake.listFilesArgsForCall[i].path
}

func (fake *FakeRepository) ListFilesReturns(result1 string, result2 error) {
	fake.ListFilesStub = nil
	fake.listFilesReturns = struct {
		result1 string
		result2 error
	}{result1, result2}
}

func (fake *FakeRepository) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.listFilesMutex.RLock()
	defer fake.listFilesMutex.RUnlock()
	return fake.invocations
}

func (fake *FakeRepository) recordInvocation(key string, args []interface{}) {
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

var _ appfiles.Repository = new(FakeRepository)
