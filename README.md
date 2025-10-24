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
- [Acknowledgements](#acknowledgements)

## Features

- **Manage the Entire Filesystem:** Manages files anywhere, from files in `/etc` to
  `/Library/LaunchDaemons`.
- **Isolated Profiles:** Manage multiple, independent configurations (e.g., from
  different repositories) using the --profile flag, which keeps all profile data
  neatly organized in subdirectories within the ~/.local/share/chezroot/ directory.
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
info: Fixing ownership of /Users/james/.local/share/chezroot/default...
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

`chezroot` is configured primarily via command-line flags and environment variables.

### Command-Line Flags

- `--profile <name>`

   Specifies the profile to use. This overrides the `$CHEZROOT_PROFILE` environment
   variable if it is also set.

   See the "Using Profiles" section for details.

- `--config <path>`

   Specifies an explicit path to your user-defined `chezmoi` configuration file. You
   can give this option multiple times.

   If this option is not given, `chezroot` will automatically search the current
   profile's config directory (e.g., `~/.config/chezroot/default` for the default
   profile) for a config file (like `config.toml`, `config.yaml`, etc.), mimicking
   chezmoi's behavior.

### Environment Variables

- `$CHEZROOT_PROFILE`

   **Default:** (unset)

   Specifies the profile to use if the `--profile` flag is not provided. This is a
   convenient way to work with a specific profile for an entire terminal session.

   If neither the `--profile` flag nor this variable is set, `chezroot` will use the
   hardcoded default profile name of `default`.

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

The current profile in use is controlled using the `--profile` flag or, as a
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
   environment variable for your session instead of typing the flag repeatedly. The
   `--profile` flag will always override this variable if you need to switch
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

## Acknowledgements

- This tool is a simple wrapper and would not be possible without the excellent
  [chezmoi](https://www.chezmoi.io) by [Tom Payne](https://github.com/twpayne).
- This project borrows from the excellent
  [chezetc](https://github.com/SilverRainZ/chezetc) by [Shengyu
  Zhang](https://github.com/SilverRainZ)
