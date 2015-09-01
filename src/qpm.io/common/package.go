// Copyright 2015 Cutehacks AS. All rights reserved.
// License can be found in the LICENSE file.

package common

import (
	//	"crypto/x509/pkix"
	//	"debug/elf"
	"bytes"
	"encoding/json"
	"fmt"
	"os"
	msg "qpm.io/common/messages"
	"qpm.io/qpm/core"
	"regexp"
	"strings"
	"text/template"
	"path/filepath"
)

const (
	ERR_REQUIRED_FIELD  = "%s is a required field"
	ERR_FORMATTED_FIELD = "%s requires a specific format"
)

var (
	regexPackageName = regexp.MustCompile("^[a-zA-Z]{2,}\\.[a-zA-Z0-9][a-zA-Z0-9\\-]{0,61}[a-zA-Z0-9]?\\.[a-zA-Z0-9][a-zA-Z0-9\\-]{0,61}[a-zA-Z0-9]?$")
	regexVersion     = regexp.MustCompile("[0-9].[0-9].[0-9]*")
	regexAuthorName  = regexp.MustCompile("^[\\p{L}\\s'.-]+$")
	regexAuthorEmail = regexp.MustCompile(".+@.+\\..+")
	regexGitSha1     = regexp.MustCompile("^[a-fA-F0-9]{8,}$")
	regexRevision    = map[msg.RepoType]*regexp.Regexp{
		msg.RepoType_GITHUB: regexGitSha1,
	}
)

func dotSlash(dots string) string {
	return strings.Replace(dots, ".", "/", -1)
}

func dotUnderscore(dots string) string {
	return strings.Replace(dots, ".", "_", -1)
}

type DependencyList map[string]string

// Creates a new DependencyList (which is really a map) which takes a list of package
// names of the form "package@version" and produces a map of "package => "version".
// Passing in multiple versions of the same package will overwrite with the last one.
func NewDependencyList(packages []string) DependencyList {
	deps := DependencyList{}
	for _, dep := range packages {
		parts := strings.Split(dep, "@")
		pName := strings.ToLower(parts[0])

		if len(parts) > 1 {
			deps[pName] = strings.ToLower(parts[1])
		} else {
			deps[pName] = ""
		}
	}
	return deps
}

func NewPackage() *msg.Package {
	return &msg.Package{
		Name:        "",
		Description: "",
		Version: &msg.Package_Version{
			Label: "0.0.1",
		},
		Author: &msg.Package_Author{
			Name:  "",
			Email: "",
		},
		Dependencies: []string{},
		Repository: &msg.Package_Repository{
			Type: msg.RepoType_GITHUB,
			Url:  "",
		},
		License: msg.LicenseType_MIT,
	}
}

type PackageWrapper struct {
	*msg.Package
	ID int // only used on server
	FilePath string
}

func LoadPackage(path string) (*PackageWrapper, error) {
	pw := &PackageWrapper{Package: &msg.Package{}}

	packageFile := filepath.Join(path, core.PackageFile)

	if _, err := os.Stat(packageFile); err == nil {

		file, err := os.Open(packageFile)
		if err != nil {
			return pw, err
		}
		defer file.Close()

		dec := json.NewDecoder(file)

		if err := dec.Decode(pw.Package); err != nil {
			return pw, err
		}

		pw.FilePath, err = filepath.Abs(file.Name())
		if err != nil {
			return pw, err
		}
	}

	return pw, nil
}

func (pw PackageWrapper) Save() error {
	var file *os.File
	var err error

	// re-create the core.PackageFile file
	file, err = os.Create(core.PackageFile)

	if err != nil {
		return err
	}
	defer file.Close()

	b, err := json.Marshal(pw.Package)
	if err != nil {
		return err
	}

	var buf bytes.Buffer
	json.Indent(&buf, b, "", "  ")
	buf.WriteTo(file)
	return nil
}

func (pw PackageWrapper) ParseDependencies() DependencyList {
	return NewDependencyList(pw.Dependencies)
}

func (pw PackageWrapper) Validate() error {
	if pw.Name == "" {
		return fmt.Errorf(ERR_REQUIRED_FIELD, "name")
	} else {
		// Validate name
		if !regexPackageName.MatchString(pw.Name) {
			return fmt.Errorf(ERR_FORMATTED_FIELD, "name")
		}
	}
	if pw.Version == nil {
		return fmt.Errorf(ERR_REQUIRED_FIELD, "version")
	} else {
		// Validate version label
		if !regexVersion.MatchString(pw.Version.Label) {
			return fmt.Errorf(ERR_FORMATTED_FIELD, "version label")
		}
		// Validate version revision
		if pw.Version.Revision == "" {
			return fmt.Errorf(ERR_REQUIRED_FIELD, "version revision")
		} else {
			if !regexRevision[pw.Repository.Type].MatchString(pw.Version.Revision) {
				return fmt.Errorf(ERR_FORMATTED_FIELD, "version revision")
			}
		}
	}
	if pw.Author == nil {
		return fmt.Errorf(ERR_REQUIRED_FIELD, "author")
	} else {
		// Validate author name
		if !regexAuthorName.MatchString(pw.Author.Name) {
			return fmt.Errorf(ERR_FORMATTED_FIELD, "author name")
		}
		//Validate author email
		if !regexAuthorEmail.MatchString(pw.Author.Email) {
			return fmt.Errorf(ERR_FORMATTED_FIELD, "author email")
		}
	}

	return nil
}

func (pw PackageWrapper) PriFile() string {
	return dotUnderscore(pw.Package.Name) + ".pri"
}

func (pw PackageWrapper) QrcFile() string {
	return dotUnderscore(pw.Package.Name) + ".qrc"
}

func (pw PackageWrapper) QrcPrefix() string {
	return dotSlash(pw.Package.Name)
}

var (
	vendorPri = template.Must(template.New("vendorPri").Parse(`
DEFINES += QPM_INIT\\(E\\)=\"E.addImportPath(QStringLiteral(\\\"qrc:/\\\"));\"
{{range $dep := .}}include($$PWD/{{$dep.Name}}/{{$dep.Name}}.pri){{"\n"}}{{end}}
`))
)

func (pw PackageWrapper) writePri(filename string, tpl *template.Template, data interface{}) error {

	var file *os.File
	var err error

	// create the vendor directory if needed
	if _, err = os.Stat(core.Vendor); err != nil {
		err = os.Mkdir(core.Vendor, 0755)
	}

	// re-create the .pri file
	file, err = os.Create(filename)

	if err != nil {
		return err
	}
	defer file.Close()

	err = tpl.Execute(file, data)
	if err != nil {
		return err
	}

	return nil
}

func (pw PackageWrapper) UpdatePri(dependencies []*msg.Dependency) error {
	return pw.writePri(core.Vendor+"/"+core.Vendor+".pri", vendorPri, dependencies)
}

// FIXME: attach this to Dependency
func (pw PackageWrapper) GetDependencySignature(d *msg.Dependency) string {
	if d != nil {
		if d.Version != nil {
			return d.Name + "@" + d.Version.Label
		}
		return d.Name
	}
	return ""
}
