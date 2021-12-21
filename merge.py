import cv2

if __name__ == "__main__":
    video_name = "daa05d6c-a6dc-11ea-9825-246e96398ca5-000001.ts"
    video_name2 = "daa05d6c-a6dc-11ea-9825-246e96398ca5-000002.ts"
    video_name3 = "daa05d6c-a6dc-11ea-9825-246e96398ca5-000003.ts"
    # vc = cv2.VideoCapture("./{}".format(video_name))  # 读入视频文件
    # vc2 = cv2.VideoCapture("./{}".format(video_name2))

    # rval, frame = vc.read()
    # rval2, frame2 = vc2.read()

    # video_len = int(vc.get(cv2.CAP_PROP_FRAME_COUNT))  # 视频总帧数
    # video_width = int(vc.get(cv2.CAP_PROP_FRAME_WIDTH))  # 视频宽度
    # video_height = int(vc.get(cv2.CAP_PROP_FRAME_HEIGHT))  # 视频高度
    # video_fps = int(vc.get(cv2.CAP_PROP_FPS))


    # print(video_len, video_width, video_height, video_fps)

    # VideoWriter = cv2.VideoWriter ("merge.avi", cv2.VideoWriter_fourcc('P','I','M','1'), video_fps,(video_width,video_height))
    # VideoWriter.write (frame)
    # # rval2, frame2 = vc2.read()
    # VideoWriter.write (frame2)

    # # eec4c02d-c4ce-11ea-8c09-246e963a41ed-000000.ts/__init__

    # VideoWriter.release ()
    # # cv2.destroyAllWindows()
    f = open("merge.ts", 'wb+')
    f.write(open(video_name, 'rb').read())
    f.write(open(video_name2, 'rb').read())
    f.write(open(video_name3, 'rb').read())

    # video_name3 = "daa05d6c-a6dc-11ea-9825-246e96398ca5-000003.ts"
    f.close()