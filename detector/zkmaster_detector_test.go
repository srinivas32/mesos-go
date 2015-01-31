package detector

import (
	log "github.com/golang/glog"
	mesos "github.com/mesos/mesos-go/mesosproto"
	"github.com/samuel/go-zookeeper/zk"
	"github.com/stretchr/testify/assert"
	"net/url"
	"strconv"
	"sync"
	"testing"
	"time"
)

var zkurl = "zk://127.0.0.1:2181/mesos"
var zkurl_bad = "zk://127.0.0.1:2181"

func TestZkMasterDetectorNew(t *testing.T) {
	md, err := NewZkMasterDetector(zkurl)
	assert.NoError(t, err)
	assert.Equal(t, 1, len(md.zkHosts))
	u, _ := url.Parse(zkurl)
	assert.True(t, u.String() == md.url.String())
	assert.Equal(t, "/mesos", md.zkPath)
}

func TestZkMasterDetectorStart(t *testing.T) {
	c, err := makeZkClient()
	assert.False(t, c.connected)
	md, err := NewZkMasterDetector(zkurl)
	assert.NoError(t, err)
	md.zkClient = c // override zk.Conn with our own.
	err = md.Start()
	assert.NoError(t, err)
	assert.True(t, c.connected)
}

func TestZkMasterDetectorChildrenChanged(t *testing.T) {
	wCh := make(chan struct{}, 1)

	c, err := makeZkClient()
	assert.NoError(t, err)
	assert.False(t, c.connected)

	md, err := NewZkMasterDetector(zkurl)
	assert.NoError(t, err)
	// override zk.Conn with our own.
	md.zkClient = c
	md.zkClient.childrenWatcher = zkChildrenWatcherFunc(md.childrenChanged)

	err = md.Start()
	assert.NoError(t, err)
	assert.True(t, c.connected)

	md.Detect(func(master *mesos.MasterInfo) {
		assert.NotNil(t, master)
		assert.Equal(t, master.GetId(), "master@localhost:5050")
		wCh <- struct{}{}
	})

	select {
	case <-wCh:
	case <-time.After(time.Second * 3):
		panic("Waited too long...")
	}
}

func TestZkMasterDetectMultiple(t *testing.T) {
	ch0 := make(chan zk.Event, 1)
	ch1 := make(chan zk.Event, 1)
	var wg sync.WaitGroup

	ch0 <- zk.Event{
		State: zk.StateConnected,
		Path:  test_zk_path,
	}

	c, err := newZkClient(test_zk_hosts, test_zk_path)
	assert.NoError(t, err)

	connector := makeMockConnector(test_zk_path, ch1)
	c.connFactory = zkConnFactoryFunc(func() (zkConnector, <-chan zk.Event, error) {
		log.V(2).Infof("**** Using zk.Conn adapter ****")
		return connector, ch0, nil
	})

	md, err := NewZkMasterDetector(zkurl)
	assert.NoError(t, err)
	md.zkClient = c
	md.zkClient.childrenWatcher = zkChildrenWatcherFunc(md.childrenChanged)

	err = md.Start()
	assert.NoError(t, err)
	assert.True(t, c.connected)

	md.Detect(func(master *mesos.MasterInfo) {
		log.V(2).Infoln("Leader change detected.")
		wg.Done()
	})

	// **** Test 4 consecutive ChildrenChangedEvents ******
	// setup event changes
	sequences := [][]string{
		[]string{"info_005", "info_010", "info_022"},
		[]string{"info_014", "info_010", "info_005"},
		[]string{"info_005", "info_004", "info_022"},
		[]string{"info_017", "info_099", "info_200"},
	}
	wg.Add(3) // leader changes 3 times.

	go func() {
		conn := NewMockZkConnector()
		md.zkClient.conn = conn

		for i := range sequences {
			path := "/test" + strconv.Itoa(i)
			conn.On("ChildrenW", path).Return([]string{path}, &zk.Stat{}, (<-chan zk.Event)(ch1), nil)
			conn.On("Children", path).Return(sequences[i], &zk.Stat{}, nil)
			conn.On("Get", path).Return(makeTestMasterInfo(), &zk.Stat{}, nil)
			md.zkClient.rootPath = path
			ch1 <- zk.Event{
				Type: zk.EventNodeChildrenChanged,
				Path: path,
			}
		}
	}()

	wg.Wait()
}
