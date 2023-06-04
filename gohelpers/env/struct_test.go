package env

import (
	"fmt"
	"reflect"
	"testing"
)

type Butterfly int32

func (e Butterfly) MarshalText() ([]byte, error) {
	return []byte(fmt.Sprintf("customized/%X/marshalling", e)), nil
}

type Moth int32

func (e Moth) MarshalText() ([]byte, error) {
	return nil, fmt.Errorf("moths are not butterflies")
}

func TestFromStructRecursive(t *testing.T) {
	type args struct {
		s      interface{}
		prefix string
	}
	tests := []struct {
		name    string
		args    args
		want    []string
		wantErr bool
	}{
		{
			name: "simple_types",
			args: args{
				struct {
					Letters string
					Number  int32
					George  bool
				}{
					"characters",
					42,
					true,
				},
				"simple",
			},
			want: []string{
				// sorted alphabetically
				"SIMPLE_GEORGE=true",
				"SIMPLE_LETTERS=characters",
				"SIMPLE_NUMBER=42",
			},
		},
		{
			name: "map_type",
			args: args{
				map[string]interface{}{
					"Authorization": "Bearer of the ring",
					"Nested": map[string]string{
						"HoneyComb": "BumbleBee",
					},
				},
				"mapper",
			},
			want: []string{
				"MAPPER_AUTHORIZATION=Bearer of the ring",
				"MAPPER_NESTED_HONEY_COMB=BumbleBee",
			},
		},
		{
			name: "skip_fields",
			args: args{
				struct {
					ignoredLowercase string
					Included         string
					CreatedAt        int32
					UpdatedAt        int32
				}{
					"i-am-not",
					"here-i-am",
					1234,
					5678,
				},
				"skipper",
			},
			want: []string{
				"SKIPPER_INCLUDED=here-i-am",
			},
		},
		{
			name: "pointers",
			args: args{
				&struct {
					Concrete string
					NilValue *struct{}
				}{
					"asphalt",
					nil,
				},
				"ptr",
			},
			want: []string{
				"PTR_CONCRETE=asphalt",
			},
		},
		{
			name: "custom_text_marshal",
			args: args{
				struct {
					Creature Butterfly
				}{
					61453, // equals F00D in hexa
				},
				"fragile",
			},
			want: []string{
				"FRAGILE_CREATURE=customized/F00D/marshalling",
			},
		},
		{
			name: "multiple_words",
			args: args{
				struct {
					RunningWater struct {
						StaticStone string
					}
				}{
					struct {
						StaticStone string
					}{
						"otherwise",
					},
				},
				"people",
			},
			want: []string{
				"PEOPLE_RUNNING_WATER_STATIC_STONE=otherwise",
			},
		},
		{
			name: "overwrite_collisions",
			args: args{
				struct {
					SunFlowerTree string
					SunFlower     struct{ Tree string }
					Sun           struct{ FlowerTree string }
				}{
					"brown",
					struct{ Tree string }{"green"},
					struct{ FlowerTree string }{"yellow"},
				},
				"earth",
			},
			want: []string{
				"EARTH_SUN_FLOWER_TREE=yellow",
			},
		},
		{
			name: "errors_inside",
			args: args{
				struct {
					Creature Moth
				}{
					72747,
				},
				"inadmissible",
			},
			want:    []string{},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotMap, err := FromStructRecursive(tt.args.s, tt.args.prefix)
			got := ToSlice(gotMap)
			if (err != nil) != tt.wantErr {
				t.Errorf("FromStructRecursive() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FromStructRecursive() = %#v, want %#v", got, tt.want)
			}
		})
	}
}
