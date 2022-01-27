package test_helpers

import "github.com/stretchr/testify/mock"

type ApplicationConfigurationCacheMock struct {
	mock.Mock
}

func (applicationConfigurationCache *ApplicationConfigurationCacheMock) ReloadCache() {

}

func (applicationConfigurationCache *ApplicationConfigurationCacheMock) GetConfigurationValue(configurationName string) string {
	args := applicationConfigurationCache.Called(configurationName)
	return args.Get(0).(string)
}

func (applicationConfigurationCache *ApplicationConfigurationCacheMock) GetConfigurationValueAsInt(configurationName string) int {
	args := applicationConfigurationCache.Called(configurationName)
	return args.Get(0).(int)
}
