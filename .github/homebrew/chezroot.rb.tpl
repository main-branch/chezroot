class Chezroot < Formula
  desc "A sudo wrapper for chezmoi to manage root-owned files"
  homepage "https.github.com/main-branch/chezroot"
  # This section will be for the Intel (amd64) binary
  url "https://github.com/main-branch/chezroot/releases/download/${TAG}/chezroot_${VERSION}_darwin_amd64.tar.gz"
  sha256 "${SHA_AMD64}"
  license "MIT"
  version "${TAG}" # Homebrew uses the tag for the version

  # This block handles the Apple Silicon (arm64) binary
  on_arm do
    url "https://github.com/main-branch/chezroot/releases/download/${TAG}/chezroot_${VERSION}_darwin_arm64.tar.gz"
    sha256 "${SHA_ARM64}"
  end

  # We no longer need "go" to build, only "chezmoi" to run
  depends_on "chezmoi"

  def install
    # This just installs the binary from the downloaded tar.gz
    bin.install "chezroot"
  end

  test do
    system bin/"chezroot", "--version"
  end
end
