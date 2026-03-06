class Syswatch < Formula
  desc "CLI system monitoring tool"
  homepage "https://github.com/joaaomanooel/syswatch"
  version "0.0.1"
  license "MIT"

  on_macos do
    if Hardware::CPU.arm?
      url "https://github.com/joaaomanooel/syswatch/releases/download/v0.0.1/syswatch-darwin-arm64.zip"
      sha256 "REPLACE_WITH_SHA256"
    else
      url "https://github.com/joaaomanooel/syswatch/releases/download/v0.0.1/syswatch-darwin-amd64.zip"
      sha256 "REPLACE_WITH_SHA256"
    end
  end

  on_linux do
    if Hardware::CPU.arm64?
      url "https://github.com/joaaomanooel/syswatch/releases/download/v0.0.1/syswatch-linux-arm64.zip"
      sha256 "REPLACE_WITH_SHA256"
    else
      url "https://github.com/joaaomanooel/syswatch/releases/download/v0.0.1/syswatch-linux-amd64.zip"
      sha256 "REPLACE_WITH_SHA256"
    end
  end

  def install
    bin.install "syswatch"
  end

  test do
    system "#{bin}/syswatch", "--version"
  end
end
