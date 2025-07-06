#!/bin/bash
temporal workflow start \
    --task-queue default \
    --type RetrieveClipsWorkflow \
    --input '{
        "streamer":"adapt",
        "limit":10,
        "filter":"LAST_WEEK"
    }'