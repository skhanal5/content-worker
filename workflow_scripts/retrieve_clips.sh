#!/bin/bash
temporal workflow start \
    --task-queue default \
    --type RetrieveClipsWorkflow \
    --input '{
        "streamer":"jasontheween",
        "days_ago":1,
        "top_n":3
    }'