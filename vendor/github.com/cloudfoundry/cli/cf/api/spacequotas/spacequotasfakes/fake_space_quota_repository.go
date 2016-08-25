// This file was generated by counterfeiter
package spacequotasfakes

import (
	"sync"

	"github.com/cloudfoundry/cli/cf/api/spacequotas"
	"github.com/cloudfoundry/cli/cf/models"
)

type FakeSpaceQuotaRepository struct {
	FindByNameStub        func(name string) (quota models.SpaceQuota, apiErr error)
	findByNameMutex       sync.RWMutex
	findByNameArgsForCall []struct {
		name string
	}
	findByNameReturns struct {
		result1 models.SpaceQuota
		result2 error
	}
	FindByOrgStub        func(guid string) (quota []models.SpaceQuota, apiErr error)
	findByOrgMutex       sync.RWMutex
	findByOrgArgsForCall []struct {
		guid string
	}
	findByOrgReturns struct {
		result1 []models.SpaceQuota
		result2 error
	}
	FindByGUIDStub        func(guid string) (quota models.SpaceQuota, apiErr error)
	findByGUIDMutex       sync.RWMutex
	findByGUIDArgsForCall []struct {
		guid string
	}
	findByGUIDReturns struct {
		result1 models.SpaceQuota
		result2 error
	}
	FindByNameAndOrgGUIDStub        func(spaceQuotaName string, orgGUID string) (quota models.SpaceQuota, apiErr error)
	findByNameAndOrgGUIDMutex       sync.RWMutex
	findByNameAndOrgGUIDArgsForCall []struct {
		spaceQuotaName string
		orgGUID        string
	}
	findByNameAndOrgGUIDReturns struct {
		result1 models.SpaceQuota
		result2 error
	}
	AssociateSpaceWithQuotaStub        func(spaceGUID string, quotaGUID string) error
	associateSpaceWithQuotaMutex       sync.RWMutex
	associateSpaceWithQuotaArgsForCall []struct {
		spaceGUID string
		quotaGUID string
	}
	associateSpaceWithQuotaReturns struct {
		result1 error
	}
	UnassignQuotaFromSpaceStub        func(spaceGUID string, quotaGUID string) error
	unassignQuotaFromSpaceMutex       sync.RWMutex
	unassignQuotaFromSpaceArgsForCall []struct {
		spaceGUID string
		quotaGUID string
	}
	unassignQuotaFromSpaceReturns struct {
		result1 error
	}
	CreateStub        func(quota models.SpaceQuota) error
	createMutex       sync.RWMutex
	createArgsForCall []struct {
		quota models.SpaceQuota
	}
	createReturns struct {
		result1 error
	}
	UpdateStub        func(quota models.SpaceQuota) error
	updateMutex       sync.RWMutex
	updateArgsForCall []struct {
		quota models.SpaceQuota
	}
	updateReturns struct {
		result1 error
	}
	DeleteStub        func(quotaGUID string) error
	deleteMutex       sync.RWMutex
	deleteArgsForCall []struct {
		quotaGUID string
	}
	deleteReturns struct {
		result1 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeSpaceQuotaRepository) FindByName(name string) (quota models.SpaceQuota, apiErr error) {
	fake.findByNameMutex.Lock()
	fake.findByNameArgsForCall = append(fake.findByNameArgsForCall, struct {
		name string
	}{name})
	fake.recordInvocation("FindByName", []interface{}{name})
	fake.findByNameMutex.Unlock()
	if fake.FindByNameStub != nil {
		return fake.FindByNameStub(name)
	} else {
		return fake.findByNameReturns.result1, fake.findByNameReturns.result2
	}
}

func (fake *FakeSpaceQuotaRepository) FindByNameCallCount() int {
	fake.findByNameMutex.RLock()
	defer fake.findByNameMutex.RUnlock()
	return len(fake.findByNameArgsForCall)
}

func (fake *FakeSpaceQuotaRepository) FindByNameArgsForCall(i int) string {
	fake.findByNameMutex.RLock()
	defer fake.findByNameMutex.RUnlock()
	return fake.findByNameArgsForCall[i].name
}

func (fake *FakeSpaceQuotaRepository) FindByNameReturns(result1 models.SpaceQuota, result2 error) {
	fake.FindByNameStub = nil
	fake.findByNameReturns = struct {
		result1 models.SpaceQuota
		result2 error
	}{result1, result2}
}

func (fake *FakeSpaceQuotaRepository) FindByOrg(guid string) (quota []models.SpaceQuota, apiErr error) {
	fake.findByOrgMutex.Lock()
	fake.findByOrgArgsForCall = append(fake.findByOrgArgsForCall, struct {
		guid string
	}{guid})
	fake.recordInvocation("FindByOrg", []interface{}{guid})
	fake.findByOrgMutex.Unlock()
	if fake.FindByOrgStub != nil {
		return fake.FindByOrgStub(guid)
	} else {
		return fake.findByOrgReturns.result1, fake.findByOrgReturns.result2
	}
}

func (fake *FakeSpaceQuotaRepository) FindByOrgCallCount() int {
	fake.findByOrgMutex.RLock()
	defer fake.findByOrgMutex.RUnlock()
	return len(fake.findByOrgArgsForCall)
}

func (fake *FakeSpaceQuotaRepository) FindByOrgArgsForCall(i int) string {
	fake.findByOrgMutex.RLock()
	defer fake.findByOrgMutex.RUnlock()
	return fake.findByOrgArgsForCall[i].guid
}

func (fake *FakeSpaceQuotaRepository) FindByOrgReturns(result1 []models.SpaceQuota, result2 error) {
	fake.FindByOrgStub = nil
	fake.findByOrgReturns = struct {
		result1 []models.SpaceQuota
		result2 error
	}{result1, result2}
}

func (fake *FakeSpaceQuotaRepository) FindByGUID(guid string) (quota models.SpaceQuota, apiErr error) {
	fake.findByGUIDMutex.Lock()
	fake.findByGUIDArgsForCall = append(fake.findByGUIDArgsForCall, struct {
		guid string
	}{guid})
	fake.recordInvocation("FindByGUID", []interface{}{guid})
	fake.findByGUIDMutex.Unlock()
	if fake.FindByGUIDStub != nil {
		return fake.FindByGUIDStub(guid)
	} else {
		return fake.findByGUIDReturns.result1, fake.findByGUIDReturns.result2
	}
}

func (fake *FakeSpaceQuotaRepository) FindByGUIDCallCount() int {
	fake.findByGUIDMutex.RLock()
	defer fake.findByGUIDMutex.RUnlock()
	return len(fake.findByGUIDArgsForCall)
}

func (fake *FakeSpaceQuotaRepository) FindByGUIDArgsForCall(i int) string {
	fake.findByGUIDMutex.RLock()
	defer fake.findByGUIDMutex.RUnlock()
	return fake.findByGUIDArgsForCall[i].guid
}

func (fake *FakeSpaceQuotaRepository) FindByGUIDReturns(result1 models.SpaceQuota, result2 error) {
	fake.FindByGUIDStub = nil
	fake.findByGUIDReturns = struct {
		result1 models.SpaceQuota
		result2 error
	}{result1, result2}
}

func (fake *FakeSpaceQuotaRepository) FindByNameAndOrgGUID(spaceQuotaName string, orgGUID string) (quota models.SpaceQuota, apiErr error) {
	fake.findByNameAndOrgGUIDMutex.Lock()
	fake.findByNameAndOrgGUIDArgsForCall = append(fake.findByNameAndOrgGUIDArgsForCall, struct {
		spaceQuotaName string
		orgGUID        string
	}{spaceQuotaName, orgGUID})
	fake.recordInvocation("FindByNameAndOrgGUID", []interface{}{spaceQuotaName, orgGUID})
	fake.findByNameAndOrgGUIDMutex.Unlock()
	if fake.FindByNameAndOrgGUIDStub != nil {
		return fake.FindByNameAndOrgGUIDStub(spaceQuotaName, orgGUID)
	} else {
		return fake.findByNameAndOrgGUIDReturns.result1, fake.findByNameAndOrgGUIDReturns.result2
	}
}

func (fake *FakeSpaceQuotaRepository) FindByNameAndOrgGUIDCallCount() int {
	fake.findByNameAndOrgGUIDMutex.RLock()
	defer fake.findByNameAndOrgGUIDMutex.RUnlock()
	return len(fake.findByNameAndOrgGUIDArgsForCall)
}

func (fake *FakeSpaceQuotaRepository) FindByNameAndOrgGUIDArgsForCall(i int) (string, string) {
	fake.findByNameAndOrgGUIDMutex.RLock()
	defer fake.findByNameAndOrgGUIDMutex.RUnlock()
	return fake.findByNameAndOrgGUIDArgsForCall[i].spaceQuotaName, fake.findByNameAndOrgGUIDArgsForCall[i].orgGUID
}

func (fake *FakeSpaceQuotaRepository) FindByNameAndOrgGUIDReturns(result1 models.SpaceQuota, result2 error) {
	fake.FindByNameAndOrgGUIDStub = nil
	fake.findByNameAndOrgGUIDReturns = struct {
		result1 models.SpaceQuota
		result2 error
	}{result1, result2}
}

func (fake *FakeSpaceQuotaRepository) AssociateSpaceWithQuota(spaceGUID string, quotaGUID string) error {
	fake.associateSpaceWithQuotaMutex.Lock()
	fake.associateSpaceWithQuotaArgsForCall = append(fake.associateSpaceWithQuotaArgsForCall, struct {
		spaceGUID string
		quotaGUID string
	}{spaceGUID, quotaGUID})
	fake.recordInvocation("AssociateSpaceWithQuota", []interface{}{spaceGUID, quotaGUID})
	fake.associateSpaceWithQuotaMutex.Unlock()
	if fake.AssociateSpaceWithQuotaStub != nil {
		return fake.AssociateSpaceWithQuotaStub(spaceGUID, quotaGUID)
	} else {
		return fake.associateSpaceWithQuotaReturns.result1
	}
}

func (fake *FakeSpaceQuotaRepository) AssociateSpaceWithQuotaCallCount() int {
	fake.associateSpaceWithQuotaMutex.RLock()
	defer fake.associateSpaceWithQuotaMutex.RUnlock()
	return len(fake.associateSpaceWithQuotaArgsForCall)
}

func (fake *FakeSpaceQuotaRepository) AssociateSpaceWithQuotaArgsForCall(i int) (string, string) {
	fake.associateSpaceWithQuotaMutex.RLock()
	defer fake.associateSpaceWithQuotaMutex.RUnlock()
	return fake.associateSpaceWithQuotaArgsForCall[i].spaceGUID, fake.associateSpaceWithQuotaArgsForCall[i].quotaGUID
}

func (fake *FakeSpaceQuotaRepository) AssociateSpaceWithQuotaReturns(result1 error) {
	fake.AssociateSpaceWithQuotaStub = nil
	fake.associateSpaceWithQuotaReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeSpaceQuotaRepository) UnassignQuotaFromSpace(spaceGUID string, quotaGUID string) error {
	fake.unassignQuotaFromSpaceMutex.Lock()
	fake.unassignQuotaFromSpaceArgsForCall = append(fake.unassignQuotaFromSpaceArgsForCall, struct {
		spaceGUID string
		quotaGUID string
	}{spaceGUID, quotaGUID})
	fake.recordInvocation("UnassignQuotaFromSpace", []interface{}{spaceGUID, quotaGUID})
	fake.unassignQuotaFromSpaceMutex.Unlock()
	if fake.UnassignQuotaFromSpaceStub != nil {
		return fake.UnassignQuotaFromSpaceStub(spaceGUID, quotaGUID)
	} else {
		return fake.unassignQuotaFromSpaceReturns.result1
	}
}

func (fake *FakeSpaceQuotaRepository) UnassignQuotaFromSpaceCallCount() int {
	fake.unassignQuotaFromSpaceMutex.RLock()
	defer fake.unassignQuotaFromSpaceMutex.RUnlock()
	return len(fake.unassignQuotaFromSpaceArgsForCall)
}

func (fake *FakeSpaceQuotaRepository) UnassignQuotaFromSpaceArgsForCall(i int) (string, string) {
	fake.unassignQuotaFromSpaceMutex.RLock()
	defer fake.unassignQuotaFromSpaceMutex.RUnlock()
	return fake.unassignQuotaFromSpaceArgsForCall[i].spaceGUID, fake.unassignQuotaFromSpaceArgsForCall[i].quotaGUID
}

func (fake *FakeSpaceQuotaRepository) UnassignQuotaFromSpaceReturns(result1 error) {
	fake.UnassignQuotaFromSpaceStub = nil
	fake.unassignQuotaFromSpaceReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeSpaceQuotaRepository) Create(quota models.SpaceQuota) error {
	fake.createMutex.Lock()
	fake.createArgsForCall = append(fake.createArgsForCall, struct {
		quota models.SpaceQuota
	}{quota})
	fake.recordInvocation("Create", []interface{}{quota})
	fake.createMutex.Unlock()
	if fake.CreateStub != nil {
		return fake.CreateStub(quota)
	} else {
		return fake.createReturns.result1
	}
}

func (fake *FakeSpaceQuotaRepository) CreateCallCount() int {
	fake.createMutex.RLock()
	defer fake.createMutex.RUnlock()
	return len(fake.createArgsForCall)
}

func (fake *FakeSpaceQuotaRepository) CreateArgsForCall(i int) models.SpaceQuota {
	fake.createMutex.RLock()
	defer fake.createMutex.RUnlock()
	return fake.createArgsForCall[i].quota
}

func (fake *FakeSpaceQuotaRepository) CreateReturns(result1 error) {
	fake.CreateStub = nil
	fake.createReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeSpaceQuotaRepository) Update(quota models.SpaceQuota) error {
	fake.updateMutex.Lock()
	fake.updateArgsForCall = append(fake.updateArgsForCall, struct {
		quota models.SpaceQuota
	}{quota})
	fake.recordInvocation("Update", []interface{}{quota})
	fake.updateMutex.Unlock()
	if fake.UpdateStub != nil {
		return fake.UpdateStub(quota)
	} else {
		return fake.updateReturns.result1
	}
}

func (fake *FakeSpaceQuotaRepository) UpdateCallCount() int {
	fake.updateMutex.RLock()
	defer fake.updateMutex.RUnlock()
	return len(fake.updateArgsForCall)
}

func (fake *FakeSpaceQuotaRepository) UpdateArgsForCall(i int) models.SpaceQuota {
	fake.updateMutex.RLock()
	defer fake.updateMutex.RUnlock()
	return fake.updateArgsForCall[i].quota
}

func (fake *FakeSpaceQuotaRepository) UpdateReturns(result1 error) {
	fake.UpdateStub = nil
	fake.updateReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeSpaceQuotaRepository) Delete(quotaGUID string) error {
	fake.deleteMutex.Lock()
	fake.deleteArgsForCall = append(fake.deleteArgsForCall, struct {
		quotaGUID string
	}{quotaGUID})
	fake.recordInvocation("Delete", []interface{}{quotaGUID})
	fake.deleteMutex.Unlock()
	if fake.DeleteStub != nil {
		return fake.DeleteStub(quotaGUID)
	} else {
		return fake.deleteReturns.result1
	}
}

func (fake *FakeSpaceQuotaRepository) DeleteCallCount() int {
	fake.deleteMutex.RLock()
	defer fake.deleteMutex.RUnlock()
	return len(fake.deleteArgsForCall)
}

func (fake *FakeSpaceQuotaRepository) DeleteArgsForCall(i int) string {
	fake.deleteMutex.RLock()
	defer fake.deleteMutex.RUnlock()
	return fake.deleteArgsForCall[i].quotaGUID
}

func (fake *FakeSpaceQuotaRepository) DeleteReturns(result1 error) {
	fake.DeleteStub = nil
	fake.deleteReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeSpaceQuotaRepository) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.findByNameMutex.RLock()
	defer fake.findByNameMutex.RUnlock()
	fake.findByOrgMutex.RLock()
	defer fake.findByOrgMutex.RUnlock()
	fake.findByGUIDMutex.RLock()
	defer fake.findByGUIDMutex.RUnlock()
	fake.findByNameAndOrgGUIDMutex.RLock()
	defer fake.findByNameAndOrgGUIDMutex.RUnlock()
	fake.associateSpaceWithQuotaMutex.RLock()
	defer fake.associateSpaceWithQuotaMutex.RUnlock()
	fake.unassignQuotaFromSpaceMutex.RLock()
	defer fake.unassignQuotaFromSpaceMutex.RUnlock()
	fake.createMutex.RLock()
	defer fake.createMutex.RUnlock()
	fake.updateMutex.RLock()
	defer fake.updateMutex.RUnlock()
	fake.deleteMutex.RLock()
	defer fake.deleteMutex.RUnlock()
	return fake.invocations
}

func (fake *FakeSpaceQuotaRepository) recordInvocation(key string, args []interface{}) {
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

var _ spacequotas.SpaceQuotaRepository = new(FakeSpaceQuotaRepository)
