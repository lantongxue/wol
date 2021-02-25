# wol
Wake On Lan for Golang

# Usage
```shell
Usage of wol:
  -host string
        主机地址 (default "255.255.255.255")
  -mac string
        MAC地址
        例如：
        00:00:00:00:00:00
        00-00-00-00-00-00
        AABBCCDDEEFF
  -port int
        端口号 (default 9)
```

# Examples

wol 00-00-00-00-00-00

wol 00-00-00-00-00-00 127.0.0.1

wol 00-00-00-00-00-00 127.0.0.1 9

wol -mac 00-00-00-00-00-00

wol -mac 00-00-00-00-00-00 -host 127.0.0.1

wol -mac 00-00-00-00-00-00 -host 127.0.0.1 -port 9
