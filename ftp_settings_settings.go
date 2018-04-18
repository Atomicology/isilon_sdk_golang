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

type FtpSettingsSettings struct {

	// The timeout, in seconds, for a remote client to establish a PASV style data connection.
	AcceptTimeout int32 `json:"accept_timeout,omitempty"`

	// Controls whether anonymous logins are permitted or not.
	AllowAnonAccess bool `json:"allow_anon_access,omitempty"`

	// Controls whether anonymous users will be permitted to upload files.
	AllowAnonUpload bool `json:"allow_anon_upload,omitempty"`

	// If set to false, all directory list commands will return a permission denied error.
	AllowDirlists bool `json:"allow_dirlists,omitempty"`

	// If set to false, all downloads requests will return a permission denied error.
	AllowDownloads bool `json:"allow_downloads,omitempty"`

	// Controls whether local logins are permitted or not.
	AllowLocalAccess bool `json:"allow_local_access,omitempty"`

	// This controls whether any FTP commands which change the filesystem are allowed or not.
	AllowWrites bool `json:"allow_writes,omitempty"`

	// This controls whether FTP will always initially change directories to the home directory of the user, regardless of whether it is chroot-ing.
	AlwaysChdirHomedir bool `json:"always_chdir_homedir,omitempty"`

	// This is the name of the user who is given ownership of anonymously uploaded files.
	AnonChownUsername string `json:"anon_chown_username,omitempty"`

	// A list of passwords for anonymous users.
	AnonPasswordList []string `json:"anon_password_list,omitempty"`

	// This option represents a directory in /ifs which vsftpd will try to change into after an anonymous login.
	AnonRootPath string `json:"anon_root_path,omitempty"`

	// The value that the umask for file creation is set to for anonymous users.
	AnonUmask int32 `json:"anon_umask,omitempty"`

	// Controls whether ascii mode data transfers are honored for various types of requests.
	AsciiMode string `json:"ascii_mode,omitempty"`

	// A list of users that are not chrooted when logging in.
	ChrootExceptionList []string `json:"chroot_exception_list,omitempty"`

	// If set to 'all', all local users will be (by default) placed in a chroot() jail in their home directory after login. If set to 'all-with-exceptions', all local users except those listed in the chroot exception list (isi ftp chroot-exception-list) will be placed in a chroot() jail in their home directory after login. If set to 'none', no local users will be chrooted by default. If set to 'none-with-exceptions', only the local users listed in the chroot exception list (isi ftp chroot-exception-list) will be place in a chroot() jail in their home directory after login.
	ChrootLocalMode string `json:"chroot_local_mode,omitempty"`

	// The timeout, in seconds, for a remote client to respond to our PORT style data connection.
	ConnectTimeout int32 `json:"connect_timeout,omitempty"`

	// The timeout, in seconds, which is roughly the maximum time we permit data transfers to stall for with no progress. If the timeout triggers, the remote client is kicked off.
	DataTimeout int32 `json:"data_timeout,omitempty"`

	// A list of uses that will be denied access.
	DeniedUserList []string `json:"denied_user_list,omitempty"`

	// If enabled, display directory listings with the time in your local time zone. The default is to display GMT. The times returned by the MDTM FTP command are also affected by this option.
	DirlistLocaltime bool `json:"dirlist_localtime,omitempty"`

	// When set to 'hide',  all user and group information in directory listings will be displayed as 'ftp'. When set to 'textual', textual names are shown in the user and group fields of directory listings. When set to 'numeric', numeric IDs are show in the user and group fields of directory listings.
	DirlistNames string `json:"dirlist_names,omitempty"`

	// The permissions with which uploaded files are created. Umasks are applied on top of this value.
	FileCreatePerm int32 `json:"file_create_perm,omitempty"`

	// This field determines whether the anon_password_list is used.
	LimitAnonPasswords bool `json:"limit_anon_passwords,omitempty"`

	// This option represents a directory in /ifs which vsftpd will try to change into after a local login.
	LocalRootPath string `json:"local_root_path,omitempty"`

	// The value that the umask for file creation is set to for local users.
	LocalUmask int32 `json:"local_umask,omitempty"`

	// If enabled, allow server-to-server (FXP) transfers.
	ServerToServer bool `json:"server_to_server,omitempty"`

	// This field controls whether the FTP daemon is running.
	Service bool `json:"service,omitempty"`

	// If enabled, maintain login sessions for each user through Pluggable Authentication Modules (PAM). Disabling this option prevents the ability to do automatic home directory creation if that functionality were otherwise available.
	SessionSupport bool `json:"session_support,omitempty"`

	// The timeout, in seconds, which is roughly the maximum time we permit data transfers to stall for with no progress. If the timeout triggers, the remote client is kicked off.
	SessionTimeout int32 `json:"session_timeout,omitempty"`

	// Specifies the directory where per-user config overrides can be found.
	UserConfigDir string `json:"user_config_dir,omitempty"`
}
