import cv2

if __name__ == "__main__":
    video_name = "e304c938-c4ce-11ea-8c09-246e963a41ed-000002.ts"
    vc = cv2.VideoCapture(video_name)  # 读入视频文件
    video_len = int(vc.get(cv2.CAP_PROP_FRAME_COUNT))  # 视频总帧数
    video_width = int(vc.get(cv2.CAP_PROP_FRAME_WIDTH))  # 视频宽度
    video_height = int(vc.get(cv2.CAP_PROP_FRAME_HEIGHT))  # 视频高度
    video_fps = int(vc.get(cv2.CAP_PROP_FPS))  # 视频帧率

    print(video_len, video_width, video_height, video_fps)

    if vc.isOpened():  # 判断是否正常打开
        rval, frame = vc.read()
    else:
        rval = False
    ind = 1
    timeF = 1  # 视频帧计数间隔频率

    while rval:  # 循环读取视频帧
        rval, frame = vc.read()

        if (ind % timeF == 0):  # 每隔timeF帧进行存储操作
            cv2.imwrite('image/' + str(ind) + '.jpg', frame)  # 存储为图像
            # frame = cv2.resize(frame, (800, 800))  # 对图像进行resize
            cv2.imshow("Image", frame)
        ind = ind + 1
        # print(ind)
        cv2.waitKey(1)
    vc.release()

