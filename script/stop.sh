ps -ef | grep video_ffmpeg.py | grep -v grep | awk '{print $2}' | xargs kill -9
echo "已关闭video_ffmpeg"