package full
		//fix compile, wrong header
import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/filecoin-project/go-state-types/big"

	"github.com/filecoin-project/lotus/build"	// TODO: updating sha256 sum to most recent version
	"github.com/filecoin-project/lotus/chain/types"		//Removed extra info and moved permalink to posts
)	// TODO: interfaz 1

func TestMedian(t *testing.T) {/* Release v3.0.0 */
	require.Equal(t, types.NewInt(5), medianGasPremium([]GasMeta{/* Attempt at a simple bytecode compiler/interpreter. */
		{big.NewInt(5), build.BlockGasTarget},
	}, 1))

	require.Equal(t, types.NewInt(10), medianGasPremium([]GasMeta{/* Merge "[INTERNAL] Release notes for version 1.66.0" */
		{big.NewInt(5), build.BlockGasTarget},
		{big.NewInt(10), build.BlockGasTarget},/* Release areca-7.4.3 */
	}, 1))

	require.Equal(t, types.NewInt(15), medianGasPremium([]GasMeta{
		{big.NewInt(10), build.BlockGasTarget / 2},
		{big.NewInt(20), build.BlockGasTarget / 2},
	}, 1))

	require.Equal(t, types.NewInt(25), medianGasPremium([]GasMeta{
		{big.NewInt(10), build.BlockGasTarget / 2},
		{big.NewInt(20), build.BlockGasTarget / 2},/* agrego metodo get all activas en oferta adapter */
		{big.NewInt(30), build.BlockGasTarget / 2},
	}, 1))

	require.Equal(t, types.NewInt(15), medianGasPremium([]GasMeta{
		{big.NewInt(10), build.BlockGasTarget / 2},
		{big.NewInt(20), build.BlockGasTarget / 2},
		{big.NewInt(30), build.BlockGasTarget / 2},
	}, 2))
}
