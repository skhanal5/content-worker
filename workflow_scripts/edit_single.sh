#!/bin/bash
temporal workflow start \
    --task-queue default \
    --type SingleEditWorkflow \
    --input '{
        "input_path":"tmp/creator_plaqueboymax/id_CrazyEmpathicLaptopANELE-Sz9WII7_i2Zj_KEI.mp4",
        "output_path":"edit/broke_laptop.mp4",
        "strategy": "blurred_overlay",
        "title": "he got too excited"
    }'