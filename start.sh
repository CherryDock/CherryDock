#/bin/bash

# Start cherrydock api
./cherrydock &
status=$?
if [ $status -ne 0 ]; then
  echo "Fail to start cherrydock api: $status"
  exit $status
fi

# Start react ui
serve -s build
status=$?
if [ $status -ne 0 ]; then
  echo "Fail to start cherrydock api: $status"
  exit $status
fi
