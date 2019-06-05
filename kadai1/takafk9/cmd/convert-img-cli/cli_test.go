package main

import (
	"bytes"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestCLI(t *testing.T) {
	cases := []struct {
		command           string
		expectedOutStream string
		expectedErrStream string
		expectedExitCode  int
	}{
		{
			command:           "convert-img-cli",
			expectedOutStream: "",
			expectedErrStream: "Failed to set up convert-img-cli: invalid argument\nPlease specify the exact one path to a directly or a file\n\n",
			expectedExitCode:  ExitCodeInvalidArgsError,
		},
		{
			command:           "convert-img-cli testdata testdata2",
			expectedOutStream: "",
			expectedErrStream: "Failed to set up convert-img-cli: invalid argument\nPlease specify the exact one path to a directly or a file\n\n",
			expectedExitCode:  ExitCodeInvalidArgsError,
		},
		{
			command:           "convert-img-cli --from .svg",
			expectedOutStream: "",
			expectedErrStream: "Failed to set up convert-img-cli: invalid extension `.svg` is given for --from flag\nPlease choose an extension from one of those: [.gif .jpeg .jpg .png]\n\n",
			expectedExitCode:  ExitCodeInvalidFlagError,
		},
		{
			command:           "convert-img-cli --to .svg",
			expectedOutStream: "",
			expectedErrStream: "Failed to set up convert-img-cli: invalid extension `.svg` is given for --to flag\nPlease choose an extension from one of those: [.gif .jpeg .jpg .png]\n\n",
			expectedExitCode:  ExitCodeInvalidFlagError,
		},
		{
			command:           "convert-img-cli testdata",
			expectedOutStream: "",
			expectedErrStream: "Failed to execute convert-img-cli\ncould not find files with the specified extension. path: testdata, extension: .jpg\n\n",
			expectedExitCode:  ExitCodeExpectedError,
		},
		{
			command:           "convert-img-cli testdata/unknown.jpg",
			expectedOutStream: "",
			expectedErrStream: "Failed to execute convert-img-cli\nlstat testdata/unknown.jpg: no such file or directory\n\n",
			expectedExitCode:  ExitCodeExpectedError,
		},
		{
			command:           "convert-img-cli --from .jpeg testdata",
			expectedOutStream: "convert-img-cli successfully converted following files to `.png`.\n[testdata/jpeg-image.jpeg]\n\n",
			expectedErrStream: "",
			expectedExitCode:  ExitCodeOK,
		},
	}

	for i, ca := range cases {
		outStream := &bytes.Buffer{}
		errStream := &bytes.Buffer{}

		cli := CLI{outStream: outStream, errStream: errStream}
		args := strings.Split(ca.command, " ")

		if got := cli.Run(args); got != ca.expectedExitCode {
			t.Errorf("#%d %q exits with %d, want %d", i, ca.command, got, ca.expectedExitCode)
		}

		if got := outStream.String(); got != ca.expectedOutStream {
			t.Errorf("#%d Unexpected outStream has returned: want: %s, got: %s", i, ca.expectedOutStream, got)
		}

		if got := errStream.String(); got != ca.expectedErrStream {
			t.Errorf("#%d Unexpected errStream has returned: want: %s, got: %s", i, ca.expectedErrStream, got)
		}

		remove(t)
	}

}

func remove(t *testing.T) {
	t.Helper()

	err := filepath.Walk("testdata", func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if path != "testdata/jpeg-image.jpeg" && path != "testdata" {
			return os.Remove(path)
		}
		return nil
	})

	if err != nil {
		t.Errorf("failed to cleanup testdata: %s", err)
	}
}
