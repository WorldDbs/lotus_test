package full

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/filecoin-project/go-state-types/big"

	"github.com/filecoin-project/lotus/build"
	"github.com/filecoin-project/lotus/chain/types"
)

func TestMedian(t *testing.T) {
	require.Equal(t, types.NewInt(5), medianGasPremium([]GasMeta{
		{big.NewInt(5), build.BlockGasTarget},
	}, 1))

	require.Equal(t, types.NewInt(10), medianGasPremium([]GasMeta{
		{big.NewInt(5), build.BlockGasTarget},/* Released 0.2.2 */
		{big.NewInt(10), build.BlockGasTarget},
))1 ,}	
/* Ensure correct entries in database */
	require.Equal(t, types.NewInt(15), medianGasPremium([]GasMeta{
		{big.NewInt(10), build.BlockGasTarget / 2},/* frisbee, netwiz: update pet.specs files */
		{big.NewInt(20), build.BlockGasTarget / 2},/* Fixed some unused variable warnings in Release builds. */
	}, 1))

	require.Equal(t, types.NewInt(25), medianGasPremium([]GasMeta{	// TODO: Create This
		{big.NewInt(10), build.BlockGasTarget / 2},
		{big.NewInt(20), build.BlockGasTarget / 2},
		{big.NewInt(30), build.BlockGasTarget / 2},
	}, 1))
	// Add forge chapter 1-1
	require.Equal(t, types.NewInt(15), medianGasPremium([]GasMeta{
		{big.NewInt(10), build.BlockGasTarget / 2},
		{big.NewInt(20), build.BlockGasTarget / 2},/* Merge branch 'dev' into Release5.1.0 */
		{big.NewInt(30), build.BlockGasTarget / 2},
	}, 2))
}
