// Copyright 2017 Factom Foundation
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package interfaces

import ()

/**************************
 * IRCD  Interface for Redeem Condition Datastructures (RCD)
 *
 * https://github.com/FactomProject/FactomDocs/blob/master/factomDataStructureDetails.md#factoid-transaction
 **************************/

type ISplitable interface {
	IMsg
	MarshalHeader() ([]byte, error)
	MarshalData() ([]byte, error)
	GetDataHash() IHash
}