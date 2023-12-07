package Q134

import (
	"reflect"
	"testing"
)

type parameters struct {
	head *ListNode
}

func Test1(t *testing.T) {

	params := parameters{
		head: &ListNode{Val: 3,
			Next: &ListNode{Val: 2,
				Next: &ListNode{Val: 0,
					Next: &ListNode{Val: -4},
				},
			},
		},
	}

	params.head.Next.Next.Next.Next = params.head.Next

	res := &ListNode{Val: 2,
		Next: &ListNode{Val: 0,
			Next: &ListNode{Val: -4},
		},
	}

	res.Next.Next.Next = res
	expected := res

	if testResult := detectCycle(params.head); reflect.DeepEqual(expected, testResult) {
		t.Log("sucess")
	} else {
		t.Error("fail coz expectec is ", expected, " and test result is ", testResult)
	}
}

func Test2(t *testing.T) {

	params := parameters{
		head: &ListNode{Val: 1,
			Next: &ListNode{Val: 2},
		},
	}

	params.head.Next.Next = params.head

	res := &ListNode{Val: 1,
		Next: &ListNode{Val: 2},
	}
	res.Next.Next = res
	expected := res

	if testResult := detectCycle(params.head); reflect.DeepEqual(expected, testResult) {
		t.Log("sucess")
	} else {
		t.Error("fail coz expectec is ", expected, " and test result is ", testResult)
	}
}

func Test3(t *testing.T) {

	params := parameters{
		head: &ListNode{Val: -1,
			Next: &ListNode{Val: -7,
				Next: &ListNode{Val: 7,
					Next: &ListNode{Val: -4,
						Next: &ListNode{Val: 19,
							Next: &ListNode{Val: 6,
								Next: &ListNode{Val: -9,
									Next: &ListNode{Val: -5,
										Next: &ListNode{Val: -2,
											Next: &ListNode{Val: -5},
										},
									},
								},
							},
						},
					},
				},
			},
		},
	}

	params.head.Next.Next.Next.Next.Next.Next.Next.Next.Next.Next = params.head.Next.Next.Next.Next.Next.Next

	res := &ListNode{Val: -9,
		Next: &ListNode{Val: -5,
			Next: &ListNode{Val: -2,
				Next: &ListNode{Val: -5},
			},
		},
	}
	res.Next.Next.Next.Next = res
	expected := res

	if testResult := detectCycle(params.head); reflect.DeepEqual(expected, testResult) {
		t.Log("sucess")
	} else {
		t.Error("fail coz expectec is ", expected, " and test result is ", testResult)
	}
}
