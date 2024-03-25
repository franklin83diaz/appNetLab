# NetLab README
## Overview
NetLab is a Go-based program designed to create isolated network environments using network namespaces. It provides the capability to limit bandwidth and monitor traffic within these environments. This tool is particularly useful for network engineers and developers who need to test applications or services in a controlled network setting.

## Features
Isolated Network Namespaces: Create separate network namespaces to simulate different network environments.
Bandwidth Limitation: Specify upload and download bandwidth limits for each namespace.
Internet Connectivity: Option to set up NAT for external internet access from within the namespace.

Traffic Monitoring: Keep an eye on the network traffic that flows through the namespace.
Installation

Before you can use NetLab, you need to have Go installed on your system. You can download and install Go from the official website: https://golang.org/dl/.

Once Go is installed, you can install NetLab by cloning the repository and building the program:

```
git clone https://github.com/yourusername/netlab.git
cd netlab
go build
```
Usage
To create a new network namespace with NetLab, use the following command:

```
sudo ./netlab create --ip 192.168.137.5 --name test02 --int true -u 100 -d 100
```
Command-Line Arguments
--ip: Assign an IP address to the namespace.
--name: Set a name for the namespace.
--int: Enable NAT for internet access (set to true or false).
-u: Set the upload bandwidth limit in kbps (Kilobits per second).
-d: Set the download bandwidth limit in kbps (Kilobits per second).
Example
The following example command creates a network namespace named test02 with an IP address of 192.168.137.5, NAT enabled for internet access, and bandwidth limits of 100 kbps for both upload and download:

```
sudo ./netlab create --ip 192.168.137.5 --name test02 --int true -u 100 -d 100
```
### Dependencies
NetLab requires the following dependencies to be installed on your system:

iproute2: For network namespace and traffic control.
iptables: For setting up NAT if internet access is required.

### Limitations
NetLab must be run with superuser privileges to manage network namespaces and iptables.
Bandwidth limits are approximate and depend on the underlying traffic control system.

### Contributing
Contributions to NetLab are welcome! Please submit pull requests to the repository or report any issues you encounter.

### License
NetLab is licensed under the MIT License. See the LICENSE file for more details.

Please note that this README assumes that the NetLab program and its features are already implemented. If you are still in the development phase, you will need to adjust the installation and usage instructions accordingly.
