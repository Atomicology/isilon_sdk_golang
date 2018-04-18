# FtpSettingsExtended

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AcceptTimeout** | **int32** | The timeout, in seconds, for a remote client to establish a PASV style data connection. | [optional] [default to null]
**AllowAnonAccess** | **bool** | Controls whether anonymous logins are permitted or not. | [optional] [default to null]
**AllowAnonUpload** | **bool** | Controls whether anonymous users will be permitted to upload files. | [optional] [default to null]
**AllowDirlists** | **bool** | If set to false, all directory list commands will return a permission denied error. | [optional] [default to null]
**AllowDownloads** | **bool** | If set to false, all downloads requests will return a permission denied error. | [optional] [default to null]
**AllowLocalAccess** | **bool** | Controls whether local logins are permitted or not. | [optional] [default to null]
**AllowWrites** | **bool** | This controls whether any FTP commands which change the filesystem are allowed or not. | [optional] [default to null]
**AlwaysChdirHomedir** | **bool** | This controls whether FTP will always initially change directories to the home directory of the user, regardless of whether it is chroot-ing. | [optional] [default to null]
**AnonChownUsername** | **string** | This is the name of the user who is given ownership of anonymously uploaded files. | [optional] [default to null]
**AnonPasswordList** | **[]string** | A list of passwords for anonymous users. | [optional] [default to null]
**AnonRootPath** | **string** | This option represents a directory in /ifs which vsftpd will try to change into after an anonymous login. | [optional] [default to null]
**AnonUmask** | **int32** | The value that the umask for file creation is set to for anonymous users. | [optional] [default to null]
**AsciiMode** | **string** | Controls whether ascii mode data transfers are honored for various types of requests. | [optional] [default to null]
**ChrootExceptionList** | **[]string** | A list of users that are not chrooted when logging in. | [optional] [default to null]
**ChrootLocalMode** | **string** | If set to &#39;all&#39;, all local users will be (by default) placed in a chroot() jail in their home directory after login. If set to &#39;all-with-exceptions&#39;, all local users except those listed in the chroot exception list (isi ftp chroot-exception-list) will be placed in a chroot() jail in their home directory after login. If set to &#39;none&#39;, no local users will be chrooted by default. If set to &#39;none-with-exceptions&#39;, only the local users listed in the chroot exception list (isi ftp chroot-exception-list) will be place in a chroot() jail in their home directory after login. | [optional] [default to null]
**ConnectTimeout** | **int32** | The timeout, in seconds, for a remote client to respond to our PORT style data connection. | [optional] [default to null]
**DataTimeout** | **int32** | The timeout, in seconds, which is roughly the maximum time we permit data transfers to stall for with no progress. If the timeout triggers, the remote client is kicked off. | [optional] [default to null]
**DeniedUserList** | **[]string** | A list of uses that will be denied access. | [optional] [default to null]
**DirlistLocaltime** | **bool** | If enabled, display directory listings with the time in your local time zone. The default is to display GMT. The times returned by the MDTM FTP command are also affected by this option. | [optional] [default to null]
**DirlistNames** | **string** | When set to &#39;hide&#39;,  all user and group information in directory listings will be displayed as &#39;ftp&#39;. When set to &#39;textual&#39;, textual names are shown in the user and group fields of directory listings. When set to &#39;numeric&#39;, numeric IDs are show in the user and group fields of directory listings. | [optional] [default to null]
**FileCreatePerm** | **int32** | The permissions with which uploaded files are created. Umasks are applied on top of this value. | [optional] [default to null]
**LimitAnonPasswords** | **bool** | This field determines whether the anon_password_list is used. | [optional] [default to null]
**LocalRootPath** | **string** | This option represents a directory in /ifs which vsftpd will try to change into after a local login. | [optional] [default to null]
**LocalUmask** | **int32** | The value that the umask for file creation is set to for local users. | [optional] [default to null]
**ServerToServer** | **bool** | If enabled, allow server-to-server (FXP) transfers. | [optional] [default to null]
**Service** | **bool** | This field controls whether the FTP daemon is running. | [optional] [default to null]
**SessionSupport** | **bool** | If enabled, maintain login sessions for each user through Pluggable Authentication Modules (PAM). Disabling this option prevents the ability to do automatic home directory creation if that functionality were otherwise available. | [optional] [default to null]
**SessionTimeout** | **int32** | The timeout, in seconds, which is roughly the maximum time we permit data transfers to stall for with no progress. If the timeout triggers, the remote client is kicked off. | [optional] [default to null]
**UserConfigDir** | **string** | Specifies the directory where per-user config overrides can be found. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


