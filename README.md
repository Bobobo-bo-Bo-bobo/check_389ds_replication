**_Note:_** Because I'm running my own servers for several years, main development is done at at https://git.ypbind.de/cgit/check_389ds_replication/

----

# Preface
In larger LDAP setups (with 389 directory sever as LDAP server) the data is replicated to one or more LDAP master and/or slave servers.
To ensure data integrity and consistency the state of the replication should be checked at regular intervals.

This check allows for checking the state of replication of an 389ds master or slave server and can be integrated
in a Nagios based monitoring solution like [Icinga2](https://icinga.com/products/)

# Build requirements
This check is implemented in Go so, obviously, a Go compiler is required.

Building this programm requires the [go-ldap.v3](https://github.com/go-ldap/ldap/) library.

# Command line parameters
:heavy_exclamation_mark: **This check only works with 389 directory server** :heavy_exclamation_mark:

This check requires a recent version of the 389 directory server with support for the `nsds5replicaLastUpdateStatusJSON` attribute.

The design has been described in [Replication Agreement Status Message Improvements](http://www.port389.org/docs/389ds/design/repl-agmt-status-design.html) and was implemented in [#49602 Improve replication status messages](https://pagure.io/389-ds-base/issue/49602)


| *Paraemter* | *Description* | *Default* | *Comment* |
|:------------|:--------------|:---------:|:----------|
| `--ca-cert=<file>` | CA certificate for validation of the server SSL certificate | - | - |
| `--domain=<domain>` | LDAP domain | - | **manatory** |
| `--help` | Show help text | - | - |
| `--insecure` | Skip SSL verification of server certificate | - | Do not use in a production environment |
| `--password-file=<file>` | Authenticate using password read from <file> | - | Only the first line of the file will be interpreted as password. |
| `--password=<pass>` | Authenticate using password of <pass> | - |  Password for authentication - either provided from a file or on the command line - is **mandatory** |
| `--server=<uri>` | URI of LDAP server | - | **mandatory** |
| `--timeout=<sec>` | LDAP connection and search timeout in seconds | 15 | - |
| `--user=<user>` | Authenticate as <user> | - | **mandatory** |
| `--version` | Show version information | - | - |

# Licenses
## check_389ds_replication
This program is free software: you can redistribute it and/or modify
it under the terms of the GNU General Public License as published by
the Free Software Foundation, either version 3 of the License, or
(at your option) any later version.

This program is distributed in the hope that it will be useful,
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
GNU General Public License for more details.

You should have received a copy of the GNU General Public License
along with this program.  If not, see <https://www.gnu.org/licenses/>.

## go-ldap.v3 (https://github.com/go-ldap/ldap/)
The MIT License (MIT)

Copyright (c) 2011-2015 Michael Mitton (mmitton@gmail.com)
Portions copyright (c) 2015-2016 go-ldap Authors

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.

