package xway

func (r *XWAYRequest) Encode() error {
	var request []byte

	buffer, params, err := r.Sender.encode(sender)
	if err != nil {
		return err
	}
	request = append(request, buffer...)
	r.params = append(r.params, params...)

	buffer, params, err = r.Receiver.encode(receiver)
	if err != nil {
		return err
	}
	request = append(request, buffer...)
	r.params = append(r.params, params...)

	r.npdu = 0xf0
	if len(r.params) > 0 {
		r.npdu += 1
	}
	if r.Refused {
		r.npdu += 2
	}
	// TODO: service level
	request = append([]byte{r.npdu}, request...)

	r.Header = append(request, r.params.encode()...)
	return nil
}

func (a Address) encode(typeAddr byte) ([]byte, parameters, error) {
	if a.Station > 63 && a.Station != 255 {
		return nil, nil, ErrWrongStationNumber
	}

	buffer := []byte{a.Station, GATE_LEVEL5}
	params := []parameter{}
	if a.Network <= 15 {
		buffer[1] = a.Network << 4
	} else {
		params = append(params, parameter{
			id:    2 + typeAddr,
			lg:    1,
			value: []byte{a.Network},
		})
	}

	if a.Network > 15 || a.Gate > 15 {
		params = append(params, parameter{
			id:    typeAddr,
			lg:    1,
			value: []byte{a.Gate},
		})
	} else {
		buffer[1] += a.Gate
	}
	return buffer, params, nil
}

func (p parameters) encode() (buffer []byte) {
	for i, param := range p {
		code := param.id<<4 + param.lg
		if i == len(p)-1 {
			code += 8
		}
		buffer = append(buffer, code)
		buffer = append(buffer, param.value...)
	}
	return
}
