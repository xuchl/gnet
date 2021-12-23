import cv2
 
capture=cv2.VideoCapture('./decrypted.ts')  
print(capture.isOpened())
num=0
while True: 
    ret,img=capture.read()  
    if not ret:
        break
    cv2.imwrite('video_cap/%s.jpg'%('pic_'+str(num)),img)  
    # if num==12:                                 
        # break
    print(num)
    num=num+1
    
capture.release()