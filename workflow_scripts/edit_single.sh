#!/bin/bash
temporal workflow start \
    --task-queue default \
    --type EditSingleWorkflow \
    --input '{
        "input_path":"./tmp/creator_jasontheween/id_584c47c4-8bec-4a01-aee2-25e3246fe0c9.mp4",
        "output_directory":"edit",
        "strategy": "blurred_overlay_stretched",
        "title": "hallo!"
    }'