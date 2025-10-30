# chezroot

A `sudo` wrapper for `chezmoi` to manage root-owned files across your entire
filesystem.

`chezmoi` is the premier tool for managing your personal dotfiles. `chezroot` is a
companion wrapper that extends `chezmoi`'s power to your system-level files, allowing
you to manage files in `/etc`, `/Library/LaunchDaemons`, or anywhere else on your
root filesystem with the same `chezmoi` workflow.

- [Features](#features)
- [Platform Support](#platform-support)
- [Installation](#installation)
- [Usage](#usage)
- [Configuration](#configuration)
- [Using Profiles](#using-profiles)
- [Passwordless sudo](#passwordless-sudo)
- [Design](#design)
- [Items to do](#items-to-do)
- [Acknowledgements](#acknowledgements)

## Features

- **Manage the Entire Filesystem:** Manages files anywhere, from files in `/etc` to
  `/Library/LaunchDaemons`.
- **Isolated Profiles:** Manage multiple, independent configurations (e.g., from
  different repositories) using the --profile option. Each profile's source, config,
  state, and cache are stored in separate, dedicated directories, ensuring they never
  conflict.
- **Safe `sudo` Execution:** Intelligently wraps `chezmoi` commands in `sudo` only
  when needed.
- **User-Level Actions:** Automatically drops privileges for commands like `edit` and
  `cd`, so you always edit configuration files as your normal user, not as root.
- **Safe by Default:** `chezroot init` automatically creates a `.chezmoiignore` file
  that ignores everything. You must explicitly un-ignore the files you wish to
  manage.
- **Automatic Shell Completion:** Full `chezmoi` completions for Bash and Zsh that
  work automatically.
- **No Dependencies:** `chezroot` is a single POSIX-compliant shell script. It only
  requires `chezmoi`, `sudo`, `bash`, and `sed`.
- **Cross-Platform:** Works seamlessly on both **Linux and macOS**.

## Platform Support

### Linux and macOS

`chezroot` is designed exclusively for POSIX-like systems (Linux and macOS) that use
`sudo` for privilege escalation. It relies on `bash`, `sudo`, and standard Unix
utilities to manage file ownership and drop privileges.

### Windows (Not Supported)

Windows is **not** a supported platform. The Windows administrative model (UAC, "Run
as Administrator") is fundamentally different from `sudo`, and `chezroot`'s `bash`
scripts are not compatible with PowerShell or the Windows filesystem.

While `chezmoi` itself works perfectly on Windows, `chezroot` cannot be used to wrap
it there.

## Installation

Installation is handled by a package manager appropriate for your platform.

### macOS (via Homebrew)

You can install `chezroot` from the official Homebrew tap:

```bash
brew tap jcouball/tap
brew install chezroot
```

### Linux

Packages for popular Linux distributions are in progress.

#### Arch Linux (AUR)

```bash
# Package coming soon
yay -S chezroot
```

#### Debian/Ubuntu (via PPA)

```bash
# PPA coming soon
sudo add-apt-repository ppa:jcouball/ppa
sudo apt update
sudo apt install chezroot
```

#### Fedora (via COPR)

```bash
# COPR repository coming soon
sudo dnf copr enable jcouball/chezroot
sudo dnf install chezroot
```

## Usage

The workflow is nearly identical to `chezmoi`, with one extra safety step. By
default, you will be operating on the `default` profile, which is stored in
`~/.local/share/chezroot/default`.

### 1. Initialize `chezroot`

This creates your source directory at `~/.local/share/chezroot/default`.

```bash
$ chezroot init
info: Running post-init setup...
info: Creating safe-by-default .chezmoiignore...
info: Creating .chezmoiversion...
```

### 2. Configure Your Allow-List (Important!)

`chezroot` is **safe by default**. It will ignore all files until you explicitly
allow them.

Edit the newly created ignore file at
`~/.local/share/chezroot/default/.chezmoiignore`.

The file will contain a single `*` (ignore all). To manage `/etc/hosts` and all files
under `/etc/nginx`, you would change it to:

```gitignore
# .chezmoiignore
# 1. Ignore everything by default
*

# 2. Un-ignore the specific files and directories
#    you explicitly want to manage.
!/etc/hosts
!/etc/nginx
```

### 3. Add and Manage Files

Now you can use the standard `chezmoi` workflow.

```bash
# Add a file from the root filesystem to your source repo
$ chezroot add /etc/hosts

# Edit the file (this runs $EDITOR as your user, not root)
$ chezroot edit /etc/hosts

# See what changes will be applied
$ chezroot diff

# Apply the changes to your system (this runs with sudo)
$ chezroot -v apply
```

## Configuration

`chezroot` is configured primarily via command-line options and environment variables.

### Command-Line Options

- `--profile <name>`

   Specifies the profile to use. This overrides the `$CHEZROOT_PROFILE` environment
   variable if it is also set.

   See the "Using Profiles" section for details.

- `--config` and `--config-format` are not supported

   `chezmoi`'s `--config` and `--config-format` options are **not supported**.

   `chezroot`'s core "Profile" feature works by strictly managing all paths for you.
   Each profile's configuration is always loaded from its dedicated directory (e.g.,
   `~/.config/chezroot/$PROFILE/`) to ensure complete isolation. Allowing an external
   config file would break this model and create a conflict with the `--profile`
   option.

### Environment Variables

- `$CHEZROOT_PROFILE`

   **Default:** (unset)

   Specifies the profile to use if the `--profile` option is not provided. This is a
   convenient way to work with a specific profile for an entire terminal session.

   If neither the `--profile` option nor this variable are set, `chezroot` will use the
   profile name `default`.

- `$VISUAL` / `$EDITOR`

   Used by the `chezroot edit ...` to determine which editor program to use. The
   editor is chosen by checking the following, in this order:

  - **$VISUAL Environment Variable:** It first checks if the `$VISUAL` environment
    variable is set.

  - **$EDITOR Environment Variable:** If `$VISUAL` is not set, it falls back to
    checking for the `$EDITOR` environment variable.

  - **System Default:** If none of the above are set, `chezroot` uses a hardcoded
    default of `vi`.

   **NOTE:** specifying `edit.command` in the config file is not supported as it is
   in chezmoi.

## Using Profiles

By default, `chezroot` manages a single profile (a self-contained set of
configurations). The default profile is named `default` and is stored in
`~/.local/share/chezroot/default`.

The current profile in use is controlled using the `--profile` option or, as a
convenient default, the `$CHEZROOT_PROFILE` environment variable. If neither of these
is specified the `default` profile is used.

This is useful if you want to manage different sets of root file configurations from
different git repositories. For example, you could have one repository for your
system-wide shell settings and a completely different repository for managing `nginx`
configurations.

### Example

Let's set up a new profile called nginx to manage web server files.

1. Initialize the new profile:

   This command will use the profile name `nginx` to create a new source directory at
   `~/.local/share/chezroot/nginx` and link it to your new repository.

   ```bash
   chezroot --profile nginx init https://github.com/user/nginx-configs.git
   ```

2. Add files to the new profile:

   All subsequent commands must use the `--profile nginx` option to operate on this
   new instance.

   ```bash
   chezroot --profile nginx add /etc/nginx/nginx.conf
   ```

3. Use the default profile:

   To go back to managing your default files, simply run `chezroot` without the
   option.

   ```bash
   chezroot apply
   ```

   This works because each profile maintains its own separate configuration, source
   directory, and state file, so they do not interfere with each other.

4. Using the Environment Variable for Convenience

   If you are going to work on the nginx profile for a while, you can set the
   environment variable for your session instead of typing the option repeatedly. The
   `--profile` option will always override this variable if you need to switch
   temporarily.

   ```shell
   export CHEZROOT_PROFILE=nginx
   chezroot add /etc/nginx/sites-available/default # Uses 'nginx' profile
   chezroot apply # Also uses 'nginx' profile
   ```

## Passwordless sudo

To avoid typing your password for every `apply` or `diff`, it is highly recommended
to add a `sudoers` rule.

Create a new file at `/etc/sudoers.d/chezroot` (as root) with the following content.
**Be sure to replace `YOUR_USER` and use the correct *absolute path* to the
`chezroot` script** (which you can find with `which chezroot`).

```shell
# /etc/sudoers.d/chezroot
YOUR_USER ALL=(root) NOPASSWD: /usr/local/bin/chezroot
```

This rule allows your user to run *only* the `chezroot` script as root without a
password.

## Design

The `chezroot` command line tool is a wrapper around `chezmoi`. It aims to be fully
compatible with `chezmoi` with a few exceptions.

### Profiles

chezroot implements *profiles* to allow multiple `chezmoi`-type sources which can be
shared and can configure specific things within the host operating system. For
example, one source might configure a webserver, another might configure system-wide
shell settings, etc.

The current profile is set at the command line with the `--profile` option or
via the environment variable $CHEZROOT_PROFILE. If a profile is not specified, The
profile `default` is used.

Profiles are used to create isolation from each other. Each profile has its own
source directory, config file, persistent state file, and cache as follows:

- Source Directory: `$HOME/.local/share/chezroot/$PROFILE/`
- Config File: `$HOME/.config/chezroot/$PROFILE/config.*`
- Persistent State File: `$HOME/.local/state/chezroot/$PROFILE/state.boltdb`
- Cache Directory: `$HOME/.cache/chezroot/cache/$PROFILE/`

### Configuration File

The path to the configuration file passed to `chezmoi` is determined by the profile.
The configuration file is loaded from the directory
`$HOME/.config/chezroot/$PROFILE/`. The configuration file is named `config.$FORMAT`
where format (following the same guidelines as `chezmoi`) is one of `json`, `jsonc`,
`toml`, or `yaml`.

The user can not change the path to the configuration file. Use of the `chezmoi`
command line options `-c`, `--config`, and `--config-format` are not allowed.

### Source Directory

The path to the source directory passed to `chezmoi` is also determined by the profile.
The source directory is `$HOME/.local/share/chezroot/PROFILE/`.

The user cannot change the path to the source directory. Use of the `chezmoi` command
line options `-S` and `--source-directory` are not allowed. If `sourceDir` is given
in the configuration file, it is ignored.

The source directory and its contents are assumed to be owned by the user
invoking the `chezroot` command.

### Destination Directory

The destination directory path is set to "/".

The user cannot change the destination directory path. Use of the `chezmoi` command line options `-D` and `--destination-directory` are not allowed. If `destDir` is given in the configuration file, it is ignored.

The destination files are assumed to be owned by root.

### Execution Environment

#### Current User

`chezroot` is designed to be invoked as a non-root user. Internally, `chezroot` will
run `chezmoi` and other commands with `sudo` as needed.

#### Variables

`chezroot` explicitly unsets any `CHEZMOI_*` environment variables in the execution
environment before invoking the underlying chezmoi command. This prevents environment
variables from the user's `chezmoi` setup from interfering with `chezroot`.

`chezroot` then reconstructs the `CHEZMOI_*` variables for the `chezmoi` subprocess
by mapping `CHEZROOT_*` variables in two stages:

1. **Global variables:**
   - Variables prefixed with `CHEZROOT_` are mapped
   - For each `CHEZROOT_VAR=value`, the corresponding `CHEZMOI_VAR=value` is set.
     *(Example: `CHEZROOT_VERBOSE=true` sets `CHEZMOI_VERBOSE=true`)*.

2. **Profile-Specific variables:**
   - Variables prefixed with `CHEZROOT_PROFILE_${PROFILE}_` (where `${PROFILE}` is
     the *name of the currently active profile*) are mapped.
   - For each `CHEZROOT_PROFILE_${PROFILE}_VAR=value`, the corresponding
      `CHEZMOI_VAR=value` is set, *overwriting* any value set by a global variable.
      *(Example: If the active profile is `web`, then
      `CHEZROOT_PROFILE_WEB_VERBOSE=false` sets `CHEZMOI_VERBOSE=false`, taking
      precedence over any `CHEZROOT_VERBOSE` setting)*.

Profile-specific variables always take presedence over global variables.

#### Path

`chezroot` does not manipulate the `PATH` variable when running `sudo chezmoi`.

Most sudo configurations reset the `PATH` to a predefined, secure set of directories
(often defined by secure_path in `/etc/sudoers`). This is a security measure to
prevent users from executing arbitrary commands as root just because those commands
happen to be in their personal `PATH`.

Users should configure absolute paths for any external commands (diff tools, merge
tools, git, password managers, encryption tools, etc.) within their chezroot
profile's configuration file if those commands are needed during operations that
require sudo.

#### Current Working Directory

`chezroot` invokes the underlying `chezmoi` command without explicitly changing the
current working directory (CWD) from where `chezroot` itself was run.

- Command line arguments, including any relative paths (e.g., `chezroot add ./file`),
  are passed **as-is** by `chezroot` to the underlying `chezmoi` command.
- `chezmoi` will interpret these relative paths based on **its own CWD** at the time
  of execution.
- **Important:** When `chezroot` uses `sudo` to run `chezmoi`, `sudo` **may change
  the CWD** before `chezmoi` executes (often to `/root`). This means relative paths
  provided on the command line might not resolve as expected in commands requiring
  `sudo` (like `add`, `apply`).
- Therefore, it is **strongly recommended to use absolute paths** when referring to
  files on the command line with `chezroot`, especially for commands that might run
  with `sudo`.
- Similarly, users should be cautious when using relative paths inside **`run_`
  scripts**, as the CWD during their execution (especially under `sudo`) might not be
  the directory where `chezroot` was invoked. Using absolute paths or determining
  paths dynamically within scripts is generally safer.

### Command Execution

#### Use of sudo

`chezroot` manages target files that are owned by root. This means that any `chezmoi`
command that *may* read or update target files in the destination directory must be
executed with `sudo`. This includes commands the evaluate templates which *may* use
template functions that read the target state (i.e, the `stat` function).

Commands executed WITH `sudo` are:

- add, apply, cat, destroy, diff, dump, execute-template, init, merge, merge-all,
  purge, re-add, status, unmanaged, update, verify

Commands executed WITHOUT `sudo` are:

- cat-config, cd, chattr, completion, data, decrypt, dump-config, edit, edit-config,
  edit-config-template, encrypt, forget, generate, git, help, ignored, import,
  license, list, managed, secret, source-path, state, target-path, unmanage

Unsupported commands:

- docker, ssh, upgrade

#### Command Specific Handling

The following are all expected to be owned by the user who invokes chezroot:

- the source directory and its contents recursively
- the cache directory and its contents recursively
- the persistent state file

Many `chezmoi` commands must be run with `sudo` which can leave modified files with
`root` ownership. To prevent this, `chezroot` will automatically "fix ownership" of
these files. This means that it will `chown` all affected files and directories back
to the invoking user after any `sudo` command completes.

`chezroot` will automatically fix ownership after running the following commands:

- add, apply, destroy, merge, merge-all, purge, re-add, update

Some commands may need additional special handling. The rest of this section gives
those details.

##### edit

The edit options `--apply` and `--watch` are not currently supported. Instead run
`chezroot edit ...` and then `chezroot apply` as two separate commands.

These options are unsupported because they create a permission conflict. The `edit`
action must run as your normal user to access your editor and source files, but the
`apply` action must run with sudo to write to the root filesystem.

##### init

`chezroot` will pass the option `--guess-repo-url=false` to `chezmoi init`. This option
is passed as a safety measure to disable `chezmoi`'s default behavior of guessing a
repository URL (like github.com/user/dotfiles). This ensures you explicitly provide
the full URL for your *system-level* configuration, rather than accidentally using your
*personal dotfiles* repository.

`chezmoi init` will be run with `sudo` only if the `--apply` or `--one-shot` option is
passed. In this case, `chezroot` will fix ownership after `sudo chezmoi init`
completes.

If `$HOME/.local/share/chezroot/$PROFILE/.chezmoiignore` does not exist, `chezroot`
will create an ignore file that ignores everything. You must explicitly un-ignore the
files you wish to manage. It also ignores the config.

If `$HOME/.local/share/chezroot/$PROFILE/.chezmoiversion` does not exist, `chezroot`
will create the file containing the minimal version required for `chezroot`. If the
file exists, `chezroot` will ensure the version is compatible with the minimal
version required.

## Items to do

Design a `chezroot profile` command to manage profiles.

## Acknowledgements

- This tool is a simple wrapper and would not be possible without the excellent
  [chezmoi](https://www.chezmoi.io) by [Tom Payne](https://github.com/twpayne).
- This project borrows from the excellent
  [chezetc](https://github.com/SilverRainZ/chezetc) by [Shengyu
  Zhang](https://github.com/SilverRainZ)
