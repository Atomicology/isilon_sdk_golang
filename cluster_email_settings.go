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

type ClusterEmailSettings struct {

	// This setting determines how notifications will be batched together to be sent by email.  'none' means each notification will be sent separately.  'severity' means notifications of the same severity will be sent together.  'category' means notifications of the same category will be sent together.  'all' means all notifications will be batched together and sent in a single email.
	BatchMode string `json:"batch_mode"`

	// The address of the SMTP server to be used for relaying the notification messages.  An SMTP server is required in order to send notifications.  If this string is empty, no emails will be sent.
	MailRelay string `json:"mail_relay"`

	// The full email address that will appear as the sender of notification messages.
	MailSender string `json:"mail_sender"`

	// The subject line for notification messages from this cluster.
	MailSubject string `json:"mail_subject"`

	// Indicates if an SMTP authentication password is set.
	SmtpAuthPasswdSet bool `json:"smtp_auth_passwd_set"`

	// The type of secure communication protocol to use if SMTP is being used.  If 'none', plain text will be used, if 'starttls', the encrypted STARTTLS protocol will be used.
	SmtpAuthSecurity string `json:"smtp_auth_security"`

	// Username to authenticate with if SMTP authentication is being used.
	SmtpAuthUsername string `json:"smtp_auth_username"`

	// The port on the SMTP server to be used for relaying the notification messages.  
	SmtpPort int32 `json:"smtp_port"`

	// If true, this cluster will send SMTP authentication credentials to the SMTP relay server in order to send its notification emails.  If false, the cluster will attempt to send its notification emails without authentication.
	UseSmtpAuth bool `json:"use_smtp_auth"`

	// Location of a custom template file that can be used to specify the layout of the notification emails.
	UserTemplate string `json:"user_template,omitempty"`
}
