package test_bench

import (
	"log"
	"os/exec"
	"testing"
)

func BenchmarkCredstash(b *testing.B) {
	for n := 0; n < b.N; n++ {
		benchexec("credstash", "get", "foobar")
	}
}

func BenchmarkGCredstash(b *testing.B) {
	for n := 0; n < b.N; n++ {
		benchexec("gcredstash", "get", "foobar")
	}
}

func benchexec(prog string, args ...string) {
	cmd := exec.Command(prog, args...)
	if err := cmd.Run(); err != nil {
		log.Println("err", err)
	}
}
