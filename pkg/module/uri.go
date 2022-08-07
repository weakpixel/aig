package module

// Autogenerated file

import (
	"fmt"
	"github.com/weakpixel/aig/pkg/types"
)

func init() {
	addModuleFactory("uri", func() types.Module {
		return NewUri()
	})
}

//
// Uri (uri) - Interacts with webservices
//
func NewUri() *Uri {
	module := Uri{}
	// Create dynamic param values
	paramValues := map[string]types.Value{}
	paramValues["body"] = types.NewStringValue(&module.Params.Body)
	paramValues["body_format"] = types.NewStringValue(&module.Params.BodyFormat)
	paramValues["ca_path"] = types.NewStringValue(&module.Params.CaPath)
	paramValues["client_cert"] = types.NewStringValue(&module.Params.ClientCert)
	paramValues["client_key"] = types.NewStringValue(&module.Params.ClientKey)
	paramValues["creates"] = types.NewStringValue(&module.Params.Creates)
	paramValues["dest"] = types.NewStringValue(&module.Params.Dest)
	paramValues["follow_redirects"] = types.NewStringValue(&module.Params.FollowRedirects)
	paramValues["force"] = types.NewBoolValue(&module.Params.Force)
	paramValues["force_basic_auth"] = types.NewBoolValue(&module.Params.ForceBasicAuth)
	paramValues["http_agent"] = types.NewStringValue(&module.Params.HttpAgent)
	paramValues["method"] = types.NewStringValue(&module.Params.Method)
	paramValues["remote_src"] = types.NewBoolValue(&module.Params.RemoteSrc)
	paramValues["removes"] = types.NewStringValue(&module.Params.Removes)
	paramValues["return_content"] = types.NewBoolValue(&module.Params.ReturnContent)
	paramValues["src"] = types.NewStringValue(&module.Params.Src)
	paramValues["timeout"] = types.NewIntValue(&module.Params.Timeout)
	paramValues["unix_socket"] = types.NewStringValue(&module.Params.UnixSocket)
	paramValues["unredirected_headers"] = types.NewStringArrayValue(&module.Params.UnredirectedHeaders)
	paramValues["url"] = types.NewStringValue(&module.Params.Url)
	paramValues["url_password"] = types.NewStringValue(&module.Params.UrlPassword)
	paramValues["url_username"] = types.NewStringValue(&module.Params.UrlUsername)
	paramValues["use_gssapi"] = types.NewBoolValue(&module.Params.UseGssapi)
	paramValues["use_proxy"] = types.NewBoolValue(&module.Params.UseProxy)
	paramValues["validate_certs"] = types.NewBoolValue(&module.Params.ValidateCerts)
	module.Params.values = paramValues

	// Create dynamic result values
	resultValues := map[string]types.Value{}

	resultValues["content"] = types.NewStringValue(&module.Result.Content)
	resultValues["cookies_string"] = types.NewStringValue(&module.Result.CookiesString)
	resultValues["elapsed"] = types.NewIntValue(&module.Result.Elapsed)
	resultValues["msg"] = types.NewStringValue(&module.Result.Msg)
	resultValues["path"] = types.NewStringValue(&module.Result.Path)
	resultValues["redirected"] = types.NewBoolValue(&module.Result.Redirected)
	resultValues["status"] = types.NewIntValue(&module.Result.Status)
	resultValues["url"] = types.NewStringValue(&module.Result.Url)
	module.Result.values = resultValues

	return &module
}

// Uri (uri) - Interacts with webservices
//
// Interacts with HTTP and HTTPS web services and supports Digest, Basic and WSSE HTTP authentication mechanisms.
//
// For Windows targets, use the M(ansible.windows.win_uri) module instead.
//
//
// Source: https://github.com/ansible/ansible/blob/v2.13.1/lib/ansible/modules/uri.py
type Uri struct {
	Params UriParams
	Result UriResult
}

