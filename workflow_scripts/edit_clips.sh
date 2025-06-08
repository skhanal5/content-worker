#!/bin/bash
temporal workflow start \
    --task-queue default \
    --type EditWorkflow \
    --input '{
        "input_directory":"tmp/creator_plaqueboymax",
        "output_directory":"edit",
        "strategy": "blurred_overlay"
    }'