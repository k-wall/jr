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

type FleetMgmtLocation struct {
	Vehicle_id int32 `json:"vehicle_id"`

	Location Location `json:"location"`

	Ts int64 `json:"ts"`
}

const FleetMgmtLocationAvroCRC64Fingerprint = "\xcc\xc2MA\xb8\xb2\x82\xa4"

func NewFleetMgmtLocation() FleetMgmtLocation {
	r := FleetMgmtLocation{}
	r.Location = NewLocation()

	return r
}

func DeserializeFleetMgmtLocation(r io.Reader) (FleetMgmtLocation, error) {
	t := NewFleetMgmtLocation()
	deser, err := compiler.CompileSchemaBytes([]byte(t.Schema()), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func DeserializeFleetMgmtLocationFromSchema(r io.Reader, schema string) (FleetMgmtLocation, error) {
	t := NewFleetMgmtLocation()

	deser, err := compiler.CompileSchemaBytes([]byte(schema), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func writeFleetMgmtLocation(r FleetMgmtLocation, w io.Writer) error {
	var err error
	err = vm.WriteInt(r.Vehicle_id, w)
	if err != nil {
		return err
	}
	err = writeLocation(r.Location, w)
	if err != nil {
		return err
	}
	err = vm.WriteLong(r.Ts, w)
	if err != nil {
		return err
	}
	return err
}

func (r FleetMgmtLocation) Serialize(w io.Writer) error {
	return writeFleetMgmtLocation(r, w)
}

func (r FleetMgmtLocation) Schema() string {
	return "{\"fields\":[{\"name\":\"vehicle_id\",\"type\":{\"arg.properties\":{\"range\":{\"max\":9999,\"min\":1000}},\"type\":\"int\"}},{\"arg.properties\":{\"options\":[{\"latitude\":37.416834,\"longitude\":-121.975002},{\"latitude\":37.664725,\"longitude\":-121.79737},{\"latitude\":38.124221,\"longitude\":-121.182135},{\"latitude\":38.623772,\"longitude\":-120.176886},{\"latitude\":39.544752,\"longitude\":-119.589118},{\"latitude\":37.725577,\"longitude\":-121.709479},{\"latitude\":38.658096,\"longitude\":-120.138434},{\"latitude\":39.752002,\"longitude\":-119.050787},{\"latitude\":39.680169,\"longitude\":-119.111212},{\"latitude\":40.445308,\"longitude\":-118.30921},{\"latitude\":40.122644,\"longitude\":-118.517951},{\"latitude\":40.907734,\"longitude\":-117.831305},{\"latitude\":39.718207,\"longitude\":-119.105719},{\"latitude\":40.853569,\"longitude\":-117.226383},{\"latitude\":40.683171,\"longitude\":-116.622809},{\"latitude\":40.82604,\"longitude\":-117.169294},{\"latitude\":37.699504,\"longitude\":-121.731452},{\"latitude\":39.912298,\"longitude\":-118.825568},{\"latitude\":39.895442,\"longitude\":-118.880499},{\"latitude\":39.827978,\"longitude\":-118.979376},{\"latitude\":38.167421,\"longitude\":-121.116217},{\"latitude\":40.628996,\"longitude\":-116.436041},{\"latitude\":38.537888,\"longitude\":-120.413092},{\"latitude\":39.756225,\"longitude\":-119.056281},{\"latitude\":39.92915,\"longitude\":-118.748663},{\"latitude\":40.924338,\"longitude\":-117.847785},{\"latitude\":40.828808,\"longitude\":-117.166632},{\"latitude\":38.499207,\"longitude\":-120.49549},{\"latitude\":39.760448,\"longitude\":-119.061774},{\"latitude\":37.391868,\"longitude\":-122.071206}]},\"name\":\"location\",\"type\":{\"fields\":[{\"name\":\"latitude\",\"type\":\"double\"},{\"name\":\"longitude\",\"type\":\"double\"}],\"name\":\"location\",\"type\":\"record\"}},{\"name\":\"ts\",\"type\":{\"arg.properties\":{\"iteration\":{\"start\":1609459200000,\"step\":100000}},\"logicalType\":\"timestamp-millis\",\"type\":\"long\"}}],\"name\":\"fleet_mgmt.FleetMgmtLocation\",\"type\":\"record\"}"
}

func (r FleetMgmtLocation) SchemaName() string {
	return "fleet_mgmt.FleetMgmtLocation"
}

func (_ FleetMgmtLocation) SetBoolean(v bool)    { panic("Unsupported operation") }
func (_ FleetMgmtLocation) SetInt(v int32)       { panic("Unsupported operation") }
func (_ FleetMgmtLocation) SetLong(v int64)      { panic("Unsupported operation") }
func (_ FleetMgmtLocation) SetFloat(v float32)   { panic("Unsupported operation") }
func (_ FleetMgmtLocation) SetDouble(v float64)  { panic("Unsupported operation") }
func (_ FleetMgmtLocation) SetBytes(v []byte)    { panic("Unsupported operation") }
func (_ FleetMgmtLocation) SetString(v string)   { panic("Unsupported operation") }
func (_ FleetMgmtLocation) SetUnionElem(v int64) { panic("Unsupported operation") }

func (r *FleetMgmtLocation) Get(i int) types.Field {
	switch i {
	case 0:
		w := types.Int{Target: &r.Vehicle_id}

		return w

	case 1:
		r.Location = NewLocation()

		w := types.Record{Target: &r.Location}

		return w

	case 2:
		w := types.Long{Target: &r.Ts}

		return w

	}
	panic("Unknown field index")
}

func (r *FleetMgmtLocation) SetDefault(i int) {
	switch i {
	}
	panic("Unknown field index")
}

func (r *FleetMgmtLocation) NullField(i int) {
	switch i {
	}
	panic("Not a nullable field index")
}

func (_ FleetMgmtLocation) AppendMap(key string) types.Field { panic("Unsupported operation") }
func (_ FleetMgmtLocation) AppendArray() types.Field         { panic("Unsupported operation") }
func (_ FleetMgmtLocation) HintSize(int)                     { panic("Unsupported operation") }
func (_ FleetMgmtLocation) Finalize()                        {}

func (_ FleetMgmtLocation) AvroCRC64Fingerprint() []byte {
	return []byte(FleetMgmtLocationAvroCRC64Fingerprint)
}

func (r FleetMgmtLocation) MarshalJSON() ([]byte, error) {
	var err error
	output := make(map[string]json.RawMessage)
	output["vehicle_id"], err = json.Marshal(r.Vehicle_id)
	if err != nil {
		return nil, err
	}
	output["location"], err = json.Marshal(r.Location)
	if err != nil {
		return nil, err
	}
	output["ts"], err = json.Marshal(r.Ts)
	if err != nil {
		return nil, err
	}
	return json.Marshal(output)
}

func (r *FleetMgmtLocation) UnmarshalJSON(data []byte) error {
	var fields map[string]json.RawMessage
	if err := json.Unmarshal(data, &fields); err != nil {
		return err
	}

	var val json.RawMessage
	val = func() json.RawMessage {
		if v, ok := fields["vehicle_id"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Vehicle_id); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for vehicle_id")
	}
	val = func() json.RawMessage {
		if v, ok := fields["location"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Location); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for location")
	}
	val = func() json.RawMessage {
		if v, ok := fields["ts"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Ts); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for ts")
	}
	return nil
}
