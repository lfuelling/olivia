package modules

import (
	"fmt"
	"github.com/gookit/color"
	"math/rand"
	"strings"
	"time"
)

// HelpTag is the intent tag for its module
var HelpTag = "help"

// HelpReplacer returns the help text
// See modules/modules.go#Module.Replacer() for more details.
func HelpReplacer(locale, entry, response, _ string) (string, string) {

	filteredModules := modules[locale][:0]

	for _, module := range modules[locale] {
		// module.title doesn't exist so I'm using the tag in this example
		// Tags should never be empty but in my fork Titles can be
		// so filtering is needed in that case.
		// Original code:
		//if module.Title != "" {
		if module.Tag != "" {
			filteredModules = append(filteredModules, module)
		}
	}

	requestedModule := FindCapability(entry, filteredModules)

	var sb strings.Builder

	// same here, using the tag instead of the non-existing title
	//if requestedModule.Title == "" {
	if requestedModule.Tag == "" {
		sb.WriteString("I was unable to find that capability!")
	// help texts also don't exist
	//} else if requestedModule.HelpText == "" {
	} else if requestedModule.Tag == "" {
		sb.WriteString("The capability has no help text (yet)...")
	} else {
		// again using tag instead of help text
		//sb.WriteString(requestedModule.HelpText)
		sb.WriteString(requestedModule.Tag)
		sb.WriteString(" ")

		examplePattern := requestedModule.Patterns[0]
		if len(requestedModule.Patterns) > 1 {
			rand.Seed(time.Now().UnixNano())
			examplePattern = requestedModule.Patterns[rand.Intn(len(requestedModule.Patterns))]
		}

		sb.WriteString("For example you can say '" + examplePattern + "'.")
	}
	return HelpTag, fmt.Sprintf(response, sb.String())
}

func FindCapability(sentence string, filteredModules []Module) Module {
	green := color.FgGreen.Render
	red := color.FgRed.Render

	sentenceToLower := strings.ToLower(sentence)

	for _, module := range filteredModules {

		// and again using tag instead of title
		//moduleTitle := strings.ToLower(module.Title)
		moduleTitle := strings.ToLower(module.Tag)

		fmt.Printf("    %s %s - %s (%s)\n", green("->︎"), "checking module", red(moduleTitle), sentenceToLower)

		if strings.HasSuffix(sentenceToLower, moduleTitle) {
			// Returns the right module
			return module
		}
	}

	// Returns an empty module if none has been found
	return Module{}
}
