package appfile

import (
	"fmt"
)

// File is the structure of a single Appfile.
type File struct {
	// Path is the path to the root file that was loaded. This might be
	// empty if the appfile was parsed from an io.Reader.
	Path string

	Application    *Application
	Project        *Project
	Infrastructure []*Infrastructure
}

// Application is the structure of an application definition.
type Application struct {
	Name         string
	Type         string
	Dependencies []*Dependency `mapstructure:"dependency"`
}

// Dependency is another Appfile that an App depends on
type Dependency struct {
	Source string
}

// Project is the structure of a project that many applications
// can belong to.
type Project struct {
	Name           string
	Infrastructure string
	Stack          *Stack
}

// Infrastructure is the structure of defining the infrastructure
// that an application must run on.
type Infrastructure struct {
	Name   string
	Type   string
	Flavor string
}

// Stack is the structure that defines the stack that a project is
// built on.
type Stack struct {
	Name string
}

//-------------------------------------------------------------------
// Helper Methods
//-------------------------------------------------------------------

// ActiveInfrastructure returns the Infrastructure that is being
// used for this Appfile.
func (f *File) ActiveInfrastructure() *Infrastructure {
	for _, i := range f.Infrastructure {
		if i.Name == f.Project.Infrastructure {
			return i
		}
	}

	return nil
}

//-------------------------------------------------------------------
// GoStringer
//-------------------------------------------------------------------

func (v *Infrastructure) GoString() string {
	return fmt.Sprintf("*%#v", *v)
}

func (v *Project) GoString() string {
	return fmt.Sprintf("*%#v", *v)
}
