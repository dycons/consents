#!/bin/bash

# Run this script within the consents-app-dev container to
# attach the go delve debugger to the app
dlv attach --headless --listen=:2345 $(pgrep app) .
