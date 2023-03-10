//  Copyright 2021 Authors of Cilium
//
//  Licensed under the Apache License, Version 2.0 (the "License");
//  you may not use this file except in compliance with the License.
//  You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
//  Unless required by applicable law or agreed to in writing, software
//  distributed under the License is distributed on an "AS IS" BASIS,
//  WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//  See the License for the specific language governing permissions and
//  limitations under the License.

package watchers

import (
	"sync"

	"github.com/cilium/cilium/pkg/k8s"
	cilium_v2a1 "github.com/cilium/cilium/pkg/k8s/apis/cilium.io/v2alpha1"
	"github.com/cilium/cilium/pkg/k8s/informer"
	"github.com/cilium/cilium/pkg/k8s/watchers/subscriber"
	"github.com/cilium/cilium/pkg/kvstore"

	v1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/fields"
	"k8s.io/client-go/tools/cache"
)

var (
	cesNotify = subscriber.NewCES()
	// cepMap maps CEPName to CEBName.
	cepMap = newCEPToCESMap()
)

func (k *K8sWatcher) ciliumEndpointSliceInit(client *k8s.K8sCiliumClient, asyncControllers *sync.WaitGroup) {
	log.Info("Initializing CES controller")
	var once sync.Once

	// Register for all ces updates.
	cesNotify.Register(newCESSubscriber(k))

	for {
		_, cesInformer := informer.NewInformer(
			cache.NewListWatchFromClient(client.CiliumV2alpha1().RESTClient(),
				cilium_v2a1.CESPluralName, v1.NamespaceAll, fields.Everything()),
			&cilium_v2a1.CiliumEndpointSlice{},
			0,
			cache.ResourceEventHandlerFuncs{
				AddFunc: func(obj interface{}) {
					if ces := k8s.ObjToCiliumEndpointSlice(obj); ces != nil {
						cesNotify.NotifyAdd(ces)
					}
				},
				UpdateFunc: func(oldObj, newObj interface{}) {
					if oldCES := k8s.ObjToCiliumEndpointSlice(oldObj); oldCES != nil {
						if newCES := k8s.ObjToCiliumEndpointSlice(newObj); newCES != nil {
							if oldCES.DeepEqual(newCES) {
								return
							}
							cesNotify.NotifyUpdate(oldCES, newCES)
						}
					}
				},
				DeleteFunc: func(obj interface{}) {
					if ces := k8s.ObjToCiliumEndpointSlice(obj); ces != nil {
						cesNotify.NotifyDelete(ces)
					}
				},
			},
			nil,
		)
		isConnected := make(chan struct{})
		// once isConnected is closed, it will stop waiting on caches to be
		// synchronized.
		k.blockWaitGroupToSyncResources(
			isConnected,
			nil,
			cesInformer.HasSynced,
			k8sAPIGroupCiliumEndpointSliceV2Alpha1,
		)

		once.Do(func() {
			// Signalize that we have put node controller in the wait group
			// to sync resources.
			asyncControllers.Done()
		})
		k.k8sAPIGroups.AddAPI(k8sAPIGroupCiliumEndpointSliceV2Alpha1)
		go cesInformer.Run(isConnected)

		<-kvstore.Connected()
		close(isConnected)

		log.Info("Connected to key-value store, stopping CiliumEndpointSlice watcher")
		k.k8sAPIGroups.RemoveAPI(k8sAPIGroupCiliumEndpointSliceV2Alpha1)
		k.cancelWaitGroupToSyncResources(k8sAPIGroupCiliumEndpointSliceV2Alpha1)
		<-kvstore.Client().Disconnected()
		log.Info("Disconnected from key-value store, restarting CiliumEndpointSlice watcher")
	}
}
