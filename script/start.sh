#!/bin/bash
nohup python3 video_ffmpeg.py >> ./video_ffmpeg.out 2>&1 & echo $! > video_ffmpeg.pid