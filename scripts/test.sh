#!/bin/sh

# Find and run all test files
for OUTPUT in $(find . -name '*_test.go')
do
    echo "Running tests in: $OUTPUT"
    RESULT=$(go test $OUTPUT)
    
    if echo "$RESULT" | grep "FAIL"; then
        echo "\n❌ Test failed:"
        echo "$RESULT"
        exit 1
    else 
        echo "✅ Tests passed\n"
    fi
done

echo "✨ All tests passed successfully!"