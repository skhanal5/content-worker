#!/bin/bash
temporal workflow start \
    --task-queue default \
    --type EditWorkflow \
    --input '{
        "input_path":"./tmp/creator_plaqueboymax/id_BlatantGlamorousSharkSoonerLater-fowLZFq4QkVdSnrd.mp4",
        "output_path":"./edit/video.mp4",
        "strategy": "blurred_overlay"
    }'