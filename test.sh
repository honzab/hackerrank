#!/usr/bin/env bash

# set -e

ERROR=false

function test_output() {
    O=$(diff "$1" "$2")
    if [[ $? != 0 ]]; then
        ERROR=true
        echo "ERROR"
        echo $O
    else
        echo "OK"
    fi
}

function py_test() {
    echo -n "Testing $1... "
    tmp=$(mktemp)
    cat "$2" | python "$1" > "$tmp"
    test_output "$3" "$tmp"
    rm "$tmp"
}

function java_test() {
    echo -n "Testing $1... "
    tmp=$(mktemp)
    filename=$(basename $1)
    class="${filename%.*}"
    javac $1
    cat $2 | java $class > "$tmp"
    test_output "$3" "$tmp"
    rm "$tmp"
}

function go_test() {
    echo -n "Testing $1... "
    tmp=$(mktemp)
    cat $2 | go run "$1" > "$tmp"
    test_output "$3" "$tmp"
    rm "$tmp"
}

function run_tests() {
    files=$(find $1 -type f -name "*.py" -or -name "*.go" -or -name "*.java")
    for f in $files; do
        basedir=$(dirname $f)
        basename=$(basename $f)
        pushd $basedir > /dev/null
        if [[ -f "tests/$basename.in" && -f "tests/$basename.out" ]]; then
            extension="${basename##*.}"
            if [[ $extension == "py" ]]; then
                py_test $f "tests/$basename.in" "tests/$basename.out"
            elif [[ $extension == "java" ]]; then
                java_test $f "tests/$basename.in" "tests/$basename.out"
            elif [[ $extension == "go" ]]; then
                go_test $f "tests/$basename.in" "tests/$basename.out"
            fi
        fi
        popd > /dev/null
    done
}

run_tests $PWD

if $ERROR; then
    exit 1
else
    exit 0
fi
