/*
 * Isilon SDK
 *
 * Isilon SDK - Language bindings for the OneFS API
 *
 * API version: 5
 * Contact: sdk@isilon.com
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package isi_sdk_8_1_0

type JobJobExtended struct {

	// State to which the job is transitioning; if control_state is identical to state, the job's state is stable.
	ControlState string `json:"control_state,omitempty"`

	// The time the job was queued, in seconds since the epoch.
	CreateTime int32 `json:"create_time"`

	// The current phase of the job.
	CurrentPhase int32 `json:"current_phase,omitempty"`

	// A text representation of the job.
	Description string `json:"description,omitempty"`

	// The time the job ended, in seconds since the Epoch.
	EndTime int32 `json:"end_time,omitempty"`

	// The ID of the job.
	Id int32 `json:"id"`

	// The current impact of the job.
	Impact string `json:"impact"`

	// The set of devids working on the job.
	Participants []int32 `json:"participants,omitempty"`

	// Paths for which the job was queued.
	Paths []string `json:"paths,omitempty"`

	// Current impact policy of the job.
	Policy string `json:"policy"`

	// Current priority of the job; lower numbers preempt higher numbers.
	Priority int32 `json:"priority"`

	// A text representation of the job's progress.
	Progress string `json:"progress,omitempty"`

	// The number of retries remaining if the job fails.
	RetriesRemaining int32 `json:"retries_remaining"`

	// The number of seconds the job has executed.
	RunningTime int32 `json:"running_time,omitempty"`

	// The time the job started, in seconds since the Epoch.
	StartTime int32 `json:"start_time,omitempty"`

	// Current state of the job.
	State string `json:"state"`

	// The total number of phases of the job type.
	TotalPhases int32 `json:"total_phases"`

	// The job type.
	Type_ string `json:"type"`

	// The ID of a job for which this job is waiting.
	WaitingOn int32 `json:"waiting_on,omitempty"`

	// The reason the job is waiting.
	WaitingReason string `json:"waiting_reason,omitempty"`
}