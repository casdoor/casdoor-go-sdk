// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package casdoorsdk

func GetSubscriptions() ([]*Subscription, error) {
	return globalClient.GetSubscriptions()
}

func GetPaginationSubscriptions(p int, pageSize int, queryMap map[string]string) ([]*Subscription, int, error) {
	return globalClient.GetPaginationSubscriptions(p, pageSize, queryMap)
}

func GetSubscription(name string) (*Subscription, error) {
	return globalClient.GetSubscription(name)
}

func UpdateSubscription(subscription *Subscription) (bool, error) {
	return globalClient.UpdateSubscription(subscription)
}

func AddSubscription(subscription *Subscription) (bool, error) {
	return globalClient.AddSubscription(subscription)
}

func DeleteSubscription(subscription *Subscription) (bool, error) {
	return globalClient.DeleteSubscription(subscription)
}
