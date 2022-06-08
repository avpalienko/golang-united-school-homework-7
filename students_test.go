package coverage

import (
	"errors"
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

func TestPeopleLen(t *testing.T) {
	p := People{{"An", "Ivan", time.Now()}}
	t.Run("TestPeopleLen/1", func(t *testing.T) {
		got := p.Len()
		want := 1
		if !reflect.DeepEqual(got, want) {
			t.Errorf("Got = %v, want %v", got, want)
		}
	})
}

func TestPeopleLess(t *testing.T) {
	tm:=time.Now()
	p := People{
		{"1", "2", tm},
		{"2", "1", time.Now().Truncate(time.Minute)},
		{"1", "1", time.Now().Truncate(time.Minute)},
		{"1", "1", tm},
	}
	tests:=[]struct {
		name string
		i    int
		j    int
		want bool		
	}{
		{"TestPeopleLess/1", 0, 1, true},
		{"TestPeopleLess/2", 1, 0, false},
		{"TestPeopleLess/2", 1, 2, false},
		{"TestPeopleLess/2", 2, 1, true},
		{"TestPeopleLess/2", 2, 0, false},
		{"TestPeopleLess/2", 2, 3, false},
		{"TestPeopleLess/2", 0, 3, false},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := p.Less(test.i, test.j)
			if !reflect.DeepEqual(got, test.want) {
				t.Errorf("Got = %v, want %v", got, test.want)
			}
		})
	}
}

func TestPeopleSwap(t *testing.T) {
	p := People{
		{"1", "2", time.Now()},
		{"2", "1", time.Now().Truncate(time.Minute)},
		{"1", "1", time.Now().Truncate(time.Minute)},
		{"1", "1", time.Now()},
	}
	tests:=[]struct {
		name string
		i    int
		j    int
		want bool		
	}{
		{"TestPeopleSwap/1", 0, 1, true},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			pi := p[test.i]
			p.Swap(test.i, test.j)
			p.Swap(test.i, test.j)
			got := pi == p[test.i] 
			if !reflect.DeepEqual(got, test.want) {
				t.Errorf("Got = %v, want %v", got, test.want)
			}
		})
	}
}

var (
	errAny = errors.New("Error")
	m = Matrix{ 2, 3, []int{0,1,2,3,4,5}}
)
func TestMatrixNew(t *testing.T) {
		
	tests:=[]struct {
		name string
		str string
		want *Matrix
		err  error
	}{
		{"TestMatrixNew/1", "0 1 2\n3 4 5", &m, nil},
		{"TestMatrixNew/2", "0 1 2\n3 4 ", nil, errAny},
		{"TestMatrixNew/3", "0 1 2\n3 4 X", nil, errAny},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got, err := New(test.str)
			if !reflect.DeepEqual(got, test.want) || err == nil && test.err != nil || err != nil && test.err == nil  {
				t.Errorf("Got = %v,%v want %v, %v", got, err, test.want, test.err)
			}
		})
	}
}

func TestMatrixRows(t *testing.T) {
	tests:=[]struct {
		name string
		want [][]int
	}{
		{"1", [][]int{{0,1,2},{3,4,5}}},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := m.Rows()
			if !reflect.DeepEqual(got, test.want) {
				t.Errorf("Got = %v, want %v", got, test.want)
			}
		})
	}
}

func TestMatrixCols(t *testing.T) {
	tests:=[]struct {
		name string
		want [][]int
	}{
		{"1", [][]int{{0,3},{1,4}, {2,5}}},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := m.Cols()
			if !reflect.DeepEqual(got, test.want) {
				t.Errorf("Got = %v, want %v", got, test.want)
			}
		})
	}
}

func TestMatrixSet(t *testing.T) {
	tests:=[]struct {
		name string
		i,j int
		want int
		wantBool bool
	}{
		{"1", 0,0, 99, true},
		{"2", 0,100, 99, false},
		{"3", 100,0, 99, false},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := m.Set(test.i,test.j,test.want)
			if got != test.wantBool  {
				t.Errorf("Got = %v, want %v", got, test.want)
			}
		})
	}
}
