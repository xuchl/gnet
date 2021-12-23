# -*- coding: utf-8 -*-
from Crypto.Cipher import AES
# import os

key=bytes(b'0871be8fcdc9fef8')
# kf=open("./key","r")
# key=kf.read().encode()
# key=
aes=AES.new(key,AES.MODE_CBC)
fpath="./Jyk2ofRV.ts"
output=open("decrypted.ts",'wb+')


# def readSize(path,size):
#     f = open(path,"rb")
#     try:
#         while True:
#             con = f.read(size)
#             yield con  # read()也会造成文件指针偏移
#             if con == '':
#                 return
#     except Exception as err:
#     	print(err)
#         # pass
#     f.close()
 
# for i in readSize(fpath,16):
#     content=aes.decrypt(i)
#     print(content)
#     output.write(content)
f=open(fpath,"rb")
i=f.read()
content=aes.decrypt(i)
output.write(content)
output.close();