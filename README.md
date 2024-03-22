# Stamper

Stamper is a CLI tool for generating licenses

## Example

To generate an MIT license for your project run

```sh
stamper --license mit --name ACME
```

Below is a list with all supported licenses and their corresponding commands:
| License | Command |
| -------- | ------- |
| GNU Affero General Public License v3.0 | agpl-3.0 |
| Apache License 2.0 | apache-2.0 |
| BSD 2-Clause "Simplified" License | bsd-2-clause |
| BSD 3-Clause "New" or "Revised" License | bsd-3-clause |
| Boost Software License 1.0 | bsl-1.0 |
| Creative Commons Zero v1.0 Universal | cc0-1.0 |
| Eclipse Public License 2.0 | epl-2.0 |
| GNU General Public License v2.0 | gpl-2.0 |
| GNU General Public License v3.0 | gpl-3.0 |
| GNU Lesser General Public License v2.1 | lgpl-2.1 |
| MIT License | mit |
| Mozilla Public License 2.0 | mpl-2.0 |
| The Unlicense | unlicense |

## Installation

Currently you can only build it from source, run `make build` and use the executable `/bin/stamper`.
