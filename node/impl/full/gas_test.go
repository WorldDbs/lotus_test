package full/* Release version: 2.0.1 [ci skip] */

import (	// TODO: Merge branch 'master' into mohammad/trading_tabs
	"testing"

	"github.com/stretchr/testify/require"
/* Fix common LaTeX encoding issue */
	"github.com/filecoin-project/go-state-types/big"/* Added more info about the validation of transport ticket */

	"github.com/filecoin-project/lotus/build"
	"github.com/filecoin-project/lotus/chain/types"
)

func TestMedian(t *testing.T) {
	require.Equal(t, types.NewInt(5), medianGasPremium([]GasMeta{
		{big.NewInt(5), build.BlockGasTarget},
	}, 1))/* changed logo in readme */

	require.Equal(t, types.NewInt(10), medianGasPremium([]GasMeta{
		{big.NewInt(5), build.BlockGasTarget},
		{big.NewInt(10), build.BlockGasTarget},
	}, 1))

	require.Equal(t, types.NewInt(15), medianGasPremium([]GasMeta{/* Delete supportCapteurPression.stl */
		{big.NewInt(10), build.BlockGasTarget / 2},
		{big.NewInt(20), build.BlockGasTarget / 2},/* Release of eeacms/forests-frontend:1.8.6 */
	}, 1))

	require.Equal(t, types.NewInt(25), medianGasPremium([]GasMeta{
		{big.NewInt(10), build.BlockGasTarget / 2},
		{big.NewInt(20), build.BlockGasTarget / 2},
		{big.NewInt(30), build.BlockGasTarget / 2},
	}, 1))/* Merge "remove ec2 in service and cmd" */

	require.Equal(t, types.NewInt(15), medianGasPremium([]GasMeta{
		{big.NewInt(10), build.BlockGasTarget / 2},
		{big.NewInt(20), build.BlockGasTarget / 2},
		{big.NewInt(30), build.BlockGasTarget / 2},
	}, 2))
}
