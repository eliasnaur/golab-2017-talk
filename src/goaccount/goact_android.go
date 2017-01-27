package goaccount

import (
	"Java/android/os"             // HLimport
	"Java/android/support/v7/app" // HLimport

	// The Java package "goaccount".
	gopkg "Java/goaccount" // HLgopkg
)

type GoActivity struct {
	app.AppCompatActivity // HLembed
}

func (a *GoActivity) OnCreate(this gopkg.GoActivity, b os.Bundle) { // HLoncreate
	this.Super().OnCreate(b) // HLsuper
	setupActivity(this)
}
