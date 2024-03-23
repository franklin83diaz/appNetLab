netlab create -int -ip 192.168.12.2 -u 100 -d 100 ./program

-int :set nat for exit to internet
-ip ip of program
-u speed upload
-d speed download

this create name space 2 interface peer, enable nat and limit with tc

ip netns add nsappnetlab01

ip link add veth0 type veth peer name veth1
ip link set veth1 netns nsappnetlab01
ip link set veth0 up
ip addr add 192.168.1.1/24 dev veth0
ip netns exec nsappnetlab01 ip addr add 192.168.1.2/24 dev veth1
ip netns exec nsappnetlab01 ip link set veth1 up
ip netns exec nsappnetlab01 ip route add default via 192.168.1.1

tc qdisc add dev veth0 root tbf rate 1mbit burst 32kbit latency 400ms
