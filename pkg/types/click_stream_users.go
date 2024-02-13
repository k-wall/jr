// Code generated by github.com/actgardner/gogen-avro/v10. DO NOT EDIT.
/*
 * SOURCES:
 *     campaign_finance.avsc
 *     click_stream.avsc
 *     click_stream_users.avsc
 *     clickstreamcodes.avsc
 *     credit_cards.avsc
 *     device_information.avsc
 *     fleet_mgmt_description.avsc
 *     fleet_mgmt_location.avsc
 *     fleet_mgmt_sensors.avsc
 *     gaming_games.avsc
 *     gaming_player_activity.avsc
 *     gaming_players.avsc
 *     insurance_customer_activity.avsc
 *     insurance_customers.avsc
 *     insuranceoffers.avsc
 *     inventory.avsc
 *     map_dumb_schema.avsc
 *     net_device.avsc
 *     orders.avsc
 *     page_views.avsc
 *     payrollbonus.avsc
 *     payrollemployee.avsc
 *     payrollemployeelocation.avsc
 *     pizzaorders.avsc
 *     pizzaorderscancelled.avsc
 *     pizzaorderscompleted.avsc
 *     product.avsc
 *     purchase.avsc
 *     ratings.avsc
 *     shoe.avsc
 *     shoe_clickstream.avsc
 *     shoe_customer.avsc
 *     shoe_order.avsc
 *     siemlogs.avsc
 *     stock_trades.avsc
 *     stores.avsc
 *     sysloglogs.avsc
 *     transactions.avsc
 *     user.avsc
 *     users.avsc
 */
package types

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/actgardner/gogen-avro/v10/compiler"
	"github.com/actgardner/gogen-avro/v10/vm"
	"github.com/actgardner/gogen-avro/v10/vm/types"
)

var _ = fmt.Printf

type ClickStreamUsers struct {
	User_id int32 `json:"user_id"`

	Username string `json:"username"`

	Registered_at int64 `json:"registered_at"`

	First_name string `json:"first_name"`

	Last_name string `json:"last_name"`

	City string `json:"city"`

	Level string `json:"level"`
}

const ClickStreamUsersAvroCRC64Fingerprint = "\x81\xab7-(\x85\xa3*"

func NewClickStreamUsers() ClickStreamUsers {
	r := ClickStreamUsers{}
	return r
}

