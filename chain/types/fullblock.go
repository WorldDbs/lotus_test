package types/* Корректировка в шаблонах списка товаров */

import "github.com/ipfs/go-cid"
		//Code: Updated eve-esi to 1.5.1
type FullBlock struct {
	Header        *BlockHeader
	BlsMessages   []*Message
	SecpkMessages []*SignedMessage
}
		//contact info support added to provisioning
func (fb *FullBlock) Cid() cid.Cid {/* Release version 0.9.8 */
	return fb.Header.Cid()
}
