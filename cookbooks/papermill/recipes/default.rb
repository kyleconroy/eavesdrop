package 'git'
package 'libpcap-dev'
package 'mercurial'

template "go-profile" do
  path "/etc/profile.d/go.sh"
  source "go.sh.erb"
  mode 0755
end

# Database
tar_extract 'https://go.googlecode.com/files/go1.2.linux-amd64.tar.gz' do
  target_dir '/usr/local'
  creates '/usr/local/go/bin'
end

directory "/usr/local/gopath"

GO = {
  "PATH" => "#{ENV['PATH']}:/usr/local/go/bin",
  "GOPATH" => "/usr/local/gopath",
}
