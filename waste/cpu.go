package waste

import (
	"fmt"
	"math/rand"
	"time"

	"golang.org/x/crypto/chacha20"
)

func CPU(interval time.Duration) {
	var buffer []byte
	if len(Buffers) > 0 {
		buffer = Buffers[0].B[:4*MiB]
	} else {
		buffer = make([]byte, 4*MiB)
	}
	rand.Read(buffer)

	// construct XChaCha20 stream cipher
	cipher, err := chacha20.NewUnauthenticatedCipher(buffer[:32], buffer[:24])
	if err != nil {
		panic(cipher)
	}

	for {
		for i := 0; i < 8; i++ {
			go func() {
				for i := 0; i < 64; i++ {
					cipher.XORKeyStream(buffer, buffer)
				}
			}()
		}

		fmt.Println("[CPU] Successfully wasted on", time.Now())

		// try to construct a new cipher
		newCipher, err := chacha20.NewUnauthenticatedCipher(buffer[:32], buffer[:24])
		if err == nil {
			cipher = newCipher
		}

		time.Sleep(interval)
	}
}
