// DO NOT EDIT!
// Code generated by ffjson <https://github.com/pquerna/ffjson>
// source: allocator/allocator.pb.go
// DO NOT EDIT!

package allocator

import (
	"bytes"
	"fmt"
	fflib "github.com/pquerna/ffjson/fflib/v1"
)

func (mj *InverseOfferStatus) MarshalJSON() ([]byte, error) {
	var buf fflib.Buffer
	if mj == nil {
		buf.WriteString("null")
		return buf.Bytes(), nil
	}
	err := mj.MarshalJSONBuf(&buf)
	if err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}
func (mj *InverseOfferStatus) MarshalJSONBuf(buf fflib.EncodingBuffer) error {
	if mj == nil {
		buf.WriteString("null")
		return nil
	}
	var err error
	var obj []byte
	_ = obj
	_ = err
	buf.WriteByte('{')
	if mj.Status != nil {
		if true {
			buf.WriteString(`"status":`)

			{

				obj, err = mj.Status.MarshalJSON()
				if err != nil {
					return err
				}
				buf.Write(obj)

			}
			buf.WriteByte(',')
		}
	}
	buf.WriteString(`"framework_id":`)

	{

		err = mj.FrameworkID.MarshalJSONBuf(buf)
		if err != nil {
			return err
		}

	}
	buf.WriteString(`,"timestamp":`)

	{

		err = mj.Timestamp.MarshalJSONBuf(buf)
		if err != nil {
			return err
		}

	}
	buf.WriteByte('}')
	return nil
}

const (
	ffj_t_InverseOfferStatusbase = iota
	ffj_t_InverseOfferStatusno_such_key

	ffj_t_InverseOfferStatus_Status

	ffj_t_InverseOfferStatus_FrameworkID

	ffj_t_InverseOfferStatus_Timestamp
)

var ffj_key_InverseOfferStatus_Status = []byte("status")

var ffj_key_InverseOfferStatus_FrameworkID = []byte("framework_id")

var ffj_key_InverseOfferStatus_Timestamp = []byte("timestamp")

func (uj *InverseOfferStatus) UnmarshalJSON(input []byte) error {
	fs := fflib.NewFFLexer(input)
	return uj.UnmarshalJSONFFLexer(fs, fflib.FFParse_map_start)
}

