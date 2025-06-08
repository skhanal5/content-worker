#!/bin/bash

temporal workflow start \
    --task-queue default \
    --type PublishClipsWorkflow \
    --input '{"Streamer":"plaqueboymax"}'