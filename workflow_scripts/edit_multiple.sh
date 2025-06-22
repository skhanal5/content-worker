#!/bin/bash
temporal workflow start \
    --task-queue default \
    --type EditWorkflow \
    --input '{
        "output_directory":"edit",
        "videos": [
            {
                "input_directory":"tmp/creator_plaqueboymax",
                "strategy": "blurred_overlay",
                "title": "",
            }     
        ]
    }'