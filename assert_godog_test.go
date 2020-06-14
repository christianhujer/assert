package assert

import (
	"flag"
	"github.com/cucumber/godog"
	"github.com/cucumber/godog/colors"
	"os"
	"testing"
)

var hotdogs int

func iHaveHotdogs(initialHotdogs int) error {
	hotdogs = initialHotdogs
	return nil
}

func iEatHotdogs(hotdogsToEat int) error {
	hotdogs -= hotdogsToEat
	return nil
}

func iMUSTHaveHotdogs(expectedHotdogsRemaining int) error {
	return Equals(nil, expectedHotdogsRemaining, hotdogs)
}

func FeatureContext(s *godog.Suite) {
	s.Step(`^I have (\d+) hotdogs$`, iHaveHotdogs)
	s.Step(`^I eat (\d+) hotdogs$`, iEatHotdogs)
	s.Step(`^I MUST have (\d+) hotdogs$`, iMUSTHaveHotdogs)
}

var opt = godog.Options{Output: colors.Colored(os.Stdout)}

func init() {
	godog.BindFlags("godog.", flag.CommandLine, &opt)
}

func TestMain(m *testing.M) {
	flag.Parse()
	opt.Paths = flag.Args()
	status := godog.RunWithOptions("acceptance", func(s *godog.Suite) {
		FeatureContext(s)
	}, opt)
	if st := m.Run(); st > status {
		status = st
	}
	os.Exit(status)
}
