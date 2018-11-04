#!/bin/bash
for i in {1..10}; do go run traffic-lights.go input/$i.txt; done