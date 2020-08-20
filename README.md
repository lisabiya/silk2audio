# 📽Silk2Audio
`silk转wav,mp3的go实现`

 ## 感恩 😘 [silk-v3-decoder](https://github.com/kn007/silk-v3-decoder)提供的silk转码


> ### 工作流程

- 使用[silk-v3-decoder](https://github.com/kn007/silk-v3-decoder)编译对应平台的静态库
- cgo调用silk-v3-decoder将silk文件编译为pcm
- 使用[ffmpeg](https://github.com/FFmpeg/FFmpeg) 将pcm转码为wav,mp3实现


>### 备注
```
模板中实现了windows版本的编译，便于学习

linux版本需要在linux平台编译silk-v3-decoder，获得libSKP_SILK_SDK.a，替换掉原有的

ffmpeg.exe也替换成linux下的ffmpeg执行文件即可
```
