package xway

func Decode(b []byte) (*XWAYRequest, []byte) {
	myXWAY := &XWAYRequest{
		Sender: Address{
			Station: b[1],
			Network: b[2] >> 4,
			Gate:    b[2] % 16,
		},
		Receiver: Address{
			Station: b[3],
			Network: b[3] >> 4,
			Gate:    b[3] % 16,
		},
	}
	if b[0]&0x02 > 0 {
		myXWAY.Refused = true
	}

	i := 4
	if b[0]%2 > 0 {
		for {

			switch b[i] / 16 {
			case 0:
				myXWAY.Sender.Gate = b[i+1]
			case 1:
				myXWAY.Receiver.Gate = b[i+1]
			case 2:
				myXWAY.Sender.Network = b[i+1]
			case 3:
				myXWAY.Receiver.Network = b[i+1]
			case 4:
				i++ //unimplemented
			case 5:
				i++ //unimplemented
			default:
			}
			i += 2
			if b[i-2]&0x08 > 0 {
				break
			}
		}
	}
	return myXWAY, b[i:]
}
