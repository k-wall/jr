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
 *     insurance_offers.avsc
 *     inventory.avsc
 *     map_dumb_schema.avsc
 *     net_device.avsc
 *     orders.avsc
 *     page_views.avsc
 *     payroll_bonus.avsc
 *     payroll_employee.avsc
 *     payroll_employee_location.avsc
 *     pizza_orders.avsc
 *     pizza_orders_cancelled.avsc
 *     pizza_orders_completed.avsc
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

type Ratings struct {
	Rating_id int64 `json:"rating_id"`

	User_id int32 `json:"user_id"`

	Stars int32 `json:"stars"`

	Route_id int32 `json:"route_id"`

	Rating_time int64 `json:"rating_time"`

	Channel string `json:"channel"`

	Message string `json:"message"`
}

const RatingsAvroCRC64Fingerprint = "\x14\xea\xdb;\xcc\xd4\xe8r"

func NewRatings() Ratings {
	r := Ratings{}
	return r
}

func DeserializeRatings(r io.Reader) (Ratings, error) {
	t := NewRatings()
	deser, err := compiler.CompileSchemaBytes([]byte(t.Schema()), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func DeserializeRatingsFromSchema(r io.Reader, schema string) (Ratings, error) {
	t := NewRatings()

	deser, err := compiler.CompileSchemaBytes([]byte(schema), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func writeRatings(r Ratings, w io.Writer) error {
	var err error
	err = vm.WriteLong(r.Rating_id, w)
	if err != nil {
		return err
	}
	err = vm.WriteInt(r.User_id, w)
	if err != nil {
		return err
	}
	err = vm.WriteInt(r.Stars, w)
	if err != nil {
		return err
	}
	err = vm.WriteInt(r.Route_id, w)
	if err != nil {
		return err
	}
	err = vm.WriteLong(r.Rating_time, w)
	if err != nil {
		return err
	}
	err = vm.WriteString(r.Channel, w)
	if err != nil {
		return err
	}
	err = vm.WriteString(r.Message, w)
	if err != nil {
		return err
	}
	return err
}

func (r Ratings) Serialize(w io.Writer) error {
	return writeRatings(r, w)
}

func (r Ratings) Schema() string {
	return "{\"fields\":[{\"name\":\"rating_id\",\"type\":{\"arg.properties\":{\"iteration\":{\"start\":1}},\"type\":\"long\"}},{\"name\":\"user_id\",\"type\":{\"arg.properties\":{\"range\":{\"max\":20,\"min\":-1}},\"type\":\"int\"}},{\"name\":\"stars\",\"type\":{\"arg.properties\":{\"range\":{\"max\":5,\"min\":1}},\"type\":\"int\"}},{\"name\":\"route_id\",\"type\":{\"arg.properties\":{\"range\":{\"max\":9999,\"min\":1}},\"type\":\"int\"}},{\"name\":\"rating_time\",\"type\":{\"arg.properties\":{\"iteration\":{\"start\":1,\"step\":12}},\"format_as_time\":\"unix_long\",\"type\":\"long\"}},{\"name\":\"channel\",\"type\":{\"arg.properties\":{\"options\":[\"web\",\"iOS\",\"ios\",\"iOS-test\",\"android\"]},\"type\":\"string\"}},{\"name\":\"message\",\"type\":{\"arg.properties\":{\"options\":[\"worst. flight. ever. #neveragain\",\"Surprisingly good, maybe you are getting your mojo back at long last!\",\"Exceeded all my expectations. Thank you !\",\"meh\",\"(expletive deleted)\",\"is this as good as it gets? really ?\",\"airport refurb looks great, will fly outta here more!\",\"why is it so difficult to keep the bathrooms clean ?\",\"thank you for the most friendly, helpful experience today at your new lounge\",\"more peanuts please\",\"your team here rocks!\"]},\"type\":\"string\"}}],\"name\":\"ksql.Ratings\",\"type\":\"record\"}"
}

func (r Ratings) SchemaName() string {
	return "ksql.Ratings"
}

func (_ Ratings) SetBoolean(v bool)    { panic("Unsupported operation") }
func (_ Ratings) SetInt(v int32)       { panic("Unsupported operation") }
func (_ Ratings) SetLong(v int64)      { panic("Unsupported operation") }
func (_ Ratings) SetFloat(v float32)   { panic("Unsupported operation") }
func (_ Ratings) SetDouble(v float64)  { panic("Unsupported operation") }
func (_ Ratings) SetBytes(v []byte)    { panic("Unsupported operation") }
func (_ Ratings) SetString(v string)   { panic("Unsupported operation") }
func (_ Ratings) SetUnionElem(v int64) { panic("Unsupported operation") }

func (r *Ratings) Get(i int) types.Field {
	switch i {
	case 0:
		w := types.Long{Target: &r.Rating_id}

		return w

	case 1:
		w := types.Int{Target: &r.User_id}

		return w

	case 2:
		w := types.Int{Target: &r.Stars}

		return w

	case 3:
		w := types.Int{Target: &r.Route_id}

		return w

	case 4:
		w := types.Long{Target: &r.Rating_time}

		return w

	case 5:
		w := types.String{Target: &r.Channel}

		return w

	case 6:
		w := types.String{Target: &r.Message}

		return w

	}
	panic("Unknown field index")
}

func (r *Ratings) SetDefault(i int) {
	switch i {
	}
	panic("Unknown field index")
}

func (r *Ratings) NullField(i int) {
	switch i {
	}
	panic("Not a nullable field index")
}

func (_ Ratings) AppendMap(key string) types.Field { panic("Unsupported operation") }
func (_ Ratings) AppendArray() types.Field         { panic("Unsupported operation") }
func (_ Ratings) HintSize(int)                     { panic("Unsupported operation") }
func (_ Ratings) Finalize()                        {}

func (_ Ratings) AvroCRC64Fingerprint() []byte {
	return []byte(RatingsAvroCRC64Fingerprint)
}

func (r Ratings) MarshalJSON() ([]byte, error) {
	var err error
	output := make(map[string]json.RawMessage)
	output["rating_id"], err = json.Marshal(r.Rating_id)
	if err != nil {
		return nil, err
	}
	output["user_id"], err = json.Marshal(r.User_id)
	if err != nil {
		return nil, err
	}
	output["stars"], err = json.Marshal(r.Stars)
	if err != nil {
		return nil, err
	}
	output["route_id"], err = json.Marshal(r.Route_id)
	if err != nil {
		return nil, err
	}
	output["rating_time"], err = json.Marshal(r.Rating_time)
	if err != nil {
		return nil, err
	}
	output["channel"], err = json.Marshal(r.Channel)
	if err != nil {
		return nil, err
	}
	output["message"], err = json.Marshal(r.Message)
	if err != nil {
		return nil, err
	}
	return json.Marshal(output)
}

func (r *Ratings) UnmarshalJSON(data []byte) error {
	var fields map[string]json.RawMessage
	if err := json.Unmarshal(data, &fields); err != nil {
		return err
	}

	var val json.RawMessage
	val = func() json.RawMessage {
		if v, ok := fields["rating_id"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Rating_id); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for rating_id")
	}
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
		if v, ok := fields["stars"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Stars); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for stars")
	}
	val = func() json.RawMessage {
		if v, ok := fields["route_id"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Route_id); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for route_id")
	}
	val = func() json.RawMessage {
		if v, ok := fields["rating_time"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Rating_time); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for rating_time")
	}
	val = func() json.RawMessage {
		if v, ok := fields["channel"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Channel); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for channel")
	}
	val = func() json.RawMessage {
		if v, ok := fields["message"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Message); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for message")
	}
	return nil
}
