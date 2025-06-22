#!/bin/bash
temporal workflow start \
    --task-queue default \
    --type RetrieveEditAndPublishWorkflow \
    --input '{
        "input_path":"./tmp/creator_adapt/id_0ea5e1d7-fcb3-4715-85d8-f5060277e4dc.mp4",
        "output_directory":"edit",
        "strategy": "blurred_overlay",
        "title": "hello world!"
    }'