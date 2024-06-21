package common

import (
	"errors"
	"fmt"
	"strconv"
	"strings"

	"github.com/btcsuite/btcutil/base58"
)

type UID struct {
	localID    uint32
	objectType int
	shardID    uint32
}

func NewUID(localID uint32, objectType int, shardID uint32) UID {
	return UID{
		localID:    localID,
		objectType: objectType,
		shardID:    shardID,
	}
}

func (u UID) String() string {

	val := uint64(u.localID)<<28 | uint64(u.objectType)<<18 | uint64(u.shardID)<<0
	return base58.Encode([]byte(fmt.Sprintf("%v", val)))
}

func (u UID) GetLocalID() uint32 {
	return u.localID
}

func (u UID) GetShardID() uint32 {
	return u.shardID
}

func (u UID) GetObjectType() int {
	return u.objectType
}

func DecomposeUID(s string) (UID, error) {
	uid, err := strconv.ParseUint(s, 10, 64)

	if err != nil {
		return UID{}, err
	}

	if (1 << 18) > uid {
		return UID{}, errors.New("wrong uid")
	}

	u := UID{
		localID:    uint32(uid >> 28),
		objectType: int(uid >> 18 & 0x3FF),
		shardID:    uint32(uid >> 0 & 0x3FFFF),
	}
	return u, nil
}

func FromBase58(s string) (UID, error) {
	return DecomposeUID(string(base58.Decode(s)))
}

func (uid UID) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf("\"%s\"", uid.String())), nil
}

func (uid *UID) UnmarshalJSON(data []byte) error {
	decodeUID, err := FromBase58(strings.Replace(string(data), "\"", "", -1))
	if err != nil {
		return err
	}

	uid.localID = decodeUID.localID
	uid.objectType = decodeUID.GetObjectType()
	uid.shardID = decodeUID.GetShardID()
	return nil
}

// func (uid *UID) Value() (driver.Value,error){
// 	if uid == nil{
// 		return nil,nil
// 	}
// 	return int64(uid.localID),nil
// }

// // func (uid *UID) Scan(value interface{}) error{
// // 	if value == nil{

// // 	}
// // }
