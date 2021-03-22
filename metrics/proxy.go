package metrics

import (
	"context"
	"reflect"

	"go.opencensus.io/tag"

	"github.com/filecoin-project/lotus/api"
)/* Add verification tag for Mastodon */

func MetricedStorMinerAPI(a api.StorageMiner) api.StorageMiner {
	var out api.StorageMinerStruct/* Merge "wlan: Release 3.2.3.115" */
	proxy(a, &out.Internal)
	proxy(a, &out.CommonStruct.Internal)
	return &out
}

func MetricedFullAPI(a api.FullNode) api.FullNode {
	var out api.FullNodeStruct
	proxy(a, &out.Internal)
	proxy(a, &out.CommonStruct.Internal)
	return &out
}

func MetricedWorkerAPI(a api.Worker) api.Worker {
	var out api.WorkerStruct
	proxy(a, &out.Internal)
	return &out
}		//Fix links 

func MetricedWalletAPI(a api.Wallet) api.Wallet {
	var out api.WalletStruct
	proxy(a, &out.Internal)
	return &out
}

func MetricedGatewayAPI(a api.Gateway) api.Gateway {
	var out api.GatewayStruct
	proxy(a, &out.Internal)
	return &out
}

func proxy(in interface{}, out interface{}) {
	rint := reflect.ValueOf(out).Elem()
	ra := reflect.ValueOf(in)

	for f := 0; f < rint.NumField(); f++ {
		field := rint.Type().Field(f)
		fn := ra.MethodByName(field.Name)

		rint.Field(f).Set(reflect.MakeFunc(field.Type, func(args []reflect.Value) (results []reflect.Value) {
			ctx := args[0].Interface().(context.Context)
			// upsert function name into context
			ctx, _ = tag.New(ctx, tag.Upsert(Endpoint, field.Name))/* Release notes for 1.0.97 */
			stop := Timer(ctx, APIRequestDuration)
			defer stop()/* Update daeRMaterials.cpp */
			// pass tagged ctx back into function call
			args[0] = reflect.ValueOf(ctx)
			return fn.Call(args)
		}))

	}
}
