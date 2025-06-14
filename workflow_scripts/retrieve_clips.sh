#!/bin/bash
temporal workflow start \
    --task-queue default \
    --type RetrieveClipsWorkflow \
    --input '{"Streamer":"silky"}'