# jocasta 
[![Release](https://img.shields.io/github/release/marema31/jocasta.svg?style=for-the-badge)](https://github.com/marema31/jocasta/releases/latest)
[![Build Status](https://travis-ci.com/marema31/jocasta.svg?branch=master)](https://travis-ci.com/marema31/jocasta) [![BCH compliance](https://bettercodehub.com/edge/badge/marema31/jocasta?branch=master)](https://bettercodehub.com/)[![codecov](https://codecov.io/gh/marema31/jocasta/branch/master/graph/badge.svg)](https://codecov.io/gh/marema31/jocasta)

Small utiliy that save stdin and stderr of the provided commands to log file with log rotation

### Usage

```jocasta [-c configurationFileWithoutExtension] command arg1 arg2 ... arg```

Configuration file can be in JSON, TOML or YAML. By default jocasta will lookup for .jocasta.json, .jocasta.toml or .jocasta.yaml in the current directory. If the -c option is provided, the provided path must be provided without the  extension, jocasta will try all compatible extension.

### Configuration file format

The configuration file can contains six keys with default value. The value can be overrided by a environment variable.

Key in config file | Environment variable | default value
-------------------|----------------------|--------------
out_file | JOCASTA_OUT_FILE | /tmp/jocasta_{CommandName}_stdout.log
out_maxsize | JOCASTA_OUT_MAXSIZE| 0
out_backups | JOCASTA_OUT_BACKUPS | 0 
err_file | JOCASTA_ERR_FILE | /tmp/jocasta_{CommandName}_stderr.log
err_maxsize | JOCASTA_ERR_MAXSIZE| 0
err_backups | JOCASTA_ERR_BACKUPS | 0 

All keys are prefixed by "out" or "err" that concerns STDOUT and STDERR of the command.

#### out_file and err_file

Path to the STDOUT and STDERR log file. The path can contains "{{.App}}" that will be replaced by the command name.

#### out_maxsize and err_maxsize

Approximative maximum size of the STDOUT and STDERR log file. The value can be suffixed by kb, mb, gb. The rotation will occurs at the first carriage return after this size is reached to avoid to split a line.

#### out_backups and err_backups

Number of kept rotation files of the STDOUT and STDERR log file. The rotated files will have the name precised by out_file or err_file suffixed by the file number.

