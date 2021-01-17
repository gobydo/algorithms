package tree

import (
	"reflect"
	"testing"
)

func TestBinarySearch_Insert(t *testing.T) {

	bst := NewBinarySearch()

	bst.Insert(8)
	bst.Insert(4)
	bst.Insert(2)
	bst.Insert(12)
	bst.Insert(9)
	bst.Insert(10)
	bst.Insert(5)

	tests := []struct {
		name string
		got  int
		want int
	}{
		{
			name: "root",
			got:  bst.root.key,
			want: 8,
		},
		{
			name: "root->left->left",
			got:  bst.root.left.left.key,
			want: 2,
		},
		{
			name: "root->left->right",
			got:  bst.root.left.right.key,
			want: 5,
		},
		{
			name: "root->right->left",
			got:  bst.root.right.left.key,
			want: 9,
		},
		{
			name: "root->right->left->right",
			got:  bst.root.right.left.right.key,
			want: 10,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.got != tt.want {
				t.Errorf("travers = %v, want = %v", tt.got, tt.want)
			}
		})
	}
}

func TestBinarySearch_Search(t *testing.T) {

	bst := NewBinarySearch()

	bst.Insert(8)
	bst.Insert(4)
	bst.Insert(2)
	bst.Insert(12)
	bst.Insert(9)
	bst.Insert(10)
	bst.Insert(5)

	type args struct {
		key int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
		{
			name: "search existing element",
			args: args{
				key: 12,
			},
			want: true,
		},
		{
			name: "search non existing element",
			args: args{
				key: 40,
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := bst.Search(tt.args.key); got != tt.want {
				t.Errorf("Search() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBinarySearch_Remove(t *testing.T) {
	type args struct {
		key int
	}
	tests := []struct {
		name    string
		args    args
		factory func() *BinarySearch
		travers func(start *Node) *Node
		want    int
	}{
		{
			name: "remove element from non empty tree",
			args: args{
				key: 9,
			},
			factory: func() *BinarySearch {
				bst := NewBinarySearch()

				bst.Insert(8)
				bst.Insert(4)
				bst.Insert(2)
				bst.Insert(12)
				bst.Insert(9)
				bst.Insert(10)
				bst.Insert(5)
				return bst
			},
			travers: func(start *Node) *Node { return start.right.left },
			want:    10,
		},
		{
			name: "remove node with two children",
			args: args{
				key: 12,
			},
			factory: func() *BinarySearch {
				bst := NewBinarySearch()

				bst.Insert(8)
				bst.Insert(4)
				bst.Insert(2)
				bst.Insert(12)
				bst.Insert(16)
				bst.Insert(14)
				bst.Insert(15)
				bst.Insert(9)
				bst.Insert(10)
				bst.Insert(5)
				return bst
			},
			travers: func(start *Node) *Node { return start.right },
			want:    14,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			bst := tt.factory()
			bst.Remove(tt.args.key)
			got := tt.travers(bst.root).key

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("got = %v, want = %v", got, tt.want)
			}
		})
	}
}
