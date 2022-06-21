package kendra_test

import "testing"

// Serialize to limit service quota exceeded errors.
func TestAccKendra_serial(t *testing.T) {
	testCases := map[string]map[string]func(t *testing.T){
		"Experience": {
			"basic":       testAccExperience_basic,
			"disappears":  testAccExperience_disappears,
			"Description": testAccExperience_Description,
			"Name":        testAccExperience_Name,
			"RoleARN":     testAccExperience_roleARN,
			"Configuration_ContentSourceConfiguration_DirectPutContent":                    testAccExperience_Configuration_ContentSourceConfiguration_DirectPutContent,
			"Configuration_UserIdentityConfiguration":                                      testAccExperience_Configuration_UserIdentityConfiguration,
			"Configuration_ContentSourceConfigurationAndUserIdentityConfiguration":         testAccExperience_Configuration_ContentSourceConfigurationAndUserIdentityConfiguration,
			"Configuration_ContentSourceConfigurationWithUserIdentityConfigurationRemoved": testAccExperience_Configuration_ContentSourceConfigurationWithUserIdentityConfigurationRemoved,
			"Configuration_UserIdentityConfigurationWithContentSourceConfigurationRemoved": testAccExperience_Configuration_UserIdentityConfigurationWithContentSourceConfigurationRemoved,
		},
		"Index": {
			"basic":                testAccIndex_basic,
			"disappears":           testAccIndex_disappears,
			"tags":                 testAccIndex_updateTags,
			"CapacityUnits":        testAccIndex_updateCapacityUnits,
			"Description":          testAccIndex_updateDescription,
			"Name":                 testAccIndex_updateName,
			"RoleARN":              testAccIndex_updateRoleARN,
			"ServerSideEncryption": testAccIndex_serverSideEncryption,
			"UserTokenJSON":        testAccIndex_updateUserTokenJSON,
		},
		"Thesaurus": {
			"basic":        testAccThesaurus_basic,
			"disappears":   testAccThesaurus_disappears,
			"tags":         testAccThesaurus_tags,
			"Description":  testAccThesaurus_description,
			"Name":         testAccThesaurus_name,
			"RoleARN":      testAccThesaurus_roleARN,
			"SourceS3Path": testAccThesaurus_sourceS3Path,
		},
	}

	for group, m := range testCases {
		m := m
		t.Run(group, func(t *testing.T) {
			for name, tc := range m {
				tc := tc
				t.Run(name, func(t *testing.T) {
					tc(t)
				})
			}
		})
	}
}
