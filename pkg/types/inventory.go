// Code generated by github.com/actgardner/gogen-avro/v10. DO NOT EDIT.
/*
 * SOURCES:
 *     campaignfinance.avsc
 *     click_stream.avsc
 *     click_stream_users.avsc
 *     clickstreamcodes.avsc
 *     creditcards.avsc
 *     device_information.avsc
 *     fleet_mgmt_description.avsc
 *     fleet_mgmt_location.avsc
 *     fleet_mgmt_sensors.avsc
 *     gaming_games.avsc
 *     gaming_player_activity.avsc
 *     gaming_players.avsc
 *     insurancecustomeractivity.avsc
 *     insurancecustomers.avsc
 *     insuranceoffers.avsc
 *     inventory.avsc
 *     map_dumb_schema.avsc
 *     net_device.avsc
 *     orders.avsc
 *     pageviews.avsc
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
 *     stockTrades.avsc
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

type Inventory struct {
	Id int64 `json:"id"`

	Quantity int64 `json:"quantity"`

	Productid int64 `json:"productid"`
}

const InventoryAvroCRC64Fingerprint = "]>\x8b%\x96!\xc5l"

func NewInventory() Inventory {
	r := Inventory{}
	return r
}

func DeserializeInventory(r io.Reader) (Inventory, error) {
	t := NewInventory()
	deser, err := compiler.CompileSchemaBytes([]byte(t.Schema()), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func DeserializeInventoryFromSchema(r io.Reader, schema string) (Inventory, error) {
	t := NewInventory()

	deser, err := compiler.CompileSchemaBytes([]byte(schema), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func writeInventory(r Inventory, w io.Writer) error {
	var err error
	err = vm.WriteLong(r.Id, w)
	if err != nil {
		return err
	}
	err = vm.WriteLong(r.Quantity, w)
	if err != nil {
		return err
	}
	err = vm.WriteLong(r.Productid, w)
	if err != nil {
		return err
	}
	return err
}

func (r Inventory) Serialize(w io.Writer) error {
	return writeInventory(r, w)
}

func (r Inventory) Schema() string {
	return "{\"fields\":[{\"name\":\"id\",\"type\":{\"arg.properties\":{\"iteration\":{\"start\":0}},\"type\":\"long\"}},{\"name\":\"quantity\",\"type\":{\"arg.properties\":{\"iteration\":{\"start\":0}},\"type\":\"long\"}},{\"name\":\"productid\",\"type\":{\"arg.properties\":{\"iteration\":{\"start\":0}},\"type\":\"long\"}}],\"name\":\"ksql.inventory\",\"type\":\"record\"}"
}

func (r Inventory) SchemaName() string {
	return "ksql.inventory"
}

func (_ Inventory) SetBoolean(v bool)    { panic("Unsupported operation") }
func (_ Inventory) SetInt(v int32)       { panic("Unsupported operation") }
func (_ Inventory) SetLong(v int64)      { panic("Unsupported operation") }
func (_ Inventory) SetFloat(v float32)   { panic("Unsupported operation") }
func (_ Inventory) SetDouble(v float64)  { panic("Unsupported operation") }
func (_ Inventory) SetBytes(v []byte)    { panic("Unsupported operation") }
func (_ Inventory) SetString(v string)   { panic("Unsupported operation") }
func (_ Inventory) SetUnionElem(v int64) { panic("Unsupported operation") }

func (r *Inventory) Get(i int) types.Field {
	switch i {
	case 0:
		w := types.Long{Target: &r.Id}

		return w

	case 1:
		w := types.Long{Target: &r.Quantity}

		return w

	case 2:
		w := types.Long{Target: &r.Productid}

		return w

	}
	panic("Unknown field index")
}

func (r *Inventory) SetDefault(i int) {
	switch i {
	}
	panic("Unknown field index")
}

func (r *Inventory) NullField(i int) {
	switch i {
	}
	panic("Not a nullable field index")
}

func (_ Inventory) AppendMap(key string) types.Field { panic("Unsupported operation") }
func (_ Inventory) AppendArray() types.Field         { panic("Unsupported operation") }
func (_ Inventory) HintSize(int)                     { panic("Unsupported operation") }
func (_ Inventory) Finalize()                        {}

func (_ Inventory) AvroCRC64Fingerprint() []byte {
	return []byte(InventoryAvroCRC64Fingerprint)
}

func (r Inventory) MarshalJSON() ([]byte, error) {
	var err error
	output := make(map[string]json.RawMessage)
	output["id"], err = json.Marshal(r.Id)
	if err != nil {
		return nil, err
	}
	output["quantity"], err = json.Marshal(r.Quantity)
	if err != nil {
		return nil, err
	}
	output["productid"], err = json.Marshal(r.Productid)
	if err != nil {
		return nil, err
	}
	return json.Marshal(output)
}

func (r *Inventory) UnmarshalJSON(data []byte) error {
	var fields map[string]json.RawMessage
	if err := json.Unmarshal(data, &fields); err != nil {
		return err
	}

	var val json.RawMessage
	val = func() json.RawMessage {
		if v, ok := fields["id"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Id); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for id")
	}
	val = func() json.RawMessage {
		if v, ok := fields["quantity"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Quantity); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for quantity")
	}
	val = func() json.RawMessage {
		if v, ok := fields["productid"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Productid); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for productid")
	}
	return nil
}
