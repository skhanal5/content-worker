#!/bin/bash

temporal workflow start \
    --task-queue default \
    --type HelloWorldWorkflow \
    --input '{"name":"Bob"}'