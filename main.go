package main

import (
	"archive/tar"
	"bufio"
	"bytes"
	"encoding/base64"
	"fmt"
	"io"

	"github.com/xi2/xz"
)

const testData = "/Td6WFoAAATm1rRGAgAhARYAAAB0L+Wj4Cf/AVFdADmZCEEvE0GXPrrDSoKhJ+9OJAlqw1u9nlhzayQyaw99CBEiY8yvxlRt65vIuCmskS53IR7EQC1SoMP9VtgQO0eQTgFrli8Lvear5enfc93+BY09eW4medlFSZDxhkXwyP99V1KBgAuqvc29rRiKW0FRuJiWRGXp8gEWxumdVWgba5iRJbpP9wgKp8OhTRSQKMjRgW9hHiyAu4WX1H8WRLs0QShms+8EZ7pp9k7K0jza1Lcj7uwfX57quluJkSJMwDLB0+hHFf6j6IhZY9flDpxusFWFXxzZRnwUpEc4V8DqCzy9i+AI+zQMesJPCGMRwxrOV5fzORl31f+uhp7NKEJOAzIvNsaE0zFWC4NEFAgoARpq6BXMT4D+QwBQEFhnPRc3xas4thzxSY8ggI9cKP61bamNcL1jMkRUZQx5zHVpZ8rh2XVpdNfWKbWesyQWpQAAAAAA9vZaAcq5Z5QAAe0CgFAAAGoXqpKxxGf7AgAAAAAEWVo="

func main() {

	var sdar, sdbr bytes.Buffer
	var sda, sdb []byte

	b64 := base64.NewDecoder(base64.StdEncoding, bytes.NewBufferString(testData))
	gz, _ := xz.NewReader(b64, 0)
	tw := tar.NewReader(bufio.NewReader(gz))

	for {
		header, _ := tw.Next()

		if header.Name == "sda" {
			io.Copy(&sdar, tw)
			sda = sdar.Bytes()
		} else {
			io.Copy(&sdbr, tw)
			sdb = sdbr.Bytes()
			break
		}
	}

	var output bytes.Buffer
	for i := range sda {
		stripe := []byte{sda[i], sdb[i], (sda[i] ^ sdb[i])}
		p := (i + 2) % 3
		output.WriteByte(stripe[(p+1)%3])
		output.WriteByte(stripe[(p+2)%3])
	}
	fmt.Println(output.String())
}
