WIP

# **VPN-Hole** 

A self-contained GoLang library that is capable of blocking ads,malwares and trackers using [DNS Sinkholing](https://en.wikipedia.org/wiki/DNS_sinkhole). This libray can be integrated with different pre-existing cross-platform applications.

You can learn more about the project from the [blogs](https://leap.se/#blog) at [leap.se](https://leap.se).

## Table of Contents

* [Installation](#installation)
    * [Go](#Go)
    * [GoMobile & Gobind](#GoMobile & Gobind)
    * [VPN-Hole](#VPN-Hole)
* [Instructions](#Instructions)
* [Build & Run](#Build & Run)
* [Test](#Test)
* [Integration](#Integration)
* [Adapting to Bitmask_core](#Adapting to Bitmask_core)
* [License](#License)

## Installation <a name="installation"></a>

### Go <a name="Go"></a>
The detailed download and installation guide can be found at official website of [**Go**](https://go.dev/doc/install).

For Linux:
```
# 1. Download Go Binary Archive
wget https://golang.org/dl/go1.18.4.linux-amd64.tar.gz

# 2. Extract it
tar -xzf go1.18.4.linux-amd64.tar.gz -C /usr/local/

# 3. Add PATH Variables
sudo nano /etc/profile
export PATH=$PATH:/usr/local/go/bin
source /etc/profile

# 4. Check Go Version
go version
```
You can refer to the official documentation for setting **GOPATH**.

### GoMobile & Gobind <a name="GoMobile & Gobind"></a>
After installing Go, we need gobind. **Gobind** is a tool that generates language bindings that make it possible to call Go functions from Java and Objective-C. It is called internally by **gomobile** which can help us build cross-platform applications.
```
go install golang.org/x/mobile/cmd/gomobile@latest

# To compile Android APK and IOS Apps
gomobile bind [-target android|ios|iossimulator|macos|maccatalyst] [-bootclasspath <path>] [-classpath <path>] [-o output] [build flags] [package]
```

## VPN-Hole <a name="VPN-Hole"></a>
To clone and build VPN-Hole Library:
```
git clone https://0xacab.org/leap/vpn-hole.git

cd vpn-hole
```

### Instructions <a name="Instructions"></a>
To load all of the packages in the main module:
```
# If go.mod and go.sum are present then,
# Run 
go mod tidy
```
If go.mod and go.sum are missing from the cloned repository, then:
```
# go mod init [module-path]
go mod init main
``` 

## Build & Run <a name="Build & Run"></a>

```
# build the library
go build
```
An executable binary file would be generated for the library.
In Windows, a executable named- main.exe would be created. 

On MacOSX Monterey you may encounter an build error like this
```
/Library/Developer/CommandLineTools/SDKs/MacOSX.sdk/usr/include/pthread.h:328:6: error: macro expansion producing 
'defined' has undefined behavior [-Werror,-Wexpansion-to-defined]
/Library/Developer/CommandLineTools/SDKs/MacOSX.sdk/usr/include/pthread.h:197:2: note: expanded from macro 
'_PTHREAD_SWIFT_IMPORTER_NULLABILITY_COMPAT'
```
You can fix it by adapting the compiler flags, e.g. by prefixing `go build` with `CGO_CPPFLAGS="-Wno-error 
-Wno-nullability-completeness -Wno-expansion-to-defined -Wbuiltin-requires-header"`.

# executing the binary
main.exe
```
This would start the network wide DNS-level blocker on port 53.

## Test the blocker <a name="Test"></a>
To check if the blocker is blocking malicious domains from the provided blacklist, run **nslookup** command and see the response:

```
# For example: 'adbuddiz.com' is one of the domain in the blacklist
nslookup -port=53 adbuddiz.com localhost

Server:  UnKnown
Address:  ::1

Non-authoritative answer:
Name:    adbuddiz.com
Addresses:  ::
          0.0.0.0
```

## Integration <a name="Integration"></a>
To integrate the VPN-Hole Library with existing cross-platform applications, we will be using gomobile to compile and bind it.
```

```

### Adapting to Bitmask_core <a name="Adapting to Bitmask_core"></a>
```

```

## License <a name="License"></a>
This project is released under GNU GPLv3 License.[See LICENSE file](https://0xacab.org/leap/vpn-hole/-/blob/no-masters/LICENSE) for details.
