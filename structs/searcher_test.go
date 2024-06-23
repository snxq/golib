package structs_test

import (
	"testing"

	"github.com/snxq/golib/structs"
)

type TestA struct {
	A string
	B *TestB
	C []*TestB
}

type TestB struct {
	A string
}

func Test_searcher_SearchField(t *testing.T) {
	type fields struct {
		Origin     interface{}
		delitimter string
	}
	type args struct {
		path string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   []string
	}{
		{
			name: "1_ok_no_nested",
			fields: fields{
				Origin: &TestA{A: "a", B: &TestB{A: "b"}},
			},
			args: args{
				path: "A",
			},
			want: []string{"a"},
		}, {
			name: "2_ok_with_nested",
			fields: fields{
				Origin: &TestA{A: "a", B: &TestB{A: "b"}},
			},
			args: args{
				path: "B.A",
			},
			want: []string{"b"},
		}, {
			name: "3_ok_without_ptr",
			fields: fields{
				Origin: TestA{A: "a", B: &TestB{A: "b"}},
			},
			args: args{
				path: "B.A",
			},
			want: []string{"b"},
		}, {
			name: "4_ok_with_delimiter",
			fields: fields{
				Origin:     TestA{A: "a", B: &TestB{A: "b"}},
				delitimter: "-",
			},
			args: args{
				path: "B-A",
			},
			want: []string{"b"},
		}, {
			name: "5_ok_with_nested_list_*",
			fields: fields{
				Origin: TestA{A: "a", C: []*TestB{{A: "b"}, {A: "c"}}},
			},
			args: args{
				path: "C.*.A",
			},
			want: []string{"b", "c"},
		}, {
			name: "6_ok_with_nested_list_number",
			fields: fields{
				Origin: TestA{A: "a", C: []*TestB{{A: "b"}, {A: "c"}}},
			},
			args: args{
				path: "C.1.A",
			},
			want: []string{"c"},
		}, {
			name: "7_ok_with_nested_list_number_but_out_of_range",
			fields: fields{
				Origin: TestA{A: "a", C: []*TestB{{A: "b"}, {A: "c"}}},
			},
			args: args{
				path: "C.10.A",
			},
			want: []string{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := structs.NewSearcher(tt.fields.Origin, structs.WithDelitimter(tt.fields.delitimter))
			got := s.SearchField(tt.args.path)
			for idx, v := range got {
				if v.String() != tt.want[idx] {
					t.Errorf("searcher.SearchField()[%d] = %v, want %v", idx, got, tt.want)
				}
			}
		})
	}
}
