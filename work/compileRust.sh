#!/bin/bash

cd "Rust/"

rustc "Rust$1.rs" -o "Rust$1" 2>./"Rust$1.err"
