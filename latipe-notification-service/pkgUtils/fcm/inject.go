package fcm

import "github.com/google/wire"

var Set = wire.NewSet(NewFirebaseSDK)
