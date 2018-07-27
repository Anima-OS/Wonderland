// Copyright (c) 2018, Mark "Happy-Ferret" Bauermeister
//
// This software may be modified and distributed under the terms
// of the BSD license.  See the LICENSE file for details.

/*
package hook defines, reads and executes hooks in Wingo. This package must
be initialized with a Gribble execution environment, and a file path to a
configuration file specifying user defined hooks.

The hook package defines only a set number of hook groups that the user can use
to execute arbitrary commands.

Whenever Wingo enters a state described by a hook group, Fire is called with
the appropriate hook group (and arguments if relevant). Adding a new hook group
only requires a new constant and a new entry in the unexported 'groups'
variable.

Please see config/hooks.wini in the Wingo project directory for an explanation
of how user defined hooks can be specified:
https://github.com/BurntSushi/wingo/blob/master/config/hooks.wini
*/
package hook
