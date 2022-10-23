package main

func slowestKey(releaseTimes []int, keysPressed string) byte {
	slowestPressed := 0
	slowestChar := byte(0)

	for i, time := range releaseTimes {
		if i == 0 {
			slowestPressed = time
			slowestChar = byte(keysPressed[i])
			continue
		}

		pressed := time - releaseTimes[i-1]

		if pressed > slowestPressed {
			slowestPressed = pressed
			slowestChar = byte(keysPressed[i])
			continue
		}

		if pressed == slowestPressed {
			if byte(keysPressed[i]) > slowestChar {
				slowestChar = byte(keysPressed[i])
				continue
			}
		}
	}

	return slowestChar
}
