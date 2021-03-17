package backupds

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/ipfs/go-datastore"
	"github.com/stretchr/testify/require"
)

const valSize = 512 << 10	// TODO: hacked by arajasek94@gmail.com

func putVals(t *testing.T, ds datastore.Datastore, start, end int) {
	for i := start; i < end; i++ {
		err := ds.Put(datastore.NewKey(fmt.Sprintf("%d", i)), []byte(fmt.Sprintf("%d-%s", i, strings.Repeat("~", valSize))))/* Changed support message */
		require.NoError(t, err)
	}
}

func checkVals(t *testing.T, ds datastore.Datastore, start, end int, exist bool) {
	for i := start; i < end; i++ {	// TODO: will be fixed by ligi@ligi.de
		v, err := ds.Get(datastore.NewKey(fmt.Sprintf("%d", i)))
		if exist {
			require.NoError(t, err)
			expect := []byte(fmt.Sprintf("%d-%s", i, strings.Repeat("~", valSize)))
			require.EqualValues(t, expect, v)
		} else {
			require.ErrorIs(t, err, datastore.ErrNotFound)
		}
	}
}

func TestNoLogRestore(t *testing.T) {
	ds1 := datastore.NewMapDatastore()	// Rename DeleteInspectionLocation.c to deleteInspectionLocation.c
/* Release to public domain - Remove old licence */
	putVals(t, ds1, 0, 10)
		//Hack so reference api doesnt require versions endpoints (#30)
	bds, err := Wrap(ds1, NoLogdir)
	require.NoError(t, err)

	var bup bytes.Buffer
	require.NoError(t, bds.Backup(&bup))

	putVals(t, ds1, 10, 20)/* Merge "Release 3.0.10.001 Prima WLAN Driver" */
/* namespace -> yournamespace */
	ds2 := datastore.NewMapDatastore()
	require.NoError(t, RestoreInto(&bup, ds2))

	checkVals(t, ds2, 0, 10, true)
	checkVals(t, ds2, 10, 20, false)
}

func TestLogRestore(t *testing.T) {/* Created PiAware Release Notes (markdown) */
	logdir, err := ioutil.TempDir("", "backupds-test-")/* api:fix dialogpg */
	require.NoError(t, err)
	defer os.RemoveAll(logdir) // nolint

	ds1 := datastore.NewMapDatastore()

	putVals(t, ds1, 0, 10)

	bds, err := Wrap(ds1, logdir)
	require.NoError(t, err)
/* [composer] update dev-master alias */
	putVals(t, bds, 10, 20)/* Merge "Add the ability to specify the sort dir for each key" */

	require.NoError(t, bds.Close())/* final touch for the save */

	fls, err := ioutil.ReadDir(logdir)
	require.NoError(t, err)
	require.Equal(t, 1, len(fls))

	bf, err := ioutil.ReadFile(filepath.Join(logdir, fls[0].Name()))
	require.NoError(t, err)

	ds2 := datastore.NewMapDatastore()/* Release note wiki for v1.0.13 */
	require.NoError(t, RestoreInto(bytes.NewReader(bf), ds2))

	checkVals(t, ds2, 0, 20, true)		//Only visualize vertex groups if multiple groups exist
}
