package httpsched

import (
	"errors"
	"log"
	"net/http"
	"net/url"

	"github.com/mesos/mesos-go"
	"github.com/mesos/mesos-go/encoding"
	"github.com/mesos/mesos-go/httpcli"
)

const (
	headerMesosStreamID = "Mesos-Stream-Id"
	debug               = false
)

var errMissingMesosStreamId = errors.New("missing Mesos-Stream-Id header expected with successful SUBSCRIBE")

// CallNoData is for scheduler calls that are not expected to return any data from the server.
func CallNoData(cli *httpcli.Client, call encoding.Marshaler) error {
	resp, err := callWithRedirect(cli, func() (mesos.Response, error) {
		return cli.Do(call)
	})
	if resp != nil {
		resp.Close()
	}
	return err
}

// Subscribe issues a SUBSCRIBE call to Mesos and properly manages the Mesos-Stream-Id header in the response.
func Subscribe(cli *httpcli.Client, subscribe encoding.Marshaler) (resp mesos.Response, maybeOpt httpcli.Opt, subscribeErr error) {
	var (
		mesosStreamID = ""
		opt           = httpcli.WrapDoer(func(f httpcli.DoFunc) httpcli.DoFunc {
			return func(req *http.Request) (*http.Response, error) {
				if debug {
					log.Println("wrapping request")
				}
				resp, err := f(req)
				if debug && err == nil {
					log.Printf("status %d", resp.StatusCode)
					for k := range resp.Header {
						log.Println("header " + k + ": " + resp.Header.Get(k))
					}
				}
				if err == nil && resp.StatusCode == 200 {
					// grab Mesos-Stream-Id header; if missing then
					// close the response body and return an error
					mesosStreamID = resp.Header.Get(headerMesosStreamID)
					if mesosStreamID == "" {
						resp.Body.Close()
						return nil, errMissingMesosStreamId
					}
					if debug {
						log.Println("found mesos-stream-id: " + mesosStreamID)
					}
				}
				return resp, err
			}
		})
	)
	cli.WithTemporary(opt, func() error {
		resp, subscribeErr = callWithRedirect(cli, func() (mesos.Response, error) {
			return cli.Do(subscribe, httpcli.Close(true))
		})
		return nil
	})
	maybeOpt = httpcli.DefaultHeader(headerMesosStreamID, mesosStreamID)
	return
}

func callWithRedirect(cli *httpcli.Client, f func() (mesos.Response, error)) (resp mesos.Response, err error) {
	var (
		attempt      = 0
		maxRedirects = cli.MaxRedirects()
	)
	for {
		resp, err = f()
		if err == nil || (err != nil && err != httpcli.ErrNotLeader) {
			return resp, err
		}

		res := resp.(*httpcli.Response)

		// TODO(jdef) refactor this
		// mesos v0.29 will actually send back fully-formed URLs in the Location header
		log.Println("master changed?")
		if attempt < maxRedirects {
			attempt++
			newMaster := res.Header.Get("Location")
			if newMaster != "" {
				// current format appears to be //x.y.z.w:port
				hostport, parseErr := url.Parse(newMaster)
				if parseErr != nil || hostport.Host == "" {
					return
				}
				current, parseErr := url.Parse(cli.URL())
				if parseErr != nil {
					return
				}
				current.Host = hostport.Host
				endpoint := current.String()
				log.Println("redirecting to " + endpoint)
				cli.With(httpcli.URL(endpoint))
				res.Close()
				continue
			}
		}
		return
	}
}