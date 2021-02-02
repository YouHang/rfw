package rfw

import (
	"reflect"
	"sort"
	"testing"
	"time"
)

func TestGetOutdatedPath(t *testing.T) {
	basepath := "test"
	tm, err := time.Parse("2006-01-02 15:04:05 UTC", "2017-12-12 18:40:00 UTC")
	if err != nil {
		t.Fatalf("Parse time failed: %v\n", err)
	}
	paths := []string{"test-20171212", "test-20171211", "test-20171113", "test-20171210", "test-20171207"}
	outdated := getOutdatedPath(basepath, paths, tm, 1, false)
	actual := []string{"test-20171113", "test-20171207", "test-20171210", "test-20171211", "test-20171212"}
	sort.Strings(outdated)
	if !reflect.DeepEqual(actual[:3], outdated) {
		t.Fatalf("expected: %v    get: %v\n", actual[:3], outdated)
	}
	outdated = getOutdatedPath(basepath, paths, tm, 10, false)
	if !reflect.DeepEqual(actual[:1], outdated) {
		t.Fatalf("expected: %v    get: %v\n", actual[:1], outdated)
	}
	outdated = getOutdatedPath(basepath, paths, tm, 100, false)
	if !reflect.DeepEqual([]string{}, outdated) {
		t.Fatalf("expected: %v    get: %v\n", []string{}, outdated)
	}

	basepath = "test-%Y%m%d.log"
	paths = []string{"test-20171212.log", "test-20171211.log", "test-20171113.log", "test-20171210.log", "test-20171207.log"}
	outdated = getOutdatedPath(basepath, paths, tm, 1, true)
	actual = []string{"test-20171113.log", "test-20171207.log", "test-20171210.log", "test-20171211.log", "test-20171212.log"}
	sort.Strings(outdated)
	if !reflect.DeepEqual(actual[:3], outdated) {
		t.Fatalf("expected: %v    get: %v\n", actual[:3], outdated)
	}
	outdated = getOutdatedPath(basepath, paths, tm, 10, true)
	if !reflect.DeepEqual(actual[:1], outdated) {
		t.Fatalf("expected: %v    get: %v\n", actual[:1], outdated)
	}
	outdated = getOutdatedPath(basepath, paths, tm, 100, true)
	if !reflect.DeepEqual([]string{}, outdated) {
		t.Fatalf("expected: %v    get: %v\n", []string{}, outdated)
	}

	basepath = "/%Y/%m/%d/test.log"
	paths = []string{"/2017/12/12/test.log", "/2017/12/11/testlog", "/2017/11/13/test.log", "/2017/12/10/test.log", "/2017/12/07/test.log"}
	outdated = getOutdatedPath(basepath, paths, tm, 1, true)
	actual = []string{"/2017/11/13/test.log", "/2017/12/07/test.log", "/2017/12/10/test.log", "/2017/12/11/test.log", "/2017/12/12/test.log"}
	sort.Strings(outdated)
	if !reflect.DeepEqual(actual[:3], outdated) {
		t.Fatalf("expected: %v    get: %v\n", actual[:3], outdated)
	}
	outdated = getOutdatedPath(basepath, paths, tm, 10, true)
	if !reflect.DeepEqual(actual[:1], outdated) {
		t.Fatalf("expected: %v    get: %v\n", actual[:1], outdated)
	}
	outdated = getOutdatedPath(basepath, paths, tm, 100, true)
	if !reflect.DeepEqual([]string{}, outdated) {
		t.Fatalf("expected: %v    get: %v\n", []string{}, outdated)
	}
}

func TestGeneratePath(t *testing.T) {
	basepath := "test.log"
	tm, err := time.Parse("2006-01-02 15:04:05 UTC", "2017-12-12 18:40:00 UTC")
	if err != nil {
		t.Fatalf("Parse time failed: %v\n", err)
	}
	generatedth := generatePath(basepath, false, tm)
	if generatedth != "test.log-20171212" {
		t.Fatalf("expected: %s    get: %s\n", "test.log-20171212", generatedth)
	}

	basepath = "test-%Y%m%d"
	generatedth = generatePath(basepath, true, tm)
	if generatedth != "test-20171212" {
		t.Fatalf("expected: %s    get: %s\n", "test-20171212", generatedth)
	}

	basepath = "test-%Y%m%d.log"
	generatedth = generatePath(basepath, true, tm)
	if generatedth != "test-20171212.log" {
		t.Fatalf("expected: %s    get: %s\n", "test-20171212.log", generatedth)
	}

	basepath = "test.log-%Y%m%d"
	generatedth = generatePath(basepath, true, tm)
	if generatedth != "test.log-20171212" {
		t.Fatalf("expected: %s    get: %s\n", "test.log-20171212", generatedth)
	}

	basepath = "/%Y/%m/%d/test.log"
	generatedth = generatePath(basepath, true, tm)
	if generatedth != "/2017/12/12/test.log" {
		t.Fatalf("expected: %s    get: %s\n", "/2017/12/12/test.log", generatedth)
	}
}
