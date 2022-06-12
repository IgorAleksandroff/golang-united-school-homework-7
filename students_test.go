package coverage

import (
	"os"
	"reflect"
	"testing"
	"time"
)

// DO NOT EDIT THIS FUNCTION
func init() {
	content, err := os.ReadFile("students_test.go")
	if err != nil {
		panic(err)
	}
	err = os.WriteFile("autocode/students_test", content, 0644)
	if err != nil {
		panic(err)
	}
}

// WRITE YOUR CODE BELOW

var (
	testPerson = Person{
		firstName: "firstPerson",
		lastName:  "lastPerson",
		birthDay:  time.Date(2022, 04, 15, 13, 00, 00, 0, time.Local),
	}
	testPersonAA = Person{
		firstName: "A",
		lastName:  "A",
		birthDay:  time.Date(2022, 04, 15, 13, 00, 01, 0, time.Local),
	}
	testPersonAB = Person{
		firstName: "A",
		lastName:  "B",
		birthDay:  time.Date(2022, 04, 15, 13, 00, 01, 0, time.Local),
	}
	testPersonBA = Person{
		firstName: "B",
		lastName:  "A",
		birthDay:  time.Date(2022, 04, 15, 13, 00, 01, 0, time.Local),
	}

	testMatrix = &Matrix{
		rows: 2,
		cols: 3,
		data: []int{1, 2, 3, 4, 5, 6},
	}
)

func TestMatrix_Cols(t *testing.T) {
	tests := []struct {
		name   string
		matrix *Matrix
		want   [][]int
	}{
		{
			name:   "success",
			matrix: testMatrix,
			want: [][]int{
				{1, 4},
				{2, 5},
				{3, 6},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := Matrix{
				rows: tt.matrix.rows,
				cols: tt.matrix.cols,
				data: tt.matrix.data,
			}
			if got := m.Cols(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Cols() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMatrix_Rows(t *testing.T) {
	tests := []struct {
		name   string
		matrix *Matrix
		want   [][]int
	}{
		{
			name:   "success",
			matrix: testMatrix,
			want: [][]int{
				testMatrix.data[:testMatrix.cols],
				testMatrix.data[testMatrix.cols:],
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := Matrix{
				rows: tt.matrix.rows,
				cols: tt.matrix.cols,
				data: tt.matrix.data,
			}
			if got := m.Rows(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Rows() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMatrix_Set(t *testing.T) {
	type fields struct {
		rows int
		cols int
		data []int
	}
	type args struct {
		row   int
		col   int
		value int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &Matrix{
				rows: tt.fields.rows,
				cols: tt.fields.cols,
				data: tt.fields.data,
			}
			if got := m.Set(tt.args.row, tt.args.col, tt.args.value); got != tt.want {
				t.Errorf("Set() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNew(t *testing.T) {
	type args struct {
		str string
	}
	tests := []struct {
		name    string
		args    args
		want    *Matrix
		wantErr bool
	}{
		{
			name: "success",
			args: args{`1 2 3
											4 5 6`},
			want:    testMatrix,
			wantErr: false,
		},
		{
			name: "field cols",
			args: args{`1 2
													3 4 5`},
			want:    nil,
			wantErr: true,
		},
		{
			name: "field Atoi",
			args: args{`1 2
													A 4`},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := New(tt.args.str)
			if (err != nil) != tt.wantErr {
				t.Errorf("New() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("New() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPeople_Len(t *testing.T) {
	tests := []struct {
		name string
		p    People
		want int
	}{
		{
			name: "success empty",
			p:    People{},
			want: 0,
		},
		{
			name: "success 1",
			p:    People{testPerson},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.p.Len(); got != tt.want {
				t.Errorf("Len() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPeople_Less(t *testing.T) {
	type args struct {
		i int
		j int
	}
	tests := []struct {
		name string
		p    People
		args args
		want bool
	}{
		{
			name: "false empty",
			p: People{
				testPerson,
				testPersonAA,
			},
			args: args{0, 1},
			want: false,
		},
		{
			name: "true date",
			p: People{
				testPersonAA,
				testPerson,
			},
			args: args{0, 1},
			want: true,
		},
		{
			name: "true firstName",
			p: People{
				testPersonAA,
				testPersonBA,
			},
			args: args{0, 1},
			want: true,
		},
		{
			name: "true lastName",
			p: People{
				testPersonAA,
				testPersonAB,
			},
			args: args{0, 1},
			want: true,
		},
		{
			name: "false firstName",
			p: People{
				testPersonBA,
				testPersonAA,
			},
			args: args{0, 1},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.p.Less(tt.args.i, tt.args.j); got != tt.want {
				t.Errorf("Less() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPeople_Swap(t *testing.T) {
	type args struct {
		i int
		j int
	}
	tests := []struct {
		name string
		p    People
		args args
		want People
	}{
		{
			name: "swapped",
			p: People{
				testPersonAA,
				testPersonBA,
			},
			args: args{0, 1},
			want: People{
				testPersonAA,
				testPersonBA,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.p.Swap(tt.args.i, tt.args.j)
		})
		if tt.want[tt.args.i] != tt.p[tt.args.j] || tt.want[tt.args.j] != tt.p[tt.args.i] {
			t.Errorf("After Swap(%v, %v) , want %v and  %v, but get %v and  %v", tt.args.i, tt.args.j, tt.want[tt.args.j], tt.want[tt.args.i], tt.p[tt.args.i], tt.p[tt.args.j])
		}
	}
}