func DeserializeClickStreamUsers(r io.Reader) (ClickStreamUsers, error) {
	t := NewClickStreamUsers()
	deser, err := compiler.CompileSchemaBytes([]byte(t.Schema()), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func DeserializeClickStreamUsersFromSchema(r io.Reader, schema string) (ClickStreamUsers, error) {
	t := NewClickStreamUsers()

	deser, err := compiler.CompileSchemaBytes([]byte(schema), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func writeClickStreamUsers(r ClickStreamUsers, w io.Writer) error {
	var err error
	err = vm.WriteInt(r.User_id, w)
	if err != nil {
		return err
	}
	err = vm.WriteString(r.Username, w)
	if err != nil {
		return err
	}
	err = vm.WriteLong(r.Registered_at, w)
	if err != nil {
		return err
	}
	err = vm.WriteString(r.First_name, w)
	if err != nil {
		return err
	}
	err = vm.WriteString(r.Last_name, w)
	if err != nil {
		return err
	}
	err = vm.WriteString(r.City, w)
	if err != nil {
		return err
	}
	err = vm.WriteString(r.Level, w)
	if err != nil {
		return err
	}
	return err
}

func (r ClickStreamUsers) Serialize(w io.Writer) error {
	return writeClickStreamUsers(r, w)
}

func (r ClickStreamUsers) Schema() string {
	return "{\"fields\":[{\"name\":\"user_id\",\"type\":{\"arg.properties\":{\"iteration\":{\"start\":1}},\"type\":\"int\"}},{\"name\":\"username\",\"type\":{\"arg.properties\":{\"options\":[\"akatz1022\",\"bobk_43\",\"alison_99\",\"k_robertz_88\",\"Ferd88\",\"Reeva43\",\"Antonio_0966\",\"ArlyneW8ter\",\"DimitriSchenz88\",\"Oriana_70\",\"AbdelKable_86\",\"Roberto_123\",\"AlanGreta_GG66\",\"Nathan_126\",\"AndySims_345324\",\"GlenAlan_23344\",\"LukeWaters_23\",\"BenSteins_235\"]},\"type\":\"string\"}},{\"name\":\"registered_at\",\"type\":{\"arg.properties\":{\"range\":{\"max\":1502339792000,\"min\":1407645330000}},\"type\":\"long\"}},{\"name\":\"first_name\",\"type\":{\"arg.properties\":{\"options\":[\"Elwyn\",\"Curran\",\"Hanson\",\"Woodrow\",\"Ferd\",\"Reeva\",\"Antonio\",\"Arlyne\",\"Dimitri\",\"Oriana\",\"Abdel\",\"Greta\"]},\"type\":\"string\"}},{\"name\":\"last_name\",\"type\":{\"arg.properties\":{\"options\":[\"Vanyard\",\"Vears\",\"Garrity\",\"Trice\",\"Tomini\",\"Jushcke\",\"De Banke\",\"Pask\",\"Rockhill\",\"Romagosa\",\"Adicot\",\"Lalonde\"]},\"type\":\"string\"}},{\"name\":\"city\",\"type\":{\"arg.properties\":{\"options\":[\"Palo Alto\",\"San Francisco\",\"Raleigh\",\"London\",\"Frankfurt\",\"New York\"]},\"type\":\"string\"}},{\"name\":\"level\",\"type\":{\"arg.properties\":{\"options\":[\"Gold\",\"Silver\",\"Platinum\"]},\"type\":\"string\"}}],\"name\":\"clickstream.ClickStreamUsers\",\"type\":\"record\"}"
}

func (r ClickStreamUsers) SchemaName() string {
	return "clickstream.ClickStreamUsers"
}

func (_ ClickStreamUsers) SetBoolean(v bool)    { panic("Unsupported operation") }
func (_ ClickStreamUsers) SetInt(v int32)       { panic("Unsupported operation") }
func (_ ClickStreamUsers) SetLong(v int64)      { panic("Unsupported operation") }
func (_ ClickStreamUsers) SetFloat(v float32)   { panic("Unsupported operation") }
func (_ ClickStreamUsers) SetDouble(v float64)  { panic("Unsupported operation") }
func (_ ClickStreamUsers) SetBytes(v []byte)    { panic("Unsupported operation") }
func (_ ClickStreamUsers) SetString(v string)   { panic("Unsupported operation") }
func (_ ClickStreamUsers) SetUnionElem(v int64) { panic("Unsupported operation") }

func (r *ClickStreamUsers) Get(i int) types.Field {
	switch i {
	case 0:
		w := types.Int{Target: &r.User_id}

		return w

	case 1:
		w := types.String{Target: &r.Username}

		return w

	case 2:
		w := types.Long{Target: &r.Registered_at}

		return w

	case 3:
		w := types.String{Target: &r.First_name}

		return w

	case 4:
		w := types.String{Target: &r.Last_name}

		return w

	case 5:
		w := types.String{Target: &r.City}

		return w

	case 6:
		w := types.String{Target: &r.Level}

		return w

	}
	panic("Unknown field index")
}

func (r *ClickStreamUsers) SetDefault(i int) {
	switch i {
	}
	panic("Unknown field index")
}

func (r *ClickStreamUsers) NullField(i int) {
	switch i {
	}
	panic("Not a nullable field index")
}

func (_ ClickStreamUsers) AppendMap(key string) types.Field { panic("Unsupported operation") }
func (_ ClickStreamUsers) AppendArray() types.Field         { panic("Unsupported operation") }
func (_ ClickStreamUsers) HintSize(int)                     { panic("Unsupported operation") }
func (_ ClickStreamUsers) Finalize()                        {}

func (_ ClickStreamUsers) AvroCRC64Fingerprint() []byte {
	return []byte(ClickStreamUsersAvroCRC64Fingerprint)
}

func (r ClickStreamUsers) MarshalJSON() ([]byte, error) {
	var err error
	output := make(map[string]json.RawMessage)
	output["user_id"], err = json.Marshal(r.User_id)
	if err != nil {
		return nil, err
	}
	output["username"], err = json.Marshal(r.Username)
	if err != nil {
		return nil, err
	}
	output["registered_at"], err = json.Marshal(r.Registered_at)
	if err != nil {
		return nil, err
	}
	output["first_name"], err = json.Marshal(r.First_name)
	if err != nil {
		return nil, err
	}
	output["last_name"], err = json.Marshal(r.Last_name)
	if err != nil {
		return nil, err
	}
	output["city"], err = json.Marshal(r.City)
	if err != nil {
		return nil, err
	}
	output["level"], err = json.Marshal(r.Level)
	if err != nil {
		return nil, err
	}
	return json.Marshal(output)
}

func (r *ClickStreamUsers) UnmarshalJSON(data []byte) error {
	var fields map[string]json.RawMessage
	if err := json.Unmarshal(data, &fields); err != nil {
		return err
	}

	var val json.RawMessage
	val = func() json.RawMessage {
		if v, ok := fields["user_id"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.User_id); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for user_id")
	}
	val = func() json.RawMessage {
		if v, ok := fields["username"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Username); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for username")
	}
	val = func() json.RawMessage {
		if v, ok := fields["registered_at"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Registered_at); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for registered_at")
	}
	val = func() json.RawMessage {
		if v, ok := fields["first_name"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.First_name); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for first_name")
	}
	val = func() json.RawMessage {
		if v, ok := fields["last_name"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Last_name); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for last_name")
	}
	val = func() json.RawMessage {
		if v, ok := fields["city"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.City); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for city")
	}
	val = func() json.RawMessage {
		if v, ok := fields["level"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Level); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for level")
	}
	return nil
}
