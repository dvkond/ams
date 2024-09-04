/*
 * All rights reserved
 * Copyright © Diasoft
 * 2024
 * Address: 3/14, Polkovaya St., Moscow, 127018, Russian Federation
 * Tel.: +7 (495) 780 7575
 * Fax.: +7 (495) 780 7576
 * WEB: http://www.diasoft.com
 * Author: Dmitrii Kondratiev (dkondratiev@diasoft.ru)
 */
package amscore

import (
	"testing"
)

func TestHello(t *testing.T) {
	want := "Hello, World!"
	result := Hello()
	if result != want {
		t.Errorf("Hello() = %q, want %q", result, want)
	}
}
