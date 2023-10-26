# 
# blenderd spec file
#

Name: blenderd
Summary: Server to blend CERNBox to unforeseen realms
Version: 1.0.2
Release: 1%{?dist}
License: AGPLv3
BuildRoot: %{_tmppath}/%{name}-buildroot
Group: CERN-IT/ST
BuildArch: x86_64
Source: %{name}-%{version}.tar.gz

%description
This RPM provides an HTTP server to fake PROPFIND responses

# Don't do any post-install weirdness, especially compiling .py files
%define __os_install_post %{nil}

%prep
%setup -n %{name}-%{version}

%install
# server versioning

# installation
rm -rf %buildroot/
mkdir -p %buildroot/usr/local/bin
mkdir -p %buildroot/usr/lib/systemd/system
mkdir -p %buildroot/var/log/blenderd
install -m 755 blenderd	     %buildroot/usr/local/bin/blenderd
install -m 644 blenderd.service    %buildroot/usr/lib/systemd/system/blenderd.service

%clean
rm -rf %buildroot/

%preun

%post

%files
%defattr(-,root,root,-)
/var/log/blenderd
/usr/lib/systemd/system/blenderd.service
/usr/local/bin/*

%changelog
* Thu Oct 26 2023 Hugo Gonzalez Labrador <hugo.gonzalez.labrador@cern.ch> 1.0.2
- expose /eos/web for sync clients
* Wed May 4 2023 Hugo Gonzalez Labrador <hugo.gonzalez.labrador@cern.ch> 1.0.1
- Support new dav files routes and propfind:root 
* Tue May 26 2020 Hugo Gonzalez Labrador <hugo.gonzalez.labrador@cern.ch> 1.0.0
- Init
