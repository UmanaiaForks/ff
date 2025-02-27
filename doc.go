// Package ff provides a flags-first approach to runtime configuration.
//
// The central function is [Parse], which mirrors the Parse method of a standard
// flag.FlagSet, populating a flag set from commandline arguments, environment
// variables, and/or a config file. [Option] values control parsing behavior.
//
// [FlagSet] is provided as a standard flag set implementation, inspired by
// getopts(3). A stdlib flag.FlagSet can be adapted to a FlagSet set via
// [NewStdFlagSet]. Callers are also free to use their own implementation of the
// [Flags] interface.
//
// [Command] is provided as a tool for building CLI applications, like docker or
// kubectl, in a simple and declarative style. It's intended to be easier to
// understand and maintain than common alternatives.
package ff
