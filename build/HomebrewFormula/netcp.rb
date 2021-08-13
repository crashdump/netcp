class Netcp < Formula
  desc "Copy files and directories across systems without requiring a direct network line-of-sight. Netcp uses a cloud storage endpoint to store the data while in flight. "
  homepage "https://github.com/crashdump/netcp"
  head "https://github.com/crashdump/netcp.git"

  depends_on "go" => :build

  def install
    ENV["GOPATH"] = buildpath
    netcppath = buildpath/"src/github.com/crashdump/netcp"
    netcppath.install buildpath.children
    cd netcppath do
      system "go", "build", "-o", bin/"netcp"
      prefix.install_metafiles
    end
  end
end