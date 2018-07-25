Name:           gdaemon
Version:        1.0.1
Release:        1%{?dist}  
Summary:        gdaemon
  
Group:          Development/Tools  
License:        GPL  
Source0:        %{name}-%{version}.tar.gz  
  
#BuildRequires:    libpcap-devel
#Requires:         libpcap-devel
  
%description  
  
%prep  
%setup -q
  
%build  


%install  
mkdir -p ${RPM_BUILD_ROOT}/usr/local/gdaemon
install -m 755 gdaemon ${RPM_BUILD_ROOT}/usr/local/gdaemon

%post

%clean
rm -rf $RPM_BUILD_ROOT

%files
/usr/local/gdaemon/gdaemon