func (uj *InverseOfferStatus) UnmarshalJSONFFLexer(fs *fflib.FFLexer, state fflib.FFParseState) error {
	var err error = nil
	currentKey := ffj_t_InverseOfferStatusbase
	_ = currentKey
	tok := fflib.FFTok_init
	wantedTok := fflib.FFTok_init

mainparse:
	for {
		tok = fs.Scan()
		//	println(fmt.Sprintf("debug: tok: %v  state: %v", tok, state))
		if tok == fflib.FFTok_error {
			goto tokerror
		}

		switch state {

		case fflib.FFParse_map_start:
			if tok != fflib.FFTok_left_bracket {
				wantedTok = fflib.FFTok_left_bracket
				goto wrongtokenerror
			}
			state = fflib.FFParse_want_key
			continue

		case fflib.FFParse_after_value:
			if tok == fflib.FFTok_comma {
				state = fflib.FFParse_want_key
			} else if tok == fflib.FFTok_right_bracket {
				goto done
			} else {
				wantedTok = fflib.FFTok_comma
				goto wrongtokenerror
			}

		case fflib.FFParse_want_key:
			// json {} ended. goto exit. woo.
			if tok == fflib.FFTok_right_bracket {
				goto done
			}
			if tok != fflib.FFTok_string {
				wantedTok = fflib.FFTok_string
				goto wrongtokenerror
			}

			kn := fs.Output.Bytes()
			if len(kn) <= 0 {
				// "" case. hrm.
				currentKey = ffj_t_InverseOfferStatusno_such_key
				state = fflib.FFParse_want_colon
				goto mainparse
			} else {
				switch kn[0] {

				case 'f':

					if bytes.Equal(ffj_key_InverseOfferStatus_FrameworkID, kn) {
						currentKey = ffj_t_InverseOfferStatus_FrameworkID
						state = fflib.FFParse_want_colon
						goto mainparse
					}

				case 's':

					if bytes.Equal(ffj_key_InverseOfferStatus_Status, kn) {
						currentKey = ffj_t_InverseOfferStatus_Status
						state = fflib.FFParse_want_colon
						goto mainparse
					}

				case 't':

					if bytes.Equal(ffj_key_InverseOfferStatus_Timestamp, kn) {
						currentKey = ffj_t_InverseOfferStatus_Timestamp
						state = fflib.FFParse_want_colon
						goto mainparse
					}

				}

				if fflib.EqualFoldRight(ffj_key_InverseOfferStatus_Timestamp, kn) {
					currentKey = ffj_t_InverseOfferStatus_Timestamp
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				if fflib.EqualFoldRight(ffj_key_InverseOfferStatus_FrameworkID, kn) {
					currentKey = ffj_t_InverseOfferStatus_FrameworkID
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				if fflib.EqualFoldRight(ffj_key_InverseOfferStatus_Status, kn) {
					currentKey = ffj_t_InverseOfferStatus_Status
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				currentKey = ffj_t_InverseOfferStatusno_such_key
				state = fflib.FFParse_want_colon
				goto mainparse
			}

		case fflib.FFParse_want_colon:
			if tok != fflib.FFTok_colon {
				wantedTok = fflib.FFTok_colon
				goto wrongtokenerror
			}
			state = fflib.FFParse_want_value
			continue
		case fflib.FFParse_want_value:

			if tok == fflib.FFTok_left_brace || tok == fflib.FFTok_left_bracket || tok == fflib.FFTok_integer || tok == fflib.FFTok_double || tok == fflib.FFTok_string || tok == fflib.FFTok_bool || tok == fflib.FFTok_null {
				switch currentKey {

				case ffj_t_InverseOfferStatus_Status:
					goto handle_Status

				case ffj_t_InverseOfferStatus_FrameworkID:
					goto handle_FrameworkID

				case ffj_t_InverseOfferStatus_Timestamp:
					goto handle_Timestamp

				case ffj_t_InverseOfferStatusno_such_key:
					err = fs.SkipField(tok)
					if err != nil {
						return fs.WrapErr(err)
					}
					state = fflib.FFParse_after_value
					goto mainparse
				}
			} else {
				goto wantedvalue
			}
		}
	}

handle_Status:

	/* handler: uj.Status type=allocator.InverseOfferStatus_Status kind=int32 quoted=false*/

	{
		if tok == fflib.FFTok_null {

			uj.Status = nil

			state = fflib.FFParse_after_value
			goto mainparse
		}

		tbuf, err := fs.CaptureField(tok)
		if err != nil {
			return fs.WrapErr(err)
		}

		if uj.Status == nil {
			uj.Status = new(InverseOfferStatus_Status)
		}

		err = uj.Status.UnmarshalJSON(tbuf)
		if err != nil {
			return fs.WrapErr(err)
		}
		state = fflib.FFParse_after_value
	}

	state = fflib.FFParse_after_value
	goto mainparse

handle_FrameworkID:

	/* handler: uj.FrameworkID type=mesos.FrameworkID kind=struct quoted=false*/

	{
		if tok == fflib.FFTok_null {

			state = fflib.FFParse_after_value
			goto mainparse
		}

		err = uj.FrameworkID.UnmarshalJSONFFLexer(fs, fflib.FFParse_want_key)
		if err != nil {
			return err
		}
		state = fflib.FFParse_after_value
	}

	state = fflib.FFParse_after_value
	goto mainparse

handle_Timestamp:

	/* handler: uj.Timestamp type=mesos.TimeInfo kind=struct quoted=false*/

	{
		if tok == fflib.FFTok_null {

			state = fflib.FFParse_after_value
			goto mainparse
		}

		err = uj.Timestamp.UnmarshalJSONFFLexer(fs, fflib.FFParse_want_key)
		if err != nil {
			return err
		}
		state = fflib.FFParse_after_value
	}

	state = fflib.FFParse_after_value
	goto mainparse

wantedvalue:
	return fs.WrapErr(fmt.Errorf("wanted value token, but got token: %v", tok))
wrongtokenerror:
	return fs.WrapErr(fmt.Errorf("ffjson: wanted token: %v, but got token: %v output=%s", wantedTok, tok, fs.Output.String()))
tokerror:
	if fs.BigError != nil {
		return fs.WrapErr(fs.BigError)
	}
	err = fs.Error.ToError()
	if err != nil {
		return fs.WrapErr(err)
	}
	panic("ffjson-generated: unreachable, please report bug.")
done:

	return nil
}