type UriParams struct {

	// Body
	// The body of the http request/response to the web service. If C(body_format) is set to 'json' it will take an already formatted JSON string or convert a data structure into JSON.
	// If C(body_format) is set to 'form-urlencoded' it will convert a dictionary or list of tuples into an 'application/x-www-form-urlencoded' string. (Added in v2.7)
	// If C(body_format) is set to 'form-multipart' it will convert a dictionary into 'multipart/form-multipart' body. (Added in v2.10)
	//
	// Default: <no value>
	// Required: false
	Body string `yaml:"body,omitempty" json:"body,omitempty"`

	// BodyFormat
	// The serialization format of the body. When set to C(json), C(form-multipart), or C(form-urlencoded), encodes the body argument, if needed, and automatically sets the Content-Type header accordingly.
	// As of v2.3 it is possible to override the C(Content-Type) header, when set to C(json) or C(form-urlencoded) via the I(headers) option.
	// The 'Content-Type' header cannot be overridden when using C(form-multipart)
	// C(form-urlencoded) was added in v2.7.
	// C(form-multipart) was added in v2.10.
	//
	// Default: raw
	// Required: false
	BodyFormat string `yaml:"body_format,omitempty" json:"body_format,omitempty"`

	// CaPath
	// PEM formatted file that contains a CA certificate to be used for validation
	//
	// Default: <no value>
	// Required: false
	CaPath string `yaml:"ca_path,omitempty" json:"ca_path,omitempty"`

	// ClientCert
	// PEM formatted certificate chain file to be used for SSL client authentication.
	// This file can also include the key as well, and if the key is included, I(client_key) is not required
	//
	// Default: <no value>
	// Required: false
	ClientCert string `yaml:"client_cert,omitempty" json:"client_cert,omitempty"`

	// ClientKey
	// PEM formatted file that contains your private key to be used for SSL client authentication.
	// If I(client_cert) contains both the certificate and key, this option is not required.
	//
	// Default: <no value>
	// Required: false
	ClientKey string `yaml:"client_key,omitempty" json:"client_key,omitempty"`

	// Creates
	// A filename, when it already exists, this step will not be run.
	//
	// Default: <no value>
	// Required: false
	Creates string `yaml:"creates,omitempty" json:"creates,omitempty"`

	// Dest
	// A path of where to download the file to (if desired). If I(dest) is a directory, the basename of the file on the remote server will be used.
	//
	// Default: <no value>
	// Required: false
	Dest string `yaml:"dest,omitempty" json:"dest,omitempty"`

	// FollowRedirects
	// Whether or not the URI module should follow redirects. C(all) will follow all redirects. C(safe) will follow only "safe" redirects, where "safe" means that the client is only doing a GET or HEAD on the URI to which it is being redirected. C(none) will not follow any redirects. Note that C(yes) and C(no) choices are accepted for backwards compatibility, where C(yes) is the equivalent of C(all) and C(no) is the equivalent of C(safe). C(yes) and C(no) are deprecated and will be removed in some future version of Ansible.
	//
	// Default: safe
	// Required: false
	FollowRedirects string `yaml:"follow_redirects,omitempty" json:"follow_redirects,omitempty"`

	// Force
	// If C(yes) do not get a cached copy.
	//
	// Default: no
	// Required: false
	Force bool `yaml:"force,omitempty" json:"force,omitempty"`

	// ForceBasicAuth
	// Force the sending of the Basic authentication header upon initial request.
	// The library used by the uri module only sends authentication information when a webservice responds to an initial request with a 401 status. Since some basic auth services do not properly send a 401, logins will fail.
	//
	// Default: no
	// Required: false
	ForceBasicAuth bool `yaml:"force_basic_auth,omitempty" json:"force_basic_auth,omitempty"`

	// Headers
	// Add custom HTTP headers to a request in the format of a YAML hash. As of C(2.3) supplying C(Content-Type) here will override the header generated by supplying C(json) or C(form-urlencoded) for I(body_format).
	//
	// Default: <no value>
	// Required: false
	Headers map[string]interface{} `yaml:"headers,omitempty" json:"headers,omitempty"`

	// HttpAgent
	// Header to identify as, generally appears in web server logs.
	//
	// Default: ansible-httpget
	// Required: false
	HttpAgent string `yaml:"http_agent,omitempty" json:"http_agent,omitempty"`

	// Method
	// The HTTP method of the request or response.
	// In more recent versions we do not restrict the method at the module level anymore but it still must be a valid method accepted by the service handling the request.
	//
	// Default: GET
	// Required: false
	Method string `yaml:"method,omitempty" json:"method,omitempty"`

	// RemoteSrc
	// If C(no), the module will search for the C(src) on the controller node.
	// If C(yes), the module will search for the C(src) on the managed (remote) node.
	//
	// Default: no
	// Required: false
	RemoteSrc bool `yaml:"remote_src,omitempty" json:"remote_src,omitempty"`

	// Removes
	// A filename, when it does not exist, this step will not be run.
	//
	// Default: <no value>
	// Required: false
	Removes string `yaml:"removes,omitempty" json:"removes,omitempty"`

	// ReturnContent
	// Whether or not to return the body of the response as a "content" key in the dictionary result no matter it succeeded or failed.
	// Independently of this option, if the reported Content-type is "application/json", then the JSON is always loaded into a key called C(json) in the dictionary results.
	//
	// Default: no
	// Required: false
	ReturnContent bool `yaml:"return_content,omitempty" json:"return_content,omitempty"`

	// Src
	// Path to file to be submitted to the remote server.
	// Cannot be used with I(body).
	//
	// Default: <no value>
	// Required: false
	Src string `yaml:"src,omitempty" json:"src,omitempty"`

	// StatusCode
	// A list of valid, numeric, HTTP status codes that signifies success of the request.
	//
	// Default: [200]
	// Required: false
	StatusCode []int `yaml:"status_code,omitempty" json:"status_code,omitempty"`

	// Timeout
	// The socket level timeout in seconds
	//
	// Default: 30
	// Required: false
	Timeout int `yaml:"timeout,omitempty" json:"timeout,omitempty"`

	// UnixSocket
	// Path to Unix domain socket to use for connection
	//
	// Default: <no value>
	// Required: false
	UnixSocket string `yaml:"unix_socket,omitempty" json:"unix_socket,omitempty"`

	// UnredirectedHeaders
	// A list of header names that will not be sent on subsequent redirected requests. This list is case insensitive. By default all headers will be redirected. In some cases it may be beneficial to list headers such as C(Authorization) here to avoid potential credential exposure.
	//
	// Default: []
	// Required: false
	UnredirectedHeaders []string `yaml:"unredirected_headers,omitempty" json:"unredirected_headers,omitempty"`

	// Url
	// HTTP or HTTPS URL in the form (http|https)://host.domain[:port]/path
	//
	// Default: <no value>
	// Required: true
	Url string `yaml:"url,omitempty" json:"url,omitempty"`

	// UrlPassword
	// A password for the module to use for Digest, Basic or WSSE authentication.
	//
	// Default: <no value>
	// Required: false
	UrlPassword string `yaml:"url_password,omitempty" json:"url_password,omitempty"`

	// UrlUsername
	// A username for the module to use for Digest, Basic or WSSE authentication.
	//
	// Default: <no value>
	// Required: false
	UrlUsername string `yaml:"url_username,omitempty" json:"url_username,omitempty"`

	// UseGssapi
	// Use GSSAPI to perform the authentication, typically this is for Kerberos or Kerberos through Negotiate authentication.
	// Requires the Python library L(gssapi,https://github.com/pythongssapi/python-gssapi) to be installed.
	// Credentials for GSSAPI can be specified with I(url_username)/I(url_password) or with the GSSAPI env var C(KRB5CCNAME) that specified a custom Kerberos credential cache.
	// NTLM authentication is C(not) supported even if the GSSAPI mech for NTLM has been installed.
	//
	// Default: no
	// Required: false
	UseGssapi bool `yaml:"use_gssapi,omitempty" json:"use_gssapi,omitempty"`

	// UseProxy
	// If C(no), it will not use a proxy, even if one is defined in an environment variable on the target hosts.
	//
	// Default: yes
	// Required: false
	UseProxy bool `yaml:"use_proxy,omitempty" json:"use_proxy,omitempty"`

	// ValidateCerts
	// If C(no), SSL certificates will not be validated.
	// This should only set to C(no) used on personally controlled sites using self-signed certificates.
	// Prior to 1.9.2 the code defaulted to C(no).
	//
	// Default: yes
	// Required: false
	ValidateCerts bool `yaml:"validate_certs,omitempty" json:"validate_certs,omitempty"`

	values map[string]types.Value
}

