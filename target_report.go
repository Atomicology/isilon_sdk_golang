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

type TargetReport struct {

	// The action to be taken by this job.
	Action string `json:"action"`

	// The number of ads streams replicated by this job.
	AdsStreamsReplicated int32 `json:"ads_streams_replicated"`

	// The number of block specs replicated by this job.
	BlockSpecsReplicated int32 `json:"block_specs_replicated"`

	// The number of bytes recoverable by this job.
	BytesRecoverable int32 `json:"bytes_recoverable"`

	// The number of bytes that have been transferred by this job.
	BytesTransferred int32 `json:"bytes_transferred"`

	// The number of char specs replicated by this job.
	CharSpecsReplicated int32 `json:"char_specs_replicated"`

	// The number of LINs corrected by this job.
	CorrectedLins int32 `json:"corrected_lins"`

	// This field is true if the node running this job is dead.
	DeadNode bool `json:"dead_node"`

	// The number of directories replicated.
	DirectoriesReplicated int32 `json:"directories_replicated"`

	// The number of directories changed by this job.
	DirsChanged int32 `json:"dirs_changed"`

	// The number of directories deleted by this job.
	DirsDeleted int32 `json:"dirs_deleted"`

	// The number of directories moved by this job.
	DirsMoved int32 `json:"dirs_moved"`

	// The number of directories created by this job.
	DirsNew int32 `json:"dirs_new"`

	// The amount of time in seconds between when the job was started and when it ended.  If the job has not yet ended, this is the amount of time since the job started.  This field is null if the job has not yet started.
	Duration int32 `json:"duration,omitempty"`

	// The time the job ended in unix epoch seconds. The field is null if the job hasn't ended.
	EndTime int32 `json:"end_time,omitempty"`

	// The primary error message for this job.
	Error_ string `json:"error"`

	// The number of files with checksum errors skipped by this job.
	ErrorChecksumFilesSkipped int32 `json:"error_checksum_files_skipped"`

	// The number of files with io errors skipped by this job.
	ErrorIoFilesSkipped int32 `json:"error_io_files_skipped"`

	// The number of files with network errors skipped by this job.
	ErrorNetFilesSkipped int32 `json:"error_net_files_skipped"`

	// A list of error messages for this job.
	Errors []string `json:"errors"`

	// Tyhe number of data chunks that failed transmission.
	FailedChunks int32 `json:"failed_chunks"`

	// The number of fifos replicated by this job.
	FifosReplicated int32 `json:"fifos_replicated"`

	// The number of bytes transferred that belong to files.
	FileDataBytes int32 `json:"file_data_bytes"`

	// The number of files changed by this job.
	FilesChanged int32 `json:"files_changed"`

	// The number of files linked by this job.
	FilesLinked int32 `json:"files_linked"`

	// The number of files created by this job.
	FilesNew int32 `json:"files_new"`

	// The number of files selected by this job.
	FilesSelected int32 `json:"files_selected"`

	// The number of files transferred by this job.
	FilesTransferred int32 `json:"files_transferred"`

	// The number of files unlinked by this job.
	FilesUnlinked int32 `json:"files_unlinked"`

	// The number of files with ads replicated by this job.
	FilesWithAdsReplicated int32 `json:"files_with_ads_replicated"`

	// The number of LINs flipped by this job.
	FlippedLins int32 `json:"flipped_lins"`

	// The number of hard links replicated by this job.
	HardLinksReplicated int32 `json:"hard_links_replicated"`

	// The number of hash exceptions fixed by this job.
	HashExceptionsFixed int32 `json:"hash_exceptions_fixed"`

	// The number of hash exceptions found by this job.
	HashExceptionsFound int32 `json:"hash_exceptions_found"`

	// A unique identifier for this object.
	Id string `json:"id"`

	// The ID of the job.
	JobId int32 `json:"job_id,omitempty"`

	// The number of LINs transferred by this job.
	LinsTotal int32 `json:"lins_total"`

	// The total number of bytes sent to the source by this job.
	NetworkBytesToSource int32 `json:"network_bytes_to_source"`

	// The total number of bytes sent to the target by this job.
	NetworkBytesToTarget int32 `json:"network_bytes_to_target"`

	// The number of new files replicated by this job.
	NewFilesReplicated int32 `json:"new_files_replicated"`

	// The number of files that have been retransmitted by this job.
	NumRetransmittedFiles int32 `json:"num_retransmitted_files"`

	// Data for each phase of this job.
	Phases []SyncJobPhase `json:"phases"`

	// The ID of the policy.
	PolicyId string `json:"policy_id"`

	// The name of the policy.
	PolicyName string `json:"policy_name"`

	// The number of regular files replicated by this job.
	RegularFilesReplicated int32 `json:"regular_files_replicated"`

	// The number of LINs resynched by this job.
	ResyncedLins int32 `json:"resynced_lins"`

	// The files that have been retransmitted by this job.
	RetransmittedFiles []string `json:"retransmitted_files"`

	// The number of times the job has been retried.
	Retry int32 `json:"retry"`

	// The number of data chunks currently being transmitted.
	RunningChunks int32 `json:"running_chunks"`

	// The number of sockets replicated by this job.
	SocketsReplicated int32 `json:"sockets_replicated"`

	// The number of bytes recovered on the source.
	SourceBytesRecovered int32 `json:"source_bytes_recovered"`

	// The number of directories created on the source.
	SourceDirectoriesCreated int32 `json:"source_directories_created"`

	// The number of directories deleted on the source.
	SourceDirectoriesDeleted int32 `json:"source_directories_deleted"`

	// The number of directories linked on the source.
	SourceDirectoriesLinked int32 `json:"source_directories_linked"`

	// The number of directories unlinked on the source.
	SourceDirectoriesUnlinked int32 `json:"source_directories_unlinked"`

	// The number of directories visited on the source.
	SourceDirectoriesVisited int32 `json:"source_directories_visited"`

	// The number of files deleted on the source.
	SourceFilesDeleted int32 `json:"source_files_deleted"`

	// The number of files linked on the source.
	SourceFilesLinked int32 `json:"source_files_linked"`

	// The number of files unlinked on the source.
	SourceFilesUnlinked int32 `json:"source_files_unlinked"`

	// Hostname or IP address of sync source cluster.
	SourceHost string `json:"source_host"`

	// The number of sparse data bytes transferred by this job.
	SparseDataBytes int32 `json:"sparse_data_bytes"`

	// The time the job started in unix epoch seconds. The field is null if the job hasn't started.
	StartTime int32 `json:"start_time,omitempty"`

	// The state of the job.
	State string `json:"state"`

	// The number of subreports that are available for this job report.
	SubreportCount int32 `json:"subreport_count"`

	// The number of data chunks that have been transmitted successfully.
	SucceededChunks int32 `json:"succeeded_chunks"`

	// The number of symlinks replicated by this job.
	SymlinksReplicated int32 `json:"symlinks_replicated"`

	// The type of sync being performed by this job.
	SyncType string `json:"sync_type"`

	// The number of bytes recovered on the target.
	TargetBytesRecovered int32 `json:"target_bytes_recovered"`

	// The number of directories created on the target.
	TargetDirectoriesCreated int32 `json:"target_directories_created"`

	// The number of directories deleted on the target.
	TargetDirectoriesDeleted int32 `json:"target_directories_deleted"`

	// The number of directories linked on the target.
	TargetDirectoriesLinked int32 `json:"target_directories_linked"`

	// The number of directories unlinked on the target.
	TargetDirectoriesUnlinked int32 `json:"target_directories_unlinked"`

	// The number of files deleted on the target.
	TargetFilesDeleted int32 `json:"target_files_deleted"`

	// The number of files linked on the target.
	TargetFilesLinked int32 `json:"target_files_linked"`

	// The number of files unlinked on the target.
	TargetFilesUnlinked int32 `json:"target_files_unlinked"`

	// Absolute filesystem path on the target cluster for the sync destination.
	TargetPath string `json:"target_path"`

	// The target snapshots created by this job.
	TargetSnapshots []string `json:"target_snapshots"`

	// The total number of data chunks transmitted by this job.
	TotalChunks int32 `json:"total_chunks"`

	// The total number of bytes transferred by this job.
	TotalDataBytes int32 `json:"total_data_bytes"`

	// The number of files affected by this job.
	TotalFiles int32 `json:"total_files"`

	// The total number of bytes sent over the network by this job.
	TotalNetworkBytes int32 `json:"total_network_bytes"`

	// The total number of phases for this job.
	TotalPhases int32 `json:"total_phases"`

	// The number of bytes unchanged by this job.
	UnchangedDataBytes int32 `json:"unchanged_data_bytes"`

	// The number of up-to-date files skipped by this job.
	UpToDateFilesSkipped int32 `json:"up_to_date_files_skipped"`

	// The number of updated files replicated by this job.
	UpdatedFilesReplicated int32 `json:"updated_files_replicated"`

	// The number of files with user conflicts skipped by this job.
	UserConflictFilesSkipped int32 `json:"user_conflict_files_skipped"`

	// A list of warning messages for this job.
	Warnings []string `json:"warnings"`

	// The number of WORM committed files which needed to be reverted. Since WORM committed files cannot be reverted, this is the number of files that were preserved in the compliance store.
	WormCommittedFileConflicts int32 `json:"worm_committed_file_conflicts"`
}
