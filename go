// math_test.go
package math

import "testing"

// TestAdd 测试 Add 函数
func TestAdd(t *testing.T) {
    // 测试案例 1
    result := Add(2, 3)
    expected := 5
    if result != expected {
        t.Errorf("Add(2, 3) = %d; want %d", result, expected)
    }

    // 测试案例 2
    result = Add(-1, 1)
    expected = 0
    if result != expected {
        t.Errorf("Add(-1, 1) = %d; want %d", result, expected)
    }

    // 测试案例 3
    result = Add(0, 0)
    expected = 0
    if result != expected {
        t.Errorf("Add(0, 0) = %d; want %d", result, expected)
    }
}
