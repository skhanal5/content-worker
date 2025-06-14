#!/bin/bash
temporal workflow start \
    --task-queue default \
    --type EditSingleWorkflow \
    --input '{
        "input_path":"tmp/creator_plaqueboymax/id_CrazyEmpathicLaptopANELE-Sz9WII7_i2Zj_KEI.mp4",
        "output_directory":"edit",
        "strategy": "blurred_overlay",
        "title": "he got too excited"
    }'