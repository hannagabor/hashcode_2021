#!/bin/bash -xue
go fmt simulator.go
go run simulator.go a.txt out_a.txt
go run simulator.go b.txt out_b.txt
go run simulator.go c.txt out_c.txt
go run simulator.go d.txt out_d.txt
go run simulator.go e.txt out_e.txt
go run simulator.go f.txt out_f.txt