func (p *UriParams) Names() []string {
	names := []string{}
	for name := range p.values {
		names = append(names, name)
	}
	return []string{}
}

func (p *UriParams) Set(name string, value interface{}) error {
	v, ok := p.values[name]
	if !ok {
		return fmt.Errorf("no param with name %q", name)
	}
	return v.Set(value)
}

func (p *UriParams) Get(name string) (interface{}, error) {
	v, ok := p.values[name]
	if !ok {
		return nil, fmt.Errorf("no param with name %q", name)
	}
	return v.Get(), nil
}

type UriResult struct {
	types.CommonReturn
	Raw string

	// Content
	// The response body content.
	Content string `yaml:"content,omitempty" json:"content,omitempty"`

	// Cookies
	// The cookie values placed in cookie jar.
	Cookies map[string]interface{} `yaml:"cookies,omitempty" json:"cookies,omitempty"`

	// CookiesString
	// The value for future request Cookie headers.
	CookiesString string `yaml:"cookies_string,omitempty" json:"cookies_string,omitempty"`

	// Elapsed
	// The number of seconds that elapsed while performing the download.
	Elapsed int `yaml:"elapsed,omitempty" json:"elapsed,omitempty"`

	// Msg
	// The HTTP message from the request.
	Msg string `yaml:"msg,omitempty" json:"msg,omitempty"`

	// Path
	// destination file/path
	Path string `yaml:"path,omitempty" json:"path,omitempty"`

	// Redirected
	// Whether the request was redirected.
	Redirected bool `yaml:"redirected,omitempty" json:"redirected,omitempty"`

	// Status
	// The HTTP status code from the request.
	Status int `yaml:"status,omitempty" json:"status,omitempty"`

	// Url
	// The actual URL used for the request.
	Url string `yaml:"url,omitempty" json:"url,omitempty"`

	values map[string]types.Value
}

func (r *UriResult) Names() []string {
	names := []string{}
	for name := range r.values {
		names = append(names, name)
	}
	return []string{}
}

func (r *UriResult) Set(name string, value interface{}) error {
	v, ok := r.values[name]
	if !ok {
		return fmt.Errorf("no param with name %q", name)
	}
	return v.Set(value)
}

func (r *UriResult) Get(name string) (interface{}, error) {
	v, ok := r.values[name]
	if !ok {
		return nil, fmt.Errorf("no param with name %q", name)
	}
	return v.Get(), nil
}

func (m *Uri) GetResult() types.Result {
	return &m.Result
}

func (m *Uri) GetResultRaw() string {
	return m.Result.Raw
}

func (m *Uri) GetCommonResult() types.CommonReturn {
	return m.Result.CommonReturn
}

func (m *Uri) GetParams() types.Params {
	return &m.Params
}

func (m *Uri) GetType() string {
	return "uri"
}
