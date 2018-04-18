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

type ClusterNodeHardware struct {

	// Name of this node's chassis.
	Chassis string `json:"chassis,omitempty"`

	// Chassis code of this node (1U, 2U, etc.).
	ChassisCode string `json:"chassis_code,omitempty"`

	// Number of chassis making up this node.
	ChassisCount string `json:"chassis_count,omitempty"`

	// Chassis depth for this node if applicable (Normal, Deep, Unknown). If not supported: Unknown. If Logic to determine chassis depth fails: Unknown. If PSI_Get fails: Unknown. PSI_Get can fail if PSI not initialized, or key does not exist.
	ChassisDepth string `json:"chassis_depth,omitempty"`

	// Class of this node (storage, accelerator, etc.).
	Class string `json:"class,omitempty"`

	// Code name of this node if applicable (Minnetonka, MiniHuron, Huron, Union, Tahoe, Superior, Unknown). If not supported: Unknown. If Logic to determine code name fails: Unknown. If PSI_Get fails: Unknown. PSI_Get can fail if PSI not initialized, or key does not exist.
	CodeName string `json:"code_name,omitempty"`

	// Type of compute node if applicable (Low, Medium, High, Turbo, Ultra, Unknown). If not supported: Unknown. If Logic to determine compute type fails: Unknown. If PSI_Get fails: Unknown. PSI_Get can fail if PSI not initialized, or key does not exist.
	ComputeType string `json:"compute_type,omitempty"`

	// Node configuration ID.
	ConfigurationId string `json:"configuration_id,omitempty"`

	// Manufacturer and model of this node's CPU.
	Cpu string `json:"cpu,omitempty"`

	// Manufacturer and model of this node's disk controller.
	DiskController string `json:"disk_controller,omitempty"`

	// Manufacturer and model of this node's disk expander.
	DiskExpander string `json:"disk_expander,omitempty"`

	// Family code of this node (X, S, NL, etc.).
	FamilyCode string `json:"family_code,omitempty"`

	// Manufacturer, model, and device id of this node's flash drive.
	FlashDrive string `json:"flash_drive,omitempty"`

	// Generation code of this node.
	GenerationCode string `json:"generation_code,omitempty"`

	// Isilon hardware generation name.
	Hwgen string `json:"hwgen,omitempty"`

	// Version of this node's Isilon Management Board.
	ImbVersion string `json:"imb_version,omitempty"`

	// Infiniband card type.
	Infiniband string `json:"infiniband,omitempty"`

	// Version of the LCD panel.
	LcdVersion string `json:"lcd_version,omitempty"`

	// Isilon node model identifier string (S200, X410, Infinity-H500, etc.).
	Model string `json:"model,omitempty"`

	// Isilon node model code string (S200, X410, H500, etc.).
	ModelCode string `json:"model_code,omitempty"`

	// Manufacturer and model of this node's motherboard.
	Motherboard string `json:"motherboard,omitempty"`

	// Description of all this node's network interfaces.
	NetInterfaces string `json:"net_interfaces,omitempty"`

	// Position of node within chassis (e.g., 1-4 for Infinity chassis). -1 for error or not supported.
	NodeSlotId int32 `json:"node_slot_id,omitempty"`

	// Manufacturer and model of this node's NVRAM board.
	Nvram string `json:"nvram,omitempty"`

	// Serial number of this node's peer/buddy node.(Infinity Only)
	PeerSerialNumber string `json:"peer_serial_number,omitempty"`

	// Performance code of this node, if applicable (2, 4, 5, etc.).
	PerformanceCode string `json:"performance_code,omitempty"`

	// Description strings for each power supply on this node.
	Powersupplies []string `json:"powersupplies,omitempty"`

	// Number of processors and cores on this node.
	Processor string `json:"processor,omitempty"`

	// Isilon product name.
	Product string `json:"product,omitempty"`

	// Size of RAM in bytes.
	Ram int32 `json:"ram,omitempty"`

	// Serial number of this node.
	SerialNumber string `json:"serial_number,omitempty"`

	// Series of this node (X, I, NL, etc.).
	Series string `json:"series,omitempty"`

	// Size of drive sleds in node, if applicable. Expected values: 3, 4, 6. 0 if unable to determine sled size. -1 for error or not supported. If PSI_Get fails: -1. PSI_Get can fail if PSI not initialized, or key does not exist.
	SledDriveCount int32 `json:"sled_drive_count,omitempty"`

	// Storage class of this node (storage or diskless).
	StorageClass string `json:"storage_class,omitempty"`

	// Platform tier level of this node if applicable (1-4 are defined, 0 for unknown or not supported, -1 for error). If not supported: 0. If Logic to determine tier fails: 0 for unknown. If PSI_Get fails: -1 for error. PSI_Get can fail if PSI not initialized, or key does not exist.
	Tier int32 `json:"tier,omitempty"`

	// Serial number of the top level assembly of this node.(Infinity Only)
	TopLevelAssemblySerialNumber string `json:"top_level_assembly_serial_number"`
}