// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

// DefenderPromptForSampleSubmission undocumented
type DefenderPromptForSampleSubmission int

const (
	// DefenderPromptForSampleSubmissionVUserDefined undocumented
	DefenderPromptForSampleSubmissionVUserDefined DefenderPromptForSampleSubmission = 0
	// DefenderPromptForSampleSubmissionVAlwaysPrompt undocumented
	DefenderPromptForSampleSubmissionVAlwaysPrompt DefenderPromptForSampleSubmission = 1
	// DefenderPromptForSampleSubmissionVPromptBeforeSendingPersonalData undocumented
	DefenderPromptForSampleSubmissionVPromptBeforeSendingPersonalData DefenderPromptForSampleSubmission = 2
	// DefenderPromptForSampleSubmissionVNeverSendData undocumented
	DefenderPromptForSampleSubmissionVNeverSendData DefenderPromptForSampleSubmission = 3
	// DefenderPromptForSampleSubmissionVSendAllDataWithoutPrompting undocumented
	DefenderPromptForSampleSubmissionVSendAllDataWithoutPrompting DefenderPromptForSampleSubmission = 4
)

// DefenderPromptForSampleSubmissionPUserDefined returns a pointer to DefenderPromptForSampleSubmissionVUserDefined
func DefenderPromptForSampleSubmissionPUserDefined() *DefenderPromptForSampleSubmission {
	v := DefenderPromptForSampleSubmissionVUserDefined
	return &v
}

// DefenderPromptForSampleSubmissionPAlwaysPrompt returns a pointer to DefenderPromptForSampleSubmissionVAlwaysPrompt
func DefenderPromptForSampleSubmissionPAlwaysPrompt() *DefenderPromptForSampleSubmission {
	v := DefenderPromptForSampleSubmissionVAlwaysPrompt
	return &v
}

// DefenderPromptForSampleSubmissionPPromptBeforeSendingPersonalData returns a pointer to DefenderPromptForSampleSubmissionVPromptBeforeSendingPersonalData
func DefenderPromptForSampleSubmissionPPromptBeforeSendingPersonalData() *DefenderPromptForSampleSubmission {
	v := DefenderPromptForSampleSubmissionVPromptBeforeSendingPersonalData
	return &v
}

// DefenderPromptForSampleSubmissionPNeverSendData returns a pointer to DefenderPromptForSampleSubmissionVNeverSendData
func DefenderPromptForSampleSubmissionPNeverSendData() *DefenderPromptForSampleSubmission {
	v := DefenderPromptForSampleSubmissionVNeverSendData
	return &v
}

// DefenderPromptForSampleSubmissionPSendAllDataWithoutPrompting returns a pointer to DefenderPromptForSampleSubmissionVSendAllDataWithoutPrompting
func DefenderPromptForSampleSubmissionPSendAllDataWithoutPrompting() *DefenderPromptForSampleSubmission {
	v := DefenderPromptForSampleSubmissionVSendAllDataWithoutPrompting
	return &v
}
